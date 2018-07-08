package service

import (
    "fmt"
    "sync"

    "github.com/bigbignerd/GoRESTful/model"
    "github.com/bigbignerd/GoRESTful/util"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
    infos := make([]*model.UserInfo, 0)
    users, count, err := model.ListUser(username, offset, limit)

    if err != nil {
        return nil, count, err
    }
    ids := []uint64{}
    for _, user := range users {
        ids = append(ids, user.Id)
    }
    //并行处理
    wg := sync.WaitGroup{}
    userList := model.UserList{
        Lock: new(sync.Mutex),
        IdMap: make(map[uint64]*model.UserInfo, len(users)),
    }

    errChan := make(chan error, 1)
    finished := make(chan bool, 1)

    //imporve query efficiency
    for _, u := range users {
        wg.Add(1)
        go func(u *model.UserModel) {
            defer wg.Done()

            shortId, err := util.GenShortId()
            if err != nil {
                errChan <- err
                return
            }
            //并行处理保证数据的一致性
            userList.Lock.Lock()
            defer userList.Lock.Unlock()
            userList.IdMap[u.Id] = &model.UserInfo{
                Id: u.Id,
                Username: u.Username,
                SayHello: fmt.Sprintf("Hello,%s", shortId),
                Password: u.Password,
                CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:04"),
                UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:04"),
            }
        }(u)
    }

    go func() {
        wg.Wait()
        close(finished)
    }()

    select {
        case <- finished:
        case err := <- errChan:
            return nil, count, err
    }

    for _, id := range ids {
        infos = append(infos, userList.IdMap[id])
    }
    return infos, count, nil

}
