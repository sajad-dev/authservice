# 🛡️ Auth Service (Authentication & Authorization)
This project is an Authentication and Authorization Service built using Casbin, gRPC, and Envoy Proxy.
It allows you to manage user access control and service authentication efficiently.

## ⚙️ Features
- Authentication: User login and registration management
- Authorization: Access control using Casbin
- gRPC API: Expose services via gRPC
- Envoy Proxy: Request routing and management

## 🚀 Run the Project
1. First, prepare the project by running:

```bash
make
```

2. Then start the services using Docker Compose:
```bash
docker compose up --build
```
