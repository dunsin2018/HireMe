package auth

import "testing"

func TestCookieDomainForHost(t *testing.T) {
	domains := []string{"hireme.dev", "hiremetoday.ca"}
	cases := []struct {
		host string
		want string
	}{
		{"hireme.dev", ".hireme.dev"},       // apex
		{"apply.hireme.dev", ".hireme.dev"}, // subdomain shares
		{"hiremetoday.ca", ".hiremetoday.ca"},         // second domain
		{"agent.hiremetoday.ca", ".hiremetoday.ca"},   // its subdomain
		{"hiremetoday.ca:8080", ".hiremetoday.ca"},    // :port ignored
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
	if got := CookieDomainForHost("hiremetoday.ca", nil); got != "" {
		t.Errorf("got %q, want empty when no domains configured", got)
	}
}
