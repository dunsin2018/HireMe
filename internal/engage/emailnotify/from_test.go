package emailnotify_test

import (
	"testing"

	"github.com/dunsin2018/HireMe/internal/engage/emailnotify"
)

func TestFrom(t *testing.T) {
	for _, tc := range []struct {
		name, display, address, want string
	}{
		{"bare address gains a name", "hireme", "notifications@hiremetoday.ca", `"hireme" <notifications@hiremetoday.ca>`},
		{"existing name is replaced", "Ilya from hireme", `hireme <notifications@hiremetoday.ca>`, `"Ilya from hireme" <notifications@hiremetoday.ca>`},
		{"no name leaves the address alone", "", "notifications@hiremetoday.ca", "notifications@hiremetoday.ca"},
		{"unparseable value passes through", "hireme", "not-an-address", "not-an-address"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := emailnotify.From(tc.display, tc.address); got != tc.want {
				t.Errorf("From(%q, %q) = %q, want %q", tc.display, tc.address, got, tc.want)
			}
		})
	}
}
