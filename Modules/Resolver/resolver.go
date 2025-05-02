package resolver

import (
	"context"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
	logutil "github.com/official-biswadeb941/HopZero/Modules/Logs"
)

type cacheEntry struct {
	Record interface{}
	Expiry time.Time
}

var dnsCache = struct {
	sync.RWMutex
	data map[string]cacheEntry
}{data: make(map[string]cacheEntry)}

// getResolver returns a custom DNS resolver using a specific DNS server, or the system resolver by default.
func getResolver(dnsServer string) *net.Resolver {
	if dnsServer == "" {
		return net.DefaultResolver
	}

	// Ensure the DNS server includes port
	if _, _, err := net.SplitHostPort(dnsServer); err != nil {
		dnsServer = net.JoinHostPort(dnsServer, "53")
	}

	// Custom resolver using provided DNS server
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", dnsServer)
		},
	}
}

// Check cache and return record if it's still valid
func getCachedRecord(domain, recordType string) (interface{}, bool) {
	cacheKey := domain + ":" + recordType

	// Lock for reading
	dnsCache.RLock()
	defer dnsCache.RUnlock()

	entry, exists := dnsCache.data[cacheKey]
	if !exists || time.Now().After(entry.Expiry) {
		return nil, false // Cache miss or expired cache
	}
	return entry.Record, true
}

// Cache the result
func setCache(domain, recordType string, record interface{}) {
	cacheKey := domain + ":" + recordType
	dnsCache.Lock()
	defer dnsCache.Unlock()
	dnsCache.data[cacheKey] = cacheEntry{
		Record: record,
		Expiry: time.Now().Add(5 * time.Minute), // Cache for 5 minutes
	}
}

// withTimeoutAndRetry runs DNS lookups with timeout and retry support.
func withTimeoutAndRetry(attempts int, timeout time.Duration, fn func(ctx context.Context) error) error {
	var lastErr error
	for i := 0; i < attempts; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		lastErr = fn(ctx)
		if lastErr == nil {
			return nil
		}
		logutil.Logger.Printf("Attempt %d failed: %v", i+1, lastErr)
	}
	return lastErr
}

// ResolveRecord resolves various DNS record types
func ResolveRecord(domain string, recordType string, dnsServer string, timeout int, debug bool, cache bool) {
	resolver := getResolver(dnsServer)

	switch recordType {
	case "A":
		resolveA(domain, resolver, cache)
	case "AAAA":
		resolveAAAA(domain, resolver, cache)
	case "MX":
		resolveMX(domain, resolver, cache)
	case "TXT":
		resolveTXT(domain, resolver, cache)
	case "NS":
		resolveNS(domain, resolver, cache)
	case "CNAME":
		resolveCNAME(domain, resolver, cache)
	case "SOA":
		resolveSOA(domain, dnsServer)
	default:
		logutil.Logger.Printf("Unsupported record type: %s", recordType)
	}
}

// A record resolver with caching
func resolveA(domain string, r *net.Resolver, cache bool) {
	logutil.Logger.Println("Resolving A records...")

	// Check cache first
	if cache {
		if cached, found := getCachedRecord(domain, "A"); found {
			logutil.Logger.Println("Cache hit for A record:", cached)
			return
		}
	}

	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		ips, err := r.LookupIPAddr(ctx, domain)
		if err != nil {
			return err
		}
		var ipsToStore []net.IP
		for _, ip := range ips {
			if ip.IP.To4() != nil {
				ipsToStore = append(ipsToStore, ip.IP)
				logutil.Logger.Println(ip.IP)
			}
		}

		// Cache the result
		if cache && len(ipsToStore) > 0 {
			setCache(domain, "A", ipsToStore)
		}
		return nil
	})
}

// AAAA record resolver with caching
func resolveAAAA(domain string, r *net.Resolver, cache bool) {
	logutil.Logger.Println("Resolving AAAA records...")

	// Check cache first
	if cache {
		if cached, found := getCachedRecord(domain, "AAAA"); found {
			logutil.Logger.Println("Cache hit for AAAA record:", cached)
			return
		}
	}

	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		ips, err := r.LookupIPAddr(ctx, domain)
		if err != nil {
			return err
		}
		var ipsToStore []net.IP
		for _, ip := range ips {
			if ip.IP.To16() != nil && ip.IP.To4() == nil {
				ipsToStore = append(ipsToStore, ip.IP)
				logutil.Logger.Println(ip.IP)
			}
		}

		// Cache the result
		if cache && len(ipsToStore) > 0 {
			setCache(domain, "AAAA", ipsToStore)
		}
		return nil
	})
}

// MX record resolver with caching
func resolveMX(domain string, r *net.Resolver, cache bool) {
	logutil.Logger.Println("Resolving MX records...")

	// Check cache first
	if cache {
		if cached, found := getCachedRecord(domain, "MX"); found {
			logutil.Logger.Println("Cache hit for MX record:", cached)
			return
		}
	}

	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		mx, err := r.LookupMX(ctx, domain)
		if err != nil {
			return err
		}
		for _, record := range mx {
			logutil.Logger.Printf("%s (Pref: %d)", record.Host, record.Pref)
		}

		// Cache the result
		if cache && len(mx) > 0 {
			setCache(domain, "MX", mx)
		}
		return nil
	})
}

// TXT record resolver with caching
func resolveTXT(domain string, r *net.Resolver, cache bool) {
	logutil.Logger.Println("Resolving TXT records...")

	// Check cache first
	if cache {
		if cached, found := getCachedRecord(domain, "TXT"); found {
			logutil.Logger.Println("Cache hit for TXT record:", cached)
			return
		}
	}

	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		txts, err := r.LookupTXT(ctx, domain)
		if err != nil {
			return err
		}
		for _, txt := range txts {
			logutil.Logger.Println(txt)
		}

		// Cache the result
		if cache && len(txts) > 0 {
			setCache(domain, "TXT", txts)
		}
		return nil
	})
}

// NS record resolver with caching
func resolveNS(domain string, r *net.Resolver, cache bool) {
	logutil.Logger.Println("Resolving NS records...")

	// Check cache first
	if cache {
		if cached, found := getCachedRecord(domain, "NS"); found {
			logutil.Logger.Println("Cache hit for NS record:", cached)
			return
		}
	}

	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		ns, err := r.LookupNS(ctx, domain)
		if err != nil {
			return err
		}
		for _, n := range ns {
			logutil.Logger.Println(n.Host)
		}

		// Cache the result
		if cache && len(ns) > 0 {
			setCache(domain, "NS", ns)
		}
		return nil
	})
}

// CNAME record resolver with caching
func resolveCNAME(domain string, r *net.Resolver, cache bool) {
	logutil.Logger.Println("Resolving CNAME record...")

	// Check cache first
	if cache {
		if cached, found := getCachedRecord(domain, "CNAME"); found {
			logutil.Logger.Println("Cache hit for CNAME record:", cached)
			return
		}
	}

	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		cname, err := r.LookupCNAME(ctx, domain)
		if err != nil {
			return err
		}
		logutil.Logger.Println("CNAME:", cname)

		// Cache the result
		if cache && cname != "" {
			setCache(domain, "CNAME", cname)
		}
		return nil
	})
}

// SOA record resolver (no caching as it's done with raw DNS query)// SOA record resolver (no caching as it's done with raw DNS query)
func resolveSOA(domain string, dnsServer string) {
	logutil.Logger.Println("Resolving SOA record for domain:", domain)

	if dnsServer == "" {
		dnsServer = "8.8.8.8:53"
	} else if !strings.Contains(dnsServer, ":") {
		dnsServer += ":53"
	}

	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeSOA)

	r, _, err := c.Exchange(m, dnsServer)
	if err != nil {
		logutil.Logger.Printf("Error querying DNS server %s: %v", dnsServer, err)
		return
	}

	if len(r.Answer) > 0 {
		for _, ans := range r.Answer {
			if soa, ok := ans.(*dns.SOA); ok {
				logutil.Logger.Println("SOA:", soa.Ns, soa.Mbox, soa.Serial, soa.Refresh, soa.Retry, soa.Expire, soa.Minttl)
			}
		}
	} else {
		logutil.Logger.Println("No SOA record found for domain:", domain)
	}
}

// ReverseLookup resolves PTR records (reverse DNS lookup)
func ReverseLookup(ip string, timeout int, debug bool) {
	logutil.Logger.Printf("Performing reverse lookup for %s...", ip)
	withTimeoutAndRetry(3, time.Duration(timeout)*time.Second, func(ctx context.Context) error {
		names, err := net.DefaultResolver.LookupAddr(ctx, ip)
		if err != nil {
			return err
		}
		for _, name := range names {
			logutil.Logger.Println(name)
		}
		return nil
	})
}
