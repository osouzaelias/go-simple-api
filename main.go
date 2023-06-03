package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/api/user", getUserHandler)
	server.HandleFunc("/health", healthCheckHandler)

	// Rota para buscar usuário por ID
	server.HandleFunc("/api/user/", getUserByIDHandler)

	log.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":8080", server)

	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
