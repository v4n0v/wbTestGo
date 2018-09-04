package constants

import (
	"time"
)

const (
	FileName      = "file.txt"
	MessageCount  = "Поток %v досчитал до %v"
	MessageFinish = "Поток %v закончил %vм"

	Threads  = 5
	Goal     = 10
	Interval = time.Second
)
