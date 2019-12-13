package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// 有了omitempty后，如果结构体中改字段为空， 则生成的json中没有该字段。

// KzzMxResult 明细
type KzzMxResult struct {
	KzzMx []KZZMX `json:"KZZ_MX,omitempty"`
}

// KZZMX 明细
type KZZMX struct {
	Bondcode          string  `json:"BONDCODE,omitempty"`
	Bcode             string  `json:"BCODE,omitempty"`
	Texch             string  `json:"TEXCH,omitempty"`
	Sname             string  `json:"SNAME,omitempty"`
	Startdate         string  `json:"STARTDATE,omitempty"`
	Correscode        string  `json:"CORRESCODE,omitempty"`
	Corresname        string  `json:"CORRESNAME,omitempty"`
	Limitbuyipub      string  `json:"LIMITBUYIPUB,omitempty"`
	Swapscode         string  `json:"SWAPSCODE,omitempty"`
	Securityshortname string  `json:"SECURITYSHORTNAME,omitempty"`
	Parvalue          float64 `json:"PARVALUE,omitempty"`
	Issueprice        float64 `json:"ISSUEPRICE,omitempty"`
	Swapprice         float64 `json:"SWAPPRICE,omitempty"`
	Aissuevol         float64 `json:"AISSUEVOL,omitempty"`
	Zqhdate           string  `json:"ZQHDATE,omitempty"`
	Luckrate          float64 `json:"LUCKRATE,omitempty"`
	Listdate          string  `json:"LISTDATE,omitempty"`
	Delistdate        string  `json:"DELISTDATE,omitempty"`
	Memo              string  `json:"MEMO,omitempty"`
	Oldcorrescode     string  `json:"OLDCORRESCODE,omitempty"`
	Oldcorresname     string  `json:"OLDCORRESNAME,omitempty"`
	Issueobject       string  `json:"ISSUEOBJECT,omitempty"`
	Swapsdate         string  `json:"SWAPSDATE,omitempty"`
	Swapedate         string  `json:"SWAPEDATE,omitempty"`
	Creditrating      string  `json:"CREDITRATING,omitempty"`
	Partiesfname      string  `json:"PARTIESFNAME,omitempty"`
	Issueyear         string  `json:"ISSUEYEAR,omitempty"`
	Bondperiod        string  `json:"BONDPERIOD,omitempty"`
	Frstvaluedate     string  `json:"FRSTVALUEDATE,omitempty"`
	Lastvaluedate     string  `json:"LASTVALUEDATE,omitempty"`
	Mrtydate          string  `json:"MRTYDATE,omitempty"`
	Payday            string  `json:"PAYDAY,omitempty"`
	Ratedes           string  `json:"RATEDES,omitempty"`
	StartdateLv       string  `json:"STARTDATE_LV,omitempty"`
	EnddateLv         string  `json:"ENDDATE_LV,omitempty"`
	CashdateLv        string  `json:"CASHDATE_LV,omitempty"`
	CouponrateLv      float64 `json:"COUPONRATE_LV,omitempty"`
}

// KzzZqhResult 中签号
type KzzZqhResult struct {
	KzzZqh []KZZZQH `json:"KZZ_ZQH,omitempty"`
}

// KZZZQH 中签号
type KZZZQH struct {
	Bondcode string `json:"BONDCODE,omitempty"`
	Texch    string `json:"TEXCH,omitempty"`
	Bcode    string `json:"BCODE,omitempty"`
	Sname    string `json:"SNAME,omitempty"`
	Typecode string `json:"TYPECODE,omitempty"`
	Strtype  string `json:"TYPE,omitempty"`
	Lucknum  string `json:"LUCKNUM,omitempty"`
}

// SingelLuck 单个中签号信息
type SingelLuck struct {
	ZqType     string `json:"zq_type,omitempty"`      // 中签类型
	ZqTypeCode string `json:"zq_type_code,omitempty"` // 中签类型代号
	ZqNumber   string `json:"zq_number,omitempty"`    // 中签号
}

// LuckInfo 所有的中签号信息
type LuckInfo struct {
	ZzCode     string       `json:"zz_code,omitempty"`     // 转债代码
	ZzName     string       `json:"zz_name,omitempty"`     // 转债名称
	ZzLuckrate float64      `json:"zz_luckrate,omitempty"` // 中签率
	ZqAll      []SingelLuck `json:"zq_all,omitempty"`      // 所有中签结果
}

// QueryResult 查询的结果
type QueryResult struct {
	IsSuccess   bool           `json:"IsSuccess,omitempty"`
	Code        int            `json:"Code,omitempty"`
	Message     string         `json:"Message,omitempty"`
	TotalPage   int            `json:"TotalPage,omitempty"`
	TotalCount  int            `json:"TotalCount,omitempty"`
	Keyword     string         `json:"Keyword,omitempty"`
	Data        []TypeBaseInfo `json:"Data,omitempty"`
	RelatedWord string         `json:"RelatedWord,omitempty"`
}

// TypeBaseInfo 查询到的数据
type TypeBaseInfo struct {
	TypeCode int          `json:"Type,omitempty"`  // 类型代码
	Name     string       `json:"Name,omitempty"`  // 类型名称
	Count    int          `json:"Count,omitempty"` //计数
	Datas    []ZqBaseInfo `json:"Datas,omitempty"` //数据
}

// ZqBaseInfo 债券基本信息
type ZqBaseInfo struct {
	Code         string `json:"Code,omitempty"` // 代码
	Name         string `json:"Name,omitempty"` // 名称
	ID           string `json:"ID,omitempty"`
	MktNum       string `json:"MktNum,omitempty"`
	SecurityType string `json:"SecurityType,omitempty"`
	MarketType   string `json:"MarketType,omitempty"`
	Jys          string `json:"JYS,omitempty"`
	GubaCode     string `json:"GubaCode,omitempty"` //交易所
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
func CheckNumber(startingNum string, strLuckNum string) string {
	// 公布的中签号 将当前签号字符串切割为数组
	arrLuckNum := strings.Split(strLuckNum, ",")
	arrLuckLen := len(arrLuckNum)
	// 匹配你的配号
	var yourLuckNum string
	for i := 0; i < 1000; i++ {
		// 获取当前配号
		nStartingNum, _ := strconv.Atoi(startingNum)
		curNum := nStartingNum + i
		// 将当前配号转换为字符串
		strCurNum := strconv.Itoa(curNum)
		for j := 0; j < arrLuckLen; j++ {
			var strCurLuckNum = arrLuckNum[j]
			// 检测当前配号是否包含中签号
			if strCurLuckNum != "" && strings.Contains(strCurNum, strCurLuckNum) {
				fmt.Println("恭喜你！你的中签号码为：", strCurNum)
				yourLuckNum += strCurNum
			}
		}
	}
	return yourLuckNum
}

// DoWork 请求URL获取数据，并查询中签结果
func DoWork(code string, startingNum string) (yourLuckInfo LuckInfo) {
	// 切片，[]里面为空或者...，切片的长度和容量可以不固定
	// var yourLuckInfo LuckInfo //你的中签信息

	// url example:
	// http://dcfm.eastmoney.com/em_mutisvcexpandinterface/api/js/get?type=KZZ_ZQH&token=70f12f2f4f091e459a279469fe49eca5&filter=(BONDCODE='113549')

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	urlBase := "http://dcfm.eastmoney.com/em_mutisvcexpandinterface/api/js/get"
	token := "70f12f2f4f091e459a279469fe49eca5"
	filter := "(BONDCODE='" + code + "')"

	// 可转债明细
	urlMx := urlBase + "?type=KZZ_MX&js={\"KZZ_MX\":(x)}&token=" + token + "&filter=" + filter
	// 中签号码
	urlZqh := urlBase + "?type=KZZ_ZQH&js={\"KZZ_ZQH\":(x)}&token=" + token + "&filter=" + filter

	// 可转债明细
	cbMx, errMxGet := HTTPGet(urlMx)
	// 中签号码
	cbZqh, errZqhGet := HTTPGet(urlZqh)

	if errMxGet != nil || errZqhGet != nil {
		fmt.Printf("未查询到中签号为%s的信息...\n", code)
		return yourLuckInfo
	}

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
	// zzName := "白电转债"
	// startingNum := "100049383324"

	zzName := "新北转债"
	startingNum := "000042256125"

	zzInfo := GetCode(zzName)
	zzCode := zzInfo.Code

	fmt.Println("zzCode:", zzCode)

	// 开始爬取数据
	result := DoWork(zzCode, startingNum)
	fmt.Println("========中签结果result:=========", result)

}
