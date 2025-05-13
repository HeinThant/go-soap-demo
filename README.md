# Golang SOAP Client and Server Demo

This project demonstrates a simple **SOAP (Simple Object Access Protocol)** implementation using **Go (Golang)**. It includes both a minimal SOAP server and a client that communicates with it using XML over HTTP.

---

## 📦 Project Structure

```
go-soap-demo/
├── go.mod
├── server/
│   └── main.go            # SOAP server with basic Add operation and auth
├── client/
│   ├── main.go            # SOAP client that sends a request and parses the response
│   └── response.go        # XML parsing structs
```

---

## 🚀 How to Run

### 1. Clone this repo

```bash
git clone https://github.com/HeinThant/go-soap-demo.git
cd go-soap-demo
```

### 2. Start the SOAP Server

```bash
cd server
go run main.go
```

> Server runs on `http://localhost:8080/soap` and uses **Basic Auth**:
- **Username:** `admin`
- **Password:** `Admin@123`

---

### 3. Run the SOAP Client

In a new terminal:

```bash
cd client
go run .
```

✅ You should see:
- Raw XML response from the server
- Parsed result (e.g., `Parsed Result: 25`)

---

## 🔐 Features

- Manual SOAP request handling using standard `net/http`
- Custom XML marshaling/unmarshaling with `encoding/xml`
- Basic HTTP authentication on the server
- Client-side parsing of XML response
- Clean, GitHub-friendly folder structure

---

## 📚 Learning Purpose

This project is ideal for:

- Understanding how SOAP works at the XML + HTTP level
- Integrating with legacy SOAP systems using Go
- Building basic XML-based APIs for testing or automation

---

## 📜 License

This project is licensed under the MIT License.

---

## 🙋‍♂️ Author

**Hein Thant Thu Shein**  
Feel free to connect or contribute.