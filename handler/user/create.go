package user

import (
    "fmt"
    "github.com/bigbignerd/GoRESTful/pkg/errno"
    . "github.com/bigbignerd/GoRESTful/handler"
    "github.com/gin-gonic/gin"
    "github.com/lexkong/log"
)

func Create(c *gin.Context) {
    var r CreateRequest

    if err := c.Bind(&r); err != nil {
        fmt.Print(err)
        SendResponse(c, errno.ErrBind, nil)
        return
    }
    admin2 := c.Param("username")
    log.Infof("URL username:%s", admin2)

    desc := c.Query("desc")
    log.Infof("URL key param desc: %s", desc)

    contentType := c.GetHeader("Content-Type")
    log.Infof("Header Content-Type:%s", contentType)

    log.Debugf("username is:[%s], password is [%s]", r.Username, r.Password)

    if r.Username == "" {
        SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in xx")), nil)
        return
    }

    if r.Password == "" {
        SendResponse(c, fmt.Errorf("password is empty"), nil)
    }
    rsp := CreateResponse{
        Username : r.Username,
    }
    SendResponse(c, nil, rsp)
}
