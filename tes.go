package main

import "fmt"

type School struct {
	name string
	city string
}
type Class struct {
	tit    string
	conste int
	*School
}

func main() {
	sch := &School{"嘻嘻", "北京"}
	var lis []Class
	for {
		var ch Class
		fmt.Print("请输入班级")
		fmt.Scan(&ch.tit)
		if ch.tit == "Q" {
			break
		}
		fmt.Print("请输入人数")
		fmt.Scan(&ch.conste)
		ch.School = sch
		lis = append(lis, ch)
	}
	fmt.Println(lis)
}
