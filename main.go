package main

import "github.com/go-martini/martini"
import "github.com/thecatspaw/shopngo/shopngo"
import "strconv"

func main() {
	m := martini.Classic()

	var prdService = &shopngo.ProductService{}
	prdService.Init()
	m.Map(prdService)
	m.Get("/product/:number", func(params martini.Params, prdService *shopngo.ProductService) (int, string) {
		i, err := strconv.Atoi(params["number"])

		if err != nil {
			return 404, ""
		}

		return 200, "Displaying product " + params["number"] + "\n" + prdService.GetProduct(i)
	})
	m.Run()
}
