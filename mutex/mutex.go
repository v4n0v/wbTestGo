package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
	. "wbTestGo/constants"
)

type Counter struct {
	ChanelNumber int
	Value        int
}

var fileMutex sync.Mutex
var countMutex sync.Mutex
var finishCount = 0

func countComplete(c Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	countMutex.Lock()
	defer countMutex.Unlock()
	finishCount++
	writeFile(fmt.Sprintf(MessageFinish, c.ChanelNumber, finishCount))
}

func main() {
	wg := new(sync.WaitGroup)
	createFile()
	for i := 0; i < Threads; i++ {
		c := Counter{i + 1, 0}
		wg.Add(1)
		go c.StartCount(wg)
	}
	wg.Wait()

}

func (c Counter) StartCount(wg *sync.WaitGroup) {
	for i := 0; i < Goal; i++ {
		time.Sleep(Interval)
		writeFile(fmt.Sprintf(MessageCount, c.ChanelNumber, i+1))
		if i == Goal-1 {
			countComplete(c, wg)
		}
	}
}

func createFile() {
	err := ioutil.WriteFile(FileName, nil, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func writeFile(txt string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()
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
