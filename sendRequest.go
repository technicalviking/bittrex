package bittrex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type queryParams = map[string]string

func (c *Client) sendRequest(endpoint string, params queryParams) *baseResponse {
	fullURI := c.getFullURI(endpoint, params)

	hasher := hmac.New(sha512.New, []byte(c.apiSecret))
	hasher.Write([]byte(fullURI))

	sign := hex.EncodeToString(hasher.Sum(nil))

	var request *http.Request
	var reqErr error

	if request, reqErr = http.NewRequest("GET", fullURI, nil); reqErr != nil {
		c.setError("sendRequest - make request", reqErr.Error())
		return nil
	}

	request.Header.Add("apisign", sign)

	var resp *http.Response
	var respErr error

	done := make(chan bool, 1)

	clientTimer := time.NewTimer(c.timeout)

	go func() {
		if resp, respErr = httpClient.Do(request); respErr != nil {
			c.setError("sendRequest - do request", respErr.Error())
			done <- false
		}

		done <- true
	}()

	select {
	case d := <-done:
		if !d {
			return nil
		}
	case <-clientTimer.C:
		c.setError(
			"sendRequest - do request",
			fmt.Sprintf(
				"BittrexAPI request timeout at %d seconds",
				c.timeout,
			),
		)
		return nil
	}

	defer resp.Body.Close()

	var rawBody []byte
	var readErr error

	if rawBody, readErr = ioutil.ReadAll(resp.Body); readErr != nil {
		c.setError("sendRequest - read response", respErr.Error())
		return nil
	}

	var response baseResponse

	if rawBody == nil || len(rawBody) == 0 {
		response = baseResponse{
			Success: false,
			Message: fmt.Sprintf("Response from API endpoint %s was nil or empty", endpoint),
			Result:  rawBody,
		}
	} else if err := json.Unmarshal(rawBody, &response); err != nil {
		fmt.Printf("here's the response: %v\n", string(rawBody[:len(rawBody)]))
		c.setError("parseResponse", err.Error())
		return nil
	}

	return &response
}

func (c *Client) getFullURI(endpoint string, params queryParams) string {

	version := apiVersion
	if params["useApi2"] != "" {
		version = undocumentedAPIVersion
		delete(params, "useApi2")
	}

	u, _ := url.Parse(baseURI)

	u.Path = strings.Join([]string{u.Path, version, endpoint}, "/")

	query := u.Query()

	query.Set("nonce", fmt.Sprintf("%d", time.Now().Unix()))
	query.Set("apikey", c.apiKey)

	//prevent 304 responses.
	query.Set("_", fmt.Sprintf("%d", time.Now().Unix()))

	for param, value := range params {
		query.Set(param, value)
	}

	u.RawQuery = query.Encode()

	return u.String()
}
