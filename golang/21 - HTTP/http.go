package main

import (
	"log"
	"net/http"
)

// HTTP - Protocolo de Cominicação
//        Base da Comunicação Web
//        Cliente (Faz Requisição) / Servidor (Processa Requisição e Envia Resposta)
// 				Request / Response
// 				Rotas (Routes)
//            URI - Identificador do Recurso
//            Método - GET(Ler), POST(Incluir), PUT(Atualizar), DELETE(Excluir)

func raiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {
	http.HandleFunc("/", raiz)
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
