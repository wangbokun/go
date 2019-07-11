package types


// list 去重复

// a := "a,b,c,d,e,f,a,b,c"
// b := strings.Split(a,",")
// c := removeDuplicateElement(b)
// justString := strings.Join(c,",")
// fmt.Println(justString)

func RemoveDuplicateElement(addrs []string) []string {
    result := make([]string, 0, len(addrs))
    temp := map[string]struct{}{}
    for _, item := range addrs {
        if _, ok := temp[item]; !ok {
            temp[item] = struct{}{}
            result = append(result, item)
        }
    }
    return result
}