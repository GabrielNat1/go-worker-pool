# 🏗️ Architecture Overview

Modular, clean, and easy to navigate.

---

## 📂 Folder Structure

```go
worker-queue-system/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── task/
│   ├── queue/
│   ├── processor/
│   ├── worker/
│   └── monitor/
│
├── pkg/
│   ├── logger/
│   └── utils/
│
├── docs/
└── go.mod
```

---

## 📌 Layer Responsibilities
```
✅ cmd/api → application entrypoint 
```

```
✅ internal/task → Task model  
```

```
✅ internal/queue → FIFO queue  
```

```
✅ internal/processor → task execution  
```

```
✅ internal/worker → worker pool + DLQ  
```

```
✅ internal/monitor → monitoring endpoint  
```

```
✅ pkg → utilities and logger
```

---