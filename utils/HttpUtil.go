package utils

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type HandleHttpResponseFunc func(data []byte) (result interface{}, err error)

func HttpGet(url string, parameters map[string]string, headers map[string]string, handleHttpResponse HandleHttpResponseFunc) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	// Url
	urlParameters := buildUrlParameters(parameters)
	apiUrl := url + urlParameters

	// Request
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if nil != err {
		return result, err
	}

	// Header
	prepareRequestHeader(req, headers)

	// Response
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if nil != err {
		return result, err
	}

	// Parse response
	statusCode := resp.StatusCode
	if 200 != statusCode {
		err = errors.New(fmt.Sprintf("Invalid http status code: %d", statusCode))
		return result, err
	}

	// 内容解压缩
	var data []byte = nil
	if "gzip" == resp.Header.Get("Content-Encoding") {
		data, err = ungzip(&resp.Body)
	} else {
		data, err = ioutil.ReadAll(resp.Body)
		if nil != err {
			return result, err
		}
	}
	if nil != err {
		return result, err
	}

	if nil == data {
		err = errors.New("Invalid parameters.")
		return result, err
	}

	return handleHttpResponse(data)
}

func ungzip(compressedReadCloser *io.ReadCloser) ([]byte, error) {
	var (
		result []byte
		err    error
	)

	reader, err := gzip.NewReader(*compressedReadCloser)
	if nil != err {
		return result, err
	}

	defer reader.Close()
	result, err = ioutil.ReadAll(reader)

	return result, err
}

func prepareRequestHeader(request *http.Request, headers map[string]string) {
	if nil != headers && len(headers) > 0 {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}
}

func buildUrlParameters(parameters map[string]string) string {
	var result string = ""

	if nil != parameters && len(parameters) > 0 {
		for k, v := range parameters {
			if 0 == len(result) {
				result = "?"
			} else {
				result += "&"
			}
			result += fmt.Sprintf("%s=%s", k, v)
		}
	}

	return result
}

func HttpPost(targetUrl string, parameters map[string]string, postDataMap map[string][]string, headers map[string]string, handleHttpResponse HandleHttpResponseFunc) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	// Url
	urlParameters := buildUrlParameters(parameters)
	apiUrl := targetUrl + urlParameters

	// Request body
	requestBody := url.Values(postDataMap)

	// Request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(requestBody.Encode()))

	if nil != err {
		return result, err
	}

	// Header
	prepareRequestHeader(req, headers)

	// Response
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Failed to request url %s\n", targetUrl)
		return nil, err
	}

	// Parse response
	statusCode := resp.StatusCode
	if 200 != statusCode {
		err = errors.New(fmt.Sprintf("Invalid http status code: %d", statusCode))
		return result, err
	}

	// 内容解压缩
	var data []byte = nil
	if "gzip" == resp.Header.Get("Content-Encoding") {
		data, err = ungzip(&resp.Body)
	} else {
		data, err = ioutil.ReadAll(resp.Body)
	}
	if nil != err {
		return result, err
	}

	if nil == data {
		err = errors.New("Invalid parameters.")
		return result, err
	}

	return handleHttpResponse(data)
}
