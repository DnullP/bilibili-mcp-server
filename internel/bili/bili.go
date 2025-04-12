package bili

import (
	"os"

	"github.com/CuteReimu/bilibili/v2"
	"github.io/dnullp/bilimcp/internel/parser"
)

var client *bilibili.Client = nil

func Client() *bilibili.Client {
	if client != nil {
		return client
	}
	var client = bilibili.New()

	userConfig, err := parser.GetConfig(os.Args[1])
	if err != nil {
		panic(err)
	}

	client.SetRawCookies(userConfig.Cookie)

	return client
}
