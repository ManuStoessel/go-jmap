package client

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ManueStoessel/go-jmap/models"
)

// JMAPClient represents the client used to communicate to a JMAP mail API
type JMAPClient struct {
	Hostname string
	Username string
	Password string
}

func (c *JMAPClient) getAuthURL() string {
	return c.Hostname + "/.well-known/jmap"
}

func (c *JMAPClient) getAuthToken() string {
	token := []byte(c.Username + ":" + c.Password + ":")
	return base64.StdEncoding.EncodeToString(token)
}

// GetSession fills a Session struct with the response from a qualified JMAP
// APIs .well-known/jmap endpoint
func (c *JMAPClient) GetSession(s *models.Session) error {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", c.getAuthURL(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+c.getAuthToken())
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(s)
}

// NewJMAPClient returns a pointer to a newly created JMAPClient
func NewJMAPClient(host string, user string, pw string) *JMAPClient {
	return &JMAPClient{
		Hostname: host,
		Username: user,
		Password: pw,
	}
}
