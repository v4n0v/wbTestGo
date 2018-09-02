package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
	. "wbTestGo/constants"
)
//var locker sync.RWMutex

type Counter struct {
	ChanelNumber int
	Value        int
}

func (c Counter) count(wg *sync.WaitGroup) {
	for {
		c.Value++
		time.Sleep(1 * (time.Second / 2))
		writeFile(wg, fmt.Sprintf(MessageCount, c.ChanelNumber, c.Value))

		if c.Value == 10 {
			return
		}
	}
}

func main() {
	//createFile()
	wg := new(sync.WaitGroup)

	for i := 0; i < 11; i++ {
		go Print("Hello", wg)
	}

	//c1:= Counter{1, 0}
	//c2:= Counter{2, 0}
	//
	//go c1.count(wg)
	//go c2.count(wg)

	wg.Wait()

}
func Print(s string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	println(s)
}

func createFile() {
	err := ioutil.WriteFile(FileName, nil, 0777)
	// Обработка ошибки
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

func writeFile(wg *sync.WaitGroup, txt string) {
	wg.Add(1)
	defer wg.Done()
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
