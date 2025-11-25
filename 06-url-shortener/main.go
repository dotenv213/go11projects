package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb *redis.Client

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func setupRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", 
		Password: "",               
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis: ", err)
	}
	fmt.Println("Redis connected successfully!")
}

func main() {
	setupRedis()

	app := fiber.New()

	app.Post("/api/v1", shortenURL)

	app.Get("/:url", resolveURL)

	log.Fatal(app.Listen(":3000"))
}


func shortenURL(c *fiber.Ctx) error {
	userIP := c.IP()
	
	limitKey := "rate_limit:" + userIP
	limitWindow := 30 * time.Minute 
	maxRequests := 10

	requestCount, err := rdb.Incr(ctx, limitKey).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Rate limit error"})
	}

	if requestCount == 1 {
		rdb.Expire(ctx, limitKey, limitWindow)
	}

	ttl, _ := rdb.TTL(ctx, limitKey).Result()

	if requestCount > int64(maxRequests) {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error":            "Rate limit exceeded",
			"rate_limit_reset": ttl / time.Minute, 
		})
	}

	body := new(Request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if body.URL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "URL is required"})
	}


	id := uuid.New().String()[:6]

	err = rdb.Set(ctx, id, body.URL, 24*time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to connect to server"})
	}

	resp := Response{
		URL:             body.URL,
		CustomShort:     "http://localhost:3000/" + id,
		Expiry:          24 * time.Hour,
		XRateRemaining:  maxRequests - int(requestCount),
		XRateLimitReset: ttl / time.Minute, 
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func resolveURL(c *fiber.Ctx) error {
	urlID := c.Params("url")

	value, err := rdb.Get(ctx, urlID).Result()
	
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in database"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal error"})
	}

	return c.Redirect(value, 301)
}