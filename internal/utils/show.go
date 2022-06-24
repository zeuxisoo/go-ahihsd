package utils

import (
	"fmt"
)

func ShowInfo(name string, value interface{}, format ...interface{}) {
	if len(name) > 30 {
		fmt.Printf("%-30s  \n%-30s: ", name, " ")
	}else{
		fmt.Printf("%-30s: ", name)
	}

	switch v := value.(type) {
	case uint8, uint16, uint32:
		if len(format) > 0 {
			fmt.Printf(format[0].(string), v)	// custom format
		}else{
			fmt.Printf("%d", v)
		}
	case float32, float64:
		if len(format) > 0 {
			fmt.Printf(format[0].(string), v)	// custom format
		}else{
			fmt.Printf("%f", v)
		}
	default:
		fmt.Printf("%s", v)
	}

	fmt.Printf("\n")
}
