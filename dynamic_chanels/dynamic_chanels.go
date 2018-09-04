package dynamic_chanels

import (
	"fmt"
	"io/ioutil"
	"time"
	. "wbTestGo/constants"
)

type Counter struct {
	ChanelNumber int
	Value        int
	OutChan      chan int
}

func (c *Counter) count() {
	for {
		time.Sleep(1 * (time.Second / 2))
		c.OutChan <- c.Value
		c.Value++
	}
}

func (c *Counter) GetOutChan() <-chan int {
	return c.OutChan
}

var cntrs = make(map[int]Counter)

func main() {
	createFile()

	for i := 0; i < Threads; i++ {
		c := Counter{i + 1, 1, make(chan int, 4)}
		cntrs[i] = c
		go c.count()
	}
	beginFileWriting(cntrs)

}
func createFile() {
	err := ioutil.WriteFile(FileName, nil, 0777)
	// Обработка ошибки
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

func beginFileWriting(cs map[int]Counter) {
	var msg string
	completeCount := 0
	for {
		for i, v := range cs {
			countChan := v.GetOutChan()
			select {
			case val := <-countChan:
				msg = fmt.Sprintf(MessageCount, cs[i].ChanelNumber, val)
				writeFile(msg)
				if val == Goal {
					completeCount++
					msg = fmt.Sprintf(MessageFinish, cs[i].ChanelNumber, completeCount)
					writeFile(msg)
				}
			}
		}
		if completeCount == Threads {
			return
		}
	}

}

func writeFile(txt string) {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	oldData := string(data)
	newData := []byte(oldData + "\n" + txt)
	err = ioutil.WriteFile(FileName, newData, 0777)
	println(txt)

	if err != nil {
		fmt.Println(err)
	}
}
