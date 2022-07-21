package httpServer

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/pengzhong2010/go-server-base/app/app1/conf"
	"github.com/pengzhong2010/go-server-base/app/app1/internal/router"
	"github.com/pengzhong2010/go-server-base/common/pkg/log"

	"github.com/gin-gonic/gin"
)

func New() (s *http.Server, cf func(), err error) {
	gin.DisableConsoleColor()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.SetupRouter(r)
	s = &http.Server{
		Addr:    ":" + conf.APP_PORT,
		Handler: r,
	}
	cf = func() {
		Close(s)
	}
	if err = s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Error("listen: %s\n", err)
	}
	return
}
func Close(s *http.Server) {
	log.Info("Shutting down exit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Warn("Server forced to shutdown:", err)
	}
}
