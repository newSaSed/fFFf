package main

import "github.com/gin-gonic/gin"

func main() {

	nigr := gin.Default()

	nigr.LoadHTMLGlob("Html/*.html")

	nigr.Static("assets", "style")

	nigr.GET("/", index)
	nigr.GET("/cloth", cloth)
	nigr.GET("/niz", niz)
	nigr.GET("/golovka", golovka)
	nigr.GET("/zima", zima)
	nigr.Run("127.0.0.1:8080")

}

func index(c *gin.Context) {

	c.HTML(200, "index.html", nil)

}
func cloth(c *gin.Context) {

	c.HTML(200, "cloth.html", nil)
}
func niz(c *gin.Context) {

	c.HTML(200, "niz.html", nil)
}
func golovka(c *gin.Context) {

	c.HTML(200, "golovka.html", nil)
}
func zima(c *gin.Context) {

	c.HTML(200, "zima.html", nil)
}
