package types

import (
	"strconv"
	"net"
	"fmt"
)


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