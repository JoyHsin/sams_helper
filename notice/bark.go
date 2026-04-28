package notice

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type BarkSet struct {
	Server  string `yaml:"barkServer"`
	Token   string `yaml:"barkToken"`
	Message string `yaml:"barkMessage"`
	Sound   string `yaml:"barkSound"`
}

func BarkPush(barkSet BarkSet) error {
	return BarkPushMessage(barkSet, "")
}

func BarkPushMessage(barkSet BarkSet, message string) error {
	if message == "" {
		message = barkSet.Message
	}
	urlPath := fmt.Sprintf("%s/%s/%s?sound=%s", barkSet.Server, barkSet.Token, url.PathEscape(message), url.QueryEscape(barkSet.Sound))
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(urlPath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("[%v] %s", resp.StatusCode, body))
	}
}
