package types

import (
	"strconv"
	"net"
	"fmt"
)


func HexToIPv4(s string) net.IP {

	ipList := make([]string, 4)

	for i := range ipList {

		part, _ := strconv.ParseInt(s[i*2:i*2+2], 16, 0)
		ipList[4-1-i] = fmt.Sprintf("%d", part)
	}

	return net.ParseIP(strings.Join(ipList, "."))
}
   
func HexToIPv6(s string) net.IP {

	ipList := make([]string, 8)
	for i := range ipList {
		ipList[i] = s[i*4 : i*4+4]
	}

	return net.ParseIP(strings.Join(ipList, ":"))
}

// 方法一
func HexToIP(s string) net.IP {

	switch len(s) {

	case 8:
	 	return HexToIPv4(s)
	case 32:
	 	return HexToIPv6(s)
	default:
		 return s
		 
	}
}


//方法二
func HexToIp(hexstr string)(string){
	var ip string

	switch (len(hexstr)){
		case 8:
			i1, _ := strconv.ParseInt(hexstr[6:8], 16, 0)
			i2, _ := strconv.ParseInt(hexstr[4:6], 16, 0)
			i3, _ := strconv.ParseInt(hexstr[2:4], 16, 0)
			i4, _ := strconv.ParseInt(hexstr[0:2], 16, 0)

			ip = fmt.Sprintf("%d.%d.%d.%d", i1, i2, i3, i4)

			return ip
			
		case 32:
			hex_ip:=fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s",hexstr[0:4],
			hexstr[4:8],
			hexstr[8:12],
			hexstr[12:16],
			hexstr[16:20],
			hexstr[20:24],
			hexstr[24:28],
			hexstr[28:32])

			return net.ParseIP(hex_ip).String()
			
		default:
			return hexstr
	}
}