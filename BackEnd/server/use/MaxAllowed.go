package use

import "github.com/gin-gonic/gin"

// MaxAllowed 限制API同时请求次数 https://github.com/aviddiviner/gin-limit
func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}
