package calendly

import "fmt"

type ScheduledEvents struct{}

// GetScheduledEvent Returns information about a specified Event.
func (s *ScheduledEvents) GetScheduledEvent(client *Client, uuid string) (GetEventResponse, error) {
	response := GetEventResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s", client.baseURL, uuid)
	err := Get(client, url, &response)

	return response, err
}

func (s *ScheduledEvents) GetInvitees(client *Client, uuid string) (GetInviteeResponse, error) {
	response := GetInviteeResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s/invitees", client.baseURL, uuid)
	err := Get(client, url, &response)

	return response, err
}
