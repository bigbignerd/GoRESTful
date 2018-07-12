package middleware

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "regexp"
    "time"

    "github.com/bigbignerd/GoRESTful/handler"
    "github.com/bigbignerd/GoRESTful/pkg/errno"

    "github.com/gin-gonic/gin"
    "github.com/lexkong/log"
    "github.com/willf/pad"
)
type BodyLogWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now().UTC()
        path := c.Request.URL.Path

        reg := regexp.MustCompile("(v1/user/|/login)")
        if !reg.MatchString(path) {
            return
        }
        var bodyBytes []byte
        if c.Request.Body != nil {
            bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
        }
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
        
        method := c.Request.Method
        ip := c.ClientIP()

        blw := &BodyLogWriter{
            body: bytes.NewBufferString(""),
            ResponseWriter: c.Writer,
        }
        c.Writer = blw

        c.Next()
        
        end := time.Now().UTC()
        latency := end.Sub(start)

        code, message := -1, ""

        var response handler.Response
        if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
            log.Errorf(err, "response body can not unmarshal to model.Resonse struct")
            code = errno.InternalServerError.Code
            message = err.Error()
        } else {
            code = response.Code
            message = response.Message
        }
        log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)

        }
    }


