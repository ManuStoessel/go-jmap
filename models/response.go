package models

// JMAPResponse specifies the response object from a JMAP server
// The object definition can be found here:
// https://jmap.io/spec-core.html#the-response-object
type JMAPResponse struct {
	MethodResponses []Invocation      `json:"methodResponses"`
	CreatedIDs      map[string]string `json:"createdIds"`
	SessionState    string            `json:"sessionState"`
}

// JMAPProblemResponse is the response object of a JMAP server in case of a
// server-level error (coincides with a HTTP status != 2xx )
type JMAPProblemResponse struct {
	Type   JMAPProblemType `json:"type"`
	Status int             `json:"status"`
	Detail string          `json:"detail"`
}

// JMAPProblemType is the type of JMAP server-level error
type JMAPProblemType string

// MethodErrorType is the type of the method-level error
type MethodErrorType string

const (
	// Error types defined here: https://jmap.io/spec-core.html#errors
	// Server-level errors

	// JMAPProblemTypeUnknownCap - The client included a capability in the
	// “using” property of the request that the server does not support
	JMAPProblemTypeUnknownCap string = "urn:ietf:params:jmap:error:unknownCapability"

	// JMAPProblemTypeNotJSON - The content type of the request was not
	// application/json or the request did not parse as I-JSON.
	JMAPProblemTypeNotJSON string = "urn:ietf:params:jmap:error:notJSON"

	// JMAPProblemTypeNotRequest - The request parsed as JSON but did not match
	// the type signature of the Request object.
	JMAPProblemTypeNotRequest string = "urn:ietf:params:jmap:error:notRequest"

	// JMAPProblemTypeLimit - The request was not processed as it would have
	// exceeded one of the request limits defined on the capability object,
	// such as maxSizeRequest, maxCallsInRequest, or maxConcurrentRequests. A
	// “limit” property MUST also be present on the “problem details” object,
	// containing the name of the limit being applied.
	JMAPProblemTypeLimit string = "urn:ietf:params:jmap:error:limit"

	// Method-level errors

	// MethodErrorTypeServerUnavailable - Some internal server resource was
	// temporarily unavailable. Attempting the same operation later (perhaps
	// after a backoff with a random factor) may succeed.
	MethodErrorTypeServerUnavailable string = "serverUnavailable"

	// MethodErrorTypeServerFailure - An unexpected or unknown error occurred
	// during the processing of the call. A description property should provide
	// more details about the error. The method call made no changes to the
	// server’s state. Attempting the same operation again is expected to fail
	// again. Contacting the service administrator is likely necessary to
	// resolve this problem if it is persistent.
	MethodErrorTypeServerFailure string = "serverFail"

	// MethodErrorTypeServerPartialFailure - Some, but not all, expected
	// changes described by the method occurred. The client MUST resynchronise
	// impacted data to determine server state. As use of this error is highly
	// discouraged for server implementations, it is unlikely to encounter.
	MethodErrorTypeServerPartialFailure string = "serverPartialFail"

	// MethodErrorTypeUnknownMethod - The server does not recognise this method
	// name.
	MethodErrorTypeUnknownMethod string = "unknownMethod"

	// MethodErrorTypeInvalidArguments - One of the arguments is of the wrong
	// type or is otherwise invalid, or a required argument is missing. A
	// description property MAY be present to help debug with an explanation of
	// what the problem was. This is a non-localised string, and it is not
	// intended to be shown directly to end users.
	MethodErrorTypeInvalidArguments string = "invalidArguments"

	// MethodErrorTypeInvalidResultReference - The method used a result
	// reference for one of its arguments (see Section 3.7), but this failed to
	// resolve.
	MethodErrorTypeInvalidResultReference string = "invalidResultReference"

	// MethodErrorTypeForbidden - The method and arguments are valid, but
	// executing the method would violate an Access Control List (ACL) or other
	// permissions policy.
	MethodErrorTypeForbidden string = "forbidden"

	// MethodErrorTypeAccountNotFound - The accountId does not correspond to a
	// valid account.
	MethodErrorTypeAccountNotFound string = "accountNotFound"

	// MethodErrorTypeAccountNotSupportedByMethod - The accountId given
	// corresponds to a valid account, but the account does not support this
	// method or data type.
	MethodErrorTypeAccountNotSupportedByMethod string = "accountNotSupportedByMethod"

	// MethodErrorTypeAccountReadOnly - This method modifies state, but the
	// account is read-only (as returned on the corresponding Account object in
	// the JMAP Session resource).
	MethodErrorTypeAccountReadOnly string = "accountReadOnly"
)
