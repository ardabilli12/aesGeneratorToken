package main

import (
	"aesGeneratorToken/pkg"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "paper aes generator",
		Usage: "simple paper aes generator prompts for you",
		Commands: []*cli.Command{
			{
				Name:   "generate",
				Usage:  "generate a paper aes key",
				Action: pkg.GenerateToken,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "email",
						Aliases: []string{"e"},
						Usage:   "email",
					},
				},
			},
		},
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("starting...")
	time.Sleep(1 * time.Second)
	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
