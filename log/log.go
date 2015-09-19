package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	Ldebug = iota
	Linfo
	Lerror
	Lfatal
)

const (
	Gray = uint8(iota + 90)
	Red
	Green
	Yellow
	Blue
	Magenta
	//NRed      = uint8(31) // Normal
	EndColor = "\033[0m"
)

var levels = []string{
	"DEBUG",
	"INFO",
	"ERROR",
	"FATAL",
}

type Logger struct {
	mu         sync.Mutex
	Level      int
	statsFile  *os.File
	out        io.Writer
	buf        bytes.Buffer
	stats      map[int]map[string]int
	isNewStats bool
}

// getColorLevel returns colored level string by given level.
func getColorLevel(level string) string {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return fmt.Sprintf("\033[%dm%s\033[0m", Green, level)
	case "INFO":
		return fmt.Sprintf("\033[%dm%s\033[0m", Blue, level)
	case "ERROR":
		return fmt.Sprintf("\033[%dm%s\033[0m", Yellow, level)
	case "FATAL":
		return fmt.Sprintf("\033[%dm%s\033[0m", Red, level)
	default:
		return level
	}
}

func New(out io.Writer, statsFile *os.File) *Logger {
	var l = Logger{out: out, statsFile: statsFile}
	l.stats = make(map[int]map[string]int)
	return &l
}

var Std = New(os.Stdout, nil)

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

func (l *Logger) Output(lvl int, calldepth int, content string) error {
	if lvl < l.Level {
		return nil
	}

	now := time.Now()
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(calldepth)
	if !ok {
		return nil
	}

	module, shortfile := splitOf(file)
	year, month, day := now.Date()
	dt := fmt.Sprintf("%04d/%02d/%02d", year, month, day)
	hour, min, sec := now.Clock()
	msec := now.Nanosecond() / 1e3
	ct := fmt.Sprintf("%02d:%02d:%02d:%d", hour, min, sec, msec)
	s := fmt.Sprintf("%s, %s, %s, %s, %s, %s", dt, ct, getColorLevel(levels[lvl]), module, fmt.Sprintf("%s:%d", shortfile, line), content)

	l.mu.Lock()
	defer l.mu.Unlock()

	l.buf.Reset()
	l.buf.WriteString(s)
	if s[len(s)-1] != '\n' {
		l.buf.WriteByte('\n')
	}

	_, err := l.out.Write(l.buf.Bytes())
	if err != nil {
		return err
	}

	if lvl >= Lerror && l.statsFile != nil {
		l.isNewStats = true
		key := fmt.Sprintf("%s:%s:%d", module, shortfile, line)
		_, found := l.stats[lvl]
		if !found {
			l.stats[lvl] = make(map[string]int)
			l.stats[lvl][key] = 1
		} else {
			_, found := l.stats[lvl][key]
			if !found {
				l.stats[lvl][key] = 1
			} else {
				l.stats[lvl][key]++
			}
		}

		l.buf.Reset()
		for lvl, _ := range l.stats {
			l.buf.WriteString(fmt.Sprintf("#%s\n", levels[lvl]))
			for k, v := range l.stats[lvl] {
				var sbuf bytes.Buffer
				sbuf.WriteString(k)
				nspace := 64 - sbuf.Len()
				for i := 0; i < nspace; i++ {
					sbuf.WriteByte(' ')
				}

				sbuf.WriteString(strconv.Itoa(v))
				sbuf.WriteByte('\n')
				l.buf.Write(sbuf.Bytes())
			}

			l.buf.WriteByte('\n')
			l.buf.WriteByte('\n')
		}

		l.statsFile.Seek(0, os.SEEK_SET)
		_, err := l.statsFile.Write(l.buf.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}

// print
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Output(Linfo, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Print(v ...interface{}) {
	l.Output(Linfo, 2, fmt.Sprint(v...))
}

func (l *Logger) Println(v ...interface{}) {
	l.Output(Linfo, 2, fmt.Sprintln(v...))
}

// debug
func (l *Logger) Debugf(format string, v ...interface{}) {
	if Ldebug < l.Level {
		return
	}
	l.Output(Ldebug, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	if Ldebug < l.Level {
		return
	}
	l.Output(Ldebug, 2, fmt.Sprintln(v...))
}

// info
func (l *Logger) Infof(format string, v ...interface{}) {
	if Linfo < l.Level {
		return
	}
	l.Output(Linfo, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	if Linfo < l.Level {
		return
	}
	l.Output(Linfo, 2, fmt.Sprintln(v...))
}

// error
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(Lerror, 2, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(Lerror, 2, fmt.Sprintln(v...))
}

// fatal
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(Lfatal, 2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(Lfatal, 2, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Breakpoint() {
	if Ldebug < l.Level {
		return
	}
	l.Output(Ldebug, 3, fmt.Sprintln("breakpoint"))
}

// stats
func (l *Logger) Stats() (stats map[int]map[string]int) {
	l.mu.Lock()
	v := make(map[int]map[string]int)
	for lvl, logs := range l.stats {
		if v[lvl] == nil {
			v[lvl] = make(map[string]int)
		}

		for k, count := range logs {
			v[lvl][k] = count
		}
	}
	l.mu.Unlock()
	l.isNewStats = false
	return v
}

func (l *Logger) IsNewStats() bool {
	return l.isNewStats
}

// set output
func (l *Logger) SetLevel(lvl int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Level = lvl
}

// standard wrapper
func Printf(format string, v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprintf(format, v...))
}

func Print(v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprint(v...))
}

func Println(v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	Std.Output(Ldebug, 2, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	Std.Output(Ldebug, 2, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	Std.Output(Linfo, 2, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	Std.Output(Lerror, 2, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	Std.Output(Lerror, 2, fmt.Sprint(v...))
}

func Stack(v ...interface{}) {
	Std.Output(Lerror, 2, fmt.Sprint(v...)+"\n"+CallerStack())
}

func Fatalf(format string, v ...interface{}) {
	Std.Output(Lfatal, 2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatal(v ...interface{}) {
	Std.Output(Lfatal, 2, fmt.Sprint(v...))
	os.Exit(1)
}

func Breakpoint() {
	Std.Breakpoint()
}

func SetLevel(lvl int) {
	Std.mu.Lock()
	defer Std.mu.Unlock()
	Std.Level = lvl
}

func SetOutput(w io.Writer) {
	Std.mu.Lock()
	defer Std.mu.Unlock()
	Std.out = w
}

func SetStatsFile(file *os.File) {
	Std.mu.Lock()
	defer Std.mu.Unlock()
	Std.statsFile = file
}

func CallerStack() string {
	var caller_str string
	for skip := 2; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip) // 获取调用者的信息
		if !ok {
			break
		}
		func_name := runtime.FuncForPC(pc).Name()
		caller_str += "Func : " + func_name + "\nFile:" + file + ":" + fmt.Sprint(line) + "\n"
	}
	return caller_str
}
