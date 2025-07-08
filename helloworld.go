package main

import "fmt"

type Product struct {
	name string
	price int
	ratings float32
}

// func main() {
// 	fmt.Println("Hello world");
// }

// func fun(){
// 	fmt.Println("Hello world");
// }

 
// func newProduct(name string , price int , ratings float32) Product {
// 	p := Product{
// 		name:    name,
// 		price:   price,
// 		ratings: ratings,
// 	}
// 	return p // unlike other lang this will return a copy of the product
// }

// func fun(p Product){
// 	// this p is a copy of the product passed to this function
// 	p.price = 56000
// }

// Now instead of passing copy we can use pointers and address to avoid 

func newProduct(name string , price int , ratings float32) *Product {
	p := Product{
		name:    name,
		price:   price,
		ratings: ratings,
	}
	return &p // unlike other lang this will return a copy of the product
}


func fun(p *Product){
	p.price = 56000
	}

	// define a member function of struct
	func (p *Product) display() {
		fmt.Println(p.name)
		fmt.Println(p.price)
		fmt.Println(p.ratings)
	}


func main(){

	new_product_pointer := newProduct("iphone 15 pro" , 125000 , 4.6)

	 fmt.Println("product price is: ", new_product_pointer.price)
	 new_product_pointer.display()

	 fun(new_product_pointer)

	 fmt.Println("product price after function call is :" , new_product_pointer.price)
	 new_product_pointer.display()

//	 fun(p)

	 // fmt.Println("product price after function call: ", p.price) // the price will not change even after function call becoz go will pass copy of product
	// fun();
	// var productName string = "iphone";
	// fmt.Println("product name is :", productName);
	// loop_demo();

	// array_demo();

	// maps_demo();

	// p,q := check_even_odd(10);
	// fmt.Println(p,q)

	// pointer_demo()
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

func check_even_odd(x int) (string, int) {
	if x%2 > 0 {
		return "odd", x
	} else {
		return "even", x
	}
}



func pointer_demo(){
	i := 10
	var ptr *int = &i // *int-> it shows the datatype of pointer that can store the address of int type , &-> used to fetch the address of i

	ptr1 := &i // short hand

	fmt.Println("value of i: ",i , "pointer of i :" , ptr , "pointer1 of i :" ,ptr1)
	fmt.Println("value present at pointer ", *ptr) // access the value present in the memory address of the pointer
}