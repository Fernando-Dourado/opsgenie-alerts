package main

import (
	"context"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type (
	Config struct {
		Token string
	}

	Alert struct {
		Message     string
		Description string
	}

	Result struct {
		Result    string
		RequestId string
	}

	Plugin struct {
		Config Config
		Alert  Alert
		Result Result
	}
)

func (p *Plugin) Exec() error {

	alertClient, err := alert.NewClient(&client.Config{
		ApiKey: p.Config.Token,
	})
	if err != nil {
		return err
	}

	a := p.Alert
	createResult, err := alertClient.Create(context.TODO(), &alert.CreateAlertRequest{
		Message:     a.Message,
		Description: a.Description,
	})
	if err != nil {
		return err
	}

	p.Result = Result{
		Result:    createResult.Result,
		RequestId: createResult.RequestId,
	}
	return nil
}
