package service

import "tat_gogogo/utilities/signal"

/*
close is close server procedure.
*/
func close() {
	closeLog()
}

/*
closeLog is close log open file.
*/
func closeLog() {
	<-signal.LogSignal
}
