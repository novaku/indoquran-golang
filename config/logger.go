package config

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// SetLogger : setup the logger, output and to a file
func SetLogger() {
	// Logging to a file.
	f, _ := os.OpenFile("./gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}
