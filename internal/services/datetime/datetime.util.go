package datetime

import (
	"dshusdock/tw_prac1/internal/services/database"
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

func Prac2() {
    date := time.Now()
    fmt.Println("The ??? layout:", date.Format("???"))
    fmt.Println("The ANSIC layout:", date.Format(time.ANSIC))
    fmt.Println("The UnixDate layout:", date.Format(time.UnixDate))
    fmt.Println("The RubyDate layout:", date.Format(time.RubyDate))
    fmt.Println("The RFC822 layout:", date.Format(time.RFC822))
    fmt.Println("The RFC822Z layout:", date.Format(time.RFC822Z))
    fmt.Println("The RFC850 layout:", date.Format(time.RFC850))
    fmt.Println("The RFC1123 layout:", date.Format(time.RFC1123))
    fmt.Println("The RFC1123Z layout:", date.Format(time.RFC1123Z))
    fmt.Println("The RFC3339 layout:", date.Format(time.RFC3339))
    fmt.Println("The RFC3339Nano layout:", date.Format(time.RFC3339Nano))
}

func Prac3() {
    database.SetAppState()
}
