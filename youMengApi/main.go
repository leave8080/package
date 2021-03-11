package main

import (
	"context"
	"gofer/pkg/log"
)

func main() {
	youmeng := &Umtrack{
		Host:      "https://verify5.market.alicloudapi.com",
		AppKey:    "203917116",
		AppSecret: "il2aUjYyRL6aCRNkbiFJb1eySVbRdVsG",
		Timeout:   60,
	}
	client := New(youmeng)
	query := map[string]interface{}{"appkey": "6041c73d6ee47d382b72f82c"}
	ret, err := client.AutoLogin(context.Background(), "eyJjIjoiWlNtYW5ETXdqNmFMeHN0ZlwvcmVMNFhNSHFia0xETWFyVk9IMUkyOW5Rc0dIY2tOODNqdVpxczhFVVVVdmIyZVNcL09IOFwvTnZlQ3dQblxuNGxMMGp3cFFkWmozYThSUHAwMldUVWZmYkloUWtjS2pZRmROMEI4UGJmRThXRmVGKzNLNVl3d1djSElvK24zc3I5NWlVUllJMElIWlxuNjFMQyt6cEl5czJYS21zTTNlQnNXa1wvTVJIR2NIQ2h1OTVGN0xmMlY3TFcrRzdPcDRDY1FCY2JhZFJwNFlCRzFXVmpOdzczcDh4dzVcbkJZSnE3TlJCNmVtTUY4SzhjUXdiektPUnZ0RUdVNnczbGVkYTRWdmZXbFlaRDJra1RkaTBNZHQ5STgyUVdwb2lRSXFmSitFWTk5SVlcbnJ0UzlnNWgyMU5qaW9vVkdHSkp5M2FyUTZ4SGl0bTAwYnBPOWRrK1JNVmxiOWpxU1loWUtESHUzN2pTTWplRUhGYm5PZzAxU3l5eDVcbmFUQlJrbXd0eDVBcWZKTFozNE52S1F4ZjVqUSs0WDBra3VDTEZucUlrZWlHYU5XUVltd2lsdkpuTXlCN1hHTWhkSDd2cUc0SFhzbW5cbldUYlhoQUlhZ01tKzNcL1NSRHJObGw4REFjZzdpT1dMZXpSR2tMMmFtRDdIQW4xaW83STJKTVRqaysxOWxubnl4S0Fhbk9jS0JZeTl0XG5XOVIwN0xvayt5Vk1QSmVaeWt4NERmaXNsQXYyR04wN2l6WTlwaENVNGNvaWhZSTFKMlwvK0hVN0Rnb1Y3eWpOcHdYSzdCTVozbFZlTVxuTVRkOVRFTE56UFE9XG4iLCJrIjoiWXdHbnprYmhMUXJKRitKUDRVMEJhY096elhOTHVUUTU1N0NCczFuMjJNMXpWZWNMMzhwTHgrUzVcL2Y3cTZQbHBRdzB5TkNDY21cL0diSXhpcUFjUTBIZ09jNVMzaG40ZjhrQnQ4SmZnVEVsSTRMM2lOeU4xb3JcL1ZKTlBPa3lEWFoyVEsyVWV5RFJKYjFTcVZZa3ErNnV0WlhoZjRKRHNNR01Sc01yaUEzNU5MS1VXblwvUGZhUDV1Nm1BY0FXRVwvMVNHZmpsbmUxbDJBcmJxVzZ1VGl4NnR2a0x6YUVrbHV0VUxuU0s2MDZ2VkF1NFBBUEpNTVl6cWdtaFpQVE9NUHpFbnVYckY2XC9SWnJ1TnNtb3VPbUZUVjBYdXJxbGVKMzBsNytNZTRDRHVucks2VjNwR0hRWjBnTzh4SkRsakFXb09MNFNEaVJXQXFZSnd0Sm1DbEdONTd3PT0iLCJvIjoiQW5kcm9pZCJ9", query)
	if err != nil {
		log.Debug("err:", err)
	} else {
		log.Info("phone:", ret.Data.Mobile)
	}
}
