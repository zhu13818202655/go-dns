package utils

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)


func (a *Address2IP) GetIP2(address string) string {
	a.RLock()
	ip := uint32(0)
	ok := false
	if ip, ok = a.address2ip[address]; ok {
		a.RUnlock()
	} else {
		a.RUnlock()
		a.Lock()
		
		a.address2ip[address] = 
		ip = Get_ip_from_hosts()
		a.Unlock()

	}

	ipByte := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(ipByte, ip)
	return net.IPv4(ipByte[0], ipByte[1], ipByte[2], ipByte[3]).String()
}

func Get_ip_from_hosts() string {
	file, err := os.Open("/etc/hosts")
	if err != nil {
      panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	ip_line := strings.Split(string(content), '\n')
	for ip_domain := range ip_line {
		
	}
	return 
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