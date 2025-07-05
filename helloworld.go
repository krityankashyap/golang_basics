package main

import "fmt"

// func main() {
// 	fmt.Println("Hello world");
// }

func fun(){
	fmt.Println("Hello world");
}

func main(){
	fun();
	var productName string = "iphone";
	fmt.Println("product name is :", productName);
	loop_demo();

	array_demo();

	maps_demo();

	p,q:= check_even_odd(10);
	fmt.Println(p , q)
}

func loop_demo() {
	for i:= 0 ;i <5 ; i++ {
		fmt.Println("loop iteration ", i);
	}

	for i := range 3 {
		fmt.Println("range interation " , i);
	}

	for i,char:= range "Krityan" {
      fmt.Println("String range loop iteration" , i , char);
	}

	// In go we don't have while loop so we use for loop as while loop

	j:=10;
	for j>0{
		fmt.Println("while loop iteration ", j);
		j--;
	}
}

func array_demo() {
	var arr1 []int // array of empty integers

	arr1 = append(arr1, 1 ,2 ,3 ,4 ,5); // append doesn't work for fixed size array

	fmt.Println("Array 1 :" ,arr1)
}

func maps_demo() {
	var productprices = map[string]int{
		"iphone 15 pro" : 100000,
		"Macbook pro" : 200000,
		"ipad air" : 45000,
	}
	fmt.Println("Product prices: ", productprices)

	for product, price:= range productprices{
		fmt.Println("Product: ", product ,",", "price: ", price)
	}
	delete(productprices , "ipad air") // Remove an item from map
	fmt.Println("After deletion : " , productprices)

	emptyMap := make(map[string]int)

	fmt.Println("empty map: " , emptyMap)

	emptyMap["key1"]= 100
	emptyMap["key2"]= 200
	emptyMap["key3"]= 300

	fmt.Println("populated emptyMap : " , emptyMap)
}

func check_even_odd(x int) (string , int){
	if(x%2>0) {
		return "odd" ,x
	} else {
		return "even" ,x
	}
}