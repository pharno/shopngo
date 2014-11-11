package service

type ProductService struct {
	first  string
	second string
}

func (this *ProductService) Init() {
	this.first = "firstProduct"
	this.second = "secondProcuct"
}

func (this *ProductService) GetProduct(id int) string {
	if id == 1 {
		return this.first
	} else if id == 2 {
		return this.second
	} else {
		return "not found"
	}
}

func getProductService() *ProductService {
	var Productservice = &ProductService{}
	Productservice.Init()
	return Productservice
}

var Productservice = getProductService()
