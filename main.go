package main

import "github.com/go-martini/martini"
import "github.com/thecatspaw/shopngo/shopngo/service"
import "strconv"

import "github.com/martini-contrib/sessions"
import "github.com/flosch/pongo2"

func main() {
	m := martini.Classic()

	m.Use(sessions.Sessions("shopngo", service.Store))

	m.Map(service.Productservice)

	m.Get("/product/:number", func(params martini.Params, prdService *service.ProductService, session sessions.Session) (int, string) {
		v := session.Get("test")
		if v == nil {
			session.Set("test", 1)
		} else {
			session.Set("test", session.Get("test").(int)+1)
		}

		i, err := strconv.Atoi(params["number"])

		if err != nil {
			return 404, ""
		}
		var tpl = pongo2.Must(pongo2.FromFile("templates/index.html"))

		var rendered, erro = tpl.Execute(pongo2.Context{"productname": prdService.GetProduct(i)})

		if erro != nil {
			return 500, "internal error"
		} else {
			return 200, rendered
		}
	})

	m.Run()
}
