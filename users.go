package calendly

import "fmt"

type UsersService struct {
	client *Client
}

// Get Returns information about a specified User.
func (u *UsersService) Get(uuid string) (GetUserResponse, error) {
	response := GetUserResponse{}

	url := fmt.Sprintf("%s/users/%s", u.client.BaseURL, uuid)
	err := u.client.Get(url, &response)

	return response, err
}
