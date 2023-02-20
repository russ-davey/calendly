package calendly

import "fmt"

type EventTypes struct{}

// GetEventType Returns information about a specified Event Type.
func (e *EventTypes) GetEventType(client *Client, uuid string) (GetEventTypeResponse, error) {
	response := GetEventTypeResponse{}

	url := fmt.Sprintf("%s/event_types/%s", client.baseURL, uuid)
	err := Get(client, url, &response)

	return response, err
}
