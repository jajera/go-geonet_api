package geonetApi

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPGetResult struct {
	Data       string
	Header     string
	StatusCode int
	Url        string
}

func SetAcceptHeader(req *http.Request, acceptHeaderValue string) {
	req.Header.Set("Accept", acceptHeaderValue)
}

func HTTPGet(url string, acceptHeader string) HTTPGetResult {
	var dataValue string
	var headerValue string
	var statusCodeValue int

	req, err := http.NewRequest("GET", url, nil)

	if err == nil {
		if VerboseLogging {
			fmt.Println("Request Method:", req.Method)
			fmt.Println("Request URL:", req.URL.String())
		}
	} else {
		fmt.Println("Error creating GET request:", err)
	}

	SetAcceptHeader(req, acceptHeader)
	headerValue = req.Header.Get("Accept")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err == nil {
		if VerboseLogging {
			fmt.Println("Request Header:", req.Header.Get("Accept"))
		}
	} else {
		fmt.Println("Error sending GET request:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if VerboseLogging {
			fmt.Println("Status code:", resp.StatusCode)
		}
	} else {
		fmt.Println("Unexpected status code:", resp.StatusCode)
	}

	statusCodeValue = resp.StatusCode

	responseBody, err := io.ReadAll(resp.Body)
	if err == nil {
		dataValue = string(responseBody)
		if VerboseLogging {
			fmt.Println("Response Body:", dataValue)
		}
	} else {
		fmt.Println("Error reading response body:", err)
	}

	return HTTPGetResult{
		Data:       dataValue,
		Header:     headerValue,
		StatusCode: statusCodeValue,
		Url:        url,
	}
}
