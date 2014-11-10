package main

import "github.com/go-martini/martini"
import "github.com/thecatspaw/shopngo/shopngo/service"
import "strconv"
import "net/http"

import "github.com/martini-contrib/sessions"

import "fmt"

type Mycontext struct {
	*http.Request
	//	session sessions.Session
}

func main() {
	m := martini.Classic()

	m.Use(sessions.Sessions("shopngo", service.Store))

	var prdService = &service.ProductService{}
	prdService.Init()
	m.Map(prdService)

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

		return 200, "Displaying product " + params["number"] + "\n" + prdService.GetProduct(i) + "\n" + fmt.Sprint(session.Get("test"))

	})

	m.Run()
}
