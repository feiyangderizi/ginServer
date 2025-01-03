package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maczh/logs"
	"github.com/maczh/mgtrace"
	"github.com/maczh/utils"
	"gopkg.in/mgo.v2/bson"

	"bid-dh-cpic/global"
	"bid-dh-cpic/initialize"
)

type PostLog struct {
	ID           bson.ObjectId          `bson:"_id"`
	Time         string                 `json:"time" bson:"time"`
	RequestId    string                 `json:"requestId" bson:"requestId"`
	ResponseTime string                 `json:"responsetime" bson:"responsetime"`
	TTL          int                    `json:"ttl" bson:"ttl"`
	ApiName      string                 `json:"apiName" bson:"apiName"`
	Controller   string                 `json:"controller" bson:"controller"`
	Token        string                 `json:"token" bson:"token"`
	RequestParam map[string]string      `json:"requestparam" bson:"requestparam"`
	ResponseStr  string                 `json:"responsestr" bson:"responsestr"`
	ResponseMap  map[string]interface{} `json:"responsemap" bson:"responsemap"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var accessChannel = make(chan string, 100)

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func SetRequestLogger() gin.HandlerFunc {

	go handleAccessChannel()

	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 开始时间
		startTime := time.Now()

		data, err := c.GetRawData()
		if err != nil {
			logs.Error("GetRawData error:", err.Error())
		}
		body := string(data)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 重新设置请求的Body

		// 处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()

		var result map[string]interface{}

		// 日志格式
		if strings.Contains(c.Request.RequestURI, "/docs") {
			return
		}

		if responseBody != "" && responseBody[0:1] == "{" {
			err := json.Unmarshal([]byte(responseBody), &result)
			if err != nil {
				result = map[string]interface{}{"status": -1, "msg": "解析异常:" + err.Error()}
			}
		}

		// 结束时间
		endTime := time.Now()

		// 日志格式
		params := utils.GinParamMap(c)
		if body != "" {
			params["body"] = body
		}
		postLog := new(PostLog)
		postLog.ID = bson.NewObjectId()
		postLog.Time = startTime.Format("2006-01-02 15:04:05")
		if strings.Contains(c.Request.RequestURI, "?") {
			postLog.Controller = c.Request.RequestURI[0:strings.Index(c.Request.RequestURI, "?")]
		} else {
			postLog.Controller = c.Request.RequestURI
		}
		postLog.RequestId = mgtrace.GetRequestId()
		postLog.RequestParam = params
		postLog.ResponseTime = endTime.Format("2006-01-02 15:04:05")
		postLog.ResponseMap = result
		postLog.TTL = int(endTime.UnixNano()/1e6 - startTime.UnixNano()/1e6)

		accessLog := "|" + c.Request.Method + "|" + postLog.Controller + "|" + c.ClientIP() + "|" + endTime.Format("2006-01-02 15:04:05.012") + "|" + fmt.Sprintf("%vms", endTime.UnixNano()/1e6-startTime.UnixNano()/1e6)

		global.Logger.Debug(accessLog)
		global.Logger.Debug(fmt.Sprintf("请求参数:%s", utils.ToJSON(params)))
		global.Logger.Debug(fmt.Sprintf("接口返回:%s", utils.ToJSON(result)))

		if global.Config.Application.UseMongodb && global.Config.MongoDB.LogCollection != "" {
			accessChannel <- utils.ToJSON(postLog)
		}
	}
}

func handleAccessChannel() {
	for accessLog := range accessChannel {
		var postLog PostLog
		err := json.Unmarshal([]byte(accessLog), &postLog)
		if err != nil {
			global.Logger.Error("日志写入错误:" + err.Error())
		}
		mongoDBConn := initialize.MongoDBConnection{}
		conn := mongoDBConn.Get()
		err = conn.C(global.Config.MongoDB.LogCollection).Insert(postLog)
		if err != nil {
			global.Logger.Error("日志写入错误:" + err.Error())
		}
	}
	return
}
