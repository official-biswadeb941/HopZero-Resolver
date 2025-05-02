package main

import (
	"flag"
	"os"
	"strings"

	HelpDesk "github.com/official-biswadeb941/HopZero/Modules/HelpDesk"
	logutil "github.com/official-biswadeb941/HopZero/Modules/Logs"
	resolver "github.com/official-biswadeb941/HopZero/Modules/Resolver"
)

func main() {
	// ✅ Initialize logging
	logutil.Init(true)

	// ✅ Hook custom help banner
	flag.Usage = HelpDesk.ShowUsage

	// ✅ Define command-line flags
	domain := flag.String("domain", "", "Domain to resolve")
	recordType := flag.String("type", "A", "DNS record type")
	reverse := flag.String("reverse", "", "IP for reverse lookup")
	customDNS := flag.String("dns", "", "Custom DNS server (ip:port or ip)")
	timeout := flag.Int("timeout", 5, "Timeout for DNS queries in seconds")
	debug := flag.Bool("debug", false, "Enable debug output")
	cache := flag.Bool("cache", true, "Enable/disable caching")

	flag.Parse()

	// ✅ Show help if no actionable flag is provided
	if *domain == "" && *reverse == "" {
		HelpDesk.ShowUsage()
		os.Exit(0)
	}

	// ✅ Handle domain resolution
	if *domain != "" {
		resolver.ResolveRecord(*domain, strings.ToUpper(*recordType), *customDNS, *timeout, *debug, *cache)
	}

	// ✅ Handle reverse lookup
	if *reverse != "" {
		resolver.ReverseLookup(*reverse, *timeout, *debug)
	}
}
