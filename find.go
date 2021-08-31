package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	str1 := "eef3243++s"
	str2 := "Var---___1_int"
	str3 := "eef3243**s"
	str4 := "Var-----____1_int"
	fmt.Println("str1: ", find(str1))
	fmt.Println("str2: ", find(str2))
	fmt.Println("str3: ", find(str3))
	fmt.Println("str4: ", find(str4))

	ip1 := "10.10.10.1a"
	ip2 := "172.45.254.1"
	ip3 := "127.0.-1.1"
	ip4 := "64.233.16.00"
	ip5 := "0.233.16.1"
	ip6 := "256.233.16.1"


	fmt.Println("ip1 ", findIP(ip1))
	fmt.Println("ip2 ", findIP(ip2))
	fmt.Println("ip3 ", findIP(ip3))
	fmt.Println("ip4 ", findIP(ip4))
	fmt.Println("ip5 ", findIP(ip5))
	fmt.Println("ip6 ", findIP(ip6))


}

func find(str string) string {
	a := strings.Split(str, "")
	//fmt.Println("str ",a)
	for _,v  := range a {
		//fmt.Printf("value: %s index: %d\n", v, i)
		if _, err := strconv.Atoi(v); err == nil {
			return v
			//fmt.Printf("%q looks like a number.\n", v)
		}
	}
	return ""
}

func findIP(str string) bool  {
	a := strings.Split(str, ".")
	//fmt.Println("str ",a)
	if len(a) != 4 {
		return false
	}

	var response bool
	//fmt.Println("resp ", response)

	for i,v  := range a {
		val, _ := strconv.Atoi(v)
		//fmt.Println("val ", val)

		if (i == 0 || i == 3) && (val == 0 || val > 255){
			return false
		}
		if val == 0 || val > 255{ //contains char or ip more than 255
			return false
		}

		response = true
	}
	return response
}
