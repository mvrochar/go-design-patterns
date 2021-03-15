package main

import (
	"github.com/mvrochar/go-design-patterns/abstract_factory"
	"fmt"
	"reflect"
)

func SetupConstructor(env string) (abstract_factory.Database, abstract_factory.FileSystem){
	fs := abstract_factory.AbstractFactory("filesystem")
	db := abstract_factory.AbstractFactory("database")

	return db(env).(abstract_factory.Database),
	fs(env).(abstract_factory.FileSystem)
}

func main() {
	env1 := "production"
	env2 := "development"

	db1, fs1 := SetupConstructor(env1)
	db2, fs2 := SetupConstructor(env2)

	db1.PutData("test", "this is mongodb")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test", "this is sqlite")
	fmt.Println(db2.GetData("test"))

	fs1.CreateFile("./test.txt")
	fmt.Println(fs1.FindFile("./test.txt"))

	fs2.CreateFile("./test2.txt")
	fmt.Println(fs2.FindFile("./test2.txt"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())

}
