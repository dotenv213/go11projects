package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client
const bucketName = "golang-uploads"

func initMinio() {
	endpoint := "localhost:9000"
	accessKeyID := "admin"
	secretAccessKey := "password123"
	useSSL := false 

	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Bucket '%s' created successfully\n", bucketName)
	} else {
		log.Printf("Bucket '%s' already exists\n", bucketName)
	}
}

func main() {
	initMinio()

	app := fiber.New()

	app.Post("/upload", uploadFile)

	log.Fatal(app.Listen(":3000"))
}

func uploadFile(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("document") 
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to get file"})
	}

	fileStream, err := fileHeader.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer fileStream.Close()

	objectName := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	contentType := fileHeader.Header.Get("Content-Type")

	ctx := context.Background()

	info, err := minioClient.PutObject(ctx, bucketName, objectName, fileStream, fileHeader.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to upload to S3"})
	}

	reqParams := make(url.Values)
	presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, time.Hour*1, reqParams)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate link"})
	}

	return c.JSON(fiber.Map{
		"message":   "File uploaded successfully",
		"filename":  objectName,
		"size":      info.Size,
		"url":       presignedURL.String(), 
	})
}