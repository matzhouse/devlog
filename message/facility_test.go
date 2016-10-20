package message

import (
	"log/syslog"
	"testing"
)

func Test_ParseFacility(t *testing.T) {

	t.Log("Checking:", syslog.LOG_KERN)
	facility := parseFacility(int(syslog.LOG_KERN))

	if facility != "Kernal" {
		t.Log("Unknown response found")
		t.Log("Response:", facility)
		t.Log("Expected Response:", "Kernal")
		t.Fail()
	}

	t.Log("Recieved:", facility)

	t.Log("Checking 9999")
	facility = parseFacility(9999)

	if facility != "Unknown" {
		t.Log("Unknown response found")
		t.Log("Response:", facility)
		t.Log("Expected Response:", "Unknown")
		t.Fail()
	}

	t.Log("Recieved:", facility)

}
