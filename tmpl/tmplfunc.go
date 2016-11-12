// Package tmpl provides ...
package tmpl

import (
	"html/template"
	"strings"
	"time"
)

var TplFuncMap = make(template.FuncMap)

func init() {
	TplFuncMap["dateformat"] = DateFormat
	TplFuncMap["str2html"] = Str2html
	TplFuncMap["join"] = StringsJoin
	TplFuncMap["isnotzero"] = IsNotZero
}

func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

// DateFormat takes a time and a layout string and returns a string with the formatted date. Used by the template parser as "dateformat"
func DateFormat(t time.Time, layout string) (datestring string) {
	datestring = t.Format(layout)
	return
}

func StringsJoin(a []string, sep string) string {
	return strings.Join(a, sep)
}

func IsNotZero(t time.Time) bool {
	return !t.IsZero()
}
