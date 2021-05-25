package tests

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"reflect"
)

func TestFun() {
	//testStr()
	//testType()

}

func testStr() {
	strCamel := str.Camel("sdf-sdfl-Edfd-")
	fmt.Printf("%v\r\n", strCamel)
}

func testType() {

	m := object.HashMap{}
	kind := reflect.TypeOf(m).Kind()
	fmt.Printf("kind: %v\n", kind)
	if kind == reflect.Map {
		fmt.Println("same type")
	}

}
