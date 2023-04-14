package types

import "github.com/AtlasInsideCorp/UTMStackConfigurationClient/enum"

type Configuration struct {
	Key         string                 `json:"key"`
	Value       string                 `json:"value"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	DataType    enum.UTMConfigDataType `json:"dataType"`
	Required    bool                   `json:"required"`
}

type ConfigurationGroup struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Configurations []Configuration `json:"configurations"`
	Active         bool            `json:"active"`
}

type ConfigurationSection struct {
	ServerName          string               `json:"serverName"`
	Module              enum.UTMModule       `json:"module"`
	Active              bool                 `json:"active"`
	ConfigurationGroups []ConfigurationGroup `json:"configurationGroups"`
}
