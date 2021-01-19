package main

import (
	"fmt"
)

func main() {
	//str := map[string]string{"name": "silence"}
	//fmt.Print(str["name"])
	//str := map[[2]string][3]string{[2]string{"11", "22"}: [3]string{"11", "22", "33"}}
	//fmt.Print(str)
	//fmt.Print(["11" "22"]

	//name2 := name
	//name3 := &name
	//fmt.Print(name2)
	//name := make(map[string]string)
	//name["name"] = "silence"
	//name["age"] = "18"
	//name["name2"] = "silence"
	//fmt.Print(name["name"], "\n")
	//fmt.Print(name["age"])
	//fmt.Print(name["name2"])
	type cha struct {
		aa, bb string
	}
	type school struct {
		name  string
		age   int
		email string
		cha
	}
	var p1 = school{"11", 11, "11", cha{"aa", "cc"}}
	p1.name = "11111111"
	fmt.Println(p1.age)
	//p2 := &p1
	//
	//fmt.Print(p2, p1)
	//var aa  = "11"
	//aaa := 11
	//fmt.Println(aa,aaa)
	//var (
	//	name int
	//	str  string
	//)
	//print(name, str)
	//type name2 struct {
	//	shc int
	//
	//
	//}
	//name2[]
	//var name3 = [3]int{}
	//pa := [...]string{}
	//
	//fmt.Print(name3, pa)
	//paa := make(map[string]string)
	//paa["nameu"] = "ceshi "
	//fmt.Print(paa)
	//pe := make([]string, 2, 3)
	//pe[0] = "111"
	//fmt.Print(pe)
	//pee := &paa
	//fmt.Print(pee)
	//paa["nameu"] = "xiaodidi"
	//fmt.Print(pee, paa)
	//cfg, err := ini.Load("/Users/silence/Desktop/my.ini")
	//if err != nil {
	//	fmt.Println("错误：", err)
	//}
	//val := cfg.Section("mysqld").Key("port").Value()
	//fmt.Println(reflect.TypeOf(val))
}
