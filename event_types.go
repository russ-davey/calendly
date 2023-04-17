package calendly

import "fmt"

// GetEventType Returns information about a specified Event Type.
func (c *Client) GetEventType(uuid string) (GetEventTypeResponse, error) {
	response := GetEventTypeResponse{}

	url := fmt.Sprintf("%s/event_types/%s", c.BaseURL, uuid)
	err := c.Get(url, &response)

	return response, err
}
