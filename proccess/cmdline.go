package process

import(
	"fmt"
	"github/wangbokun/go/types"
)

func CmdLine(pid int )(cmdLine string,err error){

	fileName := fmt.Sprintf("/proc/%s/cmdline",pid)

	if !file.IsExist(fileName) {
		
		content,err	:=	types.FileToByte(fileName)

		if err != nil {
			return err
		}
		return content
	}
}