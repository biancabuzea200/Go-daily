package main	

import "fmt"

var cars []car

type car struct {
	color string;
	price int;
}

func addCar(c car){
	cars = append(cars, c)
}

func removeCar(index int) (carsupdated []car){
	carsupdated = append(cars[:index], cars[index+1:]...)
	return carsupdated
	
}

func main() {

   carObj1 :=car {color: "red", price:100}
   carObj2 :=car{color:"blue", price: 200}
   carObj3 :=car{color: "white", price:300}

   addCar(carObj1)
   addCar(carObj2)
   addCar(carObj3)

  

   fmt.Println("print all cars: ")
   for _ , aCar := range cars{
      fmt.Println(aCar)
   }
   cars = removeCar(1)
   fmt.Println("print after removal")
   for _ , aCar := range cars{
      fmt.Println(aCar)
   }


	
}
