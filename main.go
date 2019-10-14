package main

import (
	"GitSrcGo/models"
	"fmt"
)

func main() {
	fmt.Println("Hello Git!")
	//Talking about slice
	arr := []int{1, 2, 3}
	arr = append(arr, 4)
	slice := arr[:]
	arr[1] = 12
	slice[2] = 27
	fmt.Println(arr, slice)
	//using the map data type
	m := map[string]int{"foo": 42} //string is key ::  value = 42
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 50
	fmt.Println((m))

	//Delete a element in map
	delete(m, "foo")
	fmt.Println(m)
	u := models.User{
		ID:        2,
		FirstName: "Alex",
		LastName:  "May",
	}
	fmt.Println(u)

}
