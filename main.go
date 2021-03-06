package main

import (
	"fmt"
	"github.com/charlerive/library/blackscholes"
)

func main() {
	bsm := blackscholes.NewBS("c", 4600, 5000, 0.1644, 0.025, 616.78, 0.01, 3, 0.3)
	fmt.Println(fmt.Sprintf("bsm: %+v", bsm))

	bsm = blackscholes.NewBS("p", 4600, 5000, 0.1644, 0.025, 996.27, 0.01, 3, 0.3)
	fmt.Println(fmt.Sprintf("bsm: %+v", bsm))
}
