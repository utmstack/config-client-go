package main

import (
	"fmt"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient/enum"
)

func main() {
	client := UTMStackConfigurationClient.NewUTMClient("ksjkasdkads", "172.18.25.6")
	config, err := client.GetUTMConfig("master", enum.AWS_BEANSTALK)
	if err != nil {
		return
	}
	fmt.Println(config)
}
