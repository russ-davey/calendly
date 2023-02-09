package calendly

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestScheduledEvents(t *testing.T) {
	cy := Calendly{}

	os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
	client := NewClient("test")

	Convey("Given the Calendly API scheduled_events endpoint needs to be accessed", t, func() {
		Convey("When getting the scheduled events details", func() {
			results, err := cy.ScheduledEvents.GetScheduledEvents(client)
			Convey("Then the details are successfully returned", func() {
				So(err, ShouldBeNil)
				So(results.Collection[0].Name, ShouldEqual, "15 Minute Meeting")
			})
		})
	})
}

func TestGetScheduledEvent(t *testing.T) {
	cy := Calendly{}

	Convey("Given the Calendly API scheduled_events endpoint needs to be accessed", t, func() {
		Convey("When getting the scheduled_events details", func() {
			os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
			client := NewClient("test")

			results, err := cy.GetScheduledEvent(client, "b7f4a0b7-d377-47f1-810d-645ff87a2efb")
			Convey("Then the details are successfully returned", func() {
				So(err, ShouldBeNil)
				So(results.Resource.EventType, ShouldEqual, "https://api.calendly.com/event_types/GBGBDCAADAEDCRZ2")
			})
		})

		Convey("When the event details are looked up via the Calendly API but results in an error\n", func() {
			mockCalendlyAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorBody{
					Title:   "Internal Server Error",
					Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
					Details: Details{
						{
							Parameter: "string",
							Message:   "string",
						},
					},
				})
			}))
			//req.Header.Add("Prefer", "code=200")
			os.Setenv("CALENDLY_MOCK_SERVER", mockCalendlyAPI.URL)
			client := NewClient("test")

			_, err := cy.GetScheduledEvent(client, "b7f4a0b7-d377-47f1-810d-645ff87a2efb")
			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestGetInvitees(t *testing.T) {
	cy := Calendly{}

	os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
	client := NewClient("test")

	Convey("Given the Calendly API scheduled_events/<UUID>/invitees endpoint needs to be accessed", t, func() {
		Convey("When getting the invitees details", func() {
			results, err := cy.ScheduledEvents.GetInvitees(client, "8ead31de-0033-457a-8646-124e61742999")
			Convey("Then the details are successfully returned", func() {
				So(err, ShouldBeNil)
				So(results.Collection[0].Name, ShouldEqual, "John Doe")
			})
		})
	})
}
