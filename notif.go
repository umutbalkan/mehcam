package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	YoApiUrl = "https://api.justyo.co/yo/"
)

// send uses the JustYo API to send
// a push notification to my phone
// with a link to the serve image using
// its ID.
func send(id string) error {
	if len(id) == 0 {
		return nil
	}

	values := url.Values{}
	values.Add("api_token", env.yo_api_key)
	values.Add("username", env.yo)
	values.Add("link", "http://"+env.addr+"/?f="+id) // TODO(remy):

	resp, err := http.PostForm(YoApiUrl, values)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error while sending the notification:", resp.Status)
	}

	return nil
}