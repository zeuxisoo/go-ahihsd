package utils

import "fmt"

func ShowInfo(name string, value interface{}) {
	fmt.Printf("%-30s: ", name)

	switch v := value.(type) {
	case uint8, uint16, uint32:
		if name == "quality flag 3" {
			fmt.Printf("%c", value)
		}else{
			fmt.Printf("%d", v)
		}
	case float32, float64:
		fmt.Printf("%f", v)
	default:
		fmt.Printf("%s", v)
	}

	fmt.Printf("\n")
}
