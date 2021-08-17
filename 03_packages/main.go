package main

import (
	"fmt"
	"math"

	"github.com/dropkickfish/go_crash_course/03_packages/strutil"
) //multiple imports, again will error if not used

func main() {
	fmt.Println(math.Ceil(2.6))  //round up
	fmt.Println(math.Floor(2.6)) //round down
	fmt.Println(math.Sqrt(9))    //square root

	fmt.Println(strutil.Reverse("olleH"))
}
