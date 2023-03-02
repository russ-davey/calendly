package calendly

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Client struct {
	baseURL string
	token   string
}

type Calendly struct {
	Client          *Client
	EventTypes      EventTypes
	ScheduledEvents ScheduledEvents
	Users           Users
}

type Details []struct {
	Parameter string `json:"parameter,omitempty"`
	Message   string `json:"message,omitempty"`
}

// ErrorBody all Calendly non-200 status bodies are returned in this struct format
type ErrorBody struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Details `json:"details,omitempty"`
}

func NewClient(token string) *Client {
	var baseURL string
	if mockServer := os.Getenv("CALENDLY_MOCK_SERVER"); mockServer != "" {
		baseURL = mockServer
	} else {
		baseURL = "https://api.calendly.com"
	}

	return &Client{
		baseURL: baseURL,
		token:   token,
	}
}

func UnmarshallAPIError(err error) ErrorBody {
	eb := ErrorBody{}
	json.Unmarshal([]byte(err.Error()), &eb)
	return eb
}

// Get takes a client, API endpoint and a pointer to a struct then writes the API response data to that struct
func Get(client *Client, url string, response interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		log.Printf("Calendly API error: %s", body)
		return errors.New(string(body))
	}
	return json.NewDecoder(res.Body).Decode(&response)
}
