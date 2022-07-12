package pact

import (
	"bootcamp/server"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"strconv"
	"testing"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	svr := server.NewServer()
	go svr.StartServer(strconv.Itoa(port))

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "Backend",
		Consumer:                 "Frontend",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		BrokerURL:                  "https://developer.pactflow.io",
		BrokerToken:                "jXNfHe5Nu-G5s1zd3Fs5gg",
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		PactURLs: []string{
			"https://developer.pactflow.io/pacts/provider/Backend/consumer/Frontend/latest",
		},
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}
