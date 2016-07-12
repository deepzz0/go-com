#log
打印记录日志

## Install
``` sh
go get github.com/deepzz0/go-com
```

## Usage
该文件默认使用<code>os.Stdout</code>作为输出，当然你也可以设置成其他<code>Writer</code>。
```
import "github.com/deepzz0/go-com/log"

func main(){
  ...
  log.WaitFlush()
}
```

#### 常量
```
const (
	Ldebug = iota
	Linfo
	Lwarn
	Lerror
	Lfatal
)
```

#### 设置
```
# 设置打印等级，大于等于将会输出。
log.SetLevel(lvl int)
# 设置文件输出目录，不设置则不输出到文件。默认以天切割文件。
log.SetLogDir(dir string)
# 改变默认Writer。默认是标准输出。
log.SetOutput(w io.Writer)
# 设置报警邮件，不设置则不发送邮件。更多信息请查看 mail.go
log.SetEmail(email string)
# obj为文件名前缀和邮件
log.SetObj(obj string)
```

下面你可以通过这些方法记录日志：
``` 
func Printf(format string, v ...interface{})

func Print(v ...interface{})

func Debugf(format string, v ...interface{})

func Debug(v ...interface{})

func Infof(format string, v ...interface{})

func Info(v ...interface{})

func Warnf(format string, v ...interface{})

func Warn(v ...interface{})

func Errorf(format string, v ...interface{})

func Error(v ...interface{})

func Stack(v ...interface{})

func Fatalf(format string, v ...interface{})

func Fatal(v ...interface{})
```
