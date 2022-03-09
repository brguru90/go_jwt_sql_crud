// // Go program which illustrates the
// // concept of Defer while panicking
// package main

// import (
// 	"fmt"
// 	"time"
// )

// // set-cookie: access_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImVtYWlsIjoiYWEiLCJ1dWlkIjoiZTljNGQ4MWUtMWJhYy00NDU3LTlkMzctMzU3ZThiODlkM2I2In0sInVuYW1lIjoiYWEiLCJ0b2tlbl9pZCI6ImU5YzRkODFlLTFiYWMtNDQ1Ny05ZDM3LTM1N2U4Yjg5ZDNiNl8waFdrSWdXc01GR0JCaTdMc3dVYjduTWF5L3gxUjQ0VDJuWXBvNC95WE5QdU9BSEZuVURZSC9NeFVwTHdwWkVTR0NrPV8xNjQ2NjY1MDMwMjI5IiwiZXhwIjoxNjQ2NjY4NjMwMjI5LCJpYXQiOjE2NDY2NjUwMzAyMjksImNzcmZfdG9rZW4iOiJLTVZCVDU3bUhMWVhEb01DVEtXNWxiWWdyQnZ4dzNBSGFYYlR4SXd3dm9VbVRGVG9XNngwZ3VVTUtFWW0zSGpreSt4d0pFNVZIa1MzYy9DL05YZDg4THAyaEdTRTJmcjc2KzFtNXJBNjFqb2IrOEVxUTJKUHFhbHFZa1hQYk5kblJnQTl0UT09In0.4f2KxCvLEVJWA9i-M-Y6Ra_iP8T3FmS2oNTuD7NtYAo; Path=/; Expires=Mon, 07 Mar 2022 15:57:10 GMT; HttpOnly; SameSite=Strict
// func main() {
// 	_time := time.Now().UTC()
// 	fmt.Println(fmt.Sprintf(`%s, %02d %s %d %02d:%02d:%02d %s`, _time.Weekday().String()[0:3], _time.Day(), _time.Month().String()[0:3], _time.Year(), _time.Hour(), _time.Minute(), _time.Second(), _time.Location()))
// 	fmt.Println(_time.UTC())
// 	fmt.Println(_time.Zone())
// }

package main

import (
	"fmt"
	"time"
)

func main() {

	var t1 = time.Now()
	time.Sleep(time.Second * 1)
	var t2 = time.Now()
	var diff_t = t2.Sub(t1)

	// random duration in nanoseconds
	var d time.Duration = time.Duration(diff_t.Nanoseconds())

	// converts d in hour
	fmt.Println(d.Hours())

	// converts d in minutes
	fmt.Println(d.Minutes())

	// converts d in seconds
	fmt.Println(d.Seconds())

	// converts d in milliseconds
	fmt.Println(d.Milliseconds())

	// converts d in microseconds
	fmt.Println(d.Microseconds())

	// string representation go d
	fmt.Println(d.String())
}
