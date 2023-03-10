package handler

import (
	"github.com/julienschmidt/httprouter"
	"go.mod/pkg/repository"
	"html/template"
	"net/http"
)

type Handler interface {
	Register(router *httprouter.Router)
}

type handler struct {
	serv repository.Repository
}

func NewHandler(serv repository.Repository) Handler {
	return &handler{serv: serv}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/", h.Interface)
	router.POST("/order", h.PutData)

}
func (h *handler) Interface(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	temp, _ := template.ParseFiles("html_file/http_html.html")
	a := "UID"
	temp.Execute(writer, a)
}

func (h *handler) PutData(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	temp, _ := template.ParseFiles("html_file/order.html")
	uid := request.FormValue("uid")
	res, err := h.serv.GetCache().Get(uid)
	//put, _ := json.MarshalIndent(res, "", "\t")
	if err == nil {
		temp.Execute(writer, res)
	} else {
		writer.Write([]byte("Заказ с данным ID отсутствует"))
	}
}
