package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//kullanıcıdan bilgi al
//en az 7 bilgi alınmalı

//kullanıcının adı, soyadı, email, kullanıcı adı, cinsiyet(bool) --> kadın yada erkek gibi göster, doğum tarihi(yaşı hesaplayarak göster)
//bilgileri tek ve düzenli bir formda göster (strings.builder)
//validations: email, kullanıcı adı (max 15 char)

//bu bilgiler çeşitli hesaplama ve işlemlere tabi olmalı
//bu bilgiler formatlanmış bir şekilde kullanıcıya gösterilmeli

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Your surname: ")
	surname, _ := reader.ReadString('\n')
	surname = strings.TrimSpace(surname)

	var username string
	for {
		fmt.Print("username(Max 15 characters): ")
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)

		if len(username) <= 15 && len(username) > 0 {
			break
		} else {
			fmt.Println("max uzunluk 15")
		}
	}

	var email string
	for {
		fmt.Print("Your email: ")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)

		if strings.Contains(email, "@") && strings.Contains(email, ".") {
			break
		} else {
			fmt.Println("please enter a valid email")
		}
	}

	var gender string
	for {
		fmt.Print("Your gender: ")
		gender, _ = reader.ReadString('\n')
		gender = strings.TrimSpace(gender)
		gender = strings.ToUpper(gender)
		if gender == "E" || gender == "ERKEK" || gender == "K" || gender == "KADIN" {
			break
		} else {
			fmt.Println("Please enter E/Erkek or K/Kadın")
		}
	}

	var dob int
	for {
		fmt.Print("Your birth year (e.g., 1995): ")
		dobStr, _ := reader.ReadString('\n')
		dobStr = strings.TrimSpace(dobStr)

		var err error
		dob, err = strconv.Atoi(dobStr)

		currentYear := time.Now().Year()

		if err != nil {
			fmt.Println("Please enter a valid year (numbers only)")
		} else if dob < 1900 || dob > currentYear {
			fmt.Printf("Birth year must be between 1900 and %d\n", currentYear)
		} else if currentYear-dob < 0 || currentYear-dob > 150 {
			fmt.Println("Please enter a realistic birth year")
		} else {
			break
		}
	}

	currentYear := time.Now().Year()
	age := currentYear - dob

	var genderText string
	if gender == "E" || gender == "ERKEK" {
		genderText = "Erkek"
	} else {
		genderText = "Kadın"
	}

	var builder strings.Builder

	builder.WriteString("\n")
	builder.WriteString("╔════════════════════════════════════════╗\n")
	builder.WriteString("║        USER INFORMATION                ║\n")
	builder.WriteString("╠════════════════════════════════════════╣\n")
	builder.WriteString(fmt.Sprintf("║ Name          : %-22s ║\n", name+" "+surname))
	builder.WriteString(fmt.Sprintf("║ Username      : %-22s ║\n", username))
	builder.WriteString(fmt.Sprintf("║ Email         : %-22s ║\n", email))
	builder.WriteString(fmt.Sprintf("║ Gender        : %-22s ║\n", genderText))
	builder.WriteString(fmt.Sprintf("║ Birth Year    : %-22d ║\n", dob))
	builder.WriteString(fmt.Sprintf("║ Age           : %-22d ║\n", age))
	builder.WriteString("╚════════════════════════════════════════╝\n")

	fmt.Println(builder.String())
}
