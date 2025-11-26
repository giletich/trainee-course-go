package printcakes

import (
	"fmt"
	"task-8/cake"
)

func PrintCakes(in chan cake.Cake) {
	count := 0
	for val := range in {
		fmt.Printf("nорт %d: испёк %d, упаковал %d\n", val.CakeNumber, val.BakedBy, val.PackedBy)
		count++
	}
	fmt.Printf("обработано тортов: %d\n", count)
}