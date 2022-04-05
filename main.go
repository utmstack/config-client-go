package main

import (
	"fmt"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient/utm_crypto"
)

func main() {
	//Texto: Jorge es una perrisima
	//Pass: Utm.Pwd-53cr3t.5t4k!_3mpTy*
	//Result: IwwuELM1KR2WWvV+zOXZ5sBjAb5mjezXGErleHqGcHM=

	//var plainText = "Jorge es una perrisima"
	var passphrase = "Utm.Pwd-53cr3t.5t4k!_3mpTy*"
	encryptedData := "IwwuELM1KR2WWvV+zOXZ5sBjAb5mjezXGErleHqGcHM="
	decryptedText, err := utm_crypto.Decrypt(encryptedData, passphrase)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(decryptedText))
}
