package request

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type parserStatus int

const (
	initialized parserStatus = iota
	done
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
	parsedRequestLine, err := parseRequestLine(requestStr)

	if err != nil {
		return nil, fmt.Errorf("unable to parse request line: %w", err)
	}

	fmt.Println(parsedRequestLine)

	request := &Request{
		RequestLine: *parsedRequestLine,
	}

	return request, nil
}

func parseRequestLine(requestStr string) (*RequestLine, error) {
	requestLine := strings.Split(requestStr, carriageReturn)[0]

	requestLineParts := strings.Split(requestLine, space)

	if len(requestLineParts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", requestLine)
	}

	httpVersionParts := strings.Split(requestLineParts[2], forwardSlash)

	if len(httpVersionParts) != 2 {
		return nil, fmt.Errorf("invalid http version: %s", requestLineParts[2])
	}

	if httpVersionParts[0] != "HTTP" {
		return nil, fmt.Errorf("invalid http version: %s", requestLineParts[2])
	}

	if httpVersionParts[1] != "1.1" {
		return nil, fmt.Errorf("unsupported http version: %s", httpVersionParts[1])
	}

	httpVersion := httpVersionParts[1]

	parsedRequestLine := &RequestLine{
		Method:        requestLineParts[0],
		RequestTarget: requestLineParts[1],
		HttpVersion:   httpVersion,
	}

	for _, letter := range parsedRequestLine.Method {
		if !unicode.IsUpper(letter) {
			return nil, fmt.Errorf("method must contain only uppercase alphabetic characters: %s", parsedRequestLine.Method)
		}
	}

	return parsedRequestLine, nil
}
