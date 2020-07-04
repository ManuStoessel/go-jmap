package models

import (
	"encoding/json"
	"fmt"
)

type JMAPRequest struct {
	Using       []JMAPCapabilityURN `json:"using"`
	MethodCalls []Invocation        `json:"methodCalls"`
	CreatedIDs  map[string]string   `json:"createdIds,omitempty"`
}

type Invocation struct {
	Name         string
	Arguments    map[string]interface{}
	MethodCallID string
}

type JMAPCapabilityURN string

const (
	JMAPCapabilityURNCore string = "urn:ietf:params:jmap:core"
	JMAPCapabilityURNMail string = "urn:ietf:params:jmap:mail"
)

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
