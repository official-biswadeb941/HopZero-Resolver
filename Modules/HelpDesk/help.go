package helpdesk

import "fmt"

const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Cyan    = "\033[36m"
	Yellow  = "\033[33m"
	Green   = "\033[32m"
	Magenta = "\033[35m"
	White   = "\033[97m"
)

func ShowUsage() {
	// Banner with "Resolver"
	fmt.Println(Cyan + `
  _   _            _____              
 | | | | ___  _ __|__  /___ _ __ ___  
 | |_| |/ _ \| '_ \ / // _ \ '__/ _ \ 
 |  _  | (_) | |_) / /|  __/ | | (_) |
 |_| |_|\___/| .__/____\___|_|  \___/ 
             |_|                        ` + Bold + "Resolver" + Reset + Cyan + `
` + Reset)

	// Tagline
	fmt.Println(Bold + Cyan + "HopZero-Resolver – Built to resolve. Hardened to resist." + Reset)
	fmt.Println(Cyan + "A fast, secure, recursive DNS resolver written in Go.\n" + Reset)

	// Usage
	fmt.Println(Bold + Yellow + "Usage:" + Reset)
	fmt.Println(Green + "  -domain <domain>" + Reset + "         Domain to resolve (e.g., example.com)")
	fmt.Println(Green + "  -type <record>" + Reset + "           DNS record type (A, AAAA, MX, TXT, NS, CNAME, SOA)")
	fmt.Println(Green + "  -reverse <ip>" + Reset + "            Perform reverse DNS (PTR) lookup")
	fmt.Println(Green + "  -dns <ip:port>" + Reset + "           Use custom DNS server (default: system resolver)")
	fmt.Println(Green + "  -timeout <seconds>" + Reset + "       Set timeout for DNS queries (default: 5)")
	fmt.Println(Green + "  -cache" + Reset + "                   Enable/disable caching (enabled by default)")
	fmt.Println(Green + "  -debug" + Reset + "                   Print detailed resolution steps")
	fmt.Println(Green + "  -h, --help" + Reset + "               Show this help menu")
	fmt.Println()

	// Examples
	fmt.Println(Bold + Yellow + "Examples:" + Reset)
	fmt.Println(Magenta + "  ./hopzero -domain google.com -type A" + Reset)
	fmt.Println(Magenta + "  ./hopzero -reverse 8.8.8.8" + Reset)
	fmt.Println(Magenta + "  ./hopzero -domain github.com -type MX -dns 1.1.1.1:53" + Reset)
	fmt.Println(Magenta + "  ./hopzero -domain example.com -timeout 2 -cache=false" + Reset)
	fmt.Println(Magenta + "  ./hopzero -debug -domain openai.com" + Reset)
	fmt.Println()

	// Footer
	fmt.Println(White + "Author: Mr.Biswadeb Mukherjee | GitHub: @official-biswadeb941" + Reset)
	fmt.Println(White + "License: CC BY 4.0 | Project: https://github.com/official-biswadeb941/HopZero-Resolver" + Reset)
}
