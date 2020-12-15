use std::fmt::{self, Display};

#[derive(Debug)]
pub struct Item {
    pub name: String,
    pub sell_in: i32,
    pub quality: i32,
}

impl Item {
    pub fn new(name: impl Into<String>, sell_in: i32, quality: i32) -> Item {
        Item {
            name: name.into(),
            sell_in,
            quality,
        }
    }
}

impl Display for Item {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}, {}, {}", self.name, self.sell_in, self.quality)
    }
}

#[derive(Debug)]
pub struct GildedRose {
    pub items: Vec<Item>,
}

impl PartialEq for GildedRose {
    fn eq(&self, other: &Self) -> bool {
        for i in 0..self.items.len() {
            if self.items[i].name != other.items[i].name {
                return false
            } else if self.items[i].sell_in != other.items[i].sell_in {
                return false
            } else if self.items[i].quality != other.items[i].quality {
                return false
            }
        }
        true
    }
}

impl GildedRose {
    pub fn new(items: Vec<Item>) -> GildedRose {
        GildedRose { items }
    }

    pub fn update_quality(&mut self) {
        for item in &mut self.items {
            if item.name != "Aged Brie" && item.name != "Backstage passes to a TAFKAL80ETC concert"
            {
                if item.quality > 0 {
                    if item.name != "Sulfuras, Hand of Ragnaros" {
                        item.quality = item.quality - 1;
                    }
                    if item.name == "Conjured Mana Cake" {
                        item.quality = item.quality - 1;
                    }
                }
            } else {
                if item.quality < 50 {
                    item.quality = item.quality + 1;

                    if item.name == "Backstage passes to a TAFKAL80ETC concert" {
                        if item.sell_in < 11 {
                            if item.quality < 50 {
                                item.quality = item.quality + 1;
                            }
                        }

                        if item.sell_in < 6 {
                            if item.quality < 50 {
                                item.quality = item.quality + 1;
                            }
                        }
                    }
                }
            }

            if item.name != "Sulfuras, Hand of Ragnaros" {
                item.sell_in = item.sell_in - 1;
            }

            if item.sell_in < 0 {
                if item.name != "Aged Brie" {
                    if item.name != "Backstage passes to a TAFKAL80ETC concert" {
                        if item.quality > 0 {
                            if item.name != "Sulfuras, Hand of Ragnaros" {
                                item.quality = item.quality - 1;
                            }
                        }
                    } else {
                        item.quality = item.quality - item.quality;
                    }
                } else {
                    if item.quality < 50 {
                        item.quality = item.quality + 1;
                    }
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::{GildedRose, Item};

    pub struct TestData {
        pub name: String,
        pub rounds: i32,
        pub input: Vec<Item>,
        pub expected: Vec<Item>,
    }

    impl TestData {
        pub fn new(name: impl Into<String>, rounds: i32, input: Vec<Item>, expected: Vec<Item>) -> TestData {
            TestData {
                name: name.into(),
                rounds,
                input,
                expected,
            }
        }
    }

    fn run_rounds(input: &mut GildedRose, rounds: i32) {
        for _ in 0..rounds {
            input.update_quality();
        }
    }

    fn run_tests(tests: Vec<TestData>) {
        for test in tests {
            println!("{:#?}", test.name);
            let mut rose = GildedRose::new(test.input);
            run_rounds(&mut rose, test.rounds);
            assert_eq!(GildedRose::new(test.expected), rose);
        }
    }

    #[test]
    pub fn main() {
        let test_cases = vec![
            TestData::new(
                "default data set, 1 round",
                1,
                vec![
                    Item::new("+5 Dexterity Vest", 10, 20),
                    Item::new("Aged Brie", 2, 0),
                    Item::new("Elixir of the Mongoose", 5, 7),
                    Item::new("Sulfuras, Hand of Ragnaros", 0, 80),
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 15, 20),
                ],
                vec![
                    Item::new("+5 Dexterity Vest", 9, 19),
                    Item::new("Aged Brie", 1, 1),
                    Item::new("Elixir of the Mongoose", 4, 6),
                    Item::new("Sulfuras, Hand of Ragnaros", 0, 80),
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 14, 21),
                ],
            ),
            TestData::new(
                "default data set, 10 rounds",
                10,
                vec![
                    Item::new("+5 Dexterity Vest", 10, 20),
                    Item::new("Aged Brie", 2, 0),
                    Item::new("Elixir of the Mongoose", 5, 7),
                    Item::new("Sulfuras, Hand of Ragnaros", 0, 80),
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 15, 20),
                ],
                vec![
                    Item::new("+5 Dexterity Vest", 0, 10),
                    Item::new("Aged Brie", -8, 18),
                    Item::new("Elixir of the Mongoose", -5, 0),
                    Item::new("Sulfuras, Hand of Ragnaros", 0, 80),
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 5, 35),
                ],
            ),
            TestData::new(
                "default data set, 20 rounds",
                20,
                vec![
                    Item::new("+5 Dexterity Vest", 10, 20),
                    Item::new("Aged Brie", 2, 0),
                    Item::new("Elixir of the Mongoose", 5, 7),
                    Item::new("Sulfuras, Hand of Ragnaros", 0, 80),
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 15, 20),
                ],
                vec![
                    Item::new("+5 Dexterity Vest", -10, 0),
                    Item::new("Aged Brie", -18, 38),
                    Item::new("Elixir of the Mongoose", -15, 0),
                    Item::new("Sulfuras, Hand of Ragnaros", 0, 80),
                    Item::new("Backstage passes to a TAFKAL80ETC concert", -5, 0),
                ],
            ),
            TestData::new(
                "Conjured items degrade in quality twice as fast as normal items #1",
                2,
                vec![
                    Item::new("Conjured Mana Cake", 3, 6),
                ],
                vec![
                    Item::new("Conjured Mana Cake", 1, 2),
                ],
            ),
            TestData::new(
                "Conjured items degrade in quality twice as fast as normal items #2",
                20,
                vec![
                    Item::new("Conjured Mana Cake", 3, 6),
                ],
                vec![
                    Item::new("Conjured Mana Cake", -17, 0),
                ],
            ),
            TestData::new(
                "Once the sell by date has passed, quality degrades twice as fast #1",
                15,
                vec![
                    Item::new("+5 Dexterity Vest", 10, 20),
                ],
                vec![
                    Item::new("+5 Dexterity Vest", -5, 0),
                ],
            ),
            TestData::new(
                "Once the sell by date has passed, quality degrades twice as fast #2",
                5,
                vec![
                    Item::new("Conjured Mana Cake", 3, 12),
                ],
                vec![
                    Item::new("Conjured Mana Cake", -2, 0),
                ],
            ),
            TestData::new(
                "The quality of an item is never negative #1",
                10,
                vec![
                    Item::new("Conjured Mana Cake", 3, 6),
                ],
                vec![
                    Item::new("Conjured Mana Cake", -7, 0),
                ],
            ),
            TestData::new(
                "The quality of an item is never negative #1",
                10,
                vec![
                    Item::new("Elixir of the Mongoose", 5, 7),
                ],
                vec![
                    Item::new("Elixir of the Mongoose", -5, 0),
                ],
            ),
            TestData::new(
                "Aged Brie actually increases in quality the older it gets",
                10,
                vec![
                    Item::new("Aged Brie", 2, 0),
                ],
                vec![
                    Item::new("Aged Brie", -8, 18),
                ],
            ),
            TestData::new(
                "The quality of an item is never more than 50 #1",
                60,
                vec![
                    Item::new("Aged Brie", 2, 0),
                ],
                vec![
                    Item::new("Aged Brie", -58, 50),
                ],
            ),
            TestData::new(
                "Sulfuras, being a legendary item, never has to be sold or decreases in quality",
                5,
                vec![
                    Item::new("Sulfuras, Hand of Ragnaros", 0, 80),
                ],
                vec![
                    Item::new("Sulfuras, Hand of Ragnaros", 0, 80),
                ],
            ),
            TestData::new(
                "Backstage passes quality increases as its sell-in value approaches, by 2 when <= 10 days and by 3 when <= 5 days, but quality drops to 0 after the concert #1",
                5,
                vec![
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 15, 20),
                ],
                vec![
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 10, 25),
                ],
            ),
            TestData::new(
                "Backstage passes quality increases as its sell-in value approaches, by 2 when <= 10 days and by 3 when <= 5 days, but quality drops to 0 after the concert #2",
                15,
                vec![
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 15, 20),
                ],
                vec![
                    Item::new("Backstage passes to a TAFKAL80ETC concert", 0, 50),
                ],
            ),
        ];

        run_tests(test_cases)
    }
}
