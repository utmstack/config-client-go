package UTMStackConfigurationClient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/AtlasInsideCorp/AtlasInsideAES"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient/enum"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient/types"
	"github.com/AtlasInsideCorp/UTMStackConfigurationClient/util"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	passphrase = "Utm.Pwd-53cr3t.5t4k!_3mpTy*"
)

type UTMConfigClient struct {
	ConnectionKey  string `json:"connectionKey"`
	MasterLocation string `json:"masterLocation"`
}

var lock = &sync.Mutex{}

var utmClient *UTMConfigClient

func NewUTMClient(connectionKey, masterLocation string) *UTMConfigClient {
	if utmClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if utmClient == nil {
			fmt.Println("Creating UTMConfigClient instance")
			return &UTMConfigClient{
				ConnectionKey:  connectionKey,
				MasterLocation: masterLocation,
			}
		} else {
			fmt.Println("UTMConfigClient instance already created.")
		}
	} else {
		fmt.Println("UTMConfigClient instance already created.")
	}

	return utmClient
}

func (s *UTMConfigClient) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set(util.UTMInternalKeyHeaderName, s.ConnectionKey)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (s *UTMConfigClient) CreateUTMConfig(config *types.ConfigurationSection) error {
	url := fmt.Sprintf(util.RegisterConfigURL, s.MasterLocation)
	fmt.Println(url)
	config = encryptDecryptConfValues(config, "decrypt")
	j, err := json.Marshal(config)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	if config != nil {

	}
	_, err = s.doRequest(req)
	return err
}

func (s *UTMConfigClient) GetUTMConfig(serverId int, module enum.UTMModule) (*types.ConfigurationSection, error) {
	url := fmt.Sprintf(util.GetConfigURL, s.MasterLocation, serverId, module)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data types.ConfigurationSection
	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}
	dat := encryptDecryptConfValues(&data, "decrypt")
	return dat, nil
}

func encryptDecryptConfValues(utmConfig *types.ConfigurationSection, action string) *types.ConfigurationSection {
	for _, confGroup := range utmConfig.ConfigurationGroups {
		for _, conf := range confGroup.Configurations {
			if conf.ConfValue != "" && conf.ConfDataType == enum.PasswordType {
				if action != "encrypt" {
					conf.ConfValue, _ = AtlasInsideAES.AESDecrypt(conf.ConfValue, []byte(passphrase))
				} else {
					conf.ConfValue, _ = AtlasInsideAES.AESEncrypt(conf.ConfValue, []byte(passphrase))
				}
			}
		}
	}
	return utmConfig
}
