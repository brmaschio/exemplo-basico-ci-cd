package main

import (
  "math"
  "net/http"
  "fmt"
  "strconv"
)

func main() {

	http.HandleFunc("/", index)
  
	http.ListenAndServe(":8000", nil) 

}

func index(w http.ResponseWriter, req *http.Request) {

	var x float64 =0.0001
  
  for i:=1; i<100000000; i++ {
  
      x += math.Sqrt(x)
  
  }

  fmt.Printf("x = %v\n", FloatToString(x))

  message := greeting("Code.education Rocks!!!")

  w.WriteHeader(http.StatusOK)
  
  fmt.Fprintf(w, message)

}

func greeting(str string)string{
	return str 
}

func FloatToString(input_num float64) string {
  return strconv.FormatFloat(input_num, 'f', 6, 64)
}