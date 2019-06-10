package gnometimed

import (
	"fmt"
)

func ExampleConvert() {
	gtw, err := ParseXML("testdata/generated.xml")
	if err != nil {
		panic(err)
	}
	fmt.Println(gtw.Config.StartTime.Year)
	stw, err := GnomeToSimple(gtw)
	if err != nil {
		panic(err)
	}
	fmt.Println(stw)

	// Output:
	// 2018
}
