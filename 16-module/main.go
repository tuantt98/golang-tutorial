package main

import (
	"fmt"

	"gmodule/gm1"

	"github.com/leekchan/accounting"
)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))
	gm1.Greeting()
}
