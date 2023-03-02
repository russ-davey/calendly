package calendly

import "time"

type GetUserResponse struct {
	Resource struct {
		Uri                 string    `json:"uri"`
		Name                string    `json:"name"`
		Slug                string    `json:"slug"`
		Email               string    `json:"email"`
		SchedulingUrl       string    `json:"scheduling_url"`
		Timezone            string    `json:"timezone"`
		AvatarUrl           string    `json:"avatar_url"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
		CurrentOrganization string    `json:"current_organization"`
	} `json:"resource"`
}
