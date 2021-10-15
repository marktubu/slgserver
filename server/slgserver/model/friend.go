package model

import (
	"fmt"
	"slgserver/log"
	"slgserver/server/slgserver/proto"
	"time"
)

var friendDbMgr *FriendDBMgr

type FriendDBMgr struct {
	friends chan *Friend
}

func init() {
	friendDbMgr = &FriendDBMgr{friends: make(chan *Friend, 100)}
	go friendDbMgr.running()
}

func (this *FriendDBMgr) running() {
	for true {
		select {
		case friend := <-this.friends:
			if friend.Id > 0 {

			} else {
				log.DefaultLog.Warn("update friend fail, because id <= 0")
			}
		}
	}
}

func (this *FriendDBMgr) push(friend *Friend) {
	this.friends <- friend
}

type Friend struct {
	Id    int       `xorm:"id pk autoincr"`
	Name  string    `xorm:"name"`
	Ctime time.Time `xorm:"ctime"`
}

func (this *Friend) ToProto() interface{} {
	p := &proto.Friend{}
	p.Id = this.Id
	p.Name = this.Name

	return p
}

func (this *Friend) TableName() string {
	return "tb_friend" + fmt.Sprintf("_%d", ServerId)
}

type FriendLog struct {
	Id    int       `xorm:"id pk autoincr"`
	Desc  string    `xorm:"desc"`
	Ctime time.Time `xorm:"ctime"`
}
