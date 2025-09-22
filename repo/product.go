package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Store()
	Get()
	List()
	Update()
	Delete()
}

type productRepo struct {
	// db *sql.DB
	productList []Product
}

func Store(p *Product) *Product {
	p.ID = len(productRepo.productList) + 1
	productList = append(productList, *p)

	return p
}

func GetProducts() []Product {
	return productList
}

func GetProductByID(id int) *Product {
	for _, product := range productList {
		if product.ID == id {
			return &product
		}
	}

	return nil
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateProducts(repo)
	return repo
}

func (r *productRepo) Store() {

}
func (r *productRepo) Get() {

}
func (r *productRepo) List() {

}
func (r *productRepo) Update() {

}
func (r *productRepo) Delete() {

}

func generateProducts(r *productRepo) {
	r.productList = []Product{
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
