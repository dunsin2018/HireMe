package auth

import "testing"

func TestCookieDomainForHost(t *testing.T) {
	domains := []string{"hireme.dev", "hireme.example.com"}
	cases := []struct {
		host string
		want string
	}{
		{"hireme.dev", ".hireme.dev"},       // apex
		{"apply.hireme.dev", ".hireme.dev"}, // subdomain shares
		{"hireme.example.com", ".hireme.example.com"},         // second domain
		{"agent.hireme.example.com", ".hireme.example.com"},   // its subdomain
		{"hireme.example.com:8080", ".hireme.example.com"},    // :port ignored
		{"localhost", ""},                       // dev -> host-only
		{"evil.com", ""},                        // unknown -> host-only
		{"nothireme.dev", ""},                 // suffix must be dot-bounded
		{"hireme.dev.evil.com", ""},           // not our domain
	}
	for _, tc := range cases {
		if got := CookieDomainForHost(tc.host, domains); got != tc.want {
			t.Errorf("CookieDomainForHost(%q) = %q, want %q", tc.host, got, tc.want)
		}
	}
}

// With no configured domains (dev) every host is host-only.
func TestCookieDomainForHost_NoDomains(t *testing.T) {
	if got := CookieDomainForHost("hireme.example.com", nil); got != "" {
		t.Errorf("got %q, want empty when no domains configured", got)
	}
}
