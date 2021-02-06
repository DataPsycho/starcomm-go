package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Discovery is the entry point in server call
func Discovery(c *gin.Context) {
	req := defaultRequest()
	_ = c.Bind(req)
	res := reply(req)
	c.JSON(http.StatusOK, res)
}

func mapUrls() {
	router.POST("/discovery", Discovery)
}
