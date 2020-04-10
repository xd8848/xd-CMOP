package cmd

import (
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/init/initlog"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RunWindowsServer(Router *gin.Engine) {
	address := fmt.Sprintf(":%d", config.GinVueAdminconfig.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	log.L.Debug("server run success on ", address)

	_ = s.ListenAndServe()
}
