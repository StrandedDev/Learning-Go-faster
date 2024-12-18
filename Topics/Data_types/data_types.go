package Main

import "fmt"

func main() {
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println(isTrue, isFalse)

	var int8Value int8 = -128
	var int16Value int16 = 32767
	var int32Value int32 = 2147483647
	var int64Value int64 = 9223372036854775807
	var intValue int = 2147483647
	var uint8Value uint8 = 255
	var uint16Value uint16 = 65535
	var uint32Value uint32 = 4294967295
	var uint64Value uint64 = 18446744073709551615

	fmt.Println(intValue, int8Value, int16Value, int32Value, int64Value, uint8Value, uint16Value, uint32Value, uint64Value)

	var float32Value float32 = 3.14
	var float64Value float64 = 3.14159265358979323846

	fmt.Println(float32Value, float64Value)

	var stringVar string = "Hello, World!"
	stringVar2 := "Another string"

	fmt.Println(stringVar, stringVar2)

	var intArray [5]int = [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(intArray, slice)

	type Person struct {
		Name string
		Age  int
	}

	var person Person = Person{Name: "John Doe", Age: 30}

	fmt.Println(person)

	var m map[string]int = make(map[string]int)
	m["one"] = 1

	type Writer interface {
		Write([]byte) error
	}

}
