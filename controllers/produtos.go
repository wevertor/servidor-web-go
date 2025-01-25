package controllers

import (
	"html/template"
	"log"
	"net/http"
	"servidor-web/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")

		precoConvertido, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		validaErro(err, "Erro na conversão do preço:")

		quantidadeConvertida, err := strconv.Atoi(r.FormValue("quantidade"))
		validaErro(err, "Erro na conversão da quantidade:")

		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.RemoveProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.BuscaUmProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		validaErro(err, "Erro ao converter id para inteiro.")

		precoConv, err := strconv.ParseFloat(preco, 64)
		validaErro(err, "Erro ao converter preço para float.")

		quantidadeConv, err := strconv.Atoi(quantidade)
		validaErro(err, "Erro ao converter quantidade para inteiro")
		log.Println(id, quantidade, nome, descricao, preco)

		models.AtualizaProduto(idConv, quantidadeConv, nome, descricao, precoConv)
	}
	http.Redirect(w, r, "/", 301)
}

func validaErro(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}
