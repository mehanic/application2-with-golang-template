package request

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Request(reader *bufio.Reader) (method, path string, headers map[string]string, body string, err error) {
	headers = make(map[string]string)

	// 1. Request line
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)
	if len(parts) < 2 {
		err = fmt.Errorf("malformed request line: %s", line)
		return
	}
	method = parts[0]
	path = parts[1]

	// 2. Headers
	for {
		hline, errH := reader.ReadString('\n')
		if errH != nil {
			err = errH
			return
		}
		hline = strings.TrimSpace(hline)
		if hline == "" { // end of headers
			break
		}
		colonIndex := strings.Index(hline, ":")
		if colonIndex > -1 {
			key := strings.TrimSpace(hline[:colonIndex])
			val := strings.TrimSpace(hline[colonIndex+1:])
			headers[strings.ToLower(key)] = val
		}
	}

	// 3. Body (for POST)
	// if clStr, ok := headers["content-length"]; ok {
	// cl, convErr := strconv.Atoi(clStr)
	// if convErr == nil && cl > 0 {
	// bodyBytes := make([]byte, cl)
	// _, err = io.ReadFull(reader, bodyBytes)
	// if err != nil {
	// return
	// }
	// body = string(bodyBytes)
	// }
	// }

	// 3. Body (для POST, PUT, PATCH и chunked)
	if method == "POST" || method == "PUT" || method == "PATCH" {
		body, err = ReadBody(reader, headers)
		if err != nil {
			return
		}
	}

	return
}

func ReadBody(reader *bufio.Reader, headers map[string]string) (string, error) {
	var bodyBytes []byte

	if strings.ToLower(headers["transfer-encoding"]) == "chunked" {
		for {

			sizeLine, err := reader.ReadString('\n')
			if err != nil {
				return "", err
			}
			sizeLine = strings.TrimSpace(sizeLine)
			if sizeLine == "" {
				continue
			}

			size, err := strconv.ParseInt(sizeLine, 16, 64)
			if err != nil {
				return "", fmt.Errorf("invalid chunk size: %v", err)
			}

			if size == 0 {
				_, _ = reader.ReadString('\n')
				break
			}

			chunk := make([]byte, size)
			_, err = io.ReadFull(reader, chunk)
			if err != nil {
				return "", err
			}
			bodyBytes = append(bodyBytes, chunk...)

			_, err = reader.ReadString('\n')
			if err != nil {
				return "", err
			}
		}
	} else if clStr, ok := headers["content-length"]; ok {
		cl, err := strconv.Atoi(clStr)
		if err != nil {
			return "", err
		}
		if cl > 0 {
			body := make([]byte, cl)
			_, err = io.ReadFull(reader, body)
			if err != nil {
				return "", err
			}
			bodyBytes = body
		}
	}

	return string(bodyBytes), nil
}
