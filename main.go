package main

import (
	"fmt"
	"net/http"

	"github.com/kaiquegarcia/gojooble/internal/config"
	"github.com/kaiquegarcia/gojooble/internal/views/home"
	"github.com/kaiquegarcia/gojooble/internal/views/search"
	"github.com/kaiquegarcia/gojooble/internal/views/searchresult"
	"github.com/kaiquegarcia/gojooble/internal/views/setconfig"
	"github.com/kaiquegarcia/gojooble/jooble"
)

func main() {
	// configuration
	conf, err := config.Load()
	if err != nil {
		fmt.Printf("não foi possível carregar as configurações: %s\n", err)
		return
	}

	// ask config if it's empty
	if conf.IsEmpty() {
		// force setconfig view
		conf, err = setconfig.Run(conf)
		if err != nil {
			fmt.Printf("não foi possível definir a configuração: %s\n", err)
			return
		}
	}

	// initialize clients
	httpClient := http.DefaultClient
	joobleClient := jooble.New(httpClient, conf.ApiKey())

	// run home view
	for {
		choice, err := home.Run()
		if err != nil {
			fmt.Printf("não foi possível processar a escolha inicial: %s\n", err)
			return
		}

		quit := false
		switch choice {
		case home.OptSetconfig:
			c, err := setconfig.Run(conf)
			if err != nil {
				fmt.Printf("não foi possível definir a configuração: %s\n", err)
				return
			}

			conf = c
		case home.OptSearch:
			response, err := search.Run(conf, joobleClient)
			if err != nil {
				fmt.Printf("não foi possível realizar a busca: %s\n", err)
				return
			}

			if response != nil {
				err = searchresult.Run(response)
				if err != nil {
					fmt.Printf("não foi possível apresentar o resultado da busca: %s\n", err)
					return
				}
			}
		default:
			quit = true
		}

		if quit {
			break
		}
	}
}
