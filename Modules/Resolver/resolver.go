package resolver

import (
	"context"
	"net"
	"time"
	"github.com/miekg/dns"
	"strings"
    "github.official-biswadeb941/HopZero/Modules/Logs"
)

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
func ResolveRecord(domain string, recordType string, dnsServer string) {
	resolver := getResolver(dnsServer)

	switch recordType {
	case "A":
		resolveA(domain, resolver)
	case "AAAA":
		resolveAAAA(domain, resolver)
	case "MX":
		resolveMX(domain, resolver)
	case "TXT":
		resolveTXT(domain, resolver)
	case "NS":
		resolveNS(domain, resolver)
	case "CNAME":
		resolveCNAME(domain, resolver)
	case "SOA":
		resolveSOA(domain, dnsServer)
	default:
		logutil.Logger.Printf("Unsupported record type: %s", recordType)
	}
}

// A record resolver
func resolveA(domain string, r *net.Resolver) {
	logutil.Logger.Println("Resolving A records...")
	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		ips, err := r.LookupIPAddr(ctx, domain)
		if err != nil {
			return err
		}
		for _, ip := range ips {
			if ip.IP.To4() != nil {
				logutil.Logger.Println(ip.IP)
			}
		}
		return nil
	})
}

// AAAA record resolver
func resolveAAAA(domain string, r *net.Resolver) {
	logutil.Logger.Println("Resolving AAAA records...")
	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		ips, err := r.LookupIPAddr(ctx, domain)
		if err != nil {
			return err
		}
		for _, ip := range ips {
			if ip.IP.To16() != nil && ip.IP.To4() == nil {
				logutil.Logger.Println(ip.IP)
			}
		}
		return nil
	})
}

// MX record resolver
func resolveMX(domain string, r *net.Resolver) {
	logutil.Logger.Println("Resolving MX records...")
	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		mx, err := r.LookupMX(ctx, domain)
		if err != nil {
			return err
		}
		for _, record := range mx {
			logutil.Logger.Printf("%s (Pref: %d)", record.Host, record.Pref)
		}
		return nil
	})
}

// TXT record resolver
func resolveTXT(domain string, r *net.Resolver) {
	logutil.Logger.Println("Resolving TXT records...")
	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		txts, err := r.LookupTXT(ctx, domain)
		if err != nil {
			return err
		}
		for _, txt := range txts {
			logutil.Logger.Println(txt)
		}
		return nil
	})
}

// NS record resolver
func resolveNS(domain string, r *net.Resolver) {
	logutil.Logger.Println("Resolving NS records...")
	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		ns, err := r.LookupNS(ctx, domain)
		if err != nil {
			return err
		}
		for _, n := range ns {
			logutil.Logger.Println(n.Host)
		}
		return nil
	})
}

// CNAME record resolver
func resolveCNAME(domain string, r *net.Resolver) {
	logutil.Logger.Println("Resolving CNAME record...")
	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
		cname, err := r.LookupCNAME(ctx, domain)
		if err != nil {
			return err
		}
		logutil.Logger.Println("CNAME:", cname)
		return nil
	})
}

// SOA is not supported natively in net package — notify user
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
		logutil.Logger.Println("Error querying SOA:", err)
		return
	}

	// Print detailed SOA record information
	for _, ans := range r.Answer {
		if soa, ok := ans.(*dns.SOA); ok {
			logutil.Logger.Printf("SOA Record for %s:", domain)
			logutil.Logger.Printf("  Primary NS: %s", soa.Ns)
			logutil.Logger.Printf("  Admin Email: %s", soa.Mbox)
			logutil.Logger.Printf("  Serial: %d", soa.Serial)
			logutil.Logger.Printf("  Refresh: %d", soa.Refresh)
			logutil.Logger.Printf("  Retry: %d", soa.Retry)
			logutil.Logger.Printf("  Expire: %d", soa.Expire)
			logutil.Logger.Printf("  Minimum TTL: %d", soa.Minttl)
		}
	}
}

// ReverseLookup resolves PTR records from IP addresses
func ReverseLookup(ip string) {
	logutil.Logger.Printf("Performing reverse lookup for %s...", ip)
	withTimeoutAndRetry(3, 3*time.Second, func(ctx context.Context) error {
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
