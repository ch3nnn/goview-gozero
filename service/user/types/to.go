package types

import (
	carbon "github.com/golang-module/carbon/v2"

	"github.com/ch3nnn/goview-gozero/service/user/internal/dal/model"
	"github.com/ch3nnn/goview-gozero/service/user/pb"
)

func ToJWT(userId int64, jwt string) *pb.Token {
	return &pb.Token{
		TokenName:            "Authorization",
		TokenValue:           jwt,
		IsLogin:              true,
		LoginId:              userId,
		LoginType:            "login",
		TokenTimeout:         carbon.Now().AddDay().Timestamp(),
		SessionTimeout:       carbon.Now().AddDay().Timestamp(),
		TokenSessionTimeout:  carbon.Now().AddDay().Timestamp(),
		TokenActivityTimeout: -1,
		LoginDevice:          "default-device",
		Tag:                  nil,
	}
}

func ToUserInfo(user *model.SysUser) *pb.UserInfo {
	return &pb.UserInfo{
		Id:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Nickname: user.Nickname,
		DepId:    user.DepID,
		PosId:    user.PosID,
		DepName:  "",
		PosName:  "",
	}
}
