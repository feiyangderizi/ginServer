package middleware

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maczh/utils"

	"github.com/feiyangderizi/ginServer/global"
	"github.com/feiyangderizi/ginServer/model/result"
	"github.com/feiyangderizi/ginServer/service"
)

func CheckSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		appKey := c.Request.Header.Get("appKey")
		if appKey == "" {
			r := result.FailWithMsg("身份验证失败，缺少appKey信息")
			r.Response(c)
			c.Abort()
			return
		}
		params := utils.GinParamMap(c)

		if !utils.Exists(params, "sign") {
			r := result.FailWithMsg("缺少签名")
			r.Response(c)
			c.Abort()
			return
		}
		appService := service.AppService{}
		r := appService.GetSecret(appKey)
		if !r.IsOK() {
			r.Response(c)
			c.Abort()
			return
		}

		appSecret := r.Data.(string)

		if err := checkSign(params, appSecret); err != nil {
			r := result.FailWithMsg(err.Error())
			r.Response(c)
			return
		}
	}
}

func checkSign(params map[string]string, appSecret string) error {
	var keys = make([]string, 0)
	for k := range params {
		if strings.ToLower(k) != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	var strParams []string
	for _, p := range keys {
		strParams = append(strParams, p+"="+params[p])
	}

	sign := utils.MD5Encode(strings.Join(strParams, "&") + appSecret)

	if strings.ToLower(params["sign"]) != sign {
		global.Logger.Debug(fmt.Sprintf("传入签名：%s，服务端计算签名：%s，请求参数：%s，appSecret：%s", params["sign"], sign, utils.ToJSON(params), appSecret))
		return errors.New("签名验证失败")
	}
	return nil
}
