package godataeye

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//平台类型
const (
	PlatformTypeIOS = 1 //IOS
	PlatformTypeADR = 2 //Android
	PlatformTypeWP  = 3 //WP
)

//账号类型
const (
	AccountTypeAnonymous  = 0 //AccountType默认值,未知用户来源
	AccountTypeRegistered = 1 //游戏自身注册用户
	AccountTypeSinaWeibo  = 2 //新浪微博用户
	AccountTypeQQ         = 3 //QQ用户
	AccountTypeQQWeibo    = 4 //腾讯微博用户
	AccountTypeND91       = 5 //91用户
	//6~15	AccountType.Type1~10	自定义类型
	//AccountType.Type1~10
)

//性别
const (
	GenderUNKNOWN = 0 //未知,默认值
	GenderMALE    = 1 //男性
	GenderFEMALE  = 2 //女性
)

//网络类型
const (
	NetType3G    = 1  //3G
	NetTypeWIFI  = 2  //Wifi
	NetType2G    = 4  //2G
	NetTypeOTHER = 3  //其他
	NetType4G    = 14 //4G
)

//任务类型
const (
	TaskTypeGuideLine  = 1 //新手任务
	TaskTypeMainLine   = 2 //主线任务
	TaskTypeBranchLine = 3 //分支任务
	TaskTypeDaily      = 4 //日常任务
	TaskTypeActivity   = 5 //活动任务
	TaskTypeOther      = 6 //其他任务,默认值
)

var (
	Host       string //"http://ext.gdatacube.net/dc/rest/"
	AppId      string //Appid
	AppVersion string //App版本号
)

type DataEayLog interface {
	JoinRest(buffer *bytes.Buffer)
	Name() string
}

type Context struct {
	AccountType int    //账号类型(账号枚举)
	Mac         string //
	Imei        string //
	Gender      int    //性别(性别枚举)
	Age         int    //
	Resolution  string //分辨率	720*540
	OsVersion   string //操作系统版本	4.0.0
	Brand       string //机型
	Language    string //操作系统语言
	NetType     int    //网络类型(网络枚举)
	Ip          string //IP
	Country     string //国家
	Province    string //省份
	Operators   string //运营商
}

func (self *Context) JoinRest(buffer *bytes.Buffer) {
	//Context
	if self.AccountType != 0 {
		buffer.WriteString("&accountType=")
		buffer.WriteString(fmt.Sprint(self.AccountType))
	}
	if self.Mac != "" {
		buffer.WriteString("&mac=")
		buffer.WriteString(QueryEscape(self.Mac))
	}
	if self.Imei != "" {
		buffer.WriteString("&imei=")
		buffer.WriteString(QueryEscape(self.Imei))
	}
	if self.Gender != 0 {
		buffer.WriteString("&gender=")
		buffer.WriteString(fmt.Sprint(self.Gender))
	}
	if self.Age != 0 {
		buffer.WriteString("&age=")
		buffer.WriteString(fmt.Sprint(self.Age))
	}
	if self.Resolution != "" {
		buffer.WriteString("&resolution=")
		buffer.WriteString(QueryEscape(self.Resolution))
	}
	if self.OsVersion != "" {
		buffer.WriteString("&osVersion=")
		buffer.WriteString(QueryEscape(self.OsVersion))
	}
	if self.Brand != "" {
		buffer.WriteString("&brand=")
		buffer.WriteString(QueryEscape(self.Brand))
	}
	if self.Language != "" {
		buffer.WriteString("&language=")
		buffer.WriteString(QueryEscape(self.Language))
	}
	if self.NetType != 0 {
		buffer.WriteString("&netType=")
		buffer.WriteString(fmt.Sprint(self.NetType))
	}
	if self.Ip != "" {
		buffer.WriteString("&ip=")
		buffer.WriteString(QueryEscape(self.Ip))
	}
	if self.Country != "" {
		buffer.WriteString("&country=")
		buffer.WriteString(QueryEscape(self.Country))
	}
	if self.Province != "" {
		buffer.WriteString("&province=")
		buffer.WriteString(QueryEscape(self.Province))
	}
	if self.Operators != "" {
		buffer.WriteString("&operators=")
		buffer.WriteString(QueryEscape(self.Operators))
	}
}

type AccountInfo struct {
	Context

	AccountId string //帐号(必填)

	RoleId    string //角色ID
	RoleName  string //角色昵称
	RoleClass string //角色职业
	RoleRace  string //角色种族

	Platform   int    //平台类型(平台枚举)(必填)
	GameRegion string //区服(必填)
	Channel    string //渠道(必填)
}

func (self *AccountInfo) JoinRest(buffer *bytes.Buffer) {
	buffer.WriteString("&accountId=")
	buffer.WriteString(QueryEscape(self.AccountId))
	buffer.WriteString("&platform=")
	buffer.WriteString(fmt.Sprint(self.Platform))
	buffer.WriteString("&gameRegion=")
	buffer.WriteString(QueryEscape(self.GameRegion))
	buffer.WriteString("&channel=")
	buffer.WriteString(QueryEscape(self.Channel))

	if self.RoleId != "" {
		buffer.WriteString("&roleId=")
		buffer.WriteString(QueryEscape(self.RoleId))
	}
	if self.RoleName != "" {
		buffer.WriteString("&roleName=")
		buffer.WriteString(QueryEscape(self.RoleName))
	}
	if self.RoleClass != "" {
		buffer.WriteString("&roleClass=")
		buffer.WriteString(QueryEscape(self.RoleClass))
	}
	if self.RoleRace != "" {
		buffer.WriteString("&roleRace=")
		buffer.WriteString(QueryEscape(self.RoleRace))
	}
	self.Context.JoinRest(buffer)
}

//激活/注册
type ActOrReg struct {
	*AccountInfo
	ActTime int //激活时间	unix时间戳	与regTime其一必填
	RegTime int //注册时间	unix时间戳	与actTime其一必填
}

func (self *ActOrReg) Name() string {
	return "actOrReg"
}

func (self *ActOrReg) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)
	if self.ActTime != 0 {
		buffer.WriteString("&actTime=")
		buffer.WriteString(fmt.Sprint(self.ActTime))
	}
	if self.RegTime != 0 {
		buffer.WriteString("&regTime=")
		buffer.WriteString(fmt.Sprint(self.RegTime))
	}
}

//在线
type Online struct {
	*AccountInfo
	LoginTime  int //登陆时间	unix时间戳
	OnlineTime int //在线时间	unix时间戳
	Level      int //玩家等级
}

func (self *Online) Name() string {
	return "online"
}

func (self *Online) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&loginTime=")
	buffer.WriteString(fmt.Sprint(self.LoginTime))

	buffer.WriteString("&onlineTime=")
	buffer.WriteString(fmt.Sprint(self.OnlineTime))

	if self.Level != 0 {
		buffer.WriteString("&level=")
		buffer.WriteString(fmt.Sprint(self.Level))
	}
}

//真实消费统计
type Pay struct {
	*AccountInfo
	CurrencyAmount float64 //付费金额
	CurrencyType   string  //货币类型
	PayType        string  //支付方式
	Iapid          string  //付费点
	PayTime        int     //支付时间	unix时间戳
	OrderId        string  //订单ID
}

func (self *Pay) Name() string {
	return "pay"
}

func (self *Pay) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&currencyAmount=")
	buffer.WriteString(fmt.Sprint(self.CurrencyAmount))

	buffer.WriteString("&currencyType=")
	buffer.WriteString(QueryEscape(self.CurrencyType))

	if self.PayType != "" {
		buffer.WriteString("&payType=")
		buffer.WriteString(QueryEscape(self.PayType))
	}
	if self.Iapid != "" {
		buffer.WriteString("&iapid=")
		buffer.WriteString(QueryEscape(self.Iapid))
	}

	buffer.WriteString("&payTime=")
	buffer.WriteString(fmt.Sprint(self.PayTime))

	buffer.WriteString("&orderId=")
	buffer.WriteString(QueryEscape(self.OrderId))
}

//获得虚拟币
type Coin struct {
	*AccountInfo
	CoinNum   int    //获得的虚拟币数量
	CoinType  string //虚拟币类型
	Type      string //获得虚拟的途径
	IsGain    int    //是否获得	0 消耗 1 获得
	TotalCoin int    //该玩家手里最终持有的货币数量	用于计算虚拟币留存
	MsgTime   int    //实际发生时间	unix时间戳
}

func (self *Coin) Name() string {
	return "coin"
}

func (self *Coin) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&coinNum=")
	buffer.WriteString(fmt.Sprint(self.CoinNum))
	buffer.WriteString("&coinType=")
	buffer.WriteString(QueryEscape(self.CoinType))
	buffer.WriteString("&type=")
	buffer.WriteString(QueryEscape(self.Type))
	buffer.WriteString("&isGain=")
	buffer.WriteString(fmt.Sprint(self.IsGain))
	buffer.WriteString("&totalCoin=")
	buffer.WriteString(fmt.Sprint(self.TotalCoin))

	if self.MsgTime != 0 {
		buffer.WriteString("&msgTime=")
		buffer.WriteString(fmt.Sprint(self.MsgTime))
	}
}

//购买道具
type ItemBuy struct {
	*AccountInfo
	ItemId   string //道具ID
	ItemType string //道具类型
	ItemCnt  string //购买的道具数量
	CoinNum  int    //消耗虚拟币数量
	CoinType string //虚拟币种类
	MsgTime  int    //实际发生时间	unix时间戳
}

func (self *ItemBuy) Name() string {
	return "itemBuy"
}

func (self *ItemBuy) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&itemId=")
	buffer.WriteString(QueryEscape(self.ItemId))
	buffer.WriteString("&itemType=")
	buffer.WriteString(QueryEscape(self.ItemType))
	buffer.WriteString("&itemCnt=")
	buffer.WriteString(QueryEscape(self.ItemCnt))

	buffer.WriteString("&coinNum=")
	buffer.WriteString(fmt.Sprint(self.CoinNum))
	buffer.WriteString("&coinType=")
	buffer.WriteString(QueryEscape(self.CoinType))

	if self.MsgTime != 0 {
		buffer.WriteString("&msgTime=")
		buffer.WriteString(fmt.Sprint(self.MsgTime))
	}
}

//使用物品/道具
type ItemUse struct {
	*AccountInfo
	ItemId   string //道具ID
	ItemType string //道具类型
	ItemCnt  int    //道具消耗数量
	Reason   string //道具消耗的途径
	MsgTime  int    //实际发生时间	unix时间戳
}

func (self *ItemUse) Name() string {
	return "itemUse"
}

func (self *ItemUse) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&itemId=")
	buffer.WriteString(QueryEscape(self.ItemId))
	buffer.WriteString("&itemType=")
	buffer.WriteString(QueryEscape(self.ItemType))
	buffer.WriteString("&itemCnt=")
	buffer.WriteString(fmt.Sprint(self.ItemCnt))
	buffer.WriteString("&reason=")
	buffer.WriteString(QueryEscape(self.Reason))
	if self.MsgTime != 0 {
		buffer.WriteString("&msgTime=")
		buffer.WriteString(fmt.Sprint(self.MsgTime))
	}
}

//物品/道具掉落
type ItemGet struct {
	*AccountInfo
	ItemId   string //道具ID
	ItemType string //道具类型
	ItemCnt  int    //道具消耗数量
	Reason   string //道具消耗的途径
	MsgTime  int    //实际发生时间	unix时间戳
}

func (self *ItemGet) Name() string {
	return "itemGet"
}

func (self *ItemGet) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&itemId=")
	buffer.WriteString(QueryEscape(self.ItemId))
	buffer.WriteString("&itemType=")
	buffer.WriteString(QueryEscape(self.ItemType))
	buffer.WriteString("&itemCnt=")
	buffer.WriteString(fmt.Sprint(self.ItemCnt))
	buffer.WriteString("&reason=")
	buffer.WriteString(QueryEscape(self.Reason))
	if self.MsgTime != 0 {
		buffer.WriteString("&msgTime=")
		buffer.WriteString(fmt.Sprint(self.MsgTime))
	}
}

//任务完成/失败
type Task struct {
	*AccountInfo
	TaskId   string //任务ID
	TaskType int    //任务类型(任务枚举)
	Duration int    //任务耗时
	IsSucc   int    //是否失败	0 任务失败 1 任务完成
	Reason   string //任务失败原因
}

func (self *Task) Name() string {
	return "task"
}

func (self *Task) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&taskId=")
	buffer.WriteString(QueryEscape(self.TaskId))

	buffer.WriteString("&taskType=")
	buffer.WriteString(fmt.Sprint(self.TaskType))
	buffer.WriteString("&duration=")
	buffer.WriteString(fmt.Sprint(self.Duration))
	buffer.WriteString("&isSucc=")
	buffer.WriteString(fmt.Sprint(self.IsSucc))

	buffer.WriteString("&reason=")
	buffer.WriteString(QueryEscape(self.Reason))
}

//自定义事件
type Event struct {
	*AccountInfo
	EventId  string //事件ID
	Duration int    //事件耗时
	LabelMap string //事件属性列表	Json
}

func (self *Event) Name() string {
	return "event"
}

func QueryEscape(s string) string {
	if s == "" {
		return s
	}
	return url.QueryEscape(s)
}

func (self *Event) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&eventId=")
	buffer.WriteString(QueryEscape(self.EventId))

	if self.Duration != 0 {
		buffer.WriteString("&duration=")
		buffer.WriteString(fmt.Sprint(self.Duration))
	}
	if self.LabelMap != "" {
		buffer.WriteString("&labelMap=")
		buffer.WriteString(QueryEscape(self.LabelMap))
	}
}

//角色升级
type LevelUp struct {
	*AccountInfo
	StartLevel int //开始等级
	EndLevel   int //结束等级
	Interval   int //升级时长（秒）
}

func (self *LevelUp) Name() string {
	return "levelUp"
}

func (self *LevelUp) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&startLevel=")
	buffer.WriteString(fmt.Sprint(self.StartLevel))
	buffer.WriteString("&endLevel=")
	buffer.WriteString(fmt.Sprint(self.EndLevel))
	buffer.WriteString("&interval=")
	buffer.WriteString(fmt.Sprint(self.Interval))
}

//自定义标签-添加标签
type AddTag struct {
	*AccountInfo
	Tag    string //标签名
	SubTag string //子标签名
	Seq    int    //设置tag时的时间戳，后台根据时间先用判断 add/remove先后，默认为后台接收时间
}

func (self *AddTag) Name() string {
	return "addTag"
}

func (self *AddTag) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&tag=")
	buffer.WriteString(QueryEscape(self.Tag))
	if self.SubTag != "" {
		buffer.WriteString("&subTag=")
		buffer.WriteString(QueryEscape(self.SubTag))
	}
	if self.Seq != 0 {
		buffer.WriteString("&seq=")
		buffer.WriteString(fmt.Sprint(self.Seq))
	}
}

//自定义标签-移除标签
type RemoveTag struct {
	*AccountInfo
	Tag    string //标签名
	SubTag string //子标签名
	Seq    int    //设置tag时的时间戳，后台根据时间先用判断 add/remove先后，默认为后台接收时间
}

func (self *RemoveTag) Name() string {
	return "removeTag"
}

func (self *RemoveTag) JoinRest(buffer *bytes.Buffer) {
	self.AccountInfo.JoinRest(buffer)

	buffer.WriteString("&tag=")
	buffer.WriteString(QueryEscape(self.Tag))
	if self.SubTag != "" {
		buffer.WriteString("&subTag=")
		buffer.WriteString(QueryEscape(self.SubTag))
	}
	if self.Seq != 0 {
		buffer.WriteString("&seq=")
		buffer.WriteString(fmt.Sprint(self.Seq))
	}
}

func GetRestUrl(log DataEayLog) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(Host)
	buffer.WriteString(log.Name())
	buffer.WriteString("?appId=")
	buffer.WriteString(AppId)

	if AppVersion != "" {
		buffer.WriteString("&appVersion=")
		buffer.WriteString(AppVersion)
	}
	log.JoinRest(&buffer)
	r := buffer.String()
	return r
}

func Report(log DataEayLog) string {
	add := GetRestUrl(log)
	fmt.Printf("%v\r\n", add)
	return GetUrl(add)
}

func GetUrl(add string) string {
	client := http.Client{}
	reqest, err := http.NewRequest("GET", add, nil)
	if err != nil {
		return err.Error()
	}
	reqest.Header.Add("Accept-Encoding", "deflate")
	reqest.Header.Add("Connection", "close")
	response, err := client.Do(reqest)
	if err != nil {
		return err.Error()
	}
	var body string
	defer response.Body.Close()
	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		body = err.Error()
	} else {
		body = string(bodyByte)
	}
	return body
}

func GetUrl2(add string) string {
	res, err := http.Get(add)
	if err != nil {
		return err.Error()
		//log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err.Error()
		//log.Fatal(err)
	}
	//fmt.Printf("%s", body)
	return string(body)
}
