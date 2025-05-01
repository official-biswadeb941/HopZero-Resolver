# HopZero-Resolver - Built to resolve. Hardened to resist.

HopZero is a fast, secure, and lightweight recursive DNS resolver written in Go. Designed for performance and resilience, it resolves DNS queries from the root with full control over caching, timeouts, and packet handling. Ideal for developers, security researchers, and systems engineers who need reliable name resolution without the bloat.

---

## Features

- ✨ **Recursive DNS Resolution** – Resolves from root servers without relying on upstream providers
- ⚡ **High Performance** – Built in Go with native concurrency for minimal latency
- 🛡️ **Security First** – Randomized source ports, transaction IDs, and response validation to mitigate spoofing
- ⌛ **TTL-Aware Caching** – In-memory LRU cache for efficient repeated lookups
- ⚙️ **Modular Design** – Clear separation of logic for extensibility and maintainability

---

## Installation

```bash
git clone https://github.com/official-biswadeb941/HopZero.git
cd HopZero
go build -o hopzero ./cmd
```

---

## Usage

```bash
./hopzero example.com
```

### Optional Flags

- `-dns <ip>`: Specify a fallback or testing upstream resolver
- `-timeout <seconds>`: Set a timeout for outbound queries
- `-cache`: Enable or disable caching (enabled by default)

---

## Project Structure

```
/hopzero
├── Modules/         # Contain Multiple Modules for DNS
|    ├── core/       # DNS packet processing, resolver logic
|    ├── cache/      # TTL-aware in-memory caching
|    ├── utils/      # Helper utilities and constants
|
└── App.go           # Application entry
```

---

## Testing

You can test the resolver using standard tools:

```bash
dig @127.0.0.1 -p 5353 example.com
```

Or by running HopZero with debugging enabled:

```bash
./hopzero -debug example.com
```

---

## Roadmap

- ✨ Recursive DNS Resolution – Resolves from root servers without relying on upstream providers
- ⚡ High Performance – Built in Go with native concurrency for minimal latency
- 🛡️ Security First – Randomized source ports, transaction IDs, and response validation to mitigate spoofing
- ⌛ TTL-Aware Caching – In-memory LRU cache for efficient repeated lookups
- ⚙️ Modular Design – Clear separation of logic for extensibility and maintainability

---

## Contributing

We welcome contributions! Please fork the repo, submit pull requests, or open issues to suggest improvements.

---

## License

This project is licensed under the Creative Commons Attribution 4.0 International License (CC BY 4.0) License. See [License](License.md)
 for details.

---

## Author

Built by Mr.Biswadeb Mukherjee – Ethical Hacker, Malware Developer & Cybersecurity Researcher.


