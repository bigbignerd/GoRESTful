package user

import (
    "strconv"
    . "github.com/bigbignerd/GoRESTful/handler"
    "github.com/bigbignerd/GoRESTful/model"
    "github.com/bigbignerd/GoRESTful/pkg/errno"
    "github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
    userId, _ := strconv.Atoi(c.Param("id"))
    if err := model.DeleteUser(uint64(userId)); err != nil {
        SendResponse(c, errno.ErrDatabase, nil)
        return
    }
    SendResponse(c, nil, nil)
}
