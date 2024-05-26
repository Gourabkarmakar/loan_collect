package initializer

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVeriables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Env Variable not Loaded... ")
	}
}

func RunOnPort() {
	data := os.Getenv("PORT")

	fmt.Println("Server run on : ",data)

}
