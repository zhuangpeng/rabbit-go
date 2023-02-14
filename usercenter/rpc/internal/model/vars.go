package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

var UserAuthTypeUserName string = "username"  // 平台内部
var UserAuthTypeMiniApp  string = "wechatApp" // 微信小程序
var UserAuthTypeEmail    string = "email"     // 微信小程序
var UserAuthTypeMobile   string = "mobile"    // 手机号码