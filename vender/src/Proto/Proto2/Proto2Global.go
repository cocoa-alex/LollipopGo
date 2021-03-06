package Proto2

import (
	"LollipopGo/LollipopGo/player"
)

// G_GameGlobal_Proto == 9  负责全局的游戏逻辑 的子协议
// 注：server类型为单点
const (
	ININGlobal                 = iota // 0
	G2GW_ConnServerProto2             // G2GW_ConnServerProto2 == 1 Global主动链接 gateway
	GW2G_ConnServerProto2             // GW2G_ConnServerProto2 == 2 选择链接
	GW2G_HeartBeatProto2              // GW2G_HeartBeatProto2 ==  3  心跳协议  保活的操作
	G2GW_PlayerMatchProto2            // G2GW_PlayerMatchProto2 == 4 玩家发送匹配的协议
	GW2G_PlayerMatchProto2            // GW2G_PlayerMatchProto2 == 5 服务器返回数据对应匹配机制
	G2GW_PlayerEntryHallProto2        // G2GW_PlayerEntryHallProto2 == 6 玩家进入大厅，显示的数据
	GW2G_PlayerEntryHallProto2        // GW2G_PlayerEntryHallProto2 == 7

)

//------------------------------------------------------------------------------
// G2GW_PlayerEntryHallProto2
type G2GW_PlayerEntryHall struct {
	Protocol      int
	Protocol2     int
	UID           string // 用户唯一ID,app的数据库的ID信息
	OpenID        string // 用户唯一ID
	PlayerName    string // 玩家的名字
	HeadUrl       string // 头像
	Constellation string // 星座
	PlayerSchool  string // 学校
	Sex           string // 性别
	Token         string // 数据验证
}

// GW2G_PlayerEntryHallProto2 查询需要返回的协议
type GW2G_PlayerEntryHall struct {
	Protocol      int
	Protocol2     int
	OpenID        string                      // 用户唯一ID
	PlayerName    string                      // 玩家的名字
	HeadUrl       string                      // 头像
	Constellation string                      // 星座
	Sex           string                      // 性别
	GamePlayerNum map[string]interface{}      // 每个游戏的玩家的人数,global server获取
	RacePlayerNum map[string]interface{}      // 大奖赛列表
	Personal      map[string]*player.PlayerSt // 个人信息
	DefaultMsg    map[string]*player.MsgST    // 默认跑马灯消息
	DefaultAward  map[string]interface{}      // 默认兑换列表
}

//------------------------------------------------------------------------------
// G2GW_PlayerMatchProto2  玩家发送匹配的协议
type G2GW_PlayerMatch struct {
	Protocol   int
	Protocol2  int
	GWServerID int    // 网关的ID信息，主要是数据统计需要
	Type       int    // 匹配类型，1 1V1 ,2 2V2 ,3 5V5
	OpenID     string // 用户唯一ID
}

// GW2G_PlayerMatchProto2 服务器返回数据对应匹配机制
type GW2G_PlayerMatch struct {
	Protocol  int
	Protocol2 int
	Type      int    // 匹配类型，1 1V1 ,2 2V2 ,3 5V5
	RoomInfo  string // 匹配成功后的房间中的信息
}

//------------------------------------------------------------------------------
// GW2G_HeartBeatProto2  心跳协议
type GW2G_HeartBeat struct {
	Protocol  int
	Protocol2 int
	ServerID  int //全局配置 唯一的也是
}

//------------------------------------------------------------------------------
// G2GW_ConnServerProto2  去gateway去链接
type G2GW_ConnServer struct {
	Protocol  int
	Protocol2 int
	ServerID  string //全局配置 唯一的也是
}

// GW2G_ConnServerProto2 返回的数据链接
type GW2G_ConnServer struct {
	Protocol  int
	Protocol2 int
	ServerID  string
}

//------------------------------------------------------------------------------
