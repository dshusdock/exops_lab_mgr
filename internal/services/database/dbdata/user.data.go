package dbdata

import (
	con "dshusdock/tw_prac1/internal/constants"
	d "dshusdock/tw_prac1/internal/services/database"
	"fmt"
	"reflect"
)

type User struct {
	Id 	 int
	CreateTime string
	FirstName string
	LastName string
	Email string
	UserName string
	Password string
	Status string
}

type userView1 struct {
	Username string
}

type userview2 struct {
	Username string
	Password string
}

var USER_VIEWS = make (map[string]viewMap)

func init() {
	USER_VIEWS["VIEW_ALL"] = viewMap{"select * from User", reflect.TypeOf(User{})}
	USER_VIEWS["VIEW_1"] = viewMap{"select * from User", reflect.TypeOf(User{})}
	USER_VIEWS["VIEW_2"] = viewMap{`select username, password from User where username= `, reflect.TypeOf(userview2{})}
}

func GetUserNames() ([]con.RowData, error){
	rslt, err := d.ReadDBwithType[userView1](USER_VIEWS["VIEW_1"].View)	
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return rslt, nil
}

func GetUserInfo(name string) ([]con.RowData, error){
	s :=fmt.Sprintf(USER_VIEWS["VIEW_2"].View + "\"%s\"", name)
	fmt.Println("SQL: ", s)
	rslt, err := d.ReadDBwithType[userview2](s)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return rslt, nil
}

