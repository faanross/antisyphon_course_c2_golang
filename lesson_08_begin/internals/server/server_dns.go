package server

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"

	"c2framework/internals/config"
	"c2framework/internals/control"
)

// DNSServer implements the Server interface for DNS
type DNSServer struct {
	addr   string
	server *dns.Server
}

// NewDNSServer creates a new DNS server
func NewDNSServer(cfg *config.ServerConfig) *DNSServer {
	return &DNSServer{
		addr: fmt.Sprintf("%s:%s", cfg.ListeningInterface, cfg.ListeningPort),
	}
}

// Start implements Server.Start for DNS
func (s *DNSServer) Start() error {
	// Create and configure the DNS server
	s.server = &dns.Server{
		Addr:    s.addr,
		Net:     "udp",
		Handler: dns.HandlerFunc(s.handleDNSRequest),
	}

	// Start server
	return s.server.ListenAndServe()
}

// Stop implements Server.Stop for DNS
func (s *DNSServer) Stop() error {
	if s.server == nil {
		return nil
	}
	log.Println("Stopping DNS server...")
	return s.server.Shutdown()
}

// handleDNSRequest is our DNS Server's handler
func (s *DNSServer) handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	// Create response message
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	// Process each question
	for _, question := range r.Question {
		// We only handle A records for now
		if question.Qtype != dns.TypeA {
			continue
		}

		// Log the query
		log.Printf("DNS query for: %s", question.Name)

		// Check if we should transition
		// TODO: create shouldTransition, call our new method CheckAndReset()

		// TODO create responseIP as string

		if shouldTransition {
			// TODO: if bool is true, set responseIP to 69.69.69.69
			log.Printf("DNS: Sending transition signal (69.69.69.69)")
		} else {
			// TODO: if bool is false, set responseIP to 42.42.42.42
			log.Printf("DNS: Normal response (42.42.42.42)")
		}

		// Create the response with the appropriate IP
		rr := &dns.A{
			Hdr: dns.RR_Header{
				Name:   question.Name,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    300,
			},
			// TODO: set A with net.ParseIP, pass our responseIP variable
		}
		m.Answer = append(m.Answer, rr)
	}

	// Send response
	w.WriteMsg(m)
}
