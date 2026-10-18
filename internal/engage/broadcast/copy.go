package broadcast

// campaigns is the registry. A campaign is added here when it is written, and
// removed once it has been sent: the ledger in broadcast_emails is what records
// that it went out, so keeping the copy here after the fact only offers a letter
// about a date that has passed to whoever reads the worker's usage line.
var campaigns = map[string]Campaign{}
