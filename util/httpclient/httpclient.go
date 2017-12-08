package httpclient

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type HttpClient struct {
	api    string
	hasSSL bool
	// queryID bson.ObjectId
}

//NewHttpClient -
func NewHttpClient(hasSSL bool) *HttpClient {
	c := &HttpClient{}
	// c.queryID = queryID
	c.hasSSL = hasSSL
	return c
}

//Connect - connect HTTP connection with specific url and method
func (c *HttpClient) Connect(url string, method string, params map[string]string, reqBody io.Reader) ([]byte, error) {
	log.Println("url=", url)
	client := &http.Client{}
	//habdle non-ssl way for local testing.
	if !c.hasSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}

	request, err := http.NewRequest(method, url, reqBody)

	if method == "POST" {
		request.Header.Set("Content-Type", "application/json")
	}

	//handle query string
	q := request.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	request.URL.RawQuery = q.Encode()
	fmt.Println(request.URL.String())

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

//ConnectWithMultpart -
func (c *HttpClient) ConnectWithMultpart(url string, method string, params map[string]string, fileParaname, filePath string) ([]byte, error) {
	client := &http.Client{}
	request, err := newfileUploadRequest(url, params, fileParaname, filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if !c.hasSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}
	resp, err := client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func newfileUploadRequest(uri string, params map[string]string, paramName string, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// log.Println("completed request")
	return req, err
}
