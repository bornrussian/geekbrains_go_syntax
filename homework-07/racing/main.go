//
// Задача:
//
//  ** Попробуйте смоделировать с помощью ​ sync.WaitGroup гонку автомобилей. Каждый
// автомобиль должен быть представлен отдельной горутиной со случайно устанавливаемой
// скоростью (​ math/rand​ ) и случайным временем готовности к старту. Программа должна
// ожидать готовности всех машин, обеспечивать одновременный старт и фиксировать финиш
// каждой машины.
//

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var carCount int = 5

var isRaceStarted bool = false

type channelMessage struct {
	from    int
	to      int
	message string
}

func main() {
	channelRadio := make(chan channelMessage, carCount)
	mutex := &sync.Mutex{}

	go observer(channelRadio, mutex)

	track := &sync.WaitGroup{}
	for id := 1; id <= carCount; id++ {
		track.Add(1)
		go car(id, track, channelRadio, mutex)
	}

	track.Wait()
	fmt.Printf("%v Гонка завершена.\n", curTime())
}

func car(id int, wg *sync.WaitGroup, chFromCar chan channelMessage, mu *sync.Mutex) {
	defer func() {
		fmt.Printf("%v Машина №%v успешно самоуничтожилась\n", curTime(), id)
		wg.Done()
	}()
	rand.Seed(time.Now().UTC().UnixNano())
	carReadyIn := time.Duration(rand.Intn(10)) * time.Second
	carSpeed := time.Duration(rand.Intn(10)+5) * time.Second
	fmt.Printf("%v Машина №%v будет готова к старту через %v\n", curTime(), id, carReadyIn)
	time.Sleep(carReadyIn)
	chFromCar <- channelMessage{id, 255, "iamready"}
	for {
		mu.Lock()
		itIsTimeToGo:=isRaceStarted
		mu.Unlock()
		if itIsTimeToGo {
			fmt.Printf("%v Машина №%v начала гонку и обещает приехать к финишу через %v...\n", curTime(), id, carSpeed)
			time.Sleep(carSpeed)
			chFromCar <- channelMessage{id, 255, "iamfinished"}
			break
		}
		runtime.Gosched()
	}

}

func observer(chFromCar chan channelMessage, mu *sync.Mutex) {
	var howMuchCarsAreReady int = 0
	var howMuchCarsWereFinished int = 0
	for msg := range chFromCar {
		if msg.message == "iamready" {
			fmt.Printf("%v Сообщение от машины №%v : я готова!\n", curTime(), msg.from)
			howMuchCarsAreReady++
			if howMuchCarsAreReady == carCount {
				fmt.Printf("%v Все готовы! Настартвниманиемарш!!!\n", curTime())
				mu.Lock()
				isRaceStarted = true
				mu.Unlock()
			}
		}
		if msg.message == "iamfinished" {
			fmt.Printf("%v Сообщение от машины №%v : я на финише!\n", curTime(), msg.from)
			howMuchCarsWereFinished++
			if howMuchCarsWereFinished == carCount {
				fmt.Printf("%v Все успешно финишировали.\n", curTime())
				close(chFromCar)
			}
		}
	}
}

func curTime() string {
	currentTime := time.Now()
	return currentTime.Format("15:04:05")
}
