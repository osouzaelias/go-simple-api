package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/api/user", GetUserHandler)
	server.HandleFunc("/health", HealthCheckHandler)

	// Rota para buscar usuário por ID
	server.HandleFunc("/api/user/", GetUserByIDHandler)

	log.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":8080", server)

	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
