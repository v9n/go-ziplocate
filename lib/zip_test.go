package zip

import (
	"testing"
	"fmt"
)

func TestPoint(t *testing.T) {
	fmt.Println("Start to test")
	var p Point
	//p = &Point{12.33, 13.44478}
	//p = new(Point)
	p.X = 1.00
	p.Y = 2.12
	bytes,_ := p.GobEncode()
	fmt.Printf("%v", bytes)
	t.Error("Fail ", 1)
	//if p.X != 12.33 {
		//t.Fatal("Exepect X=12.33, got ", p.X)
	//}
	//if p.Y != 13.44478 {
		//t.Fatal("Exepect Y=13.44478, got ", p.Y)
	//}
}

//func TestEncode(t *testing.T) {
	//var p *Point
	//p = &Point{12.33, 13.44478}
	//t.Fatal(string(bytes), p.X)
//}
