package service

import "tat_gogogo/glob/signal"

func close() {
	closeLog()
}

func closeLog() {
	<-signal.LogSignal
}
