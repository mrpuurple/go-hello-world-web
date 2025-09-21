package handlers

import (
	"net"
	"net/http"
	"strings"

		"github.com/mrpuurple/go-hello-world-web/pkg/config"
		"github.com/mrpuurple/go-hello-world-web/pkg/models"
		"github.com/mrpuurple/go-hello-world-web/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// Parse and store IPv4 format
	ipv4 := parseIPv4Address(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ipv4", ipv4)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	remoteIPv4 := m.App.Session.GetString(r.Context(), "remote_ipv4")

	stringMap["remote_ip"] = remoteIP
	stringMap["remote_ipv4"] = remoteIPv4

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// parseIPv4Address extracts IPv4 format from the remote address
func parseIPv4Address(remoteAddr string) string {
	// Split host and port
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		// If splitting fails, use the original address
		host = remoteAddr
	}

	// Parse the IP
	ip := net.ParseIP(host)
	if ip == nil {
		return ""
	}

	// Check if it's IPv4
	if ipv4Addr := ip.To4(); ipv4Addr != nil {
		// It's IPv4
		return ipv4Addr.String()
	} else if ip.To16() != nil {
		// It's IPv6, check for special cases
		ipv6 := ip.String()
		if ipv6 == "::1" {
			// IPv6 loopback maps to IPv4 loopback
			return "127.0.0.1"
		}
		// Check if it's IPv4-mapped IPv6
		if strings.HasPrefix(ipv6, "::ffff:") {
			// Extract IPv4 from IPv4-mapped IPv6
			return strings.TrimPrefix(ipv6, "::ffff:")
		}
	}

	return ""
}
