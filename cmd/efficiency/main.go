package main

import (
	"fmt"
	"log"

	"github.com/willroberts/faeria"
)

func main() {
	cards, err := faeria.ReadCardCSV("cards.csv")
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range cards {
		if c.Type == "creature" {
			e := determineEfficiency(c)
			fmt.Println(e, c.Name)
		}
	}
}

func determineEfficiency(c faeria.Card) float64 {
	if c.ManaCost == 0 {
		return float64(0)
	}
	efficiency := float64(c.Power+c.Toughness) / float64(c.ManaCost)
	return efficiency
}
