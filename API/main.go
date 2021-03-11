package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://verify5.market.alicloudapi.com/api/v1/mobile/info?appkey=6041c73d6ee47d382b72f82c"
	method := "POST"

	payload := strings.NewReader(`{
	"token":"eyAgIm8iIDogImlPUyIsICAiayIgOiAiT0E0Yitwd1dqU3BKcU9HdFdiSjFJZTdXM3ZSM0d4VnZnVzV1aE42Z3VZcFh5M0h3SjFNblwvakhyQndTa2dRZCt3UDBWZUlQTml0Rkd1OWorcFkwV3hlTElwZnh2YXI5OWs3bnQzdDJ6bzhxbm1QaklRUTBrSmpZa2FBRm9rRmxnYlwvU1RhR0l1cnNWcm5IczJNdnVNT2kwckRuZTlmaEFiRTVpczlDZzhOSmFWNWY0MnFFM1ZicWJnT1NhZ2prejBycjZTQmplN285XC9DTGM3aFNJa25QbzNxVHluQnhVclV3d3FuS0NrK1wveXVTQk9SNndidmhoXC9WbG1vTjZHS1VqVUhkZlZsXC90T1VMeUhjNURqMjVNM1wvTnp5MWt1SnlDVkRMaUY3ZjhaQ3RwcUN6Q2pzbVhMSkNVajVcL2g3bmlxWEFBMVNjSFwvTUFQbFVzVm50cWlicFdnPT0iLCAgImMiIDogIlBXMUlFb1FyR1U2RXpmK2hnVXdGOGR4bktFVldtR0ZMTm9nd1ZzejB2MmdFQkEyc0dvQXlZOGZBUFJaeGFGVU9aRllFd2VMVFwvTSttUnBTb2FvMWxjNlNUYmE4SzY4ckFtcWp6Q1FIalROcmhYczZ3WFlvOFozYThFVDFHTmJBb1JOb1lORlJaVlZBUUxiQ0txeFY3MTJxc0FyMUxqdGpCVGl2RWZwdTFcL3NLcmxsTjBUNDNyWjMxSEtleE1RbTdCcUZ5a05zSjkyM2R4VWhaZ0lEa3ozc0xPOHI4WVluR3RBS1dpRGNOTkJTUzZcL3lobzVaSlA1YytvWlZObXpuOStEbzdPKzRxbTA3djQ2KzBCVitucnluZEEwZ3VuSGtjV3pjalVCOEcwZ0lyV2dhMExpWTVhMVpieDlvb21iUWlYUG9GU3dVU1A3ZEJlTFhqRFwvN2pKUGczTEdzSVJPSDZUT1ZhYkwyM25yWnJFS09JempsM3JsRnZzXC9IdUdKM3QyY3FmbEc4QU1menQyOFY1NUIzc1VqS3JoQWhQaGFTMjJyMjRrUUVXZmZQbVl6dGQ4OVl4WU1FR0N0dHEyMm1GNGJ3UnZHZ3Y0bExcL2dIaDhnSEd6eHFSQ2xXZE5LRVZcLzFINHBEQTBBSE51MDBlXC9GMXU4ZHE5SG4rRHVFR09kQzE1MFZCNGFIRHNndXN4bmhhcHcyMitoM3RZNzJObzFrSGFNSk11ZlE4bGV5eUFKZVJTZzQrWVVkd3FxeEg5VlplMVh0NmJqbDEwblNmTm9yZ3UrZjZScWljSjluTVNNUkRTZFRKYUhadXVBVE9wdk5rczNtaU5BZjhVa0lhOHpzOEdtNHRETVEzSFh3c3VjYXY5OEk1bWlTRzcySjd0d1RuQVY2djBDRmM1d29MK21tZnNoczFXMmN1MmdJRU85eFh4cFI2K1diT0tkZTVWWitRc2pyc3NYaTVSa2Z6VjNQaFZrU29MdFFuR3ZcL0UwQ1k2MDlCbWpBUGFDMGVsSDhKK1B0bGtsUERnRFNMXC9DY2hFa1p0T2lpZjMrRGdGRkREcVBHcWZUMnorU0RudXNSVDhIdHFYbGNLaFQ1cTlZMVlBV3NJSXRSdTF3aEZSeFZDQkRBNTBQdz09In0="
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Date", "function Date() { [native code] }")
	req.Header.Add("Content-MD5", "gBLNl1PfKKo6/RvSbxgRUQ==")
	req.Header.Add("X-Ca-Nonce", "26741143-7b9f-41eb-af46-63c293863e82")
	req.Header.Add("X-Ca-Key", "203917116")
	req.Header.Add("X-Ca-Signature", "b4h+7ouXOlGV6LQ8ebNvo0m/QTtKIfzQDY/8YTNRCUE=")
	req.Header.Add("X-Ca-SignatureMethod", "HmacSHA256")
	req.Header.Add("X-Ca-Signature-Headers", "x-ca-key,x-ca-nonce,x-ca-signaturemethod")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
