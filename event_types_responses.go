package calendly

import "time"

type GetEventTypeResponse struct {
	Resource struct {
		Active          bool      `json:"active"`
		AdminManaged    bool      `json:"admin_managed"`
		BookingMethod   string    `json:"booking_method"`
		Color           string    `json:"color"`
		CreatedAt       time.Time `json:"created_at"`
		CustomQuestions []struct {
			AnswerChoices []string `json:"answer_choices"`
			Enabled       bool     `json:"enabled"`
			IncludeOther  bool     `json:"include_other"`
			Name          string   `json:"name"`
			Position      int      `json:"position"`
			Required      bool     `json:"required"`
			Type          string   `json:"type"`
		} `json:"custom_questions"`
		DeletedAt        string `json:"deleted_at"`
		DescriptionHtml  string `json:"description_html"`
		DescriptionPlain string `json:"description_plain"`
		Duration         int    `json:"duration"`
		InternalNote     string `json:"internal_note"`
		Kind             string `json:"kind"`
		KindDescription  string `json:"kind_description"`
		Name             string `json:"name"`
		PoolingType      string `json:"pooling_type"`
		Profile          struct {
			Name  string `json:"name"`
			Owner string `json:"owner"`
			Type  string `json:"type"`
		} `json:"profile"`
		SchedulingUrl string    `json:"scheduling_url"`
		Secret        bool      `json:"secret"`
		Slug          string    `json:"slug"`
		Type          string    `json:"type"`
		UpdatedAt     time.Time `json:"updated_at"`
		URI           string    `json:"uri"`
	} `json:"resource"`
}
