package calendly

import "fmt"

type ScheduledEventsService struct {
	client *Client
}

// TODO: just a basic implementation, add more query parameters including pagination

// GetAll Returns a list of Events.
func (s *ScheduledEventsService) GetAll() (GetScheduledEventsResponse, error) {
	response := GetScheduledEventsResponse{}

	url := fmt.Sprintf("%s/scheduled_events", s.client.BaseURL)
	err := s.client.Get(url, &response)

	return response, err
}

// Get Returns information about a specified Event.
func (s *ScheduledEventsService) Get(uuid string) (GetEventResponse, error) {
	response := GetEventResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s", s.client.BaseURL, uuid)
	err := s.client.Get(url, &response)

	return response, err
}

// GetInvitees Returns a list of Invitees for an event.
func (s *ScheduledEventsService) GetInvitees(uuid string) (GetInviteeResponse, error) {
	response := GetInviteeResponse{}

	url := fmt.Sprintf("%s/scheduled_events/%s/invitees", s.client.BaseURL, uuid)
	err := s.client.Get(url, &response)

	return response, err
}
