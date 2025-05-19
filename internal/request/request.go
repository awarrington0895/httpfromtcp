package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

const carriageReturn = "\r\n"
const space = " "
const forwardSlash = "/"

func FromReader(reader io.Reader) (*Request, error) {
	rawRequest, err := io.ReadAll(reader)

	if err != nil {
		return nil, fmt.Errorf("unable to read request from parameter: %w", err)
	}

	requestStr := string(rawRequest)

	// Discarding the rest of the request for now
	requestLine := strings.Split(requestStr, carriageReturn)[0]

	requestLineParts := strings.Split(requestLine, space)

	if len(requestLineParts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", requestLine)
	}

	httpVersion := strings.Split(requestLineParts[2], forwardSlash)[1]

	parsedRequestLine := RequestLine{
		Method:        requestLineParts[0],
		RequestTarget: requestLineParts[1],
		HttpVersion:   httpVersion,
	}

	fmt.Println(parsedRequestLine)

	request := &Request{
		RequestLine: parsedRequestLine,
	}

	return request, nil
}

func parseRequestLine(requestLine string) {

}
