package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	Name      string
	OrderedIn string
	Iteration int16
}

var jobscount = 10
var count int16 = 0
var iterationLimit int16 = 3

func main() {
	ticker := time.NewTicker(time.Duration(10 * time.Second)) // Рандом определяется в начале и будет постоянным. Типа если в начале была 1 секунда то так и останется

	timeStart := time.Now()

	var wg sync.WaitGroup
	// wg.Add(jobscount)

	iterationChan := make(chan int16, 10)
	jobChan := make(chan Job, 2) // Добавил пару значений для хранения на всякий случай
	defer close(jobChan)
	var workersChan = make(chan string, 3)

	func() {
		workersChan <- "Rakhat"
		workersChan <- "Eldos"
		workersChan <- "Someone"
	}()
	close(workersChan)

	for w := range workersChan { // Переставил его наверх чтоб успевало запускать слушателя канала jobChan. До этого ticker был вечной функцией которая создавала рутины и ждала отправки значения в канал, пока никто просто не мог принять потому что этот цикл исполнялся быстрее и не мог закончиться не давая очереди worker-ам.
		go func(worker string, ch <-chan Job) {
			for work := range ch {
				time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
				log.Printf("Iteration: %d - Ordered in: %v | worker %s is done with %s", work.Iteration, work.OrderedIn, worker, work.Name)
				wg.Done()
			}
		}(w, jobChan)
	}

outer: // По совету чата использую метки (лейблы).
	for t := range ticker.C { // Бесконечный цикл который нужно закрывать по идее
		count++
		limitCounter(iterationChan, count, iterationLimit)
		for i := 1; i <= jobscount; i++ {
			go func(index int, time time.Time, ch chan<- Job, iterationCount int16) {
				newJob := Job{ // Mutex здесь максимально бесполезен ведь я убрал слайс и теперь проблемы с тем, что 2 го рутины может добавить случайно одинаковые значения отпали.
					Name:      fmt.Sprintf("job %d", index),
					OrderedIn: time.Format("2006-01-02 15:04:05"),
					Iteration: iterationCount,
				}
				// fmt.Println("В канал передаётся значение")
				wg.Add(1)
				ch <- newJob // С jobChan <- jobSlice[len(jobSlice - 1)] было неплохо так проблем. Убрал после ошибок с принятием только одного job jobscount раз внутрь слайса. Решил что зачем мне слайс если канал и так может сразу же передать значения без хранения.
				// fmt.Printf("Значение в канал передалось для работы %d\n", index)
			}(i, t, jobChan, count)
		}
		select {
		case exit := <-iterationChan:
			wg.Wait() // Чтобы при входе в 3 итерацию преждевременно не запустилось
			log.Printf("%v | End of the programm: %v", exit, time.Since(timeStart))
			close(iterationChan)
			ticker.Stop() // Здесь и завершится цикл тикера
			break outer
		default:
			continue
		}
	}
}

func limitCounter(exitChan chan<- int16, counter int16, iterationLimit int16) {
	if counter >= iterationLimit {
		exitChan <- 1
	}
}

/*
	Изучить:
		1. Лейблы
*/
