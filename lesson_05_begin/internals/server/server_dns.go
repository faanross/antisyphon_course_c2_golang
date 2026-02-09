package server

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"

	"c2framework/internals/config"
)

// DNSServer implements the Server interface for DNS
type DNSServer struct {
	// TODO: create addr of type string
	// TODO: create server of type *dns.Server
}

// NewDNSServer creates a new DNS server
func NewDNSServer(cfg *config.ServerConfig) *DNSServer {
	return &DNSServer{
		// TODO: Assign addr with ListeningInterface and ListeningPort
	}
}

// Start implements Server.Start for DNS
func (s *DNSServer) Start() error {
	// Create and configure the DNS server
	s.server = &dns.Server{
		// TODO: assign Addr field as s.addr
		// TODO: assign Net field as udp
		// TODO: assign Handler field by calling dns.HandlerFunc, passing s.handleDNSRequest as argument
	}

	// Start server
	// TODO return a call to s.server.ListenAndServe()
}

// Stop implements Server.Stop for DNS
func (s *DNSServer) Stop() error {
	// TODO: if there is no server, return nil

	log.Println("Stopping DNS server...")

	// TODO: return method call - s.server.ShutDown()
}

// handleDNSRequest is our DNS Server's handler
func (s *DNSServer) handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	// Create response message

	// TODO: create message m, call new() pass dns.Msg as argument
	// TODO: call method SetReply() on m, pass r as argument
	// TODO: set m.Authoritative to true

	// Process each question
	for _, question := range r.Question {
		// We only handle A records for now

		// TODO: Conditional asks if question.Qtype is not equal to dns.TypeA, then continue

		// Log the query
		log.Printf("DNS query for: %s", question.Name)

		// For now, always return 42.42.42.42
		rr := &dns.A{
			Hdr: dns.RR_Header{
				Name:   question.Name,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    300,
			},
			// TODO set A to value returned by net.ParseIP, pass string 42.42.42.42 as argument
		}
		// TODO: append rr to m.Answer using append()
	}

	// Send response
	// TODO: call WriteMsg on w, pass m as argument
}
