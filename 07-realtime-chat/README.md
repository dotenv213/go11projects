# Real-time Chat System (WebSockets)

A real-time messaging application replacing HTTP with **WebSockets** for bidirectional communication.

## 🚀 Key Features
- **WebSocket Protocol:** Upgrading HTTP connection for real-time data.
- **The Hub Pattern:** Managing active clients using Maps and Mutexes.
- **Broadcasting:** Distributing messages to all connected users instantly.
- **Race Condition Handling:** Using `sync.Mutex` for thread safety.

## 🛠️ Tech Stack
- **Go**, **Fiber**, **WebSockets**, **JavaScript** (Frontend)