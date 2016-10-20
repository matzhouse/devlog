package message

import (
	"log/syslog"
)

func parseSeverity(in int) (priority string) {

	switch in {

	case int(syslog.LOG_EMERG):
		return "Emergency"

	case int(syslog.LOG_ALERT):
		return "Alert"

	case int(syslog.LOG_CRIT):
		return "Critical"

	case int(syslog.LOG_ERR):
		return "Error"

	case int(syslog.LOG_WARNING):
		return "Warning"

	case int(syslog.LOG_NOTICE):
		return "Notice"

	case int(syslog.LOG_INFO):
		return "Info"

	case int(syslog.LOG_DEBUG):
		return "Debug"

	}
	return "Unknown"

}
