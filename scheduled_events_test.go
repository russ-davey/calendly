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
	Convey("Given a Calendly client", t, func() {
		os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
		client := NewClient("test")

		Convey("When the GetAll function is called", func() {
			results, err := client.ScheduledEvents.GetAll()

			Convey("Then a list of the scheduled events are returned", func() {
				So(err, ShouldBeNil)
				So(results.Collection[0].Name, ShouldEqual, "15 Minute Meeting")
			})
		})
	})
}

func TestGetScheduledEvent(t *testing.T) {
	Convey("Given a client and a UUID", t, func() {
		os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
		client := NewClient("test")

		Convey("When the Get function is called", func() {
			results, err := client.ScheduledEvents.Get("b7f4a0b7-d377-47f1-810d-645ff87a2efb")

			Convey("Then the event details are returned", func() {
				So(err, ShouldBeNil)
				So(results.Resource.EventType, ShouldEqual, "https://api.calendly.com/event_types/GBGBDCAADAEDCRZ2")
			})
		})

		Convey("When an error response is returned back from the Calendly API\n", func() {
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

			_, err := client.ScheduledEvents.Get("b7f4a0b7-d377-47f1-810d-645ff87a2efb")
			Convey("Then the error is handled and returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestGetInvitees(t *testing.T) {
	Convey("Given a client and a UUID", t, func() {
		os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
		client := NewClient("test")
		Convey("When the GetInvitees function is called", func() {
			results, err := client.ScheduledEvents.GetInvitees("8ead31de-0033-457a-8646-124e61742999")

			Convey("Then the invitees details are returned", func() {
				So(err, ShouldBeNil)
				So(results.Collection[0].Name, ShouldEqual, "John Doe")
			})
		})
	})
}
