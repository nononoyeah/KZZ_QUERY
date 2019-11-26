package main

// package

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// KzzMxResult 明细
type KzzMxResult struct {
	KzzMx []KZZMX `json:"KZZ_MX"`
}

// KZZMX 明细
type KZZMX struct {
	Bondcode          string  `json:"BONDCODE"`
	Bcode             string  `json:"BCODE"`
	Texch             string  `json:"TEXCH"`
	Sname             string  `json:"SNAME"`
	Startdate         string  `json:"STARTDATE"`
	Correscode        string  `json:"CORRESCODE"`
	Corresname        string  `json:"CORRESNAME"`
	Limitbuyipub      string  `json:"LIMITBUYIPUB"`
	Swapscode         string  `json:"SWAPSCODE"`
	Securityshortname string  `json:"SECURITYSHORTNAME"`
	Parvalue          float64 `json:"PARVALUE"`
	Issueprice        float64 `json:"ISSUEPRICE"`
	Swapprice         float64 `json:"SWAPPRICE"`
	Aissuevol         float64 `json:"AISSUEVOL"`
	Zqhdate           string  `json:"ZQHDATE"`
	Luckrate          float64 `json:"LUCKRATE"`
	Listdate          string  `json:"LISTDATE"`
	Delistdate        string  `json:"DELISTDATE"`
	Memo              string  `json:"MEMO"`
	Oldcorrescode     string  `json:"OLDCORRESCODE"`
	Oldcorresname     string  `json:"OLDCORRESNAME"`
	Issueobject       string  `json:"ISSUEOBJECT"`
	Swapsdate         string  `json:"SWAPSDATE"`
	Swapedate         string  `json:"SWAPEDATE"`
	Creditrating      string  `json:"CREDITRATING"`
	Partiesfname      string  `json:"PARTIESFNAME"`
	Issueyear         string  `json:"ISSUEYEAR"`
	Bondperiod        string  `json:"BONDPERIOD"`
	Frstvaluedate     string  `json:"FRSTVALUEDATE"`
	Lastvaluedate     string  `json:"LASTVALUEDATE"`
	Mrtydate          string  `json:"MRTYDATE"`
	Payday            string  `json:"PAYDAY"`
	Ratedes           string  `json:"RATEDES"`
	StartdateLv       string  `json:"STARTDATE_LV"`
	EnddateLv         string  `json:"ENDDATE_LV"`
	CashdateLv        string  `json:"CASHDATE_LV"`
	CouponrateLv      float64 `json:"COUPONRATE_LV"`
}

// KzzZqhResult 中签号
type KzzZqhResult struct {
	KzzZqh []KZZZQH `json:"KZZ_ZQH"`
}

// KZZZQH 中签号
type KZZZQH struct {
	Bondcode string `json:"BONDCODE"`
	Texch    string `json:"TEXCH"`
	Bcode    string `json:"BCODE"`
	Sname    string `json:"SNAME"`
	Typecode string `json:"TYPECODE"`
	Strtype  string `json:"TYPE"`
	Lucknum  string `json:"LUCKNUM"`
}

// SingelLuck 单个中签号信息
type SingelLuck struct {
	ZqType     string // 中签类型
	ZqTypeCode string // 中签类型代号
	ZqNumber   string // 中签号
}

// LuckInfo 所有的中签号信息
type LuckInfo struct {
	ZzCode     string       // 转债代码
	ZzName     string       // 转债名称
	ZzLuckrate float64      // 中签率
	ZqAll      []SingelLuck // 所有中签结果
}

// QueryResult 查询的结果
type QueryResult struct {
	IsSuccess   bool           `json:"IsSuccess"`
	Code        int            `json:"Code"`
	Message     string         `json:"Message"`
	TotalPage   int            `json:"TotalPage"`
	TotalCount  int            `json:"TotalCount"`
	Keyword     string         `json:"Keyword"`
	Data        []TypeBaseInfo `json:"Data"`
	RelatedWord string         `json:"RelatedWord"`
}

// TypeBaseInfo 查询到的数据
type TypeBaseInfo struct {
	TypeCode int          `json:"Type"`  // 类型代码
	Name     string       `json:"Name"`  // 类型名称
	Count    int          `json:"Count"` //计数
	Datas    []ZqBaseInfo `json:"Datas"` //数据
}

// ZqBaseInfo 债券基本信息
type ZqBaseInfo struct {
	Code         string `json:"Code"` // 代码
	Name         string `json:"Name"` // 名称
	ID           string `json:"ID"`
	MktNum       string `json:"MktNum"`
	SecurityType string `json:"SecurityType"`
	MarketType   string `json:"MarketType"`
	Jys          string `json:"JYS"`
	GubaCode     string `json:"GubaCode"` //交易所
}

// GetCode 通过债券名字查询债券信息
func GetCode(zzName string) (zzInfo ZqBaseInfo) {

	url := "http://api.so.eastmoney.com/bussiness/web/QuotationLabelSearch?token=32A8A21716361A5A387B0D85259A0037&type=0&pi=1&ps=30&_=1574674323184&keyword=" + zzName

	result, _ := HTTPGet(url)

	zzResult := QueryResult{}
	_ = json.Unmarshal(result, &zzResult)

	zzInfo = zzResult.Data[0].Datas[0]

	return
}

// HTTPGet HTTP的GET请求
func HTTPGet(url string) (result []byte, err error) {
	resp, errHTTP := http.Get(url)
	if errHTTP != nil {
		err = errHTTP
		return
	}
	// defer用于资源的释放，会在函数之前进行调用
	// 先给返回值复制，然后调用defer表达式，最后才是返回调用函数中
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

// CheckNumber 中签号比对
func CheckNumber(startingNum int, strLuckNum string) string {
	// 公布的中签号 将当前签号字符串切割为数组
	arrLuckNum := strings.Split(strLuckNum, ",")
	arrLuckLen := len(arrLuckNum)
	// 匹配你的配号
	var yourLuckNum string
	for i := 0; i < 1000; i++ {
		// 获取当前配号
		curNum := startingNum + i
		// 将当前配号转换为字符串
		strCurNum := strconv.Itoa(curNum)
		for j := 0; j < arrLuckLen; j++ {
			var strCurLuckNum = arrLuckNum[j]
			// 检测当前配号是否包含中签号
			if strings.Contains(strCurNum, strCurLuckNum) {
				fmt.Println("恭喜你！你的中签号码为：", strCurNum)
				yourLuckNum += strCurNum
			}
		}
	}
	return yourLuckNum
}

// DoWork 请求URL获取数据，并查询中签结果
func DoWork(code string, startingNum int) (yourLuckInfo LuckInfo) {
	// 切片，[]里面为空或者...，切片的长度和容量可以不固定
	// var yourLuckInfo LuckInfo //你的中签信息

	// url example:
	// http://dcfm.eastmoney.com/em_mutisvcexpandinterface/api/js/get?type=KZZ_ZQH&token=70f12f2f4f091e459a279469fe49eca5&filter=(BONDCODE='113549')

	urlBase := "http://dcfm.eastmoney.com/em_mutisvcexpandinterface/api/js/get"
	token := "70f12f2f4f091e459a279469fe49eca5"
	filter := "(BONDCODE='" + code + "')"

	// 可转债明细
	urlMx := urlBase + "?type=KZZ_MX&js={\"KZZ_MX\":(x)}&token=" + token + "&filter=" + filter
	// 中签号码
	urlZqh := urlBase + "?type=KZZ_ZQH&js={\"KZZ_ZQH\":(x)}&token=" + token + "&filter=" + filter

	// 可转债明细
	cbMx, _ := HTTPGet(urlMx)
	// 中签号码
	cbZqh, _ := HTTPGet(urlZqh)

	kzzmx := KzzMxResult{}
	kzzzqh := KzzZqhResult{}

	// 解析复杂json
	errMx := json.Unmarshal(cbMx, &kzzmx)
	errZqh := json.Unmarshal(cbZqh, &kzzzqh)

	if errMx != nil || errZqh != nil {
		fmt.Println("转债明细解析错误1...", errMx, errZqh)
		return yourLuckInfo
	}

	oKzzmx := kzzmx.KzzMx[0]
	aKzzzqh := kzzzqh.KzzZqh

	bondCode := oKzzmx.Bondcode // 代码
	sname := oKzzmx.Sname       // 名称
	luckrate := oKzzmx.Luckrate // 中签率

	yourLuckInfo.ZzCode = bondCode     //代码
	yourLuckInfo.ZzName = sname        //名字
	yourLuckInfo.ZzLuckrate = luckrate //中签率

	cbZqhLen := len(aKzzzqh)
	for i := 0; i < cbZqhLen; i++ {
		sKzzzqh := aKzzzqh[i]
		var singleZq SingelLuck // 某一种中签情况

		singleZq.ZqType = sKzzzqh.Strtype
		singleZq.ZqTypeCode = sKzzzqh.Typecode

		yourLuckNum := CheckNumber(startingNum, sKzzzqh.Lucknum)
		singleZq.ZqNumber = yourLuckNum

		yourLuckInfo.ZqAll = append(yourLuckInfo.ZqAll, singleZq)
	}

	return yourLuckInfo
}

func main() {
	fmt.Println("中签查询...")

	// 通过中文名获取代码
	zzName := "白电转债"
	startingNum := 100049383324

	zzInfo := GetCode(zzName)
	zzCode := zzInfo.Code

	fmt.Println("zzCode:", zzCode)

	// 开始爬取数据
	result := DoWork("113549", startingNum)
	fmt.Println("========中签结果result:=========", result)

}
