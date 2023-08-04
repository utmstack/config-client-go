package main

import (
	"fmt"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient/enum"
)

func main() {
	client := UTMStackConfigurationClient.NewUTMClient("EliHR1ZjPsaVdwg4CNI0X7c5QmCLOaJR", "192.168.1.59")
	config, err := client.GetUTMConfig(enum.O365)
	if err != nil {
		fmt.Sprintf("Error s%", err.Error())
	}
	fmt.Println(config)
}
