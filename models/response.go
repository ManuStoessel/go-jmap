package models

type JMAPResponse struct {
	MethodResponses []Invocation      `json:"methodResponses"`
	CreatedIDs      map[string]string `json:"createdIds"`
	SessionState    string            `json:"sessionState"`
}

type JMAPProblemResponse struct {
	Type   JMAPProblemType `json:"type"`
	Status int             `json:"status"`
	Detail string          `json:"detail"`
}

type JMAPProblemType string

const (
	JMAPProblemTypeUnknownCap string = "urn:ietf:params:jmap:error:unknownCapability"
	JMAPProblemTypeNotJSON    string = "urn:ietf:params:jmap:error:notJSON"
	JMAPProblemTypeNotRequest string = "urn:ietf:params:jmap:error:notRequest"
	JMAPProblemTypeLimit      string = "urn:ietf:params:jmap:error:limit"
)
