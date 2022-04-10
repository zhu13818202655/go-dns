package main

import (
	"bytes"
	"fmt"
	. "dnsdemo/dnsKit"
	"net"
)

func main() {

	dnsMsg1 := DNSMessage{
		Header: &DNSHeader{
			ID:                  0x0010,
			QR:                  0,
			OperationCode:       0,
			AuthoritativeAnswer: 0,
			Truncation:          0,
			RecursionDesired:    1,
			RecursionAvailable:  0,
			Zero:                0,
			ResponseCode:        0,
			QuestionCount:       1,
			AnswerRRs:           0,
			AuthorityRRs:        0,
			AdditionalRRs:       0,
		},
		Questions: []*DNSQuestion{
			&DNSQuestion{
				QuestionName:  "www.baidu.com",
				QuestionType:  1,
				QuestionClass: 1,
			},
		},
	}

	dnsMsg2 := DNSMessage{
		Header: &DNSHeader{
			ID:                  0x0010,
			QR:                  0,
			OperationCode:       0,
			AuthoritativeAnswer: 0,
			Truncation:          0,
			RecursionDesired:    1,
			RecursionAvailable:  0,
			Zero:                0,
			ResponseCode:        0,
			QuestionCount:       1,
			AnswerRRs:           0,
			AuthorityRRs:        0,
			AdditionalRRs:       0,
		},
		Questions: []*DNSQuestion{
			&DNSQuestion{
				QuestionName:  "www.test2.com",
				QuestionType:  1,
				QuestionClass: 1,
			},
		},
	}

	dnsServer := "172.18.178.253:53"
	var conn net.Conn
	var err error
	if conn, err = net.Dial("udp", dnsServer); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	if _, err := conn.Write(dnsMsg1.ToBytes()); err != nil {
		fmt.Println(err.Error())
		return
	}
	buf := make([]byte, 1024)

	if length, err := conn.Read(buf); err == nil {
		result := NewDNSMessage(bytes.NewBuffer(buf[0:length]))
		fmt.Println("query:", result.Questions[0].QuestionName, ", get ip:", result.ResourceRecodes[0].RData)
	} else {
		fmt.Println(err.Error())
	}

	if _, err = conn.Write(dnsMsg2.ToBytes()); err != nil {
		fmt.Println(err.Error())
		return
	}

	buf = make([]byte, 1024)

	if length, err := conn.Read(buf); err == nil {
		result := NewDNSMessage(bytes.NewBuffer(buf[0:length]))
		fmt.Println("query:", result.Questions[0].QuestionName, ", get ip:", result.ResourceRecodes[0].RData)
	} else {
		fmt.Println(err.Error())
	}

}
