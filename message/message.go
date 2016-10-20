package message

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
)

type Message struct {
	Severity string
	Facility string
	Tag      string
	Content  string
}

func Parse(in io.Reader) (err error) {

	s := bufio.NewScanner(in)

	for s.Scan() {
		l := s.Bytes()

		sev, fac, tag, msg := splitLog(l)

		m := &Message{}

		m.Facility = parseFacility(fac)
		m.Severity = parseSeverity(sev)

		m.Tag = tag

		m.Content = msg

		fmt.Println(m)

	}

	// return an error if there was something wrong with
	// the scanning
	if err = s.Err(); err != nil {
		return err
	}

	return nil

}

func splitLog(in []byte) (severity, facility int, tag, message string) {

	pri := in[bytes.Index(in, []byte("<"))+1 : bytes.Index(in, []byte(">"))]

	priInt, err := strconv.Atoi(string(pri))

	if err != nil {
		log.Println(err)
	}

	facility, severity = splitPRI(int(priInt))

	message = string(in[bytes.Index(in, []byte(": "))+1:])

	return
}

func splitPRI(pri int) (facility, severity int) {

	severity = pri & 0x07

	facility = pri >> 3

	return
}
