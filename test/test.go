package main

import (
	"fmt"
	"time"
)

type Pruebita struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Factor int `json:"factor"`
}

func worker(c chan []Pruebita, arreglito []Pruebita) {
	for i := range arreglito {
		arreglito[i].X = 69
		arreglito[i].Y = 420
	}
}

func main() {
	chan1 := make(chan []Pruebita)
	chan2 := make(chan []Pruebita)
	chan3 := make(chan []Pruebita)
	chan4 := make(chan []Pruebita)
	fmt.Println("BITIII")
	arreglito := []Pruebita{}
	for i := 0; i < 12; i++ {
		objetito := Pruebita{
			X:      i + 1,
			Y:      10 - i,
			Factor: 0,
		}
		arreglito = append(arreglito, objetito)
	}
	index1 := len(arreglito) / 4
	index2 := index1 + index1
	index3 := index2 + index1
	a1 := arreglito[:index1]
	a2 := arreglito[index1:index2]
	a3 := arreglito[index2:index3]
	a4 := arreglito[index3:]

	go worker(chan1, a1)
	go worker(chan2, a2)
	go worker(chan3, a3)
	go worker(chan4, a4)
	time.Sleep(time.Millisecond * 500)
	final := []Pruebita{}
	final = append(final, a1...)
	final = append(final, a2...)
	final = append(final, a3...)
	final = append(final, a4...)
	fmt.Println("PRUEBITA")
	fmt.Println(final)
}
