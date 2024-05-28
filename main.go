package main

import (
	"fmt"
)

type CustomLogWriter struct {
}

func main() {
	// alex := Person{firstName: "Alex", lastName: "Anderson", contactInfo: ContactInfo{zipCode: 2222, email: "test@mailinator.com"}}
	// alex.print()

	// alex.updateFirstName("alexander")
	// alex.print()
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// }

	// fmt.Println(colors)

	// newColors := make(map[string]string)

	// newColors["red"] = "#ff0000"

	// fmt.Println(newColors)

	// eb := EnglishBot{}
	// sb := SpanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)

	// resp, err := http.Get("https://google.com")

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// lw := CustomLogWriter{}

	// io.Copy(lw, resp.Body)

	newSquare := Square{sideLength: 12}
	newTriangle := Triangle{height: 12, base: 12}

	printArea(newSquare)
	printArea(newTriangle)
}

func (CustomLogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
