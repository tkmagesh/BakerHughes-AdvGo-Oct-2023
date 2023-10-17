package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	dataWg := &sync.WaitGroup{}

	dataWg.Add(1)
	go source("data1.dat", dataCh, dataWg)
	dataWg.Add(1)
	go source("data2.dat", dataCh, dataWg)

	processWg := &sync.WaitGroup{}
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenSumCh := make(chan int)
	oddSumCh := make(chan int)

	processWg.Add(1)
	go splitter(dataCh, evenCh, oddCh, processWg)

	processWg.Add(1)
	go sum(evenCh, evenSumCh, processWg)

	processWg.Add(1)
	go sum(oddCh, oddSumCh, processWg)

	processWg.Add(1)
	go merger(evenSumCh, oddSumCh, processWg)

	dataWg.Wait()
	close(dataCh)

	processWg.Wait()
	fmt.Println("Done")
}

func source(fileName string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if no, err := strconv.Atoi(txt); err == nil {
			ch <- no
		}
	}
}

func splitter(dataCh chan int, evenCh chan int, oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(evenCh)
	defer close(oddCh)
	for val := range dataCh {
		if val%2 == 0 {
			evenCh <- val
		} else {
			oddCh <- val
		}
	}
}

func sum(ch chan int, resultCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for val := range ch {
		result += val
	}
	resultCh <- result
}

func merger(evenSumCh, oddSumCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for i := 0; i < 2; i++ {
		select {
		case evenSum := <-evenSumCh:
			fmt.Fprintf(file, "Even Total : %d\n", evenSum)
		case oddSum := <-oddSumCh:
			fmt.Fprintf(file, "Odd Total : %d\n", oddSum)
		}
	}
}
