package main

import "bytes"

//对可变参数列表进行遍历
func joinStrings(slist ...string) string {
	var buf bytes.Buffer
	for _, s := range slist {
		buf.WriteString(s)
	}
	return buf.String()
}
