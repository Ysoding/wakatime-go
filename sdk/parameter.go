package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	urllib "net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/Ysoding/wakatime-go/sdk/request"
)

var baseRequestFields []string

func init() {
	req := request.BaseRequest{}
	reqType := reflect.TypeOf(req)
	for i := 0; i < reqType.NumField(); i++ {
		baseRequestFields = append(baseRequestFields, reqType.Field(i).Name)
	}

	for i := 0; i < len(noEscape); i++ {
		// expects every character except these to be escaped
		noEscape[i] = (i >= 'A' && i <= 'Z') ||
			(i >= 'a' && i <= 'z') ||
			(i >= '0' && i <= '9') ||
			i == '-' ||
			i == '.' ||
			i == '_' ||
			i == '~'
	}
}

type ParameterBuilder interface {
	BuildURL(url string, paramJSON []byte) (string, error)
	BuildBody(paramJSON []byte) (string, error)
}

func GetParameterBuilder(method string) ParameterBuilder {
	if method == MethodGet || method == MethodDelete || method == MethodHead {
		return &WithoutBodyBuilder{}
	}
	return &WithBodyBuilder{}
}

type WithBodyBuilder struct {
}

func (b *WithBodyBuilder) BuildURL(url string, paramJSON []byte) (string, error) {
	return buildURL(url, paramJSON, false)
}

func (b *WithBodyBuilder) BuildBody(paramJSON []byte) (string, error) {
	paramMap := make(map[string]interface{})
	err := json.Unmarshal(paramJSON, &paramMap)
	if err != nil {
		return "", err
	}

	for k := range paramMap {
		if includes(baseRequestFields, k) {
			delete(paramMap, k)
		}
	}

	body, _ := json.Marshal(paramMap)
	return string(body), nil
}

type WithoutBodyBuilder struct {
}

func (b *WithoutBodyBuilder) BuildURL(url string, paramJSON []byte) (string, error) {
	return buildURL(url, paramJSON, true)
}

func (b *WithoutBodyBuilder) BuildBody(paramJSON []byte) (string, error) {
	return "", nil
}

func buildURL(reqURL string, paramJSON []byte, hashParam bool) (string, error) {
	paramMap := make(map[string]interface{})
	err := json.Unmarshal(paramJSON, &paramMap)
	if err != nil {
		return "", err
	}

	replacedURL, err := replaceURLWithPathParam(reqURL, paramMap)
	if err != nil {
		return "", err
	}

	var vals urllib.Values
	if hashParam {
		vals = buildQueryParams(paramMap, reqURL)
	}

	encodedURL, err := encodeURL(replacedURL, vals)
	if err != nil {
		return "", err
	}

	return encodedURL, nil
}

func buildQueryParams(paramMap map[string]interface{}, url string) urllib.Values {
	values := urllib.Values{}
	accessMap(paramMap, url, "", values)
	return values
}

func accessMap(paramMap map[string]interface{}, url, prefix string, values urllib.Values) {
	for k, v := range paramMap {
		if shouldIgnoreField(url, k) {
			continue
		}

		switch e := v.(type) {
		case []interface{}:
			for i, n := range e {
				switch f := n.(type) {
				case map[string]interface{}:
					subPrefix := fmt.Sprintf("%s.%d.", k, i+1)
					accessMap(f, url, subPrefix, values)
				case nil:
				default:
					values.Set(fmt.Sprintf("%s%s.%d", prefix, k, i+1), fmt.Sprintf("%s", n))
				}
			}
		case nil:
		default:
			values.Set(fmt.Sprintf("%s%s", prefix, k), fmt.Sprintf("%v", v))
		}
	}
}

func shouldIgnoreField(url, field string) bool {
	// path paramter
	flag := "{" + field + "}"
	if strings.Contains(url, flag) {
		return true
	}

	if includes(baseRequestFields, field) {
		return true
	}

	return false
}

func replaceURLWithPathParam(url string, params map[string]interface{}) (string, error) {
	r, _ := regexp.Compile("{[a-zA-Z0-9-_]+}")
	matches := r.FindAllString(url, -1)
	for _, match := range matches {
		field := strings.TrimLeft(match, "{")
		field = strings.TrimRight(field, "}")
		value, ok := params[field]
		if !ok {
			return "", errors.New("Can't find path paramter: " + field)
		}

		valueStr := fmt.Sprintf("%v", value)
		url = strings.Replace(url, match, valueStr, -1)
	}
	return url, nil
}

func encodeURL(url string, values urllib.Values) (string, error) {
	urlObj, err := urllib.Parse(url)
	if err != nil {
		return "", err
	}
	urlObj.RawPath = EscapePath(urlObj.Path, false)
	uri := urlObj.EscapedPath()

	if values != nil {
		queryParam := values.Encode()
		// RFC 3986, ' ' should be encoded to 20%, '+' to 2B%
		queryParam = strings.Replace(queryParam, "+", "%20", -1)
		if queryParam != "" {
			uri += "?" + queryParam
		}
	}
	return uri, nil
}

var noEscape [256]bool

func EscapePath(path string, encodeSep bool) string {
	var buf bytes.Buffer
	for i := 0; i < len(path); i++ {
		c := path[i]
		if noEscape[c] || (c == '/' && !encodeSep) {
			buf.WriteByte(c)
		} else {
			fmt.Fprintf(&buf, "%%%02X", c)
		}
	}
	return buf.String()
}
