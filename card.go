package faeria

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"strconv"
)

type Card struct {
	Number        int64
	Faction       string
	Name          string
	Type          string
	Gold          string
	ManaCost      int64
	BlueCost      int64
	GreenCost     int64
	RedCost       int64
	YellowCost    int64
	Power         int64
	Toughness     int64
	Text          string
	Codex         string
	NumberInCodex int64
	CodexID       string
	Rarity        string
}

func NewCard(attrs []string) (Card, error) {
	c := Card{}
	var err error

	c.Number, err = parseInt(attrs[0])
	if err != nil {
		return c, err
	}

	c.Faction = attrs[1]
	c.Name = attrs[2]
	c.Type = attrs[3]
	c.Gold = attrs[4]

	c.ManaCost, err = parseInt(attrs[5])
	if err != nil {
		return c, err
	}

	c.BlueCost, err = parseInt(attrs[6])
	if err != nil {
		return c, err
	}

	c.GreenCost, err = parseInt(attrs[7])
	if err != nil {
		return c, err
	}

	c.RedCost, err = parseInt(attrs[8])
	if err != nil {
		return c, err
	}

	c.YellowCost, err = parseInt(attrs[9])
	if err != nil {
		return c, err
	}

	c.Power, err = parseInt(attrs[10])
	if err != nil {
		return c, err
	}

	c.Toughness, err = parseInt(attrs[11])
	if err != nil {
		return c, err
	}

	c.Text = attrs[12]
	c.Codex = attrs[13]

	c.NumberInCodex, err = parseInt(attrs[14])
	if err != nil {
		return c, err
	}

	c.CodexID = attrs[15]
	c.Rarity = attrs[16]

	return c, nil
}

func parseInt(input string) (int64, error) {
	var i int64
	if input == "" {
		return i, nil
	}
	return strconv.ParseInt(input, 10, 64)
}

func ReadCardCSV(filename string) ([]Card, error) {
	cards := make([]Card, 0)

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return cards, err
	}

	reader := bytes.NewReader(b)
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = 17

	var cardAttrs [][]string
	cardAttrs, err = csvReader.ReadAll()
	if err != nil {
		return cards, err
	}

	for _, a := range cardAttrs {
		c, err := NewCard(a)
		if err != nil {
			return cards, err
		}
		cards = append(cards, c)
	}

	return cards, nil
}
