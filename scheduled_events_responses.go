package calendly

import "time"

type GetScheduledEventsResponse struct {
	Collection []struct {
		Uri       string    `json:"uri"`
		Name      string    `json:"name"`
		Status    string    `json:"status"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
		EventType string    `json:"event_type"`
		Location  struct {
			Type     string `json:"type"`
			Location string `json:"location"`
		} `json:"location"`
		InviteesCounter struct {
			Total  int `json:"total"`
			Active int `json:"active"`
			Limit  int `json:"limit"`
		} `json:"invitees_counter"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		EventMemberships []struct {
			User string `json:"user"`
		} `json:"event_memberships"`
		EventGuests []struct {
			Email     string    `json:"email"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"event_guests"`
		CalendarEvent struct {
			Kind       string `json:"kind"`
			ExternalId string `json:"external_id"`
		} `json:"calendar_event"`
	} `json:"collection"`
	Pagination struct {
		Count             int    `json:"count"`
		NextPage          string `json:"next_page"`
		PreviousPage      string `json:"previous_page"`
		NextPageToken     string `json:"next_page_token"`
		PreviousPageToken string `json:"previous_page_token"`
	} `json:"pagination"`
}

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

type GetInviteeResponse struct {
	Collection []struct {
		CancelUrl           string    `json:"cancel_url"`
		CreatedAt           time.Time `json:"created_at"`
		Email               string    `json:"email"`
		Event               string    `json:"event"`
		Name                string    `json:"name"`
		FirstName           string    `json:"first_name"`
		LastName            string    `json:"last_name"`
		NewInvitee          string    `json:"new_invitee"`
		OldInvitee          string    `json:"old_invitee"`
		QuestionsAndAnswers []struct {
			Answer   string `json:"answer"`
			Position int    `json:"position"`
			Question string `json:"question"`
		} `json:"questions_and_answers"`
		RescheduleUrl      string `json:"reschedule_url"`
		Rescheduled        bool   `json:"rescheduled"`
		Status             string `json:"status"`
		TextReminderNumber string `json:"text_reminder_number"`
		Timezone           string `json:"timezone"`
		Tracking           struct {
			UtmCampaign    interface{} `json:"utm_campaign"`
			UtmSource      interface{} `json:"utm_source"`
			UtmMedium      interface{} `json:"utm_medium"`
			UtmContent     interface{} `json:"utm_content"`
			UtmTerm        interface{} `json:"utm_term"`
			SalesforceUuid interface{} `json:"salesforce_uuid"`
		} `json:"tracking"`
		UpdatedAt             time.Time `json:"updated_at"`
		Uri                   string    `json:"uri"`
		Canceled              bool      `json:"canceled"`
		RoutingFormSubmission string    `json:"routing_form_submission"`
		Payment               struct {
			ExternalId string  `json:"external_id"`
			Provider   string  `json:"provider"`
			Amount     float64 `json:"amount"`
			Currency   string  `json:"currency"`
			Terms      string  `json:"terms"`
			Successful bool    `json:"successful"`
		} `json:"payment"`
		NoShow         interface{} `json:"no_show"`
		Reconfirmation struct {
			CreatedAt   time.Time `json:"created_at"`
			ConfirmedAt time.Time `json:"confirmed_at"`
		} `json:"reconfirmation"`
	} `json:"collection"`
	Pagination struct {
		Count             int    `json:"count"`
		NextPage          string `json:"next_page"`
		PreviousPage      string `json:"previous_page"`
		NextPageToken     string `json:"next_page_token"`
		PreviousPageToken string `json:"previous_page_token"`
	} `json:"pagination"`
}
