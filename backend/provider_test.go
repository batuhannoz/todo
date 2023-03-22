package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
)

type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	ConsumerName    string
	ConsumerVersion string // a git sha, semantic version number
	ConsumerTag     string // dev, staging, prod
	ProviderVersion string
}

func (s *Settings) create() {
	s.Host = "127.0.0.1"
	s.ProviderName = "provider-todo"
	s.ConsumerName = "consumer-todo"
	s.BrokerBaseURL = "http://localhost"
	s.ProviderVersion = "2.0.0"
	s.ConsumerVersion = "2.0.0"
}

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()

	settings := Settings{}
	settings.create()

	pact := dsl.Pact{
		Host:                     settings.Host,
		Provider:                 settings.ProviderName,
		Consumer:                 settings.ConsumerName,
		DisableToolValidityCheck: true,
	}

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://%s:%d", settings.Host, port),
		ProviderVersion: settings.ProviderVersion,
		PactURLs:        []string{"./pacts"},
		StateHandlers: map[string]types.StateHandler{
			"sends a todo and expects 200 in return": func() error {
				return nil
			},
		},
		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
	}

	verifyResponses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}

	pp.Println(len(verifyResponses), "pact tests run")
}
