package calendly

import "fmt"

type EventTypesService struct {
	client *Client
}

// Get Returns information about a specified Event Type.
func (e *EventTypesService) Get(uuid string) (GetEventTypeResponse, error) {
	response := GetEventTypeResponse{}

	url := fmt.Sprintf("%s/event_types/%s", e.client.BaseURL, uuid)
	err := e.client.Get(url, &response)

	return response, err
}
