package main

import (
	"fmt"
	"io/ioutil"
	"time"
	. "wbTestGo/constants"
)

type Counter struct {
	Value   int
	OutChan chan int
}

func (c *Counter) count() {
	for {
		time.Sleep(1 * (time.Second / 2))
		c.OutChan <- c.Value
		c.Value++
	}
}

func (c *Counter) getOutChan() <-chan int {
	return c.OutChan
}

func main() {
	createFile()
	c1 := Counter{
		1,
		make(chan int, 4),
	}
	c2 := Counter{
		1,
		make(chan int, 4),
	}

	go c1.count()
	go c2.count()

	beginFileWriting(c1, c2)

	//printFile(FileName)
}
func createFile() {
	err := ioutil.WriteFile(FileName, nil, 0777)
	// Обработка ошибки
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

func beginFileWriting(c1 Counter, c2 Counter) {
	var msg string
	countChan1 := c1.getOutChan()
	countChan2 := c2.getOutChan()
	completeCount := 0
	for {
		select {
		case a := <-countChan1:
			if a <= 10 {
				msg = fmt.Sprintf(MessageCount, 1, a)
				writeFile(msg)
				if a == 10 {
					completeCount++
					msg = fmt.Sprintf(MessageFinish, 1, completeCount)
					writeFile(msg)
				}
			}

		case b := <-countChan2:
			if b <= 10 {
				msg = fmt.Sprintf(MessageCount, 2, b)
				writeFile(msg)
				if b == 10 {
					completeCount++
					msg = fmt.Sprintf(MessageFinish, 2, completeCount)
					writeFile(msg)

				}
			}
		}
		if completeCount == 2 {
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
