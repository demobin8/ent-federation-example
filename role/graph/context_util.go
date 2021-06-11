package graph

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io/ioutil"
	"os"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

/**
请求之前
*/
func RestLogAop() func(c *gin.Context) {
	return func(c *gin.Context) {

		// 开始时间
		startTime := time.Now()

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// 取一下
		data, _ := ioutil.ReadAll(c.Request.Body)
		// 放回去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		// 处理请求
		c.Next()
		params := data

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		log := zerolog.New(os.Stdout).With().Timestamp().Logger()

		log.Info().Int("status_code", statusCode).
			Str("req_uri", reqUri).
			Str("req_method", reqMethod).
			Str("client_ip", c.Request.Host).
			Dur("latency_time", latencyTime).
			RawJSON("req_params", params).
			Str("response:", w.body.String()).Send()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}