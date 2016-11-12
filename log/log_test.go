package log

import (
	logg "log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	Printf("Print: foo\n")
	Print("Print: foo")

	SetLevel(Ldebug)

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Errorf("Error: foo")
	Error("Error: foo")

	SetLevel(Lerror)

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Errorf("Error: foo")
	Error("Error: foo")
}

func BenchmarkLogFileChan(b *testing.B) {
	log := New(LogOption{
		Flag:       LAsync | Ldate | Ltime | Lshortfile,
		LogDir:     "testdata",
		ChannelLen: 1000,
	})

	for i := 0; i < b.N; i++ {
		log.Print("testing this is a testing about benchmark")
	}
	log.WaitFlush()
}

func BenchmarkLogFile(b *testing.B) {
	f, _ := os.OpenFile("testdata/onlyfile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log := New(LogOption{
		Out:        f,
		Flag:       Ldate | Ltime | Lshortfile,
		LogDir:     "testdata",
		ChannelLen: 1000,
	})

	for i := 0; i < b.N; i++ {
		log.Print("testing this is a testing about benchmark")
	}
	log.WaitFlush()
}

func BenchmarkStandardFile(b *testing.B) {
	f, _ := os.OpenFile("testdata/logfile.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	log := logg.New(f, "", logg.LstdFlags)
	for i := 0; i < b.N; i++ {
		log.Print("testing this is a testing about benchmark")
	}
}

func BenchmarkLogFileChanMillion(b *testing.B) {
	log := New(LogOption{
		Flag:       LAsync | Ldate | Ltime | Lshortfile,
		LogDir:     "testdata",
		ChannelLen: 1000,
	})
	b.N = 1000000
	for i := 0; i < b.N; i++ {
		log.Print("testing this is a testing about benchmark")
	}
	log.WaitFlush()
}

func BenchmarkLogFileMillion(b *testing.B) {
	f, _ := os.OpenFile("testdata/onlyfilemillion.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log := New(LogOption{
		Out:        f,
		Flag:       Ldate | Ltime | Lshortfile,
		LogDir:     "testdata",
		ChannelLen: 1000,
	})
	b.N = 1000000
	for i := 0; i < b.N; i++ {
		log.Print("testing this is a testing about benchmark")
	}
	log.WaitFlush()
}

func BenchmarkStandardFileMillion(b *testing.B) {
	f, _ := os.OpenFile("testdata/logfilemillion.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	log := logg.New(f, "", logg.LstdFlags)
	b.N = 1000000
	for i := 0; i < b.N; i++ {
		log.Print("testing this is a testing about benchmark")
	}
}
