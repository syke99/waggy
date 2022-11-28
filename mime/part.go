package mime

import (
	"bufio"
	"bytes"
	"github.com/syke99/waggy/header"
	"strings"
	"sync"
)

// Part is used for accessing individual parts of a MultipartForm Request body
type Part struct {
	body    []byte
	headers *header.Header
}

// ParsePart takes the bytes representation of a MultipartForm Part
// separates the header.Header sections from the body section,
// automatically parses the header.Header values, and then
// creates and returns a *Part holding these values
func ParsePart(part []byte) *Part {
	hBytes, bBytes := separateHeaderAndBodyBytes(part)

	headers := parseHeaderFromBytes(hBytes)

	p := Part{
		body:    bBytes,
		headers: headers,
	}

	return &p
}

func separateHeaderAndBodyBytes(payload []byte) ([][]byte, []byte) {
	buf := bytes.NewBuffer(make([]byte, 0))

	scanner := bufio.NewScanner(buf)

	headerBytes := make([][]byte, 0)

	bodyBuffer := bytes.NewBuffer(make([]byte, 0))

	headersDone := false

	for scanner.Scan() {
		// if the break between headers and the body hasn't been reached,
		// add the line to the headerBytes to be turned into individual
		// *header.Headers
		if scanner.Text() != "\n" && !headersDone {
			headerBytes = append(headerBytes, scanner.Bytes())
		}

		// once the break has been reached, signal this so that the body can be separated
		if scanner.Text() == "\n" && !headersDone {
			headersDone = true
		}

		// after the break has been reached, write each line of bytes to a buffer to then be
		// placed into the *Part's body
		if headersDone {
			bodyBuffer.Write(scanner.Bytes())
		}
	}

	return headerBytes, bodyBuffer.Bytes()
}

func parseHeaderFromBytes(headerBytes [][]byte) *header.Header {
	h := header.Header{}

	for _, headerLine := range headerBytes {
		splitHeader := strings.Split(string(headerLine), ":")
		headerKey := splitHeader[0]
		headerValues := strings.Split(splitHeader[1], ";")

		var wg sync.WaitGroup

		for _, value := range headerValues {
			go func() {
				wg.Add(1)

				defer wg.Done()

				h.Add(headerKey, value)
			}()
		}

		wg.Wait()
	}

	return &h
}
