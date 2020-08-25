package main

import (
	"fmt"
	"github.com/PGshen/june/common/setting"
	"github.com/PGshen/june/common/validator"
	"github.com/PGshen/june/routers"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

func main() {
	binding.Validator = new(validator.DefaultValidator)
	router := routers.InitRouter()
	conf := setting.Config.Server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Port),
		Handler:        router,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
