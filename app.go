package main

import (
	"flag"
	"os"
	"strings"

	HelpDesk "github.official-biswadeb941/HopZero/Modules/HelpDesk"
	logutil "github.official-biswadeb941/HopZero/Modules/Logs"
	resolver "github.official-biswadeb941/HopZero/Modules/Resolver"
)

func main() {
	// ✅ Initialize logging before anything else
	logutil.Init(true)

	// ✅ Hook help flag to our HelpDesk module
	flag.Usage = HelpDesk.ShowUsage

	// ✅ Define flags
	domain := flag.String("domain", "", "Domain to resolve")
	recordType := flag.String("type", "A", "DNS record type")
	reverse := flag.String("reverse", "", "IP for reverse lookup")
	customDNS := flag.String("dns", "", "Custom DNS server")

	flag.Parse()

	// ✅ Show help if no actionable input is given
	if *domain == "" && *reverse == "" {
		HelpDesk.ShowUsage()
		os.Exit(0)
	}

	// ✅ Handle forward lookups (A, AAAA, MX, etc.)
	if *domain != "" {
		resolver.ResolveRecord(*domain, strings.ToUpper(*recordType), *customDNS)
	}

	// ✅ Handle reverse lookups (PTR)
	if *reverse != "" {
		resolver.ReverseLookup(*reverse)
	}
}
