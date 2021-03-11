package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

func (c *Client) Sign(appSecret, method, path string, header http.Header, signHeaderList []string, qm map[string]interface{}, req *RequestBody) string {
	var sb strings.Builder
	sb.WriteString(strings.ToUpper(method))
	sb.WriteByte(LF)
	if header != nil {
		if v := header.Get(HTTP_HEADER_ACCEPT); v != "" {
			sb.WriteString(v)
		}
		sb.WriteByte(LF)
		if req != nil {
			// HTTP_HEADER_CONTENT_MD5
			md5Body := md5RequestBody(req)
			header.Set(HTTP_HEADER_CONTENT_MD5, md5Body)
			sb.WriteString(md5Body)
		}
		sb.WriteByte(LF)
		if v := header.Get(HTTP_HEADER_CONTENT_TYPE); v != "" {
			sb.WriteString(v)
		}
		sb.WriteByte(LF)
		if v := header.Get(HTTP_HEADER_DATE); v != "" {
			sb.WriteString(v)
		}
	}
	sb.WriteByte(LF)
	if signHeaderList != nil && len(signHeaderList) > 0 {
		buildHeaders(&sb, header, signHeaderList)
	}
	sb.WriteString(buildParams(path, qm))
	//log.Debug(sb.String())
	h := hmac.New(sha256.New, []byte(appSecret))
	h.Write([]byte(sb.String()))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func md5RequestBody(req *RequestBody) string {
	if req != nil {
		hash := md5.New()
		bs, _ := json.Marshal(req)
		hash.Write(bs)
		return base64.StdEncoding.EncodeToString(hash.Sum(nil))
	}
	return ""
}

func buildHeaders(sb *strings.Builder, header http.Header, signHeader []string) {
	for _, v := range signHeader {
		sb.WriteString(v + ":" + header.Get(v))
		sb.WriteByte(LF)
	}
}

func buildParams(path string, qm map[string]interface{}) string {
	var sb strings.Builder
	if path != "" {
		sb.WriteString(path)
	}
	mMap := make(map[string]string)
	for k, v := range qm {
		vv, ok := v.(string)
		if !ok {
			vv = convertToString(v)
		}
		mMap[k] = vv
	}
	if len(mMap) > 0 {
		sb.WriteByte('?')
		sb.WriteString(encodeParams(mMap))
	}
	return sb.String()
}

// ("bar=baz&foo=quux") sorted by key.
func encodeParams(bm map[string]string) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for i, k := range keyList {
		if i > 0 {
			buf.WriteByte('&')
		}
		if v := bm[k]; v != "" {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
