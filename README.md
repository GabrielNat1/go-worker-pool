# Go Worker Queue Demo

![Go Version](https://img.shields.io/badge/Go-1.20+-blue.svg)
![Status](https://img.shields.io/badge/Status-Active-brightgreen)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Platform](https://img.shields.io/badge/Platform-Backend-lightgrey)

A clean and educational demonstration of how a **worker queue system** works internally using Go.  
This project implements a **task producer API**, an **in-memory queue**, a **worker pool**, a **retry system**, a **dead-letter queue**, and a **monitoring endpoint** — all in a simple and easy-to-understand architecture.

---

# 📌 Overview

This project was built for **learning purposes**.  
It demonstrates:

- How asynchronous processing works
- How worker pools are implemented
- How queues dispatch tasks
- How retries + DLQ work
- How to structure a small Go application

No external dependencies.  
No database.  
No message brokers.  
Just clean Go.

---

# 📖 Documentation

All documentation is organized inside the `docs/` folder:

| Section | Description |
|--------|-------------|
| ✅ [Features](./docs/features.md) | Full list of system capabilities |
| ✅ [Architecture](./docs/architecture.md) | Folder structure + diagrams |
| ✅ [API Guide](./docs/api.md) | How to send tasks + examples |
| ✅ [Worker System](./docs/worker-system.md) | Worker pool, retries, DLQ |
| ✅ [Queue](./docs/queue.md) | In-memory queue internals |
| ✅ [Task Processor](./docs/processor.md) | Task execution logic |
| ✅ [Monitoring](./docs/monitor.md) | Status endpoint documentation |
| ✅ [Setup](./docs/setup.md) | How to run and test |

---

# 📜 License

MIT License.