package message

import (
	"log/syslog"
)

func parseFacility(in int) (facility string) {

	switch in {

	case int(syslog.LOG_KERN):
		return "Kernal"

	case int(syslog.LOG_USER):
		return "User"

	case int(syslog.LOG_MAIL):
		return "Mail"

	case int(syslog.LOG_DAEMON):
		return "Daemon"

	case int(syslog.LOG_AUTH):
		return "Authentication"

	case int(syslog.LOG_SYSLOG):
		return "Syslog"

	case int(syslog.LOG_LPR):
		return "Line Printer"

	case int(syslog.LOG_NEWS):
		return "News"

	case int(syslog.LOG_UUCP):
		return "UUCP"

	case int(syslog.LOG_CRON):
		return "cron"

	case int(syslog.LOG_AUTHPRIV):
		return "Auth Privaliges"

	case int(syslog.LOG_FTP):
		return "FTP"

	case int(syslog.LOG_LOCAL0):
		return "Local 0"

	case int(syslog.LOG_LOCAL1):
		return "Local 1"

	case int(syslog.LOG_LOCAL2):
		return "Local 2"

	case int(syslog.LOG_LOCAL3):
		return "Local 3"

	case int(syslog.LOG_LOCAL4):
		return "Local 4"

	case int(syslog.LOG_LOCAL5):
		return "Local 5"

	case int(syslog.LOG_LOCAL6):
		return "Local 6"

	case int(syslog.LOG_LOCAL7):
		return "Local 7"
	}

	return "Unknown"

}
