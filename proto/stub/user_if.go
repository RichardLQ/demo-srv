package stub

import (
	"context"
)

var UserService = &userServiceServer{}

type userServiceServer struct {
	UserServiceServer
}

//GetUserToken 获取token
func (this *userServiceServer) GetUserToken(ctx context.Context, req *UserReq) (rsp *UserRsp, err error) {
	//m := auth.MyClaims{}
	rsp = &UserRsp{
		Code: 200,
		Data: req.Account,
	}
	return
}

//GetUserInfo 获取用户信息
func (this *userServiceServer) GetUserInfo(ctx context.Context, req *UserInfoReq) (rsp *UserInfoRsp, err error) {
	rsp = &UserInfoRsp{
		Code: 200,
		User: &User{
			Id:       req.UserId,
			Username: "Jerome",
			Desc:     "这是一条数据",
		},
	}
	return
}

//DelUserInfo 删除用户信息
func (this *userServiceServer) DelUserInfo(ctx context.Context, req *UserInfoReq) (rsp *UserRsp, err error) {
	rsp = &UserRsp{
		Code: 200,
		Msg:  "删除成功！",
		Data: "一大堆杂数据",
	}
	return
}
