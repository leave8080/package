package main

import (
	"fmt"
	"net/url"
)

func UrlJoin(host, path string) (string, error) {
	base, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	return base.ResolveReference(u).String(), nil
}

func main() {
	ss, err := UrlJoin("https://oss.cloud-dev.mxchip.com/", "https://vbs9010dev.oss-cn-shanghai.aliyuncs.com/images/docs/1626139375295.png")
	fmt.Println(ss, err)
}
