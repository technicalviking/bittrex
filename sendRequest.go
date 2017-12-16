package bittrex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (c *Client) sendRequest(endpoint string, params map[string]string) *baseResponse {
	nonce := time.Now().Unix()

	version := apiVersion

	if params["useApi2"] != "" {
		version = undocumentedApiVersion
		delete(params, "useApi2")
	}

	endpoint = strings.Join([]string{baseURI, version, endpoint}, "/")

	fullURI := fmt.Sprintf("%s?nonce=%d", endpoint, nonce)

	params["apikey"] = c.apiKey

	for param, value := range params {
		fullURI = fmt.Sprintf("%s&%s=%s", fullURI, param, value)
	}

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

	if resp, respErr = httpClient.Do(request); respErr != nil {
		c.setError("sendRequest - do request", respErr.Error())
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

	if err := json.Unmarshal(rawBody, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil
	}

	return &response
}
