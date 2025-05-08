# HopZero-Resolver – Built to Resolve. Hardened to Resist.

**HopZero** is a fast, secure, and lightweight recursive DNS resolver written in Go. It gives you complete control over DNS resolution by querying root servers directly. Built with performance, modularity, and security in mind, it’s the perfect tool for developers, security researchers, and systems engineers.

---

## ✨ Features

* 🔁 **Recursive DNS Resolution** – Queries root and authoritative servers directly, no upstream dependency
* ⚡ **High Performance** – Built with Go's native concurrency for speed and scalability
* 🛡️ **Security First** – Implements randomized source ports, transaction IDs, and strict response validation
* 🧠 **TTL-Aware Caching** – In-memory LRU caching speeds up repeated queries
* 🛠️ **Modular Architecture** – Easily extensible and maintainable codebase

---

## 📦 Installation

Clone and build the project:

```bash
git clone https://github.com/official-biswadeb941/HopZero.git
cd HopZero
go build -o hopzero ./cmd
```

Or for quick testing:

```bash
go run app.go
```

---

## 🚀 Basic Usage

Run HopZero with a target domain:

```bash
go run app.go -domain example.com
```

### 📘 Available Flags

| Flag               | Description                                                        |
| ------------------ | ------------------------------------------------------------------ |
| `-domain <domain>` | Domain to resolve (e.g., `example.com`)                            |
| `-type <record>`   | DNS record type (`A`, `AAAA`, `MX`, `TXT`, `NS`, `CNAME`, `SOA`)   |
| `-reverse <ip>`    | Perform reverse DNS (PTR) lookup (e.g., `8.8.8.8`)                 |
| `-dns <ip:port>`   | Use a specific upstream DNS server instead of recursive resolution |
| `-timeout <secs>`  | Set the timeout for outbound queries (default: `5`)                |
| `-cache`           | Enable or disable TTL-based caching (`true` by default)            |
| `-debug`           | Print detailed resolution steps and debug output                   |
| `-h, --help`       | Show help message                                                  |

---

## 🔍 Command Examples Explained

### 1. 🔎 A Record Lookup

```bash
go run app.go -domain google.com -type A
```

Resolves the IPv4 address (A record) for `google.com`.

---

### 2. 🔁 Reverse DNS (PTR) Lookup

```bash
go run app.go -reverse 8.8.8.8
```

Looks up the PTR record for `8.8.8.8`, typically to find the domain name of an IP address.

---

### 3. 📱 Query Using Custom DNS

```bash
go run app.go -domain github.com -type MX -dns 1.1.1.1:53
```

Resolves `MX` (mail exchange) records for `github.com` using Cloudflare’s DNS.

---

### 4. ⏱ Timeout with Cache Disabled

```bash
go run app.go -domain example.com -timeout 2 -cache=false
```

Runs a query with a 2-second timeout and bypasses the in-memory cache for a fresh result.

---

### 5. 🐞 Debug Mode

```bash
go run app.go -debug -domain openai.com
```

Enables debug mode to print every resolution step taken by HopZero.

---

## 📁 Project Structure

```
/hopzero
├── Modules/         # Core resolver modules
│   ├── core/        # DNS packet processing and resolution logic
│   ├── cache/       # TTL-aware LRU caching implementation
│   └── utils/       # Utility functions and constants
└── app.go           # Application entry point
```

---

## 🧪 Testing

Test using `dig` on a custom port (if applicable):

```bash
dig @127.0.0.1 -p 5353 example.com
```

Or simply use the binary with debugging:

```bash
go run app.go -debug -domain example.com
```

---

## 🗺️ Roadmap

* ✅ Recursive root resolution
* ✅ TTL-aware caching system
* ✅ Support for major DNS record types
* ✅ Flexible command-line usage
* 🕵️ DNSSEC validation (planned)
* 🌐 DoH / DoT (planned)
* 📊 Web-based dashboard (planned)
* 📊 Prometheus metrics export (planned)

---

## 🤝 Contributing

We welcome all contributions!

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes (`git commit -am 'Add foo feature'`)
4. Push to your branch (`git push origin feature/foo`)
5. Open a Pull Request

---

## 📜 License

This project is licensed under the [Creative Commons Attribution 4.0 International License (CC BY 4.0)](LICENSE.md).

---

## 👨‍💻 Author

**Mr. Biswadeb Mukherjee**
Ethical Hacker • Malware Developer • Cybersecurity Researcher
GitHub: [@official-biswadeb941](https://github.com/official-biswadeb941)

---

> HopZero – Built to resolve. Hardened to resist.
