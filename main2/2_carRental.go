package main

import "fmt"

type car struct{
   color string;
   price int;
}

var cars [] car

func addCar(c car){
   cars = append (cars, c)
}

func removeCar(index int) (updated[] car){
   updated = append (cars[:index], cars[index+1:]...) 
   return updated
}

func main() {
   objCar1  :=car{"green", 100}
   objCar2 :=car{"red", 200}
   objCar3 :=car {"blue", 300}
      addCar(objCar1)
      addCar(objCar2)
      addCar(objCar3)
      
      fmt.Println("print all cars: ")
      for _, aCar := range cars{
         fmt.Println(aCar)
      }
      cars = removeCar(1)
      fmt.Println("cars after removal:")
      for _, aCar := range cars{
         fmt.Println(aCar)
      }
 }
 


