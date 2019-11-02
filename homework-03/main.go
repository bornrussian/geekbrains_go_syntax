//
// Задача:
//
//	1. Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны
//	содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты
//	ли окна, насколько заполнен объем багажника.
//
//	2. Инициализировать несколько экземпляров структур. Применить к ним различные действия.
//	Вывести значения свойств экземпляров в консоль.
//
//	3. * Реализовать очередь. Это структура данных, работающая по принципу FIFO (First Input —
//	first output, или «первым зашел — первым вышел»).
//

package main

import (
	"fmt"
)

type vehicle struct {
	model string
	manufactureYear int
	cargoVolumeMax int
	cargoVolumeCurrent int
	isEngineStarted bool
	isWindowsOpened bool
}

type garage struct {
	cars []vehicle
}

// Добавляем автомобиль в гараж
func addCarToGarage (garage *garage, car vehicle) {
	garage.cars=append(garage.cars,car)
}

// Печатаем содержимое гаража
func printGarageContains (garage *garage) {
	fmt.Printf("В гараже сейчас находятся:\n")
	for _,car:=range garage.cars {
		fmt.Println("  ",car)
	}
}

// Печатаем список авто с запущенным двигателем
func printCarsWithEngineStarted(garage *garage) {
	fmt.Println("Двигатель запущен на:")
	isAnyCar:=false
	for id,car:=range garage.cars {
		if car.isEngineStarted {
			isAnyCar=true
			fmt.Println("  ", garage.cars[id])
		}
	}
	if !isAnyCar {
		fmt.Println("  --nothing found--")
	}
}

// Печатаем авто с максимальным объемом кузова/багажника
func printCarWithBiggestCargo(garage *garage) {
	maxCargo:=-1
	carID:=-1

	for id,car:=range garage.cars {
		if car.cargoVolumeMax>maxCargo {
			maxCargo=car.cargoVolumeMax
			carID=id
		}
	}
	fmt.Println("Самый большой багажник у:")
	if maxCargo<0 {
		fmt.Println("  --nothing found--")
	} else {
		fmt.Println("  ", garage.cars[carID])
	}
}

// Берем из гаража машину ту, которая оказалась там первой
func getCarFromGarageFIFO (garage *garage) vehicle {
	if len(garage.cars) > 0 {
		car := garage.cars[0]
		garage.cars = garage.cars[1:]
		fmt.Println("Из гаража забрали машину по алгоритму FIFO:")
		fmt.Println("  ", car)
		return car
	} else {
		fmt.Println("В гараже не осталось машин.")
		return vehicle{}
	}

}

// Берем из гаража машину ту, которая оказалась там последней
func getCarFromGarageLIFO (garage *garage) vehicle {
	if len(garage.cars) > 0 {
		car := garage.cars[len(garage.cars)-1]
		garage.cars = garage.cars[:len(garage.cars)-1]
		fmt.Println("Из гаража забрали машину по алгоритму LIFO:")
		fmt.Println("  ", car)
		return car
	} else {
		fmt.Println("В гараже не осталось машин.")
		return vehicle{}
	}
}

// Основной код
func main() {
	var theGarage garage

	addCarToGarage(&theGarage, vehicle{
		"Kia Rio",
			2014,
			500,
			10,
			true,
			false,
	})
	addCarToGarage(&theGarage, vehicle{
			"Renault Duster",
			2016,
			408,
			50,
			true,
			true,
	})
	addCarToGarage(&theGarage, vehicle{
			"Nissan X-Trail",
			2012,
			479,
			0,
			false,
			true,
	})
	addCarToGarage(&theGarage, vehicle{
			"Mitsubishi L200",
			2008,
			100500,
			100,
			false,
			false,
	})

	printCarsWithEngineStarted(&theGarage)
	printCarWithBiggestCargo(&theGarage)

	getCarFromGarageFIFO(&theGarage)
	printGarageContains(&theGarage)

	getCarFromGarageLIFO(&theGarage)
	printGarageContains(&theGarage)

	printGarageContains(&theGarage)
}
