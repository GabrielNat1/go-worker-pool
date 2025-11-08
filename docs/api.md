# 📡 API Documentation

---

# ✅ POST /task

Creates a new task.

### ✅ Example Request (generate_report)
```json 
{
"type": "generate_report",
"payload": { "name": "sales_q4" }
}
```

### ✅ Example Response

```json
{
"id": "uuid",
"status": "queued"
}
```

---

# ✅ GET /monitor/status

### ✅ Example
```json
{
"queue_length": 0,
"dlq_length": 1,
"queue_items": [],
"dlq_items": []
}
```