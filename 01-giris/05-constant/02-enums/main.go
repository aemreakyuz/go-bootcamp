package main

import "fmt"

type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	GOOGLE    Brand = "Google"
	MICROSOFT Brand = "Microsoft"
	SAFEBOX   Brand = "Safebox"
)

func PrintBrand(brand Brand) {
	fmt.Println(brand)
}

func main() {
	PrintBrand(GOOGLE)
	PrintBrand(SAFEBOX)
}
