package main

import (
	"reflect"
	"testing"
)

func Test_UpdateQuality(t *testing.T) {
	var testData = []struct {
		name     string
		rounds   int
		input    []*Item
		expected []*Item
	}{
		{
			name:   "default data set, 1 round",
			rounds: 1,
			input: []*Item{
				&Item{"+5 Dexterity Vest", 10, 20},
				&Item{"Aged Brie", 2, 0},
				&Item{"Elixir of the Mongoose", 5, 7},
				&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
				&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			},
			expected: []*Item{
				&Item{"+5 Dexterity Vest", 9, 19},
				&Item{"Aged Brie", 1, 1},
				&Item{"Elixir of the Mongoose", 4, 6},
				&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
				&Item{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
			},
		},
		{
			name:   "default data set, 10 rounds",
			rounds: 10,
			input: []*Item{
				&Item{"+5 Dexterity Vest", 10, 20},
				&Item{"Aged Brie", 2, 0},
				&Item{"Elixir of the Mongoose", 5, 7},
				&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
				&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			},
			expected: []*Item{
				&Item{"+5 Dexterity Vest", 0, 10},
				&Item{"Aged Brie", -8, 18},
				&Item{"Elixir of the Mongoose", -5, 0},
				&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
				&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 35},
			},
		},
		{
			name:   "default data set, 20 rounds",
			rounds: 20,
			input: []*Item{
				&Item{"+5 Dexterity Vest", 10, 20},
				&Item{"Aged Brie", 2, 0},
				&Item{"Elixir of the Mongoose", 5, 7},
				&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
				&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			},
			expected: []*Item{
				&Item{"+5 Dexterity Vest", -10, 0},
				&Item{"Aged Brie", -18, 38},
				&Item{"Elixir of the Mongoose", -15, 0},
				&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
				&Item{"Backstage passes to a TAFKAL80ETC concert", -5, 0},
			},
		},
		{
			name:   "conjured mana cake, 2 rounds",
			rounds: 2,
			input: []*Item{
				&Item{"Conjured Mana Cake", 3, 6},
			},
			expected: []*Item{
				&Item{"Conjured Mana Cake", 1, 2},
			},
		},
		{
			name:   "conjured mana cake, 10 rounds",
			rounds: 10,
			input: []*Item{
				&Item{"Conjured Mana Cake", 3, 6},
			},
			expected: []*Item{
				&Item{"Conjured Mana Cake", -7, 0},
			},
		},
		{
			name:   "conjured mana cake, 20 rounds",
			rounds: 20,
			input: []*Item{
				&Item{"Conjured Mana Cake", 3, 6},
			},
			expected: []*Item{
				&Item{"Conjured Mana Cake", -17, 0},
			},
		},
	}

	for n, test := range testData {
		for i := 0; i < test.rounds; i++ {
			UpdateQuality(test.input)
		}
		if ok := reflect.DeepEqual(test.input, test.expected); !ok {
			t.Errorf("test case #%d '%s' wanted: %v, but got: %v", n+1, test.name, test.expected, test.input)
		}
	}
}
