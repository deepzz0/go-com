package useragent

import (
	"strings"
)

const (
	///////////////////////////////// OS_PRODUCER //////////////////////////////////
	// Microsoft Corporation
	MICROSOFT = "Microsoft Corporation"
	// Apple Inc.
	APPLE = "Apple Inc."
	// Sun Microsystems, Inc.
	SUN = "Sun Microsystems, Inc."
	// Symbian Ltd.
	SYMBIAN = "Symbian Ltd."
	// Nokia Corporation
	NOKIA = "Nokia Corporation"
	// Research In Motion Limited
	BLACKBERRY = "Research In Motion Limited"
	// Hewlett-Packard Company, previously Palm
	HP = "Hewlett Packard"
	// Sony Ericsson Mobile Communications AB
	SONY_ERICSSON = "Sony Ericsson Mobile Communications AB"
	// Samsung Electronics
	SAMSUNG = "Samsung Electronics"
	// Sony Computer Entertainment, Inc.
	SONY = "Sony Computer Entertainment, Inc."
	// Nintendo
	NINTENDO = "Nintendo"
	// Opera Software ASA
	OPERA = "Opera Software ASA"
	//  Mozilla Foundation
	MOZ = " Mozilla Foundation"
	// Google Inc.
	GOOGLE = "Google Inc."
	// CompuServe Interactive Services, Inc.
	COMPUSERVE = "CompuServe Interactive Services, Inc."
	// Yahoo Inc.
	YAHOO = "Yahoo Inc."
	// AOL LLC.
	AOL = "AOL LLC."
	// Mail.com Media Corporation
	MMC = "Mail.com Media Corporation"
	// Amazon.com, Inc.
	AMAZON = "Amazon.com, Inc."
	// Roku sells home digital media products
	ROKU = "Roku, Inc."
	// Adobe Systems Inc.
	ADOBE = "Adobe Systems Inc."
	// Canonical Ltd.
	CONONICAL = "Canonical Ltd."
	// Unknow or rare manufacturer
	OHTER_PRODUCER = "Other"
	///////////////////////////////// AGENT_TYPE //////////////////////////////////
	// Standard web-browser
	WEB_BROWSER = "Browser"
	// Special web-browser for mobile devices
	MOBILE_BROWSER = "Browser (mobile)"
	// Text only browser like the good old Lynx
	TEXT_BROWSER = "Browser (text only)"
	// Email client like Thunderbird
	EMAIL_CLIENT = "Email Client"
	// Search robot, spider, crawler,...
	ROBOT = "robot"
	// Downloading tools
	TOOL = "Downloading tool"
	// Application
	APP = "Application"
	////////////////////////////////// DEVICE_TYPE //////////////////////////////////
	// Standard desktop or laptop computer
	COMPUTER = "Computer"
	// Mobile phone or similar small mobile device
	MOBILE = "Mobile"
	// Small tablet type computer.
	TABLET = "Tablet"
	// Game console like the Wii or Playstation.
	GAME_CONSOLE = "Game console"
	// Digital media receiver like the Google TV.
	DMR = "Digital media receiver"
	// Miniature device like a smart watch or interactive glasses
	WEARABLE = "Wearable computer"
	// Other or unknow type of device.
	UNKNOWN_DEVICE = "unknown"
	////////////////////////////////// AGENT_RENDERING_ENGINE //////////////////////////////////
	// EdgeHTML is a proprietary layout engine developed for the Microsoft Edge web browser, developed by Microsoft.
	EDGE_HTML = "EdgeHTML"
	// Trident is the the Microsoft layout engine, mainly used by Internet Explorer.
	TRIDENT = "Trident"
	// HTML parsing and rendering engine of Microsoft Office Word, used by some other products of the Office suite instead of Trident.
	WORD = "Microsoft Office Word"
	// Open source and cross platform layout engine, used by Firefox and many other browsers.
	GECKO = "Gecko"
	// Layout engine based on KHTML, used by Safari, Chrome and some other browsers.
	WEBKIT = "Webkit"
	// Proprietary layout engine by Opera Software ASA
	PRESTO = "Presto"
	// Original layout engine of the Mozilla browser and related products. Predecessor of Gecko.
	MOZILLA = "Mozilla"
	// Layout engine of the KDE project
	KHTML = "KHTML"
	// Layout engine developed as part ofthe Chromium project. Fored from WebKit.
	BLINK = "Blink"
	// Other or unknown layout engine.
	OTHER_ENGINE = "Other"
)

var Agent_types = map[string]map[string]string{
	WEB_BROWSER: map[string]string{ // reg-->name
		// IE
		"MSIE [\\d\\.]+":                 "Internet Explorer",
		"Trident/[\\d\\.]+":              "Internet Explorer",
		"IE 11\\.":                       "Internet Explorer",
		"xbox":                           "Xbox",
		"Edge/[\\d+\\.]+":                "Microsoft Edge",
		"(?:CriOS|CrMo|Chrome)\\/[\\d]+": "Chrome",
		"Firefox/[\\d\\.]+":              "Firefox",
		"Version/[\\d\\.]+":              "Safari",
		"Opera/[\\d\\.]+":                "Opera",
		"Version/[\\d]+\\.":              "Opera",
		"OPR/[\\d]+\\.":                  "Opera",
		"Camino/[\\d\\.]+":               "Camino",
		"Camino/2":                       "Camino",
		"Flock/[\\d\\.]+":                "Flock",
		"Vivaldi/[\\d\\.]+":              "Vivaldi",
		"Mozilla":                        "Mozilla",
		"Moozilla":                       "Moozilla",
	},
	MOBILE_BROWSER: map[string]string{
		"IEMobile/\\w+":            "Mobile",
		"IEMobile [\\d\\.]+":       "Internet Explorer(mobile)",
		"IEMobile/[\\d\\.]+":       "Internet Explorer(mobile)",
		"Edge/[\\d]+":              "Microsoft Edge(Mobile)",
		"Konqueror/[\\d\\.]+":      "Konqueror(mobile)",
		"Opera Mini":               "Opera(mini)",
		"Mobile Safari":            "Opera(mobile)",
		"Dolfin\\/2":               "Samsung(mobile)",
		"FxiOS":                    "Firefox(mobile)",
		"Version/[\\d\\.]+ Mobile": "Safari(mobile)",
	},
	TEXT_BROWSER: map[string]string{},
	EMAIL_CLIENT: map[string]string{
		"MSOffice [\\d]+":          "Outlook",
		"Microsoft Outlook [\\d]+": "Outlook",
		"Outlook-Express/7.0":      "Windows Live Mail",
		"Lotus-Notes/[\\d\\.]+":    "Lotus Notes",
		"Thunderbird/[\\d\\.]+":    "Thunderbird",
	},
	ROBOT: map[string]string{
		"Googlebot":            "Googlebot",
		"Mediapartners-Google": "Mediapartners-Google",
		"Web Preview":          "Web Preview",
		"Applebot":             "Applebot",
		"crawler":              "crawler",
		"Feedfetcher":          "Feedfetcher",
		"Slurp":                "Slurp",
		"Twiceler":             "Twiceler",
		"Nutch":                "Nutch",
		"BecomeBot":            "BecomeBot",
		"bingbot":              "bingbot",
		"BingPreview":          "BingPreview",
		"Google Web Preview":   "Google Web Preview",
		"WordPress.com mShots": "WordPress.com mShots",
		"Seznam":               "Seznam",
		"facebookexternalhit":  "facebookexternalhit",
		"YandexMarket":         "YandexMarket",
		"Teoma":                "Teoma",
		"ThumbSniper":          "ThumbSniper",
		"Phantom":              "Phantom",
		"Googlebot-Mobile":     "Googlebot-Mobile",
		"Baiduspider":          "Baiduspider",
		"Baiduspider-image":    "Baiduspider-image",
		"Baiduspider-sfkr":     "Baiduspider-sfkr",
		"YodaoBot":             "YodaoBot",
		"YoudaoBot":            "YoudaoBot",
		"Sosospider":           "Sosospider",
		"YahooSeeker":          "YahooSeeker",
		"sogou spider":         "sogou spider",
		"msnbot":               "msnbot",
		"msnbot-media":         "msnbot-media",
		"chinasospider":        "chinasospider",
	},
	TOOL: map[string]string{
		"cURL":              "Downloading Tool",
		"wget":              "Downloading Tool",
		"ggpht.com":         "Downloading Tool",
		"Apache-HttpClient": "Downloading Tool",
	},
	APP: map[string]string{
		"iTunes":      "iTunes",
		"MacAppStore": "App Store",
		"AdobeAIR":    "Adobe AIR application",
	},
}

var IE = map[string]string{
	"IE 11.":    "Internet Explorer 11",
	"Trident/7": "Internet Explorer 11",
}
var IEMobile = map[string]string{}

var Outlook = map[string]string{
	"MSOffice 12":          "Outlook 2007",
	"MSOffice 14":          "Outlook 2010",
	"Microsoft Outlook 14": "Outlook 2010",
	"Microsoft Outlook 15": "Outlook 2013",
}

var TypeToName = map[string]func(string) string{
	"Internet Explorer": func(str string) string {
		if strings.HasPrefix(str, "MSIE") {
			return "Internet Explorer " + strings.Split(str, " ")[1]
		}
		return IE[str]
	},
	"Internet Explorer(mobile)": func(str string) string {
		if strings.HasPrefix(str, "IEMobile") {
			return "IE Mobile " + strings.Split(str, "/")[1]
		}
		return IEMobile[str]
	},
	"Microsoft Edge(Mobile)": func(str string) string {
		return "Microsoft Edge Mobile"
	},
	"Xbox": func(str string) string {
		return "Xbox"
	},
	"Outlook": func(str string) string {
		return Outlook[str]
	},
	"Chrome": func(str string) string {
		return strings.Replace(str, "/", " ", 1)
	},
	"Firefox": func(str string) string {
		return strings.Replace(str, "/", " ", 1)
	},
	"Firefox(mobile)": func(str string) string {
		return "Firefox Mobile (iOS)"
	},
	"Safari(mobile)": func(str string) string {
		return strings.Replace(str, "/", " ", 1)
	},
	"Mobile": func(str string) string {
		return strings.Replace(str, "/", " ", 1)
	},
	"Safari": func(str string) string {
		return strings.Replace(str, "/", " ", 1)
	},
	"Opera": func(str string) string {
		ver := strings.Split(str, "/")
		if len(ver) < 2 {
			return "Opera unknown"
		}
		return "Opera " + ver[1]
	},
	"Opera(mini)": func(str string) string {
		return "Opera Mini"
	},
	"Opera(mobile)": func(str string) string {
		return "Opera Mobile"
	},
	"Samsung(mobile)": func(str string) string {
		return "Samsung Dolphin 2"
	},
	"Konqueror(mobile)": func(str string) string {
		return "Konqueror"
	},
	"iTunes": func(str string) string {
		return "iTunes"
	},
	"MacAppStore": func(str string) string {
		return "App Store"
	},
	"AdobeAIR": func(str string) string {
		return "Adobe AIR application"
	},
	"Camino": func(str string) string {
		return strings.Replace(str, "/", " ", 1)
	},
	"Flock": func(str string) string {
		return "Flock"
	},
	"Lotus Notes": func(str string) string {
		return "Lotus Notes"
	},
	"Thunderbird": func(str string) string {
		return strings.Replace(str, "/", " ", 1)
	},
	"Vivaldi": func(str string) string {
		return "Vivaldi"
	},
	"Mozilla": func(str string) string {
		return "Mozilla"
	},
	"Moozilla": func(str string) string {
		return "Mozilla"
	},
}

var TypeToRenderingEngine = map[string]map[string]string{
	WEB_BROWSER: map[string]string{
		"Internet Explorer": TRIDENT,
		"Xbox":              TRIDENT,
		"Microsoft Edge":    EDGE_HTML,
		"Chrome":            WEBKIT,
		"Firefox":           GECKO,
		"Safari":            WEBKIT,
		"Opera":             WEBKIT,
		"Camino":            GECKO,
		"Flock":             GECKO,
		"Vivaldi":           BLINK,
		"Mozilla":           OTHER_ENGINE,
	},
	MOBILE_BROWSER: map[string]string{
		"Internet Explorer(mobile)": TRIDENT,
		"Microsoft Edge(Mobile)":    EDGE_HTML,
		"Firefox(mobile)":           WEBKIT,
		"Safari(mobile)":            WEBKIT,
		"Opera(mini)":               PRESTO,
		"Opera(mobile)":             BLINK,
		"Samsung(mobile)":           WEBKIT,
		"Konqueror(mobile)":         OTHER_ENGINE,
		"Mobile":                    WEBKIT,
	},
	EMAIL_CLIENT: map[string]string{
		"Outlook":           WORD,
		"Windows Live Mail": TRIDENT,
		"Lotus Notes":       OTHER_ENGINE,
		"Thunderbird":       GECKO,
	},
	APP: map[string]string{
		"iTunes":      WEBKIT,
		"MacAppStore": WEBKIT,
		"AdobeAIR":    WEBKIT,
	},
}

var TypeToProducer = map[string]string{
	"Outlook":                   MICROSOFT,
	"Windows Live Mail":         MICROSOFT,
	"Internet Explorer":         MICROSOFT,
	"Internet Explorer(mobile)": MICROSOFT,
	"Microsoft Edge(Mobile)":    MICROSOFT,
	"Xbox":              MICROSOFT,
	"Chrome":            GOOGLE,
	"Firefox":           MOZ,
	"Firefox(mobile)":   MOZ,
	"Safari(mobile)":    APPLE,
	"iTunes":            APPLE,
	"MacAppStore":       APPLE,
	"Opera":             OPERA,
	"Opera(mini)":       OPERA,
	"Opera(mobile)":     OPERA,
	"Samsung(mobile)":   SAMSUNG,
	"Konqueror(mobile)": OHTER_PRODUCER,
	"AdobeAIR":          ADOBE,
	"Lotus Notes":       OHTER_PRODUCER,
	"Camino":            OHTER_PRODUCER,
	"Flock":             OHTER_PRODUCER,
	"Thunderbird":       MOZ,
	"Vivaldi":           OHTER_PRODUCER,
	"Mozilla":           MOZ,
}
