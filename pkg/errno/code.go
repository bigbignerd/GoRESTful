package errno

var (
    //Common errors
    OK = &Errno{Code: 0, Message: "OK"}
    InternalServerError = &Errno{Code: 100001, Message: "Internal server error."}
    ErrBind = &Errno{Code:100002, Message: "Error occurred while binding the request body"}
    //user errors
    ErrUserNotFound = &Errno{Code: 200001, Message: "The user was not found"}
)


