package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods in Golang!!")

	abhishek := User{"abhi", "abhi@gmail.com", true, 19}
	fmt.Println(abhishek)
	fmt.Printf("abhishek details are: %+v\n", abhishek)
	fmt.Printf("Name is %v and Email is %v\n", abhishek.Name, abhishek.Email)
	
	abhishek.GetStatus();
	abhishek.NewMail();
	
	fmt.Printf("Name is %v and Email is %v\n", abhishek.Name, abhishek.Email)
}

type User struct{
	Name string 
	Email string
	Status bool
	Age int
}

// Methods
func (u User) GetStatus(){
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is:", u.Email)
}