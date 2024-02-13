package enum

type UTMModule string

var AllV10Integrations = []UTMModule{
	WINDOWS_AGENT,
	SYSLOG,
	LINUX_LOGS,
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
	UFW,
	MIKROTIK,
	PALO_ALTO,
	CISCO_SWITCH,
	SONIC_WALL,
	DECEPTIVE_BYTES,
	GITHUB,
	BITDEFENDER,
	SOCAI,
}

const (
	FILE_INTEGRITY     UTMModule = "FILE_INTEGRITY"
	VULNERABILITIES              = "VULNERABILITIES"
	ASSET_SCANNER                = "ASSET_SCANNER"
	AD_AUDIT                     = "AD_AUDIT"
	NETFLOW                      = "NETFLOW"
	WINDOWS_AGENT                = "WINDOWS_AGENT"
	SYSLOG                       = "SYSLOG"
	LINUX_LOGS                   = "LINUX_LOGS"
	VMWARE                       = "VMWARE"
	WEBROOT                      = "WEBROOT"
	AWS_TRAFFIC_MIRROR           = "AWS_TRAFFIC_MIRROR"
	AWS_IAM_USER                 = "AWS_IAM_USER"
	AWS_CLOUDTRAIL               = "AWS_CLOUDTRAIL"
	AWS_SQL_SERVER               = "AWS_SQL_SERVER"
	AWS_POSTGRESQL               = "AWS_POSTGRESQL"
	AWS_BEANSTALK                = "AWS_BEANSTALK"
	AWS_FARGATE                  = "AWS_FARGATE"
	AWS_LAMBDA                   = "AWS_LAMBDA"
	SOPHOS                       = "SOPHOS"
	AZURE                        = "AZURE"
	O365                         = "O365"
	IIS                          = "IIS"
	GCP                          = "GCP"
	JSON                         = "JSON"
	MACOS_AGENT                  = "MACOS_AGENT"
	LINUX_AGENT                  = "LINUX_AGENT"
	APACHE                       = "APACHE"
	APACHE2                      = "APACHE2"
	AUDITD                       = "AUDITD"
	ELASTICSEARCH                = "ELASTICSEARCH"
	HAPROXY                      = "HAPROXY"
	ICINGA                       = "ICINGA"
	KAFKA                        = "KAFKA"
	KIBAN                        = "KIBAN"
	LOGSTASH                     = "LOGSTASH"
	MONGODB                      = "MONGODB"
	MYSQL                        = "MYSQL"
	NATS                         = "NATS"
	NGINX                        = "NGINX"
	OSQUERY                      = "OSQUERY"
	PENSANDO                     = "PENSANDO"
	POSTGRESQL                   = "POSTGRESQL"
	REDIS                        = "REDIS"
	SANTA                        = "SANTA"
	SYSTEM                       = "SYSTEM"
	TRAEFIK                      = "TRAEFIK"
	CISCO                        = "CISCO"
	MERAKI                       = "MERAKI"
	KASPERSKY                    = "KASPERSKY"
	ESET                         = "ESET"
	FORTIGATE                    = "FORTIGATE"
	SOPHOS_XG                    = "SOPHOS_XG"
	SENTINEL_ONE                 = "SENTINEL_ONE"
	FIRE_POWER                   = "FIRE_POWER"
	UFW                          = "UFW"
	MACOS                        = "MACOS"
	MIKROTIK                     = "MIKROTIK"
	PALO_ALTO                    = "PALO_ALTO"
	CISCO_SWITCH                 = "CISCO_SWITCH"
	SONIC_WALL                   = "SONIC_WALL"
	DECEPTIVE_BYTES              = "DECEPTIVE_BYTES"
	GITHUB                       = "GITHUB"
	SALESFORCE                   = "SALESFORCE"
	BITDEFENDER                  = "BITDEFENDER"
	IBM_AS_400                   = "IBM_AS_400"
	SOCAI                        = "SOC_AI"
)
