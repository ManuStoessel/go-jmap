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
