package main

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "provider-test",
		Consumer: "consumer-test",
	}

}

func TestCounterApi(t *testing.T) {
	pact := createPact()
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://localhost:8080/",
		PactURLs:                   []string{"https://zumber.pactflow.io/pacts/provider/provider-test/consumer/consumer-test/version/33bc00-master%2B33bc00.SNAPSHOT.trhqmac184"},
		PublishVerificationResults: true,
		ProviderVersion:            "2.0.0",
		BrokerToken:                "AhT3yO3pFlB6lZpWskVDcA",
	})
	if err != nil {
		t.Fatal(err)
	}
}
