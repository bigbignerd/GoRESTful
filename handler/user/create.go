package user

import (
    "github.com/bigbignerd/GoRESTful/pkg/errno"
    "github.com/bigbignerd/GoRESTful/model"
    . "github.com/bigbignerd/GoRESTful/handler"
    "github.com/gin-gonic/gin"
    "github.com/lexkong/log"
    "github.com/lexkong/log/lager"
    "github.com/bigbignerd/GoRESTful/util"
)

func Create(c *gin.Context) {
    log.Info("user create function called.", lager.Data{"X-Request-Id":util.GetReqID(c)})
    var r CreateRequest

    if err := c.Bind(&r); err != nil {
        SendResponse(c, errno.ErrBind, nil)
        return
    }
    u := model.UserModel{
        Username: r.Username,
        Password: r.Password,
    }
    if err := u.Validate(); err != nil {
        SendResponse(c, errno.ErrValidation, nil)
        return
    }
    //encrypt user password
    if err := u.Encrypt(); err != nil {
        SendResponse(c, errno.ErrEncrypt, nil)
        return
    }
    //save to mysql
    if err := u.Create(); err != nil {
        SendResponse(c, errno.ErrEncrypt, nil)
        return
    }
    rsp := CreateResponse{
        Username : r.Username,
    }
    SendResponse(c, nil, rsp)
}
