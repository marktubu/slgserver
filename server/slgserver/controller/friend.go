package controller

import (
	"slgserver/constant"
	"slgserver/middleware"
	"slgserver/net"
	"slgserver/server/slgserver/logic/mgr"
	"slgserver/server/slgserver/proto"

	"github.com/goinggo/mapstructure"
)

var DefaultFriend = Friend{}

type Friend struct {
}

func (this *Friend) InitRouter(r *net.Router) {
	g := r.Group("friend").Use(middleware.ElapsedTime(), middleware.Log())

	g.AddRouter("list", this.list)
	g.AddRouter("add", this.add)
	g.AddRouter("delete", this.delete)
	g.AddRouter("verify", this.verify)
	g.AddRouter("applyList", this.applyList)
}

func (this *Friend) list(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	reqObj := &proto.FriendListReq{}
	rspObj := &proto.FriendListRsp{}
	mapstructure.Decode(req.Body.Msg, reqObj)
	rsp.Body.Msg = rspObj
	rsp.Body.Code = constant.OK

	l := mgr.FriendMgr.List()
	rspObj.List = make([]proto.Friend, len(l))
	for i, u := range l {
		rspObj.List[i] = u.ToProto().(proto.Friend)
	}
}

func (this *Friend) add(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	reqObj := &proto.FriendAddReq{}
	rspObj := &proto.FriendAddRsp{}
	mapstructure.Decode(req.Body.Msg, reqObj)
	rsp.Body.Msg = rspObj
	rsp.Body.Code = constant.OK

}

func (this *Friend) delete(req *net.WsMsgReq, rsp *net.WsMsgRsp) {

}

func (this *Friend) verify(req *net.WsMsgReq, rsp *net.WsMsgRsp) {

}

func (this *Friend) applyList(req *net.WsMsgReq, rsp *net.WsMsgRsp) {

}
