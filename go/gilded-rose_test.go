package main

import (
	"reflect"
	"testing"
)

func Test_Run(t *testing.T) {
	run(2)
}

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
			name:   "Conjured items degrade in quality twice as fast as normal items #1",
			rounds: 2,
			input: []*Item{
				&Item{"Conjured Mana Cake", 3, 6},
			},
			expected: []*Item{
				&Item{"Conjured Mana Cake", 1, 2},
			},
		},
		{
			name:   "Conjured items degrade in quality twice as fast as normal items #2",
			rounds: 20,
			input: []*Item{
				&Item{"Conjured Mana Cake", 3, 6},
			},
			expected: []*Item{
				&Item{"Conjured Mana Cake", -17, 0},
			},
		},
		{
			name:   "Once the sell by date has passed, quality degrades twice as fast #1",
			rounds: 15,
			input: []*Item{
				&Item{"+5 Dexterity Vest", 10, 20},
			},
			expected: []*Item{
				&Item{"+5 Dexterity Vest", -5, 0},
			},
		},
		{
			name:   "Once the sell by date has passed, quality degrades twice as fast #2",
			rounds: 5,
			input: []*Item{
				&Item{"Conjured Mana Cake", 3, 12},
			},
			expected: []*Item{
				&Item{"Conjured Mana Cake", -2, 0},
			},
		},
		{
			name:   "The quality of an item is never negative #1",
			rounds: 10,
			input: []*Item{
				&Item{"Conjured Mana Cake", 3, 6},
			},
			expected: []*Item{
				&Item{"Conjured Mana Cake", -7, 0},
			},
		},
		{
			name:   "The quality of an item is never negative #2",
			rounds: 10,
			input: []*Item{
				&Item{"Elixir of the Mongoose", 5, 7},
			},
			expected: []*Item{
				&Item{"Elixir of the Mongoose", -5, 0},
			},
		},
		{
			name:   "Aged Brie actually increases in quality the older it gets",
			rounds: 10,
			input: []*Item{
				&Item{"Aged Brie", 2, 0},
			},
			expected: []*Item{
				&Item{"Aged Brie", -8, 18},
			},
		},
		{
			name:   "The quality of an item is never more than 50 #1",
			rounds: 60,
			input: []*Item{
				&Item{"Aged Brie", 2, 0},
			},
			expected: []*Item{
				&Item{"Aged Brie", -58, 50},
			},
		},
		{
			name:   "Sulfuras, being a legendary item, never has to be sold or decreases in quality",
			rounds: 5,
			input: []*Item{
				&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
			},
			expected: []*Item{
				&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
			},
		},
		{
			name:   "Backstage passes quality increases as its sell-in value approaches, by 2 when <= 10 days and by 3 when <= 5 days, but quality drops to 0 after the concert #1",
			rounds: 5,
			input: []*Item{
				&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			},
			expected: []*Item{
				&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
			},
		},
		{
			name:   "Backstage passes quality increases as its sell-in value approaches, by 2 when <= 10 days and by 3 when <= 5 days, but quality drops to 0 after the concert #2",
			rounds: 15,
			input: []*Item{
				&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 25},
			},
			expected: []*Item{
				&Item{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
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
