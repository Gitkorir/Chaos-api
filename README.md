# 🔥 Chaos-as-a-Service (CaaS) API

An intentionally unstable API built with **Go (Golang)** that injects failures into responses.

> Designed for developers who want to test retry logic, fallback systems, resilience strategies, and error-handling mechanisms safely.

---

## 🚀 Project Overview

Chaos-as-a-Service (CaaS) is a lightweight REST API that deliberately introduces unpredictable behavior into responses.

It simulates real-world backend instability such as:

- 🔴 Random `500 Internal Server Error`
- ⏳ Artificial latency (0–3 seconds)
- 💥 Corrupted JSON responses
- 🚦 Simulated rate limiting (`429 Too Many Requests`)

This makes it useful for:
- Testing frontend resilience
- Practicing retry logic
- Evaluating API client robustness
- Simulating disaster scenarios in development environments

---

## 🛠️ Built With

- **Go (Golang) 1.20+**
- Native `net/http` package
- Standard Go libraries only (no external dependencies)

---

## 📦 Project Structure

chaos-api/
│
├── go.mod
├── main.go
└── README.md


---

## ⚙️ System Requirements

- Go 1.20 or higher
- Windows / macOS / Linux
- Terminal or Command Prompt

Check Go installation:

go version


If not installed, download from:
https://go.dev/dl/

---

## 🔧 Installation & Setup

### 1️⃣ Clone the Repository

git clone https://github.com/your-username/chaos-api.git
cd chaos-api


### 2️⃣ Initialize Dependencies (if needed)

go mod tidy


### 3️⃣ Run the API

go run main.go


You should see:

🔥 Chaos API running on http://localhost:8080


---

## 🌐 API Endpoint

### `GET /chaos`

Example:

http://localhost:8080/chaos


You can test using:

curl http://localhost:8080/chaos


Or open it in your browser.

---

## 🎲 Possible Responses

### ✅ Normal JSON Response

```json
{
  "status": "success",
  "delay": 1423,
  "message": "Chaos response generated"
}
🔴 500 Error
500 Internal Server Error - Chaos Injected
🚦 429 Rate Limit
429 Too Many Requests - Simulated Rate Limit
💥 Corrupted JSON
{invalid_json: true,,,}
⏳ Delayed Response
The API may take 0–3 seconds before responding.

🧠 How It Works
The API introduces randomness using:

rand.Float32() for probabilistic failures

time.Sleep() for artificial latency

sync.Mutex for safe request counting

Conditional logic for simulated rate limiting

Every request has a chance of triggering a failure mode.

🧪 Use Cases
This API is ideal for:

Testing frontend error boundaries

Practicing exponential backoff strategies

Simulating production instability

Teaching API resilience concepts

DevOps experimentation

⚠️ Important Note
This project is intended only for development and testing environments.

Do not deploy this API in production unless you intentionally want chaos 😈

📈 Future Improvements
Add configurable chaos levels via query parameters

Add Docker containerization

Add logging dashboard

Add authentication simulation

Add configurable rate limit thresholds

📚 Learning Goals
This project demonstrates:

Building a REST API in Go

Handling HTTP requests and responses

Simulating real-world failure conditions

Basic concurrency with Mutex

Backend resilience concepts

📖 References
Official Go Documentation: https://go.dev/doc/

Go HTTP Package: https://pkg.go.dev/net/http

Go by Example: https://gobyexample.com/

👨‍💻 Author
Arnold Korir
Backend & Systems Development

