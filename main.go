package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "shubhammishra123"

	//to hash the password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		fmt.Println(err)
	}

	//note hashed password returns bytes so it can be converted into string like string(hashedPassword)

	fmt.Println("hashed password :: ", hashedPassword)

	//to compare weather hashedPassword and Original Password are same or not

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	if err == nil {
		fmt.Println("Samed hashed password")
	}

	//to get what Cost used to create the above hashedPassword we know its 14

	cost, er := bcrypt.Cost(hashedPassword)

	if er != nil {
		fmt.Println("invalid hashed password")
	}

	fmt.Println("Cost of hashedPassword:: ", cost)

	//there are 3 variable defined
	//MinCost=4
	//MaxCost=31
	//DefaultCost =10

	fmt.Println(bcrypt.MinCost)
	fmt.Println(bcrypt.MaxCost)
	fmt.Println(bcrypt.DefaultCost)

}
