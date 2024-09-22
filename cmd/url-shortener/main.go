package main

import (
	"Shortener_Tuzov/internal/config"
	"fmt"
)

func main() {
	//init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//todo init logger: slog
	//todo init storage: sqlite
	//todo init router: chi
	//todo run server

}
