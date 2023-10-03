package main

import (
	"fmt"
	UTMStackConfigurationClient "github.com/utmstack/config-client-go"
	"github.com/utmstack/config-client-go/enum"
	"os"
)

func main() {
	confKey := os.Getenv("INTERNAL_KEY")
	utmPanel := os.Getenv("BACKEND_HOST")
	client := UTMStackConfigurationClient.NewUTMClient(confKey, utmPanel)
	config, err := client.GetUTMConfig(enum.BITDEFENDER)
	if err != nil {
		fmt.Sprintf("Error s%", err.Error())
	}
	fmt.Println(config)
}
