package calendly

import (
	"encoding/json"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetScheduledEvent(t *testing.T) {
	cy := Calendly{}

	Convey("Given an event has been received from Calendly", t, func() {

		Convey("When the event details are looked up via the Calendly API", func() {
			os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
			client := NewClient("test")

			result, err := cy.GetScheduledEvent(client, "b7f4a0b7-d377-47f1-810d-645ff87a2efb")
			Convey("Then the details are successfully returned", func() {
				So(err, ShouldBeNil)
				So(result.Resource.EventType, ShouldEqual, "https://api.calendly.com/event_types/GBGBDCAADAEDCRZ2")
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
			os.Setenv("CALENDLY_MOCK_SERVER", mockCalendlyAPI.URL)
			client := NewClient("test")

			_, err := cy.GetScheduledEvent(client, "b7f4a0b7-d377-47f1-810d-645ff87a2efb")
			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestGetEventType(t *testing.T) {
	cy := Calendly{}

	os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
	client := NewClient("test")

	Convey("Given an event has been received from Calendly and the UUID has been looked up for the event type", t, func() {
		Convey("When the event type details are looked up via the Calendly API", func() {
			result, err := cy.GetEventType(client, "8ead31de-0033-457a-8646-124e61742999")
			Convey("Then the details are successfully returned", func() {
				So(err, ShouldBeNil)
				So(result.Resource.Name, ShouldEqual, "15 Minute Meeting")
			})
		})
	})
}

func TestMarshellAPIError(t *testing.T) {
	Convey("Given an error has been returned by the Calendly API", t, func() {
		Convey("When the error is unmarshalled", func() {
			jsonStr := "{\n  \"title\": \"Internal Server Error\",\n  \"message\": \"The server encountered an unexpected condition that prevented it from fulfilling the request.\",\n  \"details\": [\n    {\n      \"parameter\": \"string\",\n      \"message\": \"string\"\n    }\n  ]\n}"
			results := UnmarshallAPIError(errors.New(jsonStr))
			Convey("Then the error message is successfully unmarshalled", func() {
				So(results.Title, ShouldEqual, "Internal Server Error")
				So(results.Message, ShouldEqual, "The server encountered an unexpected condition that prevented it from fulfilling the request.")
			})
		})
		Convey("When a different error is unmarshalled", func() {
			results := UnmarshallAPIError(errors.New("unknown error"))
			Convey("Then the error message is successfully unmarshalled but the values are empty", func() {
				So(results.Title, ShouldEqual, "")
			})
		})
	})
}
