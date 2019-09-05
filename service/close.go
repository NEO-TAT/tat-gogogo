package service

import "tat_gogogo/utilities/signal"

func close() {
	closeLog()
}

func closeLog() {
	<-signal.LogSignal
}
