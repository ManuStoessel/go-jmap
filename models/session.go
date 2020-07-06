package models

// Session represents a JMAP session object as described here:
// https://jmap.io/spec-core.html#the-jmap-session-resource
type Session struct {
	Capabilities    map[JMAPCapabilityURN]interface{} `json:"capabilities"`
	Accounts        map[AccountID]Account             `json:"accounts"`
	PrimaryAccounts map[string]AccountID              `json:"primaryAccounts"`
	Username        string                            `json:"username"`
	APIUrl          string                            `json:"apiUrl"`
	DownloadURL     string                            `json:"downloadUrl"`
	UploadURL       string                            `json:"uploadUrl"`
	EventSourceURL  string                            `json:"eventSourceUrl"`
	State           string                            `json:"state"`
}

// SessionMailCapability is an object describing the JMAP servers capabilities
// and limits that have a direct influence on client behaviour for the mail URN
type SessionMailCapability struct {
	MaxSizeUpload         uint64   `json:"maxSizeUpload"`
	MaxConcurrentUpload   uint64   `json:"maxConcurrentUpload"`
	MaxSizeRequest        uint64   `json:"maxSizeRequest"`
	MaxConcurrentRequests uint64   `json:"maxConcurrentRequests"`
	MaxCallsInRequest     uint64   `json:"maxCallsInRequest"`
	MaxObjectsInGet       uint64   `json:"maxObjectsInGet"`
	MaxObjectsInSet       uint64   `json:"maxObjectsInSet"`
	CollationAlgorithms   []string `json:"collationAlgorithms"`
}
