package onboarding_test

import (
	"context"
	"strings"
	"testing"

	"github.com/dunsin2018/HireMe/internal/engage/emailprefs"
	"github.com/dunsin2018/HireMe/internal/engage/onboarding"
)

// The message list is the only thing most people read before deciding to open, and
// a bare address renders there as "notifications".
func TestSenderCarriesAHumanName(t *testing.T) {
	sender := &fakeSender{}
	m := onboarding.NewMailer(sender, "notifications@hireme.example.com", "ilya@example.test", "https://hireme.example.com", testLinks())
	if err := m.Send(context.Background(), onboarding.StepWelcome, 1, "someone@example.com"); err != nil {
		t.Fatalf("Send: %v", err)
	}
	if !strings.Contains(sender.sent[0].from, "Ilya") {
		t.Errorf("From = %q, want a person's name in it", sender.sent[0].from)
	}
	if !strings.Contains(sender.sent[0].from, "notifications@hireme.example.com") {
		t.Errorf("From = %q, want the sending address preserved", sender.sent[0].from)
	}
}

// testLinks signs the unsubscribe URLs these mails carry. The secret only has to
// clear emailprefs' length floor - nothing here verifies a token.
func testLinks() *emailprefs.Links {
	return emailprefs.NewLinks("mail-test-secret-padded-to-32-byte", "https://hireme.example.com")
}
