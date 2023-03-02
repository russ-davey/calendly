package calendly

import "fmt"

type Users struct{}

// GetUser Returns information about a specified User.
func (u *Users) GetUser(client *Client, uuid string) (GetUserResponse, error) {
	response := GetUserResponse{}

	url := fmt.Sprintf("%s/users/%s", client.baseURL, uuid)
	err := Get(client, url, &response)

	return response, err
}
