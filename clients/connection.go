package clients

import (
	"fmt"

	"github.com/RSOI/gateway/utils"
	"github.com/valyala/fasthttp"
)

// Connection to client
type Connection struct{}

func (conn Connection) request(url string, method string, body []byte) (int, []byte) {
	req := fasthttp.AcquireRequest()

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		utils.LOG(fmt.Sprintf("Connection error: %s", err.Error()))
		return 0, nil
	}

	return resp.Header.StatusCode(), resp.Body()
}

// GET wrapper
func (conn Connection) GET(url string) (int, []byte) {
	return conn.request(url, "GET", nil)
}

// POST wrapper
func (conn Connection) POST(url string, body []byte) (int, []byte) {
	return conn.request(url, "POST", body)
}

// PUT wrapper
func (conn Connection) PUT(url string, body []byte) (int, []byte) {
	return conn.request(url, "PUT", body)
}

// PATCH wrapper
func (conn Connection) PATCH(url string, body []byte) (int, []byte) {
	return conn.request(url, "PATCH", body)
}

// DELETE wrapper
func (conn Connection) DELETE(url string, body []byte) (int, []byte) {
	return conn.request(url, "DELETE", body)
}
