package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestCounterApi(t *testing.T) {

	pact := &dsl.Pact{
		Provider: "counter_api",
		PactDir:  "/Users/ahmet.zumberoglu/Documents/go-work/pacts",
	}

	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:3000",
		BrokerURL:       "http://localhost:9292",
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/myconsumer-myprovider.json", pact.PactDir))},
		StateHandlers: types.StateHandlers{
			// Setup any state required by the test
			// in this case, we ensure there is a "user" in the system
			"Counter API exists": func() error {
				data.counter = 0
				return nil
			},
		},
	})
}
