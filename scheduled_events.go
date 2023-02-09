package calendly

import "fmt"

type ScheduledEvents struct{}

// TODO: just a basic implementation, add more query parameters including pagination

// GetScheduledEvents Returns a list of Events.
func (s *ScheduledEvents) GetScheduledEvents(client *Client) (GetScheduledEventsResponse, error) {
	response := GetScheduledEventsResponse{}

	url := fmt.Sprintf("%s/scheduled_events", client.baseURL)
	err := Get(client, url, &response)

	return response, err
}

// GetScheduledEvent Returns information about a specified Event.
func (s *ScheduledEvents) GetScheduledEvent(client *Client, uuid string) (GetEventResponse, error) {
	response := GetEventResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s", client.baseURL, uuid)
	err := Get(client, url, &response)

	return response, err
}

// GetInvitees Returns a list of Invitees for an event.
func (s *ScheduledEvents) GetInvitees(client *Client, uuid string) (GetInviteeResponse, error) {
	response := GetInviteeResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s/invitees", client.baseURL, uuid)
	err := Get(client, url, &response)

	return response, err
}
