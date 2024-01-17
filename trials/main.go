package main

import "fmt"

type Codot struct {
	ID        string `json:"id"`
	CodotType string `json:"type"`
	Defeated  bool   `json:"defeated"`
}

//constructor

func NewCodot(id string, codotType string, defeated bool) *Codot {
	return &Codot{
		ID: id, CodotType: codotType, Defeated: defeated,
	}
}

//Setter Function 

func (c *Codot) SetCodotType (codotType string) {
	c.CodotType = codotType
}

func main() {
	codot1:= NewCodot("df4a5f4af5a", "Punisher", false)
	fmt.Println(*codot1)
	codot1.SetCodotType("Liberator")
	fmt.Println(*codot1)

}
