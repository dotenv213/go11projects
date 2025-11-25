# Cloud File Uploader (S3 Compatible)

A microservice for uploading files to Object Storage (**MinIO**), simulating AWS S3 behavior locally.

## 🚀 Key Features
- **Object Storage:** Storing binary files (Images/PDFs) in buckets.
- **Multipart Upload:** Handling file streams from HTTP forms.
- **Presigned URLs:** Generating secure, temporary download links.
- **MinIO:** Running a local S3-compatible server with Docker.

## 🛠️ Tech Stack
- **Go**, **Fiber**, **MinIO SDK**, **Docker**