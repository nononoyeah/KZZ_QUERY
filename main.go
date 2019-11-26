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
	kzzMx []KZZMX `json:"KZZ_MX"`
}

// KzzZqhResult 中签号
type KzzZqhResult struct {
	kzzZqh []KZZZQH `json:"KZZ_ZQH"`
}

// KZZMX 明细
type KZZMX struct {
	bondcode string `json:"BONDCODE"`
	bcode 	string `json:"BCODE"`
	texch string `json:"TEXCH"`
	sname string `json:"SNAME"`
	startdate string `json:"STARTDATE"`
	correscode string `json:"CORRESCODE"`
	corresname string `json:"CORRESNAME"`
	limitbuyipub string `json:"LIMITBUYIPUB"`
	swapscode string `json:"SWAPSCODE"`
	securityshortname string `json:"SECURITYSHORTNAME"`
	parvalue int    `json:"PARVALUE"`
	issueprice int    `json:"ISSUEPRICE"`
	swapprice int    `json:"SWAPPRICE"`
	aissuevol int    `json:"AISSUEVOL"`
	zqhdate string `json:"ZQHDATE"`
	luckrate int    `json:"LUCKRATE"`
	listdate string `json:"LISTDATE"`
	delistdate string `json:"DELISTDATE"`
	memo string `json:"MEMO"`
	oldcorrescode string `json:"OLDCORRESCODE"`
	oldcorresname string `json:"OLDCORRESNAME"`
	issueobject string `json:"ISSUEOBJECT"`
	swapsdate string `json:"SWAPSDATE"`
	swapedate string `json:"SWAPEDATE"`
	creditrating string `json:"CREDITRATING"`
	partiesfname string `json:"PARTIESFNAME"`
	issueyear string `json:"ISSUEYEAR"`
	bondperiod string `json:"BONDPERIOD"`
	frstvaluedate string `json:"FRSTVALUEDATE"`
	lastvaluedate string `json:"LASTVALUEDATE"`
	mrtydate string `json:"MRTYDATE"`
	payday string `json:"PAYDAY"`
	ratedes string `json:"RATEDES"`
	startdateLv      string `json:"STARTDATE_LV"`
	enddateLv        string `json:"ENDDATE_LV"`
	cashdateLv       string `json:"CASHDATE_LV"`
	couponrateLv     int    `json:"COUPONRATE_LV"`
}

// KZZZQH 中签号
type KZZZQH struct {
	bondcode string `json:"BONDCODE"`
	texch string `json:"TEXCH"`
	bcode string `json:"BCODE"`
	sname string `json:"SNAME"`
	typecode string `json:"TYPECODE"`
	strtype string `json:"TYPE"`
	lucknum string `json:"LUCKNUM"`
}

// SingelLuck 单个中签号信息
type SingelLuck struct {
	zqType string // 中签类型
	zqTypeCode string // 中签类型代号
	zqNumber string // 中签号
}

// LuckInfo 所有的中签号信息
type LuckInfo struct {
	zzCode string       // 转债代码
	zzName string       // 转债名称
	zzLuckrate int          // 中签率
	zqAll []SingelLuck // 所有中签结果
}

// QueryResult 查询的结果
type QueryResult struct {
	isSuccess  string
	code       int
	message    string
	totalPage  int
	totalCount int
	keyword    bool
	data       []TypeBaseInfo
}

// TypeBaseInfo 查询到的数据
type TypeBaseInfo struct {
	typeCode  string         // 类型代码
	name      string         // 类型名称
	count     int            //计数
	datas     []ZqBaseInfo //数据
}

// ZqBaseInfo 债券基本信息
type ZqBaseInfo struct {
	code         string `json:"Code"` // 代码
	name         string `json:"Name"` // 名称
	id           string `json:"ID"`
	mktNum       string `json:"MktNum"`
	securityType string `json:"SecurityType"`
	marketType   string `json:"MarketType"`
	jys          string `json:"JYS"`
	gubaCode     string `json:"GubaCode"` //交易所
}

// GetCode 通过债券名字查询债券信息
func GetCode(zzName string) (zzInfo ZqBaseInfo) {
	url := "http://api.so.eastmoney.com/bussiness/web/QuotationLabelSearch?token=32A8A21716361A5A387B0D85259A0037&keyword=" + zzName + "&type=0&pi=1&ps=30&_=1574674323184"

	fmt.Println("url=====", url)
	result, _ := HTTPGet(url)

	zzResult := QueryResult{}

	// 解析复杂json
	_ = json.Unmarshal(result, &zzResult)

	fmt.Println("result: ====", result)
	fmt.Println("result: ====", string(result))
	fmt.Println("zzResult: ====", zzResult)

	zzInfo = zzResult.data[0].datas[0]

	return
}

// HTTPGet HTTP的GET请求
func HTTPGet(url string) (result []byte, err error) {
	// fmt.Println("==================url: %s===================", url)
	resp, errHTTP := http.Get(url)
	if errHTTP != nil {
		err = errHTTP
		return
	}
	// defer用于资源的释放，会在函数之前进行调用
	// 先给返回值复制，然后调用defer表达式，最后才是返回调用函数中
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// fmt.Println("==================httpget: %s===================", string(body))
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
	filter := "(BONDCODE=%27 + code + %27)"

	// 可转债明细
	urlMx := urlBase + "?type=KZZ_MX + &token=" + token + "&filter=" + filter + "&js={\"KZZ_MX\":(x)}"
	// 中签号码
	urlZqh := urlBase + "?type=KZZ_ZQH" + "&token=" + token + "&filter=" + filter + "&js={\"KZZ_ZQH\":(x)}"

	// 可转债明细
	cbMx, _ := HTTPGet(urlMx)
	// 中签号码
	cbZqh, _ := HTTPGet(urlZqh)

	kzzmx := KzzMxResult{}
	kzzzqh := KzzZqhResult{}

	// 解析复杂json
	_ = json.Unmarshal(cbMx, &kzzmx)
	_ = json.Unmarshal(cbZqh, &kzzzqh)

	oKzzmx := kzzmx.kzzMx[0]
	aKzzzqh := kzzzqh.kzzZqh

	bondCode := oKzzmx.bondcode // 代码
	sname     := oKzzmx.sname    // 名称
	luckrate  := oKzzmx.luckrate // 中签率

	yourLuckInfo.zzCode = bondCode    //代码
	yourLuckInfo.zzName = sname        //名字
	yourLuckInfo.zzLuckrate = luckrate //中签率

	cbZqhLen := len(aKzzzqh)
	for i := 0; i < cbZqhLen; i++ {
		sKzzzqh := aKzzzqh[i]
		var singleZq SingelLuck // 某一种中签情况

		singleZq.zqType = sKzzzqh.strtype
		singleZq.zqTypeCode = sKzzzqh.typecode

		yourLuckNum := CheckNumber(startingNum, sKzzzqh.lucknum)
		singleZq.zqNumber = yourLuckNum

		yourLuckInfo.zqAll = append(yourLuckInfo.zqAll, singleZq)
	}

	return yourLuckInfo
}

func main() {
	// var name, code string
	fmt.Println("中签查询...")

	// 通过中文名获取代码
	zzName := "白电转债"
	// var startingNum int = 100049383324

	zzInfo := GetCode(zzName)
	zzCode := zzInfo.code

	fmt.Println("zzCode:", zzCode)

	// 开始爬取数据
	// result := DoWork(string(zzCode), startingNum)
	// fmt.Println("========中签结果result:=========", result)

}
