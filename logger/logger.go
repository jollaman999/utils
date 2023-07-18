package logger

import (
	"fmt"
	"github.com/jollaman999/utils/fileutil"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var Logger *log.Logger

var fpLog *os.File
var err error

const (
	// INFO : Informational messages printed with green.
	INFO = "INFO"
	// DEBUG : Debugging messages printed with teal.
	DEBUG = "DEBUG"
	// WARN : Warning messages printed with yellow.
	WARN = "WARN"
	// ERROR : Error messages printed with red.
	ERROR = "ERROR"
	// CRITICAL : Critical messages printed with magenta.
	CRITICAL = "CRITICAL"
	// NONE : Print normal messages with none of color and date, prefix.
	NONE = "NONE"
)

var (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	magenta = "\033[35m"
	teal    = "\033[36m"
)

var date string
var _time string

func getDateAndTime() {
	now := time.Now()

	year := fmt.Sprintf("%04d", now.Year())
	month := fmt.Sprintf("%02d", now.Month())
	day := fmt.Sprintf("%02d", now.Day())

	hour := fmt.Sprintf("%02d", now.Hour())
	minute := fmt.Sprintf("%02d", now.Minute())
	second := fmt.Sprintf("%02d", now.Second())

	date = year + "/" + month + "/" + day
	_time = hour + ":" + minute + ":" + second
}

func getPrefix(logLevel string) (prefixForPrint string, prefixForWrite string) {
	getDateAndTime()

	switch logLevel {
	case INFO:
		return date + " " + _time + " [ " + green + "INFO" + reset + " ] ", "[ INFO ] "
	case DEBUG:
		return date + " " + _time + " [ " + teal + "DEBUG" + reset + " ] ", "[ DEBUG ] "
	case WARN:
		return date + " " + _time + " [ " + yellow + "WARN" + reset + " ] ", "[ WARN ] "
	case ERROR:
		return date + " " + _time + " [ " + red + "ERROR" + reset + " ] ", "[ ERROR ] "
	case CRITICAL:
		return date + " " + _time + " [ " + magenta + "CRITICAL" + reset + " ] ", "[ CRITICAL ] "
	case NONE:
		fallthrough
	default:
		return "", ""
	}
}

func getCallLocation(printCallLocation bool) string {
	if printCallLocation {
		skip := 3
		if strings.HasSuffix(os.Args[0], ".test") {
			skip = 2
		}
		_, filepath, line, _ := runtime.Caller(skip)
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		filename := strings.Replace(filepath, pwd+"/", "", -1)
		return filename + ":" + strconv.Itoa(line) + ": "
	}

	return ""
}

// Print : Print the log with a colored level without new line
func Print(logLevel string, printCallLocation bool, msg ...interface{}) {
	prefixForPrint, prefixForWrite := getPrefix(logLevel)
	fmt.Print(prefixForPrint + getCallLocation(printCallLocation) + fmt.Sprint(msg...))
	if Logger == nil {
		return
	}

	if logLevel == NONE {
		_, _ = fmt.Fprint(io.Writer(fpLog), prefixForWrite+getCallLocation(printCallLocation)+fmt.Sprint(msg...))
	} else {
		getDateAndTime()
		_, _ = fmt.Fprint(io.Writer(fpLog), date+" "+_time+" [ "+logLevel+" ] "+getCallLocation(printCallLocation)+fmt.Sprint(msg...))
	}
}

// Println : Print the log with a colored level with new line
func Println(logLevel string, printCallLocation bool, msg ...interface{}) {
	prefixForPrint, prefixForWrite := getPrefix(logLevel)
	fmt.Println(prefixForPrint + getCallLocation(printCallLocation) + fmt.Sprint(msg...))
	if Logger == nil {
		return
	}
	if logLevel == NONE {
		_, _ = fmt.Fprintln(io.Writer(fpLog), getCallLocation(printCallLocation)+prefixForWrite+fmt.Sprint(msg...))
	} else {
		Logger.Println(prefixForWrite + getCallLocation(printCallLocation) + fmt.Sprint(msg...))
	}
}

// Printf : Print the formatted log with a colored level
func Printf(logLevel string, printCallLocation bool, format string, a ...any) {
	prefixForPrint, prefixForWrite := getPrefix(logLevel)
	fmt.Printf(prefixForPrint+getCallLocation(printCallLocation)+format, a...)
	if Logger == nil {
		return
	}
	if logLevel == NONE {
		_, _ = fmt.Fprintf(io.Writer(fpLog), prefixForWrite+getCallLocation(printCallLocation)+format, a...)
	} else {
		Logger.Printf(prefixForWrite+getCallLocation(printCallLocation)+format, a...)
	}
}

// Fatal : Print the log with a colored level then exit with return value 1
func Fatal(logLevel string, printCallLocation bool, exitCode int, msg ...interface{}) {
	Print(logLevel, printCallLocation, msg...)
	CloseLogFile()
	os.Exit(exitCode)
}

// Fatalln : Print the log with a colored level with new line then exit with return value 1
func Fatalln(logLevel string, printCallLocation bool, exitCode int, msg ...interface{}) {
	Println(logLevel, printCallLocation, msg...)
	CloseLogFile()
	os.Exit(exitCode)
}

// Fatalf : Print the formatted log with a colored level then exit with return value 1
func Fatalf(logLevel string, printCallLocation bool, exitCode int, format string, a ...any) {
	Printf(logLevel, printCallLocation, format, a...)
	CloseLogFile()
	os.Exit(exitCode)
}

// Panic : Print the log with a colored level then call panic()
func Panic(logLevel string, printCallLocation bool, msg ...interface{}) {
	Print(logLevel, printCallLocation, msg...)
	CloseLogFile()
	panic(fmt.Sprint(msg...))
}

// Panicln : Print the log with a colored level with new line then call panic()
func Panicln(logLevel string, printCallLocation bool, msg ...interface{}) {
	Println(logLevel, printCallLocation, msg...)
	CloseLogFile()
	panic(fmt.Sprintln(msg...))
}

// Panicf : Print the formatted log with a colored level then call panic()
func Panicf(logLevel string, printCallLocation bool, format string, a ...any) {
	Printf(logLevel, printCallLocation, format, a...)
	CloseLogFile()
	panic(fmt.Sprintf(format, a...))
}

// InitLogFile : Initialize log file
func InitLogFile(logPath string, logFileNamePrefix string) error {
	// Create directory if not exist
	if _, err = os.Stat(logPath); os.IsNotExist(err) {
		err = fileutil.CreateDirIfNotExist(logPath)
		if err != nil {
			return err
		}
	}

	now := time.Now()

	year := fmt.Sprintf("%d", now.Year())
	month := fmt.Sprintf("%02d", now.Month())
	day := fmt.Sprintf("%02d", now.Day())

	datePrefix := year + month + day
	fpLog, err = os.OpenFile(logPath+"/"+logFileNamePrefix+"_"+datePrefix+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		Logger = log.New(io.Writer(os.Stdout), "", log.Ldate|log.Ltime)
		return err
	}

	Logger = log.New(io.Writer(fpLog), "", log.Ldate|log.Ltime)

	return nil
}

// CloseLogFile : Close log file
func CloseLogFile() {
	if fpLog != nil {
		_ = fpLog.Close()
	}
}
