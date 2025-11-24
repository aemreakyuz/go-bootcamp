package main

import "fmt"

func main() {
	//enum gibi kullanmamız için bir değer atamamızı sağlıyo, index atıyor, 0dan başlayıp sıralı gidiyor
	//ilkini yazdıktan sonra diğerlerini yazmaya gerek yok
	const (
		A1 = iota // 0 : start at 0
		B1 = iota // 1 : increment by 1
		C1 = iota // 2 : increment by 1
	)
	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota // 0
		B2        // increment by 1
		C2        // increment by 1
	)

	const (
		A3 = iota + 1 // 1: start at 0 + 1
		B3            // 2 : increment by 1
		C3            // 3 : increment by 1
	)
}
