package proto

const (
	FriendUntreated = 0 //未处理
	FriendRefuse    = 1 //拒绝
	FriendAdopt     = 2 //通过
)

type Friend struct {
	Id   int    `json:"id"`   //id
	Name string `json:"name"` //名字
}

type FriendListReq struct {
}

type FriendListRsp struct {
	Friends []Friend `json:"friends"`
}

type FriendAddReq struct {
	Id int `json:"id"`
}

type FriendAddRsp struct {
}

type FriendDeleteReq struct {
	Id int `json:"id"`
}

type FriendDeleteRsp struct {
}

type FriendApplyItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//获取申请列表
type FriendApplyReq struct {
}

type FriendApplyRsp struct {
	Applys []FriendApplyItem `json:"applys"`
}

//审核
type FriendVerifyReq struct {
	Id     int  `json:"id"`     //申请操作的id
	Decide int8 `json:"decide"` //1是拒绝，2是通过
}

type FriendVerifyRsp struct {
	Id     int  `json:"id"`     //申请操作的id
	Decide int8 `json:"decide"` //1是拒绝，2是通过
}
