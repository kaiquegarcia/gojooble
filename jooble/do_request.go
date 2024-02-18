package jooble

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (j *jooble) do(
	method string,
	url string,
	payload []byte,
	dest interface{},
) (a *audit, err error) {
	a = &audit{
		payload: payload,
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := j.httpClient.Do(req)
	if err != nil {
		return
	}

	a.statusCode = resp.StatusCode

	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	a.response = response

	if resp.StatusCode >= http.StatusMultipleChoices && resp.StatusCode < http.StatusBadRequest {
		fmt.Println("warning: redirect status code during request")
	}

	if resp.StatusCode >= http.StatusBadRequest && resp.StatusCode < http.StatusInternalServerError {
		err = Err4xx
		return
	}

	if resp.StatusCode >= http.StatusInternalServerError {
		err = Err5xx
		return
	}

	err = json.Unmarshal(response, dest)
	return
}
