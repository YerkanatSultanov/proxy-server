package proxy

import (
	"io/ioutil"
	"net/http"
	"proxy-server/pkg/entity"
)

func SendRequest(req entity.ProxyRequest) (entity.ProxyResponse, error) {
	client := &http.Client{}
	request, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		return entity.ProxyResponse{}, err
	}
	for key, value := range req.Headers {
		request.Header.Set(key, value)
	}
	response, err := client.Do(request)
	if err != nil {
		return entity.ProxyResponse{}, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return entity.ProxyResponse{}, err
	}

	headers := make(map[string]string)
	for key, values := range response.Header {
		headers[key] = values[0]
	}

	return entity.ProxyResponse{
		Status:  response.StatusCode,
		Headers: headers,
		Length:  len(body),
	}, nil
}
