package calendly

import "time"

type GetEventResponse struct {
	Resource struct {
		CalendarEvent struct {
			ExternalId string `json:"external_id"`
			Kind       string `json:"kind"`
		} `json:"calendar_event"`
		Cancellation struct {
			CancelerType string `json:"canceler_type,omitempty"`
			CanceledBy   string `json:"canceled_by,omitempty"`
			Reason       string `json:"reason,omitempty"`
		} `json:"cancellation,omitempty"`
		CreatedAt        time.Time     `json:"created_at"`
		EndTime          time.Time     `json:"end_time"`
		EventGuests      []interface{} `json:"event_guests"`
		EventMemberships []struct {
			User string `json:"user"`
		} `json:"event_memberships"`
		EventType       string `json:"event_type"`
		InviteesCounter struct {
			Active int `json:"active"`
			Limit  int `json:"limit"`
			Total  int `json:"total"`
		} `json:"invitees_counter"`
		Location struct {
			Data struct {
				ID       int64 `json:"id"`
				Settings struct {
					GlobalDialInNumbers []struct {
						CountryName string `json:"country_name"`
						Number      string `json:"number"`
						Type        string `json:"type"`
						Country     string `json:"country"`
					} `json:"global_dial_in_numbers"`
				} `json:"settings"`
				Extra struct {
					IntlNumbersUrl string `json:"intl_numbers_url"`
				} `json:"extra"`
				Password string `json:"password"`
			} `json:"data"`
			JoinUrl string `json:"join_url"`
			Status  string `json:"status"`
			Type    string `json:"type"`
		} `json:"location"`
		Name      string    `json:"name"`
		StartTime time.Time `json:"start_time"`
		Status    string    `json:"status"`
		UpdatedAt time.Time `json:"updated_at"`
		URI       string    `json:"uri"`
	} `json:"resource"`
}

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
		DeletedAt        interface{} `json:"deleted_at"`
		DescriptionHtml  string      `json:"description_html"`
		DescriptionPlain string      `json:"description_plain"`
		Duration         int         `json:"duration"`
		InternalNote     string      `json:"internal_note"`
		Kind             string      `json:"kind"`
		KindDescription  string      `json:"kind_description"`
		Name             string      `json:"name"`
		PoolingType      string      `json:"pooling_type"`
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
