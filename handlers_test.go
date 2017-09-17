package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/schema"
)

var encoder = schema.NewEncoder()
var testPayload = SlackPayload{
	Token:       "token",
	TeamID:      "team-id",
	TeamDomain:  "team-domain",
	ChannelID:   "channel-id",
	ChannelName: "channel-name",
	UserID:      "user-id",
	UserName:    "user-name",
	Command:     "/table",
	Text:        "",
	ResponseURL: "https://www.slack.com",
}

func checkStatus(t *testing.T, code int, expected int) {
	if code != expected {
		t.Errorf("handler returned wrong status code, got: %v, want: %v", code, expected)
	}
}

func TestHealthHandler(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	checkStatus(t, rr.Code, http.StatusOK)
}

func TestTableHandler(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	tests := []struct {
		Desc          string
		Attachments   bool
		Expected      string
		SlackDataText string
		Status        int
	}{{
		Desc:          "Table flip",
		Attachments:   false,
		Expected:      "{\"response_type\":\"in_channel\",\"text\":\"(╯°□°)╯︵ ┻━┻\",\"attachments\":[]}",
		SlackDataText: "flip",
		Status:        http.StatusOK,
	}, {
		Desc:          "Table flip with attachments",
		Attachments:   true,
		Expected:      "{\"response_type\":\"in_channel\",\"text\":\"(╯°□°)╯︵ ┻━┻\",\"attachments\":[{\"text\":\"Table flipped!\"}]}",
		SlackDataText: "flip",
		Status:        http.StatusOK,
	}, {
		Desc:          "Table catch",
		Attachments:   false,
		Expected:      "{\"response_type\":\"in_channel\",\"text\":\"┬─┬ノ( º _ ºノ)\",\"attachments\":[]}",
		SlackDataText: "catch",
		Status:        http.StatusOK,
	}, {
		Desc:          "Table catch with attachments",
		Attachments:   true,
		Expected:      "{\"response_type\":\"in_channel\",\"text\":\"┬─┬ノ( º _ ºノ)\",\"attachments\":[{\"text\":\"Table caught!\"}]}",
		SlackDataText: "catch",
		Status:        http.StatusOK,
	}, {
		Desc:          "Invalid action",
		Attachments:   false,
		Expected:      "Did you want me to `flip` or `catch` the table?",
		SlackDataText: "notFlipOrCatch",
		Status:        http.StatusOK,
	}}

	for _, test := range tests {
		if test.Attachments {
			os.Setenv("ATTACHMENTS", "yes")
		} else {
			os.Setenv("ATTACHMENTS", "no")
		}
		testPayload.Text = test.SlackDataText

		form := url.Values{}
		if err := encoder.Encode(testPayload, form); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/table", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tableHandler)
		handler.ServeHTTP(rr, req)

		checkStatus(t, rr.Code, test.Status)
		if rr.Body.String() != test.Expected {
			t.Errorf("handler returned unexpected body, got: %v, want: %v", rr.Body.String(), test.Expected)
		}
	}
}
