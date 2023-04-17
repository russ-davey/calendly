package calendly

import "fmt"

// GetUser Returns information about a specified User.
func (c *Client) GetUser(uuid string) (GetUserResponse, error) {
	response := GetUserResponse{}

	url := fmt.Sprintf("%s/users/%s", c.BaseURL, uuid)
	err := c.Get(url, &response)

	return response, err
}
