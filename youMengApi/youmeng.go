package main

import (
	"context"
	"fmt"
	"gofer/pkg/log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	xhttp "gofer/pkg/http"

	"github.com/google/uuid"
)

const (
	LF = '\n'

	// http Request header key
	HTTP_HEADER_ACCEPT       = "Accept"
	HTTP_HEADER_CONTENT_MD5  = "Content-MD5"
	HTTP_HEADER_CONTENT_TYPE = "Content-Type"
	HTTP_HEADER_USER_AGENT   = "User-Agent"
	HTTP_HEADER_DATE         = "Date"

	X_CA_KEY               = "X-Ca-Key"
	X_CA_NONCE             = "X-Ca-Nonce"
	X_CA_TIMESTAMP         = "X-Ca-Timestamp"
	X_CA_SIGNATURE         = "X-Ca-Signature"
	X_CA_SIGNATURE_HEADERS = "X-Ca-Signature-Headers"

	PathGetDeviceThingInfo      = "/cloud/thing/info/get"
	PathGetDevicePropertiesInfo = "/cloud/thing/properties/get"
	PathSetDevicePropertiesInfo = "/cloud/thing/properties/set"
	PathSetDeviceStatusInfo     = "/cloud/thing/status/get"
	PathForceUnbindDevice       = "/living/cloud/user/binding/device/unbind"
	PathGetProductList          = "/cloud/thing/productList/get"
	PathGetProductInfo          = "/cloud/thing/product/get"
	PathCloudToken              = "/cloud/token"
	PathRefreshCloudToken       = "/cloud/token/refresh"
	PathGetFeiyanAccountInfo    = "/cloud/account/getByOpenId"
	PathGetFeiyanUserDeviceList = "/cloud/device/queryByUser"
	PathGetBindUserListByDevice = "/living/user/device/binding/query"
)

type Umtrack struct {
	//AuthClientId     string // 针对飞燕平台的 Oauth2 client_id
	//AuthClientSecret string // 针对飞燕平台的 Oauth2 client_secret
	Projectid  string // 飞燕项目ID
	Host       string
	AppKey     string
	AppSecret  string
	Timeout    int
	UmtrackKey string
}

type Client struct {
	Host      string // host地址，https
	AppKey    string
	AppSecret string
	Timeout   int
}

func New(c *Umtrack) (youmeng *Client) {
	youmeng = &Client{
		Host:      c.Host,
		AppKey:    c.AppKey,
		AppSecret: c.AppSecret,
		Timeout:   c.Timeout,
	}
	return youmeng
}

func (c *Client) doRequest(method, path string, qm map[string]interface{}, req *RequestBody, rsp interface{}) (err error) {
	httpCli := xhttp.NewClient().Post(c.Host + path).Type(xhttp.TypeJSON)

	signHeader := []string{X_CA_KEY, X_CA_NONCE, X_CA_TIMESTAMP}
	header := http.Header{
		HTTP_HEADER_ACCEPT:       {"application/json"},
		HTTP_HEADER_CONTENT_TYPE: {"application/json"},
		X_CA_KEY:                 {c.AppKey},
		X_CA_TIMESTAMP:           {strconv.Itoa(int(time.Now().UnixNano() / 1e6))},
		X_CA_NONCE:               {uuid.New().String()},
		X_CA_SIGNATURE_HEADERS:   {fmt.Sprintf("%s,%s,%s", X_CA_KEY, X_CA_NONCE, X_CA_TIMESTAMP)},
	}
	header.Add(X_CA_SIGNATURE, c.Sign(c.AppSecret, method, path, header, signHeader, qm, req))
	httpCli.SetHeaders(header)

	_, errs := httpCli.SendStruct(req).EndStruct(rsp)
	if len(errs) > 0 {
		return errs[0]
	}
	elem := reflect.ValueOf(rsp).Elem()
	code := elem.FieldByName("Code").Int()
	log.Info("code:", code)
	return nil
}

func (c *Client) AutoLogin(context context.Context, token string, query map[string]interface{}) (rsp *UmtrackRsp, err error) {
	path := "/api/v1/mobile/info"
	for k, v := range query {
		path += "?" + k + "=" + v.(string)
	}
	reqBody := NewRequestBody(token)
	rsp = &UmtrackRsp{}

	if err = c.doRequest("post", path, nil, reqBody, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
