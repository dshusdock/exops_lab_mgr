package logging

import (
	"fmt"
)


func Log(format string, a ...any) {
	
	fmt.Printf(format, a...)
}
