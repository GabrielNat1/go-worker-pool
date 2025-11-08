# 🛠️ Setup & Installation

---

## ✅ Requirements
- Go 1.20+
- Git

---

## ✅ Installation

### Clone:
```bash
git clone https://github.com/GabrielNat1/go-worker-pool
```

### Enter folder:
```bash
cd worker-queue-demo
```

### Install modules:
```bash
go mod tidy
```

### Run:
```bash
go run cmd/api/main.go
```

---

## ✅ Testing

### Create task:
```http request
POST http://localhost:8080/task
```

### Monitor:
```http request
GET http://localhost:8080/monitor/status
```
---

## ✅ Stop
CTRL + C