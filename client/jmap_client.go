package client

import (
	"bytes"
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
	Client   *http.Client
	Session  *models.Session
}

// NewJMAPClient returns a pointer to a newly created JMAPClient
func NewJMAPClient(host string, user string, pw string) *JMAPClient {
	return &JMAPClient{
		Hostname: host,
		Username: user,
		Password: pw,
		Client:   &http.Client{Timeout: 10 * time.Second},
		Session:  &models.Session{},
	}
}

// GetSession fills a Session struct with the response from a qualified JMAP
// APIs .well-known/jmap endpoint
func (c *JMAPClient) GetSession() error {

	req, err := http.NewRequest("GET", c.getAuthURL(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(c.Session)
}

// JMAPRequest sends a request to a JMAP API and hands over the response in a
// generic interface
func (c *JMAPClient) JMAPRequest(data []byte, response interface{}) error {
	req, err := http.NewRequest("POST", c.getAuthURL(), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func (c *JMAPClient) getAuthURL() string {
	return c.Hostname + "/.well-known/jmap"
}
