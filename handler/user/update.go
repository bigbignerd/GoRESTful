package user

import (
    . "github.com/bigbignerd/GoRESTful/handler"
    "github.com/bigbignerd/GoRESTful/model"
    "github.com/bigbignerd/GoRESTful/pkg/errno"
    "github.com/bigbignerd/GoRESTful/util"
    "github.com/gin-gonic/gin"
    "github.com/lexkong/log"
    "github.com/lexkong/log/lager"
    "strconv"
    "fmt"
)

//update user info
func Update(c *gin.Context) {
    log.Info("user update method called. ", lager.Data{"X-Request-Id":util.GetReqID(c)})
    userId, _ := strconv.Atoi(c.Param("id"))

    //bind user data to model
    var u model.UserModel
    if err := c.Bind(&u); err != nil {
        SendResponse(c, errno.ErrBind, nil)
        return
    }
    u.Id = uint64(userId)

    //validate user data
    if err := u.Validate(); err != nil {
        SendResponse(c, errno.ErrValidation, nil)
        return
    }

    //encrypt password
    if err := u.Encrypt(); err != nil {
        SendResponse(c, errno.ErrEncrypt, nil)
        return
    }

    //save change fields
    if err := u.Update(); err != nil {
        fmt.Print(err)
        SendResponse(c, errno.ErrDatabase, nil)
        return
    }

    SendResponse(c, nil, nil)
}
