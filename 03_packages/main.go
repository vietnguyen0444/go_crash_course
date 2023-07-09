package main

import (
	"example/user/hello/03_packages/strutil"
	"fmt"
	"math"
	//"github.com/ryannguyen/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor((2.7)))
	fmt.Println(math.Ceil((2.7)))
	fmt.Println(math.Sqrt((16)))
	fmt.Println(strutil.Reverse("hello"))
}
