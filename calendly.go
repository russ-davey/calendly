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
	Client *Client
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

// GetScheduledEvent Returns information about a specified Event.
func (cy *Calendly) GetScheduledEvent(client *Client, uuid string) (GetEventResponse, error) {
	response := GetEventResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s", client.baseURL, uuid)
	err := Get(client, url, &response)

	return response, err
}

// GetEventType Returns information about a specified Event Type.
func (cy *Calendly) GetEventType(client *Client, uuid string) (GetEventTypeResponse, error) {
	response := GetEventTypeResponse{}

	url := fmt.Sprintf("%s/event_types/%s", client.baseURL, uuid)
	err := Get(client, url, &response)

	return response, err
}
