package agent

import (
	"context"
	"fmt"
	"log"

	"github.com/miekg/dns"
)

// DNSAgent implements the Agent interface for DNS
type DNSAgent struct {
	// TODO: create serverAddr of type string
	// TODO: create client of type *dns.Client
}

// NewDNSAgent creates a new DNS client
func NewDNSAgent(serverIP string, serverPort string) *DNSAgent {
	return &DNSAgent{
		serverAddr: fmt.Sprintf("%s:%s", serverIP, serverPort),
		client:     new(dns.Client),
	}
}

// Send implements Agent.Send for DNS
func (c *DNSAgent) Send(ctx context.Context) ([]byte, error) {
	// Create DNS query message
	// TODO: create m by calling new(), pass dns.Msg as arg

	// For now, we'll query for a fixed domain
	domain := "www.thisdoesnotexist.com."
	// TODO: call SetQuestion() on m, pass the domain and specify A type record
	log.Printf("Sending DNS query for: %s", domain)

	// Send query
	r, _, err := c.client.Exchange(m, c.serverAddr)
	if err != nil {
		return nil, fmt.Errorf("DNS exchange failed: %w", err)
	}

	// Check if we got an answer
	// TODO: check if we got an answer by seeing if len of r.Answer is 0

	// Extract the first A record
	for _, ans := range r.Answer {
		if a, ok := ans.(*dns.A); ok {
			// Return the IP address as string
			// TODO: set ipStr as a.A.String()
			log.Printf("Received DNS response: %s -> %s", domain, ipStr)s
			return []byte(ipStr), nil
		}
	}

	return nil, fmt.Errorf("no A record in response")
}
