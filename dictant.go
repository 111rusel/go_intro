git checkout -b nameBranch

git chekout master
git merge test

git add .

git commit -m "lalala"

git push

git pull 

git clone ссылка


func main() {}

package main

var name string

namee := ""

if got != want{

}

+, -, *, /, %, ^

==, !=, <, >, >=, <=

:=, =

func twoParametrs(one string, two int) string{
	return ""
}

func TestCreate(t *testing.T){
	t.Run("testCreate2", func(t *testing.T){
		// arange 
		want := 10
		// act
		got := Create(4, 6)
		// assert

		if want != got {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

arr := [5]int{1,2,3,4,5}
arr2 := []int{1,2}
arr1 := len(arr)
arr2 = append(arr2, 3,4,5,6)

for i := 0; i < 5: i++{

}

for i := range arr2{

}

for _, elem := range arr2{

}

for i, elem := range arr2{

} 

const A = 1

import (
	"github.com/stretchr/testify"
	"fmt"
)

switch A (
case 2: 
	return 2
case 4:
	return 4
default: 
	return 0
)