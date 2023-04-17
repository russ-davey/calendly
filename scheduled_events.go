package calendly

import "fmt"

type ScheduledEvents struct{}

// TODO: just a basic implementation, add more query parameters including pagination

// GetScheduledEvents Returns a list of Events.
func (c *Client) GetScheduledEvents() (GetScheduledEventsResponse, error) {
	response := GetScheduledEventsResponse{}

	url := fmt.Sprintf("%s/scheduled_events", c.BaseURL)
	err := c.Get(url, &response)

	return response, err
}

// GetScheduledEvent Returns information about a specified Event.
func (c *Client) GetScheduledEvent(uuid string) (GetEventResponse, error) {
	response := GetEventResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s", c.BaseURL, uuid)
	err := c.Get(url, &response)

	return response, err
}

// GetInvitees Returns a list of Invitees for an event.
func (c *Client) GetInvitees(uuid string) (GetInviteeResponse, error) {
	response := GetInviteeResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s/invitees", c.BaseURL, uuid)
	err := c.Get(url, &response)

	return response, err
}
