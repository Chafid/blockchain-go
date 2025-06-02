# 🧱 Simple Blockchain in Go

This is a basic implementation of a blockchain in Go. It demonstrates core blockchain concepts like:
- Block structure
- SHA-256 hashing
- Proof-of-Work (mining)
- Block chaining via hashes

## 🛠️ Project Structure

```
blockchain-go/
├── cmd/
│   └── blockchain/
│       └── main.go         # CLI entrypoint
├── internal/
│   └── blockchain/
│       └── blockchain.go   # Blockchain logic
├── go.mod
└── README.md
```

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/blockchain-go.git
cd blockchain-go
```

### 2. Run the Program

```bash
go run ./cmd/blockchain
```

### 3. Example Output

```text
Mining block 1...
Block mined! Hash: 000a1b4c...
Mining block 2...
Block mined! Hash: 000d9f2e...

Block #1:
Data: First block after Genesis
Hash: ...
PrevHash: ...

Block #2:
Data: Another block
Hash: ...
PrevHash: ...
```

## ⚙️ Configuration

### Difficulty

In `blockchain.go`, you can change the mining difficulty by shortening or adding the target prefix.
For example, you can increase the difficulty by adding another zero in the prefix, "000"

```go
const targetPrefix = "00" // Difficulty: hash must start with this
```

## 🧪 Extend This Project

Here are a few ideas to expand this:

- 🧑‍💻 Create a web API to submit and view blocks (using `net/http` or `gin`)
- 💾 Save blockchain data to a file or database
- 📡 Connect multiple nodes with peer-to-peer networking
- 📉 Track and display mining stats or block explorer UI

---

## 📚 Learning Goals

This project is designed to help you learn:
- How blockchains work internally
- Hashing and nonce mechanisms
- Structs and methods in Go
- Command-line Go applications

---

## 📄 License

MIT — feel free to use, copy, and modify.
