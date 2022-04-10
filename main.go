package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

func Get_ip_from_hosts(targetDomain string) string{
	file, err := os.Open("/etc/hosts")
	if err != nil {
      panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	// fmt.Println(string(content))
	ip_line := strings.Split(string(content), "\n")
	for _, ip_domain := range ip_line {
		ips := strings.Split(ip_domain, " ")
		if len(ips) == 2{
			ip, domain := ips[0], ips[1]
			if domain == targetDomain{
				return ip
			}
		}
		
	}
	return ""
}

func UInt32ToIP(intIP uint32) net.IP {
    var bytes [4]byte
    bytes[0] = byte(intIP & 0xFF)
    bytes[1] = byte((intIP >> 8) & 0xFF)
    bytes[2] = byte((intIP >> 16) & 0xFF)
    bytes[3] = byte((intIP >> 24) & 0xFF)
 
    return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func IPToUInt32(ipnr net.IP) uint32 {
    bits := strings.Split(ipnr.String(), ".")
	 
    b0, _ := strconv.Atoi(bits[0])
    b1, _ := strconv.Atoi(bits[1])
    b2, _ := strconv.Atoi(bits[2])
    b3, _ := strconv.Atoi(bits[3])
	 
    var sum uint32
	 
    sum += uint32(b0) << 24
    sum += uint32(b1) << 16
    sum += uint32(b2) << 8
    sum += uint32(b3)
	 
    return sum
}


func StringIPToUInt32(ip string) uint32 {
    bits := strings.Split(ip, ".")
	 
    b0, _ := strconv.Atoi(bits[0])
    b1, _ := strconv.Atoi(bits[1])
    b2, _ := strconv.Atoi(bits[2])
    b3, _ := strconv.Atoi(bits[3])
	 
    var sum uint32
	 
    sum += uint32(b0) << 24
    sum += uint32(b1) << 16
    sum += uint32(b2) << 8
    sum += uint32(b3)
	 
    return sum
}

func main() {
	ip := Get_ip_from_hosts("github.global.ssl.fastly.net")
	uint_ip := StringIPToUInt32(ip)
	ip2 := UInt32ToIP(uint_ip)
	fmt.Println(uint_ip, ip2)
    sip, _ := net.LookupIP("www.baidu.com")
    fmt.Println(sip)
    // conn, _ := net.Dial("udp", "172.18.178.253:53")
    
    
}
