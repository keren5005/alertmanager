// Copyright 2022 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mattermost

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/notify"
	commoncfg "github.com/prometheus/common/config"
	"net/http"

	"github.com/prometheus/alertmanager/template"
	"github.com/prometheus/alertmanager/types"
)

// Telegram supports 4096 chars max - from https://limits.tginfo.me/en.
const maxMessageLenRunes = 4096

// Notifier implements a Notifier for telegram notifications.
type Notifier struct {
	conf   *config.MattermostConfig
	tmpl   *template.Template
	logger log.Logger
}

// New returns a new Telegram notification handler.
func New(conf *config.MattermostConfig, t *template.Template, l log.Logger, httpOpts ...commoncfg.HTTPClientOption) (*Notifier, error) {
	return &Notifier{
		conf:   conf,
		tmpl:   t,
		logger: l,
	}, nil
}

type requestTokenStruct struct {
	LoginId  string `json:"login_id"`
	Password string `json:"password"`
}

func obtainToken(server_url string, user_name string, password string) (string, error) {
	request := requestTokenStruct{
		LoginId:  user_name,
		Password: password,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/api/v4/users/login", server_url)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code %d", res.StatusCode)
	}
	headers := res.Header["Token"]
	if len(headers) != 1 {
		return "", fmt.Errorf("expected 1 token header")
	}
	return headers[0], nil
}

type postMessageStruct struct {
	ChannelId string `json:"channel_id"`
	Message   string `json:"message"`
}

func sendMessage(server_url string, token string, channel_id string, text string) error {
	request := postMessageStruct{
		ChannelId: channel_id,
		Message:   text,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/api/v4/posts", server_url)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return err
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("status code %d", res.StatusCode)
	}
	return nil
}

func (n *Notifier) Notify(ctx context.Context, alert ...*types.Alert) (bool, error) {
	// TODO
	// 1) Find out how to convert alert to text
	// 2) Use n.logger instead of fmt.Printf
	token, err := obtainToken(n.conf.ServerUrl, n.conf.UserName, n.conf.Password)
	if err != nil {
		fmt.Printf("Error obtaining token %v\n", err)
	} else {
		fmt.Printf("Token: %s\n", token)

		var (
			data     = notify.GetTemplateData(ctx, n.tmpl, alert, n.logger)
			tmplText = notify.TmplText(n.tmpl, data, &err)
		)
		err = sendMessage(n.conf.ServerUrl, token, n.conf.ChannelId, tmplText("KEREN"))
		if err != nil {
			fmt.Printf("Error while sending message: %v\n", err)
		}
	}

	fmt.Printf("MATTERMOST NOTIFY: %s\n", n.conf.Text)
	return false, nil
}
