package main

import "fmt"

func main() {
	var first_name string
	var last_name string
	var grade int
	var student_id int
	var login string
	var gpa float64

	fmt.Println("lütfen takip eden bilgileri doldurunuz : ")
	fmt.Print("Ad: ")
	fmt.Scanf("%s\n", &first_name)
	fmt.Print("Soyad: ")
	fmt.Scanf("%s\n", &last_name)
	fmt.Print("Sınıf (9-12) : ")
	fmt.Scanf("%d\n", &grade)
	fmt.Print("Ogrenci ID: ")
	fmt.Scanf("%d\n", &student_id)
	fmt.Print("Oturum aç: ")
	fmt.Scanf("%s\n", &login)
	fmt.Print("Not Ortalamas (0.0 - 4.0) : ")
	fmt.Scanf("%f\n", &gpa)

	fmt.Println("Bilgileriniz : ")
	fmt.Printf("\tOturum: %s\n", login)
	fmt.Printf("\tID: %d\n", student_id)
	fmt.Printf("\tAd: %s %s\n", first_name, last_name)
	fmt.Printf("\tGPA: %f\n", gpa)
	fmt.Printf("\tSınıf: %d\n", grade)

	//------

	var name string
	var age int
	var age_back int
	var age_plus int

	fmt.Println("Hello. What is your name?")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hi %s! How old are you? ", name)
	fmt.Scanf("%d", &age)

	age_back = age - 5
	fmt.Printf("five years ago, you were %d years old?\n", age_back)
	age_plus = age + 5
	fmt.Printf("... and five years from now, you will be %d years old\n", age_plus)
}
