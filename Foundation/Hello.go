package Foundation

import "fmt"

// func main() {
// 	fmt.Println("Hello World!")
// }

func Hello() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "code=%d endDate=%s\n"
	// var target_url = fmt.Sprintf(url, stockcode, enddate)
	// fmt.Println(target_url)
	fmt.Printf(url, stockcode, enddate)
}
