package types

import(
	"unsafe"
	"strconv"
	"strings"
	"regexp"
	"encoding/json"
	"reflect"
	
)

func IntToString(i int)(string){

	return strconv.Itoa(i)
}

func  Int64ToString(i int64)(string){

	return strconv.FormatInt(i,10)
}

func StringToInt(s string)(int,error){

	return strconv.Atoi(s)
}

func StringToint64(s string)(int64,error){

	return  strconv.ParseInt(s, 10, 64)
}

func Hex2dec(hexstr string) string{

    i, _ := strconv.ParseInt(hexstr, 16, 0)
    return strconv.FormatInt(i, 10)
}


// ToBytes interface => []byte
func ToBytes(v interface{}) (Bytes, error) {
	switch value := reflect.ValueOf(v); v.(type) {
	case string:
		return StringToBytes(value.String()), nil
	case Bytes: //[]byte
		return value.Bytes(), nil
	default:
		return json.Marshal(v)
	}
}

// BytesToString byte => string
// 直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string(Bytes{})的百倍以上，且转换量越大效率优势越明显。

func BytesToString(b Bytes) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes string => Bytes
// 直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string(Bytes{})的百倍以上，且转换量越大效率优势越明显。
// 转换之后若没做其他操作直接改变里面的字符，则程序会崩溃。
// 如 b:=String2bytes("xxx"); b[1]='d'; 程序将panic。

func StringToBytes(s string) Bytes {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*Bytes)(unsafe.Pointer(&h))
}


func StringToLower(s string) string{
	return strings.ToLower(s)
}

func FindString(regStr, s string) bool{
	match, _ := regexp.MatchString(regStr, s)
	if match {
		return true
	}
	return false
}