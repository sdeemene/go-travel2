package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func LoggingUri(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr, r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
func Logger() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.SetOutput(logger.Writer())
}
