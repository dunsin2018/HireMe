package prowelcome

import (
	"html/template"

	"github.com/dunsin2018/HireMe/internal/application/mailtpl"
)

// body parses the letter's markup with the shell's sign-off appended, the same shape
// internal/engage/onboarding's copy.go uses — this letter is one of that sequence's
// family (personal, first-person, Reply-To a human inbox), just triggered by a payment
// rather than by a day offset from signup.
var body = template.Must(mailtpl.Partials().New("pro-welcome").Parse(`
{{template "p" "Hi — I'm Dunsin, I run HireMe. Thank you for subscribing, and congratulations on being one of HireMe's earliest paying members!"}}
{{template "lead" (printf "You're now on %s, running through %s." .TierLabel .UntilLabel)}}
{{template "p" "Payment for the AI features is brand new, so some of it may still be rough around the edges. Feel free to ask about anything you run into — I read every message myself."}}
{{template "p" "Just reply to this email — I read every one."}}
` + "\n" + `{{template "signature" .}}`))

// text builds the plain-text alternative, spelling links out the way onboarding's own
// text bodies do.
func text(tierLabel, untilLabel string) string {
	return "Hi — I'm Dunsin, I run HireMe. Thank you for subscribing, and congratulations on\n" +
		"being one of HireMe's earliest paying members!\n\n" +
		"You're now on " + tierLabel + ", running through " + untilLabel + ".\n\n" +
		"Payment for the AI features is brand new, so some of it may still be rough around\n" +
		"the edges. Feel free to ask about anything you run into — I read every message myself.\n\n" +
		"Just reply to this email — I read every one.\n\n" +
		"— Dunsin Fakorede\n"
}
