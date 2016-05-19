package useragent

import (
	"fmt"
	// "os"
	// "strings"
	"testing"
)

var TestData = map[string]string{
	"BingBot":       "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
	"BaiduSpider":   "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
	"GoogleBot":     "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	"YahooChina":    "Mozilla/5.0 (compatible; Yahoo! Slurp China; http://misc.yahoo.com.cn/help.html)",
	"YodaoBot":      "Mozilla/5.0 (compatible; YodaoBot/1.0; http://www.yodao.com/help/webmaster/spider/; )",
	"YoudaoBot":     "Mozilla/5.0 (compatible; YoudaoBot/1.0; http://www.youdao.com/help/webmaster/spider/; )",
	"MSNBot":        "msnbot/2.1",
	"SousouSpider":  "Sosospider+(+http://help.soso.com/webspider.htm)",
	"YahooSeeker":   "YahooSeeker/1.2 (compatible; Mozilla 4.0; MSIE 5.5; yahooseeker at yahoo-inc dot com ; http://help.yahoo.com/help/us/shop/merchant/)",
	"chinasospider": "chinasospider",
	///////////////////////////////////////////////////////////////
	"ie6":  "Mozilla/4.0 (Windows; MSIE 6.0; Windows NT 5.2)",
	"ie7":  "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"ie8":  "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"ie9":  "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
	"ie10": "Mozilla/5.0 (compatible; WOW64; MSIE 10.0; Windows NT 6.2)",
	"ie11": "Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv 11.0) like Gecko",

	"ipad":   "Mozilla/5.0 (iPad; CPU OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B466 Safari/600.1.4",
	"iphone": "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A366 Safari/600.1.4",
	"ipod":   "Mozilla/5.0 (iPod; CPU iPhone OS 5_1_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9B206 Safari/7534.48.3",

	"wphone7":   "Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; LG; GW910)",
	"wphone7_5": "Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; SGH-i917)",
	"wphone8":   "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 920)",

	"nexus7":      "Mozilla/5.0 (Linux; Android 4.1.1; Nexus 7 Build/JRO03D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19",
	"galaxys3":    "Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	"galaxytable": "Mozilla/5.0 (Linux; U; Android 2.2; en-gb; GT-P1000 Build/FROYO) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"kitkat":      "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36",

	"firefoxAndroid": "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36",
	"firefoxTablet":  "Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0",
	"firefoxOsX":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0",
	"firefoxUbuntu":  "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:21.0) Gecko/20130331 Firefox/21.0",
	"firefoxWindows": "Mozilla/5.0 (Windows NT 6.2; WOW64; rv:21.0) Gecko/20100101 Firefox/21.0",

	"chromeAndroid": "Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19",
	"chromeTablet":  "Mozilla/5.0 (Linux; Android 4.1.2; Nexus 7 Build/JZ054K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19",
	"chromeMac":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36",
	"chromeUbuntu":  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.11 (KHTML, like Gecko) Ubuntu/11.10 Chromium/27.0.1453.93 Chrome/27.0.1453.93 Safari/537.36",
	"chromeWindow":  "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.94 Safari/537.36",
	"chromeIphone":  "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_4 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) CriOS/27.0.1453.10 Mobile/10B350 Safari/8536.25",

	"operaMac": "Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.9.168 Version/11.52",
	"operaWin": "Opera/9.80 (Windows NT 6.1; WOW64; U; en) Presto/2.10.229 Version/11.62",

	"Blackberry": "Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML, like Gecko) Version/7.2.1.0 Safari/536.2+",

	"safariMac":    "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_6; en-US) AppleWebKit/533.20.25 (KHTML, like Gecko) Version/5.0.4 Safari/533.20.27",
	"safariWin":    "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/533.20.25 (KHTML, like Gecko) Version/5.0.4 Safari/533.20.27",
	"safariIpad":   "Mozilla/5.0 (iPad; CPU OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
	"safariIphone": "Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
}

var SortTestData = [...]string{
	"BingBot",
	"BaiduSpider",
	"GoogleBot",
	// "YahooChina",
	// "YodaoBot",
	// "YoudaoBot",
	// "MSNBot",
	// "SousouSpider",
	// "YahooSeeker",
	// "chinasospider",
	///////////////////////////////////////////////////////////////
	// "ie6",
	// "ie7",
	// "ie8",
	// "ie9",
	// "ie10",
	// "ie11",

	"ipad",
	"iphone",
	"ipod",

	"wphone7",
	"wphone7_5",
	"wphone8",

	"nexus7",
	"galaxys3",
	"galaxytable",
	"kitkat",

	"firefoxAndroid",
	"firefoxTablet",
	"firefoxOsX",
	"firefoxUbuntu",
	"firefoxWindows",

	"chromeAndroid",
	"chromeTablet",
	"chromeMac",
	"chromeUbuntu",
	"chromeWindow",
	"chromeIphone",

	"operaMac",
	"operaWin",

	"Blackberry",

	"safariMac",
	"safariWin",
	"safariIpad",
	"safariIphone",
}

func TestTrimSpace(t *testing.T) {
	f, _ := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer f.Close()
	for _, k := range SortTestData {
		v := TestData[k]
		agent := ParseByString(v)
		f.WriteString(fmt.Sprintf("%#v\n", agent))
	}
}
