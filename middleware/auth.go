package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/feiyangderizi/ginServer/model/result"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			r := result.FailWithMsg("身份验证失败，缺少token信息")
			r.Response(c)
			c.Abort()
			return
		}

		//token验证
		if r := checkToken(token); r.IsOK() {
			r.Response(c)
			c.Abort()
			return
		}
	}
}

//token验证
func checkToken(token string) result.Result {
	if token == "1233" {
		return result.Success()
	} else {
		return result.FailWithMsg("无效token")
	}
}
