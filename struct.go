package main

import "fmt"

type Developer struct {
	Number     int
	Actorname  string
	Companions []string
}

func main() {
	Developer := Developer{
		Number:    1,
		Actorname: "Zia-ul-Haq",
		Companions: []string{
			"Abdullah ", // 	"Abdullah ",
			"Hiyat ",    // 	"Hiyat ",
			"Hussain ",  // 	"Hussain ",
		},
	}
	fmt.Println(Developer)                 //print whole struct
	fmt.Println(Developer.Companions[1:3]) //print specific using slice

}

//	teacher := struct{ name string }{name: "Zeeshan Ashfaq"}
//	fmt.Println(teacher) // not used outside or normally
