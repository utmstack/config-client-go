package types

import "github.com/utmstack/config-client-go/enum"

type ConfigurationSection struct {
	ID                  int           `json:"id"`
	ServerID            int           `json:"serverId"`
	PrettyName          string        `json:"prettyName"`
	ModuleName          string        `json:"moduleName"`
	ModuleDescription   string        `json:"moduleDescription"`
	ModuleActive        bool          `json:"moduleActive"`
	ModuleIcon          string        `json:"moduleIcon"`
	ModuleCategory      string        `json:"moduleCategory"`
	LiteVersion         bool          `json:"liteVersion"`
	NeedsRestart        bool          `json:"needsRestart"`
	ConfigurationGroups []ModuleGroup `json:"moduleGroups"`
	Activatable         bool          `json:"activatable"`
}

type ModuleGroup struct {
	ID               int             `json:"id"`
	ModuleID         int             `json:"moduleId"`
	GroupName        string          `json:"groupName"`
	GroupDescription string          `json:"groupDescription"`
	Configurations   []Configuration `json:"moduleGroupConfigurations"`
}

type Configuration struct {
	ID              int                    `json:"id"`
	GroupID         int                    `json:"groupId"`
	ConfKey         string                 `json:"confKey"`
	ConfValue       string                 `json:"confValue"`
	ConfName        string                 `json:"confName"`
	ConfDescription string                 `json:"confDescription"`
	ConfDataType    enum.UTMConfigDataType `json:"confDataType"`
	ConfRequired    bool                   `json:"confRequired"`
}
