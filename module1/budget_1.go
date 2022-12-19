package module1

// Budget stores budget information

func main() {
	type Budget struct {
		Max  float32
		Item []Item
	}
	type Item struct {
		Description string
		Price       float32
	}
}

// Item stores item information
