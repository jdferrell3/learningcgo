package httplib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const PinnedCertHash = "62f87a97e95782aac126b0adc70afdaffbdb111137effe7fa7a01a451b226e84"

// const TestURL = "https://httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l"
// const TestURLResp = "HTTPBIN is awesome"

func TestPinnedCert(t *testing.T) {
	c := PinnedClient(PinnedCertHash, false)

	// Make the API call
	request := Request{
		Method:  Get,
		BaseURL: "https://httpbin.org/get",
	}

	_, err := c.Send(request)
	if nil != err {
		t.Fatal(err)
	}
}

func TestBasicAuth(t *testing.T) {
	c := &Client{HTTPClient: http.DefaultClient}

	// /hidden-basic-auth/{user}/{passwd}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Basic " + BasicAuth("user", "pass")
	request := Request{
		Method:  Get,
		BaseURL: "https://httpbin.org//hidden-basic-auth/user/pass",
		Headers: headers,
	}

	res, err := c.Send(request)
	if nil != err {
		t.Fatal(err)
	}

	assert.Equal(t, 200, res.StatusCode)
}
