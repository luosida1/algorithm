package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var prefix string
	arrylen := len(strs)
	if arrylen == 0 {
		return ""
	}

	prefix = strs[0]
	fmt.Printf("o = %s\n", prefix)
	for i := 1; i < len(strs); i++ {
		tmp := len(prefix)
		if len(prefix) > len(strs[i]) {
			tmp = len(strs[i])
		}

		fmt.Printf("str[%d] = %s\n", i, strs[i])
		j := 0
		for j = 0; j < tmp; j++ {
			if prefix[j] != strs[i][j] {
				break
			}
		}

		fmt.Printf("j = %d,strs = %s\n", j, strs[i])
		if 0 == j {
			return ""
		}
		prefix = prefix[:j]
	}
	return prefix
}

func main() {

	strs1 := []string{"gloqqq", "gloqwdhuad", "iiglioqjaineusaf"}
	result := longestCommonPrefix(strs1)

	fmt.Printf("result = %s, len = %d, stres = %s\n", result, len(strs1), strs1)

}
