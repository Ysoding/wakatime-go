package response

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response interface {
	ParseErrorFromHTTPResponse(code int, body []byte) error
}

type ErrorResponse struct {
}

type BaseResponse struct {
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(code int, body []byte) error {
	return nil
}

func ParseFromHttpResponse(rawResponse *http.Response, response Response) error {
	defer rawResponse.Body.Close()
	body, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return err
	}

	if rawResponse.StatusCode < 200 || rawResponse.StatusCode > 299 {
		return fmt.Errorf("request fail with status: %s, with body: %s", rawResponse.Status, body)
	}

	if err := response.ParseErrorFromHTTPResponse(rawResponse.StatusCode, body); err != nil {
		return err
	}

	return json.Unmarshal(body, &response)
}
