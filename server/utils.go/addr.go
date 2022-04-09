package utils

import "sync"

const RECV_BUF_LEN = 1024

var addr2IP Address2IP

type Address2IP struct {
	lastIP uint32 //167772160 -> 10.0.0.0
	sync.RWMutex
	address2ip map[string]uint32
}