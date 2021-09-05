package tests

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"reflect"
)

func TestFun() {
	//testStr()
	//testType()

}

func testStr() {
	strCamel := object.Camel("sdf-sdfl-Edfd-")
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
