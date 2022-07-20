package main
import(
  "fmt"
  "strconv"
)
func Atoi(strArr []string) {
  fmt.Printf("Value: %v\nType: %T\n\n", strArr, strArr)
  var intArr = []int{}
  for _,i := range strArr {
    j, err := strconv.Atoi(i)

    if err != nil {
      fmt.Printf("Value: %v\nType: %T\n", intArr, intArr)
      recover()
    }else {
      intArr = append(intArr, j)
    }
  }
  fmt.Printf("Value: %v\nType: %T\n", intArr, intArr)
  for k:= 0; k < 20; k++ {
    fmt.Printf("=-")
  }
  fmt.Printf("\n")
}

func main() {
  var strArr = []string{"1", "2", "3", "4", "5", "6"}
  Atoi(strArr)
  strArr = []string{"10", "20", "30", "40", "50", "60"}
  Atoi(strArr)
  strArr = []string{"100", "200", "300", "400", "500", "600"}
  Atoi(strArr)
  strArr = []string{"4324", "234", "32423", "1231", "505453", "645"}
  Atoi(strArr)
}
