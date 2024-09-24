package main

import "fmt"

// initialize the product struct
type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"reading", "gaming", "coding"}
	fmt.Println("Hobbies array", hobbies)

	//output first element of the hobbies array
	fmt.Println("First element of the hobbies array:", hobbies[0])
	fmt.Println("Second and third element of the hobbies array:", hobbies[1:])

	// slice based on the first element that contains the first and second elements
	slice12 := hobbies[0:2]
	fmt.Println(slice12)
	fmt.Println(hobbies[:2])

	slice23 := slice12[1:3]
	fmt.Println(slice23)

	courseGoals := []string{"read", "honors", "project", "graduate"}
	fmt.Println(courseGoals) // output the dynamic array

	courseGoals[1] = "pass exams"
	fmt.Println(courseGoals) // output the dynamic array

	courseGoals = append(courseGoals, "find job")
	fmt.Println(courseGoals) // output the dynamic array

	// initialize the array of structs
	products := []Product{
		{
			title: "learn go",
			id:    1,
			price: 39.99,
		},
		{
			title: "learn rust",
			id:    2,
			price: 50.99,
		},
	}

	fmt.Println(products)

	// add product to the products array
	newProduct := Product{
		title: "learn solidity",
		id:    1,
		price: 20.19,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}
