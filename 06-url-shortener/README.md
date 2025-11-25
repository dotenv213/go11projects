# High-Performance URL Shortener

A clone of bit.ly using **Redis** for lightning-fast caching and temporary data storage.

## 🚀 Key Features
- **Redis Integration:** Using NoSQL for high-speed read/writes.
- **TTL (Time-To-Live):** Auto-expiring links (24h).
- **Rate Limiting:** Preventing abuse using Redis counters (IP-based).
- **UUID:** Generating unique short codes.

## 🛠️ Tech Stack
- **Go**, **Fiber**, **Redis**, **Docker**