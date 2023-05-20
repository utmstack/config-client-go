package main

import (
	"fmt"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient/enum"
)

func main() {
	client := UTMStackConfigurationClient.NewUTMClient("90415H8um7Q3V7pOnpGZQQWDw2gimoUO", "qa.utmstack.com")
	config, err := client.GetUTMConfig(1, enum.O365)
	if err != nil {
		fmt.Sprintf("Error s%", err.Error())
	}
	fmt.Println(config)
}
