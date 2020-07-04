package models

import (
	"encoding/json"
	"fmt"
)

// JMAPRequest represents the JMAP request object as it is defined here:
// https://jmap.io/spec-core.html#the-request-object
type JMAPRequest struct {
	Using       []JMAPCapabilityURN `json:"using"`
	MethodCalls []Invocation        `json:"methodCalls"`
	CreatedIDs  map[string]string   `json:"createdIds,omitempty"`
}

// Invocation represents the JMAP Method Invocation object which is defined
// as a three item tupel represented as a JSON array of three elements. The
// items have the datatypes corresponding to the struct field. More can be
// found here: https://jmap.io/spec-core.html#the-invocation-data-type
type Invocation struct {
	Name         string
	Arguments    map[string]interface{}
	MethodCallID string
}

// JMAPCapabilityURN is a string representing the JMAP capabilities used in a
// JMAP request
type JMAPCapabilityURN string

const (
	// JMAPCapabilityURNCore is the identifier for core capabilities in the
	// JMAP spec https://jmap.io/spec-core.html
	JMAPCapabilityURNCore string = "urn:ietf:params:jmap:core"

	// JMAPCapabilityURNMail is the identifier for mail capabilities in the
	// JMAP spec https://jmap.io/spec-mail.html
	JMAPCapabilityURNMail string = "urn:ietf:params:jmap:mail"
)

// UnmarshalJSON is a custom json unmarshaling function for the
// Invocation type since this is represented as a JSON array in the JMAP spec
func (i *Invocation) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&i.Name, &i.Arguments, &i.MethodCallID}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
	}
	return nil
}
