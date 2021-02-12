package main

import(
   "fmt"
)

func binarySearch(value int, collection []int) int {
   begin := 0
   end := len(collection)

   for begin != end {
      // disgusting boundary check for adjacent elements
      // FIXME
      if end - begin == 1 {
         if collection[end-1] == value {
            return end-1
         } else if collection[begin] == value {
            return begin
         } else {
            return -1
         }
      } else if value < collection[(end-begin)/2] {
         end -= (end-begin)/2
      } else if value > collection[(end-begin)/2] {
         begin += (end-begin)/2
      } else if collection[(end-begin)/2] == value {
         return (begin+end)/2
      }
   }
   return -1
}

func applyTest(expected, searchValue  int, inputCollection []int) bool {
   return expected == binarySearch(searchValue, inputCollection)
}

func main() {
   fmt.Println("test case: 0", applyTest(-1, 3, []int{}))
   fmt.Println("test case: 1", applyTest(-1, 3, []int{1}))
   fmt.Println("test case: 2", applyTest(0,  1, []int{1}))

   fmt.Println("test case: 3", applyTest(0,  1, []int{1, 3, 5}))
   fmt.Println("test case: 4", applyTest(1,  3, []int{1, 3, 5}))
   fmt.Println("test case: 5", applyTest(2,  5, []int{1, 3, 5}))
   fmt.Println("test case: 6", applyTest(-1, 0, []int{1, 3, 5}))
   fmt.Println("test case: 7", applyTest(-1, 2, []int{1, 3, 5}))
   fmt.Println("test case: 8", applyTest(-1, 4, []int{1, 3, 5}))
   fmt.Println("test case: 9", applyTest(-1, 6, []int{1, 3, 5}))

   fmt.Println("test case 10: ", applyTest(0,  1, []int{1, 3, 5, 7}))
   fmt.Println("test case 11: ", applyTest(1,  3, []int{1, 3, 5, 7}))
   fmt.Println("test case 12: ", applyTest(2,  5, []int{1, 3, 5, 7}))
   fmt.Println("test case 13: ", applyTest(3,  7, []int{1, 3, 5, 7}))
   fmt.Println("test case 14: ", applyTest(-1, 0, []int{1, 3, 5, 7}))
   fmt.Println("test case 15: ", applyTest(-1, 2, []int{1, 3, 5, 7}))
   fmt.Println("test case 16: ", applyTest(-1, 4, []int{1, 3, 5, 7}))
   fmt.Println("test case 17: ", applyTest(-1, 6, []int{1, 3, 5, 7}))
   fmt.Println("test case 18: ", applyTest(-1, 8, []int{1, 3, 5, 7}))
}
