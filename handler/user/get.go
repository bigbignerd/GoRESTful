package user

import (
    . "github.com/bigbignerd/GoRESTful/handler"
    "github.com/bigbignerd/GoRESTful/model"
    "github.com/bigbignerd/GoRESTful/pkg/errno"
    
    "github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
    username := c.Param("username")

    user, err := model.GetUser(username)
    if err != nil {
        SendResponse(c, errno.ErrUserNotFound, nil)
        return
    }
    SendResponse(c, nil, user)
}
