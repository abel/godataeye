package godataeye

import (
	//"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Point struct {
	X int
	Y int
}

func GetDefaultAccount() *AccountInfo {
	Host = "http://ext.gdatacube.net/dc/rest/"
	AppId = "E6F18DA5BFE3D953C0ADC274E6C14C3D"
	a := AccountInfo{}
	a.AccountId = "seonzhang"
	a.Platform = 1
	a.Channel = "channel"
	a.GameRegion = "gameRegion"
	a.OsVersion = "9.1"
	a.Imei = "99000310516212"
	a.Mac = "74:27:EA:6B:35:2F"
	a.Resolution = "720*960"
	a.Country = "中国"
	a.Ip = "58.60.188.178"
	a.NetType = 1
	a.Age = 26
	a.Province = "北京"
	a.AccountType = 1
	a.Brand = "iphone6s"
	a.Gender = 1
	a.Operators = "中国移动"
	a.Language = "cn"
	return &a
}

const (
	Success = "{\"code\":0}"
)

func TestActOrReg(t *testing.T) {
	log := ActOrReg{
		AccountInfo: GetDefaultAccount(),
		ActTime:     1415620801,
		RegTime:     1415620801,
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("ActOrReg", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestOnline(t *testing.T) {
	log := Online{
		AccountInfo: GetDefaultAccount(),
		LoginTime:   1415620801,
		OnlineTime:  3000,
		Level:       1,
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("Online", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestPay(t *testing.T) {
	log := Pay{
		AccountInfo:    GetDefaultAccount(),
		CurrencyAmount: 100,
		CurrencyType:   "cny",
		PayType:        "alipay",
		Iapid:          "buycoin",
		PayTime:        1415620801,
		OrderId:        "14156208011415620801",
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("Pay", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestCoin(t *testing.T) {
	log := Coin{
		AccountInfo: GetDefaultAccount(),
		CoinNum:     100,
		CoinType:    "gold",
		Type:        "buy",
		IsGain:      1,
		TotalCoin:   1999,
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("Coin", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestItemBuy(t *testing.T) {
	log := ItemBuy{
		AccountInfo: GetDefaultAccount(),
		ItemId:      "sword",
		ItemType:    "weapons",
		ItemCnt:     "1",
		CoinNum:     100,
		CoinType:    "gold",
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("ItemBuy", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestItemUse(t *testing.T) {
	log := ItemUse{
		AccountInfo: GetDefaultAccount(),
		ItemId:      "sword",
		ItemType:    "weapons",
		ItemCnt:     1,
		Reason:      "broken",
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("ItemUse", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestItemGet(t *testing.T) {
	log := ItemGet{
		AccountInfo: GetDefaultAccount(),
		ItemId:      "sword",
		ItemType:    "weapons",
		ItemCnt:     1,
		Reason:      "gift",
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("ItemGet", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestTask(t *testing.T) {
	log := Task{
		AccountInfo: GetDefaultAccount(),
		TaskId:      "task1",
		TaskType:    1,
		Duration:    60,
		IsSucc:      0,
		Reason:      "failreason",
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("Task", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestEvent(t *testing.T) {
	log := Event{
		AccountInfo: GetDefaultAccount(),
		EventId:     "task1",
		Duration:    60,
		LabelMap:    "{\"a\":1}",
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("Event", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestLevelUp(t *testing.T) {
	log := LevelUp{
		AccountInfo: GetDefaultAccount(),
		StartLevel:  9,
		EndLevel:    10,
		Interval:    3600,
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("LevelUp", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestAddTag(t *testing.T) {
	log := AddTag{
		AccountInfo: GetDefaultAccount(),
		Tag:         "TestTag",
		SubTag:      "TestSubTag1",
		Seq:         1422688944,
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("AddTag", func() {
			So(r, ShouldEqual, Success)
		})
	})
}

func TestRemoveTag(t *testing.T) {
	log := RemoveTag{
		AccountInfo: GetDefaultAccount(),
		Tag:         "TestTag",
		SubTag:      "TestSubTag1",
		Seq:         1422688944,
	}
	r := Report(&log)
	Convey("dataeye test", t, func() {
		Convey("RemoveTag", func() {
			So(r, ShouldEqual, Success)
		})
	})
}
