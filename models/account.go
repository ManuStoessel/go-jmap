package models

// AccountID is just a string representing the unique ID of an account
type AccountID string

// Account represents a JMAP mail server's account, this does not have to be
// the users account itself, this can also be a shared account or public one
type Account struct {
	Name                string                 `json:"name"`
	IsPersonal          bool                   `json:"isPersonal"`
	IsReadOnly          bool                   `json:"isReadOnly"`
	AccountCapabilities map[string]interface{} `json:"accountCapabilities"`
}

// AccountCapabilityMail represents the accountCapability object identified
// by the value of JMAPCapabilityURNMail in requests.go
type AccountCapabilityMail struct {
	MaxMailboxesPerEmail       *uint64  `json:"maxMailboxesPerEmail"`
	MaxMailboxDepth            *uint64  `json:"maxMailboxDepth"`
	MaxSizeMailboxName         uint64   `json:"maxSizeMailboxName"`
	MaxSizeAttachmentsPerEmail uint64   `json:"maxSizeAttachmentsPerEmail"`
	EmailQuerySortOptions      []string `json:"emailQuerySortOptions"`
	MayCreateTopLevelMailbox   bool     `json:"mayCreateTopLevelMailbox"`
}

// AccountCapabilitySubmission represents the accountCapability object identified
// by the value of JMAPCapabilityURNSubmission in requests.go
type AccountCapabilitySubmission struct {
	MaxDelayedSend       uint64              `json:"maxDelayedSend"`
	SubmissionExtensions map[string][]string `json:"submissionExtensions"`
}

// HasMailCapability asserts if an Account object has listed the URN for the
// JMAP mail capability and returns it as interface{}
func (a *Account) HasMailCapability() (interface{}, bool) {
	val, ok := a.AccountCapabilities[JMAPCapabilityURNMail]
	return val, ok
}

// HasSubmissionCapability asserts if an Account object has listed the URN for the
// JMAP submission capability and returns it as interface{}
func (a *Account) HasSubmissionCapability() (interface{}, bool) {
	val, ok := a.AccountCapabilities[JMAPCapabilityURNSubmission]
	return val, ok
}

// HasResponseCapability asserts if an Account object has listed the URN for the
// JMAP vacation responder capability and does not return it as the spec says
// it's always an empty object
func (a *Account) HasResponseCapability() bool {
	_, ok := a.AccountCapabilities[JMAPCapabilityURNResponse]
	return ok
}
