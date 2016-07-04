package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	Ldebug = iota
	Linfo
	Lwarn
	Lerror
	Lfatal

	ch = 1000
)

var levels = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}

type Logger struct {
	mu     sync.Mutex
	obj    string
	level  int
	out    io.Writer
	in     chan string
	dir    string
	emails []string
}

func New(out io.Writer) *Logger {
	wd, _ := os.Getwd()
	tmp := strings.Split(wd, "/")
	logger := &Logger{obj: tmp[len(tmp)-1], out: out, in: make(chan string, ch)}
	go logger.receive()
	return logger
}

func (l *Logger) receive() {
	today := time.Now()
	var file *os.File
	var err error
	for {
		str := <-l.in
		if l.dir != "" {
			if file == nil || today.Day() != time.Now().Day() {
				today = time.Now()
				file, err = os.OpenFile(fmt.Sprintf("%s/%s_%s.log", l.dir, l.obj, today.Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
				if err != nil {
					panic(err)
				}
			}
			file.WriteString(str)
		}
		l.out.Write([]byte(str))
	}
}

func (l *Logger) Output(lvl int, calldepth int, content string) error {
	if lvl < l.level {
		return nil
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return nil
	}

	module, shortfile := splitOf(file)
	// log format: date, time(hour:minute:second:microsecond), level, module, shortfile:line, <content>
	year, month, day := now.Date()
	dt := fmt.Sprintf("%04d/%02d/%02d", year, month, day)
	hour, min, sec := now.Clock()
	msec := now.Nanosecond() / 1e3
	ct := fmt.Sprintf("%02d:%02d:%02d:%d", hour, min, sec, msec)
	s := fmt.Sprintf("%s, %s, %s, %s, %s, %s", dt, ct, getColorLevel(levels[lvl]), module, fmt.Sprintf("%s:%d", shortfile, line), content)

	if s[len(s)-1] != '\n' {
		s += "\n"
	}

	if len(l.emails) != 0 && lvl >= Lwarn {
		go sendMail(l.obj, s, l.emails)
	}

	l.in <- s
	return nil
}

func (l *Logger) WaitFlush() {
	for {
		if len(l.in) > 0 {
			time.Sleep(time.Millisecond * 50)
		} else {
			break
		}
	}
}

// print
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Output(Linfo, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Print(v ...interface{}) {
	l.Output(Linfo, 2, fmt.Sprintf(smartFormat(v...), v...))
}

// debug
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(Ldebug, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.Output(Ldebug, 2, fmt.Sprintf(smartFormat(v...), v...))
}

// info
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(Linfo, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(Linfo, 2, fmt.Sprintf(smartFormat(v...), v...))
}

// warn
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(Lwarn, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.Output(Lwarn, 2, fmt.Sprintf(smartFormat(v...), v...))
}

// error
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(Lerror, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(Lerror, 2, fmt.Sprintf(smartFormat(v...), v...))
}

// fatal
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(Lfatal, 2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(Lfatal, 2, fmt.Sprintf(smartFormat(v...), v...))
	os.Exit(1)
}

func (l *Logger) Breakpoint() {
	l.Output(Ldebug, 3, fmt.Sprintln("breakpoint"))
}

func (l *Logger) SetLogDir(dir string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.dir = dir
}

func (l *Logger) SetObj(obj string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.obj = obj
}

func (l *Logger) SetOutput(out io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = out
}

func (l *Logger) SetLevel(lvl int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = lvl
}

func (l *Logger) SetEmail(v string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.emails = append(l.emails, v)
}

// standard wrapper
var Std = New(os.Stdout)

func Printf(format string, v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprintf(format, v...))
}

func Print(v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprintf(smartFormat(v...), v...))
}

func Debugf(format string, v ...interface{}) {
	Std.Output(Ldebug, 2, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	Std.Output(Ldebug, 2, fmt.Sprintf(smartFormat(v...), v...))
}

func Infof(format string, v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprintf(smartFormat(v...), v...))
}

func Warnf(format string, v ...interface{}) {
	Std.Output(Lwarn, 2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	Std.Output(Lwarn, 2, fmt.Sprintf(smartFormat(v...), v...))
}

func Errorf(format string, v ...interface{}) {
	body := fmt.Sprintf(format, v...)
	Std.Output(Lerror, 2, body+"\n"+CallerStack())
}

func Error(v ...interface{}) {
	body := fmt.Sprintf(smartFormat(v...), v...)
	Std.Output(Lerror, 2, body+"\n"+CallerStack())
}

func Stack(v ...interface{}) {
	Std.Output(Lerror, 2, fmt.Sprint(v...)+"\n"+CallerStack())
}

func Fatalf(format string, v ...interface{}) {
	Std.Output(Lfatal, 2, fmt.Sprintf(format, v...))
	Std.Output(Lfatal, 2, CallerStack())
	os.Exit(1)
}

func Fatal(v ...interface{}) {
	Std.Output(Lfatal, 2, fmt.Sprintf(smartFormat(v...), v...))
	Std.Output(Lfatal, 2, CallerStack())
	os.Exit(1)
}

func WaitFlush() {
	Std.WaitFlush()
}

func Breakpoint() {
	Std.Breakpoint()
}

func SetLevel(lvl int) {
	Std.SetLevel(lvl)
}

func SetLogDir(dir string) {
	Std.SetLogDir(dir)
}

func SetOutput(w io.Writer) {
	Std.SetOutput(w)
}

func SetEmail(email string) {
	Std.SetEmail(email)
}

func SetObj(obj string) {
	Std.SetObj(obj)
}

///////////////////////////////////////////////////////////////////////////////////////////
func smartFormat(v ...interface{}) string {
	format := ""
	for i := 0; i < len(v); i++ {
		format += " %v"
	}
	format += "\n"
	return format
}

func splitOf(file string) (module string, shortfile string) {
	module = "_unknown_"
	pos := strings.LastIndex(file, "/")
	shortfile = file[pos+1:]
	if pos != -1 {
		pos1 := strings.LastIndex(file[:pos], "/src/")
		if pos1 != -1 {
			module = file[pos1+5 : pos]
		}
	}
	return
}

const (
	Gray = uint8(iota + 90)
	Red
	Green
	Yellow
	Blue
	Magenta
	EndColor = "\033[0m"
)

// getColorLevel returns colored level string by given level.
func getColorLevel(level string) string {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return fmt.Sprintf("\033[%dm%s\033[0m", Green, level)
	case "INFO":
		return fmt.Sprintf("\033[%dm%s\033[0m", Blue, level)
	case "WARN":
		return fmt.Sprintf("\033[%dm%s\033[0m", Magenta, level)
	case "ERROR":
		return fmt.Sprintf("\033[%dm%s\033[0m", Yellow, level)
	case "FATAL":
		return fmt.Sprintf("\033[%dm%s\033[0m", Red, level)
	default:
		return level
	}
}

func CallerStack() string {
	var caller_str string
	for skip := 2; ; skip++ {
		// 获取调用者的信息
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		func_name := runtime.FuncForPC(pc).Name()
		caller_str += "Func : " + func_name + "\nFile:" + file + ":" + fmt.Sprint(line) + "\n"
	}
	return caller_str
}
