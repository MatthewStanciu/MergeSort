package main
import "fmt"

var arr []int = []int{35,27,16,76,43,89,38,-2,28,53,99}
var sortedList []int = sort(arr)

func main() {
  fmt.Println(sortedList)
}

func merge(left []int, right []int) []int {
  merged := []int{}
  var leftInd int
  var rightInd int

  for leftInd < len(left) && rightInd < len(right) {
    if left[leftInd] <= right[rightInd] {
      merged = append(merged, left[leftInd])
      leftInd++
    } else {
      merged = append(merged, right[rightInd])
      rightInd++
    }
  }

  merged = append(merged, left[leftInd:]...)
  merged = append(merged, right[rightInd:]...)
  return merged
}

func sort(arr []int) []int {
  if len(arr) <= 1 {return arr}
  mid := len(arr) / 2
  left := sort(arr[:mid])
  right := sort(arr[mid:])

  return merge(left, right)
}
