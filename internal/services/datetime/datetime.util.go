package dateti

import (
	t  "dshusdock/tw_prac1/internal/services/database/tables"
	f  "fmt"
	ti "time"
)

func Prac1() {
    dateStr := "2002-03-01"
    layout := "2006-01-02"
    parsedDate, err := ti.Parse(layout, dateStr)
    if err != nil {
        f.Println("Error:", err)
        return
    }
    f.Println(parsedDate)
}

func Prac2() {
    date := ti.Now()
    f.Println("The ??? layout:", date.Format("???"))
    f.Println("The ANSIC layout:", date.Format(ti.ANSIC))
    f.Println("The UnixDate layout:", date.Format(ti.UnixDate))
    f.Println("The RubyDate layout:", date.Format(ti.RubyDate))
    f.Println("The RFC822 layout:", date.Format(ti.RFC822))
    f.Println("The RFC822Z layout:", date.Format(ti.RFC822Z))
    f.Println("The RFC850 layout:", date.Format(ti.RFC850))
    f.Println("The RFC1123 layout:", date.Format(ti.RFC1123))
    f.Println("The RFC1123Z layout:", date.Format(ti.RFC1123Z))
    f.Println("The RFC3339 layout:", date.Format(ti.RFC3339))
    f.Println("The RFC3339Nano layout:", date.Format(ti.RFC3339Nano))
}

func Prac3() {
    t.SetAppState()
}
