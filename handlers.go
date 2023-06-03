package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Estrutura para representar os dados de um usuário
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Simulando um banco de dados de usuários
var users = map[int]User{
	1: {ID: 1, Username: "johndoe", Email: "johndoe@example.com"},
	2: {ID: 2, Username: "janedoe", Email: "janedoe@example.com"},
}

// Rota para buscar um usuário por ID
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Obtém o ID do usuário a partir dos parâmetros da rota
	idStr := r.URL.Path[len("/api/user/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID inválido:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Verifica se o usuário existe
	user, ok := users[id]
	if !ok {
		log.Println("Usuário não encontrado")
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	// Serializa a estrutura de usuário para JSON
	response, err := json.Marshal(user)
	if err != nil {
		log.Println("Erro ao serializar usuário para JSON:", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho de resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Escreve a resposta
	w.Write(response)
}

// Rota que retorna os dados de um usuário específico
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Simulando um usuário retornado de um banco de dados ou outra fonte de dados
	user := User{
		ID:       1,
		Username: "johndoe",
		Email:    "johndoe@example.com",
	}

	// Serializando a estrutura de usuário para JSON
	response, err := json.Marshal(user)
	if err != nil {
		log.Println("Erro ao serializar usuário para JSON:", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Definindo o cabeçalho de resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Escrevendo a resposta
	w.Write(response)
}

// Estrutura para representar o status de saúde do serviço
type HealthStatus struct {
	Status string `json:"status"`
}

// Rota para verificar o status de saúde do serviço
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o serviço está em execução corretamente
	// Você pode adicionar lógica adicional aqui, se necessário

	// Cria uma estrutura de status de saúde
	status := HealthStatus{
		Status: "UP",
	}

	// Serializa a estrutura de status para JSON
	response, err := json.Marshal(status)
	if err != nil {
		log.Println("Erro ao serializar status de saúde para JSON:", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho de resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Escreve a resposta
	w.Write(response)
}
