package utmconfig

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/utmstack/config-client-go/enum"
	"github.com/utmstack/config-client-go/types"
	"github.com/utmstack/config-client-go/util"
)

type UTMConfigClient struct {
	ConnectionKey  string `json:"connectionKey"`
	MasterLocation string `json:"masterLocation"`
}

var utmClient *UTMConfigClient

var one sync.Once

func NewUTMClient(connectionKey, masterLocation string) *UTMConfigClient {
	one.Do(func() {
		fmt.Println("Creating UTMConfigClient instance")
		utmClient = &UTMConfigClient{
			ConnectionKey:  connectionKey,
			MasterLocation: masterLocation,
		}
	})

	return utmClient
}

func (s *UTMConfigClient) doRequest(req *http.Request) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := new(http.Client)

	req.Header.Set(util.UTMInternalKeyHeaderName, s.ConnectionKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("body: %s; header: %s", body, resp.Header.Get("X-UtmStack-error"))
	}

	return body, nil
}

func (s *UTMConfigClient) CreateUTMConfig(config *types.ConfigurationSection) error {
	url := fmt.Sprintf(util.RegisterConfigURL, s.MasterLocation)
	
	j, err := json.Marshal(config)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	_, err = s.doRequest(req)
	
	return err
}

func (s *UTMConfigClient) GetUTMConfig(module enum.UTMModule) (*types.ConfigurationSection, error) {
	url := fmt.Sprintf(util.GetConfigURL, s.MasterLocation, module)
	
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

	return &data, nil
}

func (s *UTMConfigClient) GetAllUTMConfig() ([]*types.ConfigurationSection, error) {
	confs := []*types.ConfigurationSection{}

	for _, module := range enum.AllV10Integrations {
		url := fmt.Sprintf(util.GetConfigURL, s.MasterLocation, module)
		
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
		
		confs = append(confs, &data)
	}

	return confs, nil
}
