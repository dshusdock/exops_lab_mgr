package datetime

import (
	"fmt"
	"time"
)

func Prac1() {
    dateStr := "2002-03-01"
    layout := "2006-01-02"
    parsedDate, err := time.Parse(layout, dateStr)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(parsedDate)
}
