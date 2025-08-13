package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func init() {
	ProductList = []Product{
		{
			ID:          1,
			Title:       "Apple MacBook Air",
			Description: "Laptop",
			Price:       999.99,
			ImgUrl:      "https://picsum.photos/200/300",
		},
		{
			ID:          2,
			Title:       "Dell XPS 13",
			Description: "Laptop",
			Price:       1299.99,
			ImgUrl:      "https://picsum.photos/200/301",
		},
		{
			ID:          3,
			Title:       "Canon EOS 80D",
			Description: "Camera",
			Price:       1200.00,
			ImgUrl:      "https://picsum.photos/200/302",
		},
	}
}
