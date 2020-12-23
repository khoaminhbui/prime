package main

import (
	"log"

	"github.com/khoaminhbui/prime/prime-service/deliveries/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	// init config
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// init server
	e := echo.New()
	http.InitHandler(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
