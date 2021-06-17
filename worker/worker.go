package worker

import (
	"github.com/8rb/Go-API/model"
)

func CalculateFactor(c chan []model.Item, halfItems []model.Item) {
	for i := range halfItems {
		halfItems[i].Factor = halfItems[i].X * halfItems[i].Y
	}
}

func CalculateTotalFactor(c chan int, halfItems []model.Item) {
	halfFactor := 0
	for i := range halfItems {
		halfFactor += halfItems[i].Factor
	}
	c <- halfFactor
}
