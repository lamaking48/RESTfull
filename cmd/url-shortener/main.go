package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/lamaking48/RESTfull.git/internal/config"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)
}
func init() {
	godotenv.Load() // Загружает переменные из .env
}
