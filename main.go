package main

// import "runtime/trace"

//tipagem forte
//não roda se tiver váriavel sem usar

// func soma (x int, y int) int{
// 	return x + y
// }

// func soma (x int, y int) (int,bool){
// 	if x> 10 {
// 		return x + y, true
// 	}
// 	return x + y, false

// }

// func main ()  {

// 	//a:= "Hello World" declara e atribui a váriavel

// 	// println(soma(1,2))
// 	resultado, status := soma(10,20)
// }

import (
	"encoding/json"
	"fmt"      // "fmt" formatação
	"net/http" // servidor web
	"time"
)

//Struct
//C maiunculo Púbico
// c minusculo privado

type Course struct{
	Name          string `json:"course"`;
	Description   string `json:"description"`;
	Price         int    `json:"price"`
}

//metodo

func (c Course) GetFullInfo() string{
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func counter() {
	for i :=0; i <10; i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}


func worker(workerId int, data chan int){

	for x:= range data {
		fmt.Printf("Worker %d received %d\n ", workerId, x)
		time.Sleep(time.Second)
	}
}

//go routine - cria uma thread gerenciada pelo runtime da go 
//cada thread 200kb

// go - na frente significa uma nova thread
//T1 - channel 
func main() {
	//threads
	go counter() //T1 
	go counter() // T2
	counter()//T3 roda elas em paralelo(000, 111, 222 etc)

	//chanel - canal de comunicação entre as threads
	channel := make (chan int)
	go worker(1, channel)//T2
	go worker(2, channel)

	for i := 0; i <10; i++ {
		channel <- i
	}

	//output - Worker 1 received 1
	//output - Worker 1 received 2
	//output - Worker 1 received 3....


	
	// fmt.Println(course.GetFullInfo())

	// fmt.Fprintf()
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request ){
	course := Course {
		Name:  "Golang",
	    Description:   "Golang Course",
	    Price:      100,
	}

	json.NewEncoder(w).Encode(course)
}