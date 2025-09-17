package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepository interface {
	Create()
	Get()
	List()
	Update()
	Delete()
}

type productRepo struct {
	// db *sql.DB
	productList []Product
}
