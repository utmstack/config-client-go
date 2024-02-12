package UTMStackConfigurationClient_test

import (
	"fmt"
	"testing"

	UTMStackConfigurationClient "github.com/utmstack/config-client-go"
	"github.com/utmstack/config-client-go/enum"
)

func TestClient(t *testing.T) {
	for {
		confKey := "qNANPjjNm7eantt7sgld0iSWFFeGKz5i"
		utmPanel := "https://192.168.1.18"
		client := UTMStackConfigurationClient.NewUTMClient(confKey, utmPanel)
		config, err := client.GetUTMConfig(enum.O365)
		if err != nil {
			t.Errorf("Error %s", err.Error())
		}
		fmt.Println(config)
	}
	
}
