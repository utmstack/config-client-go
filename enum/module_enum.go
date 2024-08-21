package enum

type UTMModule string

var AllV10Integrations = []UTMModule{
	WINDOWS_AGENT,
	SYSLOG,
	VMWARE,
	LINUX_AGENT,
	APACHE,
	AUDITD,
	ELASTICSEARCH,
	HAPROXY,
	KAFKA,
	KIBAN,
	LOGSTASH,
	MONGODB,
	MYSQL,
	NATS,
	NGINX,
	OSQUERY,
	POSTGRESQL,
	REDIS,
	TRAEFIK,
	CISCO,
	MERAKI,
	JSON,
	IIS,
	KASPERSKY,
	ESET,
	SENTINEL_ONE,
	FORTIGATE,
	SOPHOS_XG,
	MACOS,
	FILE_INTEGRITY,
	AZURE,
	O365,
	AWS_IAM_USER,
	SOPHOS,
	GCP,
	FIRE_POWER,
	MIKROTIK,
	PALO_ALTO,
	CISCO_SWITCH,
	SONIC_WALL,
	DECEPTIVE_BYTES,
	GITHUB,
	BITDEFENDER,
	SOCAI,
	NETFLOW,
	PFSENSE,
	FORTIWEB,
	AIX,
	IBM_AS_400,
}

const (
	WINDOWS_AGENT   UTMModule = "WINDOWS_AGENT"
	SYSLOG          UTMModule = "SYSLOG"
	VMWARE          UTMModule = "VMWARE"
	LINUX_AGENT     UTMModule = "LINUX_AGENT"
	APACHE          UTMModule = "APACHE"
	AUDITD          UTMModule = "AUDITD"
	ELASTICSEARCH   UTMModule = "ELASTICSEARCH"
	HAPROXY         UTMModule = "HAPROXY"
	KAFKA           UTMModule = "KAFKA"
	KIBAN           UTMModule = "KIBANA"
	LOGSTASH        UTMModule = "LOGSTASH"
	MONGODB         UTMModule = "MONGODB"
	MYSQL           UTMModule = "MYSQL"
	NATS            UTMModule = "NATS"
	NGINX           UTMModule = "NGINX"
	OSQUERY         UTMModule = "OSQUERY"
	POSTGRESQL      UTMModule = "POSTGRESQL"
	REDIS           UTMModule = "REDIS"
	TRAEFIK         UTMModule = "TRAEFIK"
	CISCO           UTMModule = "CISCO"
	MERAKI          UTMModule = "MERAKI"
	JSON            UTMModule = "JSON"
	IIS             UTMModule = "IIS"
	KASPERSKY       UTMModule = "KASPERSKY"
	ESET            UTMModule = "ESET"
	SENTINEL_ONE    UTMModule = "SENTINEL_ONE"
	FORTIGATE       UTMModule = "FORTIGATE"
	SOPHOS_XG       UTMModule = "SOPHOS_XG"
	MACOS           UTMModule = "MACOS"
	FILE_INTEGRITY  UTMModule = "FILE_INTEGRITY"
	AZURE           UTMModule = "AZURE"
	O365            UTMModule = "O365"
	AWS_IAM_USER    UTMModule = "AWS_IAM_USER"
	SOPHOS          UTMModule = "SOPHOS"
	GCP             UTMModule = "GCP"
	FIRE_POWER      UTMModule = "FIRE_POWER"
	MIKROTIK        UTMModule = "MIKROTIK"
	PALO_ALTO       UTMModule = "PALO_ALTO"
	CISCO_SWITCH    UTMModule = "CISCO_SWITCH"
	SONIC_WALL      UTMModule = "SONIC_WALL"
	DECEPTIVE_BYTES UTMModule = "DECEPTIVE_BYTES"
	GITHUB          UTMModule = "GITHUB"
	BITDEFENDER     UTMModule = "BITDEFENDER"
	SOCAI           UTMModule = "SOC_AI"
	NETFLOW         UTMModule = "NETFLOW"
	PFSENSE         UTMModule = "PFSENSE"
	FORTIWEB        UTMModule = "FORTIWEB"
	AIX             UTMModule = "AIX"
	IBM_AS_400      UTMModule = "AS_400"
)
