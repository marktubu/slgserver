package mgr

import (
	"slgserver/db"
	"slgserver/log"
	"slgserver/server/slgserver/model"
	"sync"

	"go.uber.org/zap"
)

type friendMgr struct {
	mutex   sync.RWMutex
	friends map[int]*model.Friend
}

var FriendMgr = &friendMgr{
	friends: make(map[int]*model.Friend),
}

func (this *friendMgr) Get(id int) (*model.Friend, bool) {
	this.mutex.RLock()
	f, ok := this.friends[id]
	this.mutex.RUnlock()

	if ok {
		return f, true
	}

	m := &model.Friend{}
	ok, err := db.MasterDB.Table(new(model.Friend)).Where("id=?", id).Get(m)
	if ok {
		this.mutex.Lock()
		this.friends[id] = m
		this.mutex.Unlock()
		return m, true
	} else {
		log.DefaultLog.Warn("db error", zap.Error(err), zap.Int("id", id))
		return nil, false
	}
}

func (this *friendMgr) List() []*model.Friend {
	r := make([]*model.Friend, 0)
	this.mutex.RLock()
	for _, friend := range this.friends {
		r = append(r, friend)
	}
	this.mutex.RUnlock()
	return r
}
