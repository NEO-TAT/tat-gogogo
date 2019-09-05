package logs

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"tat_gogogo/utilities/signal"
	"time"

	"github.com/spf13/viper"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

/*
LogInit is init Log file and register componert.
*/
func LogInit() {
	// ----------------------------------------------
	output := map[string][]io.Writer{}
	logLevel := viper.GetInt("log.level")
	file := viper.GetBool("log.output")
	times := time.Now().UTC().Format("2006-01-02")
	os.MkdirAll("log", os.ModePerm)
	// -----------------------------------------------
	defaultPath := fmt.Sprintf("./log/Default_%s.log", times)
	errorPath := fmt.Sprintf("./log/Error_%s.log", times)
	warningPath := fmt.Sprintf("./log/Warning_%s.log", times)
	infoPath := fmt.Sprintf("./log/Info_%s.log", times)
	// -----------------------------------------------
	if viper.GetBool("log.print") {
		output["defaultPath"] = append(output["defaultPath"], os.Stderr)
		output["error"] = append(output["error"], os.Stderr)
		output["warning"] = append(output["warning"], os.Stderr)
		output["info"] = append(output["info"], os.Stderr)
	}
	// -----------------------------------------------[Default log]
	if viper.GetBool("log.default_log") {
		if file {
			defaultSave, _ := os.OpenFile(defaultPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			defer defaultSave.Close()
			output["defaultPath"] = append(output["defaultPath"], defaultSave)
		}
		log.SetOutput(io.MultiWriter(output["defaultPath"]...))
	}
	// -----------------------------------------------[Default set]
	Error = log.New(ioutil.Discard, "", log.LstdFlags)
	Warning = log.New(ioutil.Discard, "", log.LstdFlags)
	Info = log.New(ioutil.Discard, "", log.LstdFlags)
	// -----------------------------------------------[Error log]
	if logLevel >= 1 {
		if file {
			errSave, _ := os.OpenFile(errorPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			defer errSave.Close()
			output["error"] = append(output["error"], errSave)
		}
		Error = log.New(io.MultiWriter(output["error"]...), "【Error】 ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	// -----------------------------------------------[Warning Log]
	if logLevel >= 2 {
		if file {
			warningSave, _ := os.OpenFile(warningPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			defer warningSave.Close()
			output["warning"] = append(output["warning"], warningSave)
		}
		Warning = log.New(io.MultiWriter(output["warning"]...), "【Warning】 ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	// -----------------------------------------------[Info Log]
	if logLevel >= 3 {
		if file {
			infoSave, _ := os.OpenFile(infoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			defer infoSave.Close()
			output["info"] = append(output["info"], infoSave)
		}
		Info = log.New(io.MultiWriter(output["info"]...), "【Info】 ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	// -----------------------------------------------
	signal.LogSignal = make(chan bool)
	signal.LogSignal <- true
}
