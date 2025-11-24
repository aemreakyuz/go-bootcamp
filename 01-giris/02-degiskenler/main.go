package main

import (
	"fmt"
	"reflect"
)

var sayi = 13

const aciklama = "Ornek Go uygulaması"

var a = "this is stored in the variable a" //global scope
var b, c = "stored in b", "stored in c"    //global scope
var d string

// global scope
var (
	str2 = ""
	str3 = ""
)

//fonksiyon alanı dışında, her yapı bir anahtar kelime (var, func vs.) ile başlar
//veri tipi vermeden := olarak shortcut, ama sadece function ve metodların içinde, global scope ta olmaz
//y := 7 !!hata

func main() {
	//x := 5
	/*
		using var
	*/
	//var message string
	//message = "Merhaba Mars!"
	//fmt.Println(message)

	/*
		var var1 int //default 0
		var var2 int = 3
		var var3 float64
		var3 = 5
	*/

	// var fullName string = "Emre Akyüz"
	// fullName += " / Coder"
	// fmt.Println(fullName)

	// fmt.Printf("%v \n", a)
	// fmt.Printf("%v \n", b)
	// fmt.Printf("%v \n", c)
	// fmt.Printf("%v \n", d)

	/*
		int, uint, byte, rune
	*/

	//var var3 uint8 = 255 // 0-255 arası(dahil) veri alır
	//var var4 byte = 255  // uint8 ile aynı. Yani bir alias dır

	//uint16, uint32, uint64 şeklinde kullanımlar mümkün

	//u ile başlayan veri tipleri bir çok dilden bildiğimiz üzere unsigned(sinyalsiz) dir.
	//yani 0'dan küçük-eksi değer alamaz.
	//Eksi değer alabilen veri tipleri ise aşağıdaki gibidir
	/*
		int8, int16, int32, int64 tipleri signed dır. eksi değer alabilirler
		float için ise; float32, float64
	*/

	//ek olarak rune tipi vardır. bu da int32 için bir alias dır
	//var runeType rune = 445

	//var keyword u ve veri tipi kullanmadan değişken tanımlamak mümkündür
	// bunun için := kullanılmalı
	//bu kullanımda veri tipi belirtilmesine gerek yok, otomatik olarak veriye göre kendi tipine atayacaktır
	//var6 := 4
	//var var7 int32:= // bu kullanım hatalıdır, var var ise := kullanılmaz

	// var initialize
	//var message = "Hello World!"
	//var a, b, c int = 1, 2, 3

	// infer type
	//var message = "Hello World!"
	//var a, b, c int = 1, 2, 3
	//fmt.Println(message, a, b, c)

	//mixed
	//var message = "Hello World!"
	//var a, b, c = 1, false, 3

	//var zero-value
	// var a int
	// var b string
	// var c float64
	// var d bool

	//örtülü (implicit) ya da belirgin (explicit) ilklendirme(initialization)
	//var c int
	//var k, o string = "foo", "bar" //multiple assignment

	//var p = 42             //int
	//var s, b = "foo", true //string, bool

	/*
		Kısaltılmış değişken bildirimleri(sadece fonksiyonların içinde )
	*/

	u := 42
	//v, n := "foo", true
	//you can only do this inside a func
	//message := "hello"
	a, b, c := 1, false, 3
	d := 4
	e := true

	//a := 10
	//b := "golang"
	//c := 4.17
	//d := true
	//e := "merhaba"
	f := `golang` //work like double-quotes
	g := `M`

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	var var8 float64 = 5.4
	fmt.Println(var8)
	fmt.Println(reflect.TypeOf(var8)) //reflect import edilerek TypeOf metoduna eriştik ve verinin tipini elde ettil

	var var9 int = int(var8)          // int e dönüşümü tamamladık
	fmt.Println(reflect.TypeOf(var9)) //döünüşüm sonrası verinin tipi int
	fmt.Println(var9)

	// -------
	var myFloat32 float32 = 42.
	myComplex := complex(2, 3)

	println(myFloat32)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))

	//adresini almak mümkün
	fmt.Println(&u) // u'nun adresini yazdır (bellekte gösterdiği yer)

}
