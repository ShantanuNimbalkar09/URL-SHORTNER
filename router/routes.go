package router

import (
	"app/constant"
	"app/controller"
	"net/http"
)

var urlShortner = Routes{
	Route{"Url Shortner Service", http.MethodPost, constant.UrlShortnerPath, controller.ShortTheUrl},
	Route{"Redirect to url", http.MethodGet, constant.RedirectUrlPath, controller.RedirectURL},
}
