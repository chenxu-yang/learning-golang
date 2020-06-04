package array_and_slice

import (
	"reflect"
	"testing"
)

func Sum(numbers []int) int {
	//sum := 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}
	//return sum
	sum:=0
	for _,number:=range numbers{
		sum+=number
	}
	return sum
}
//func SumAll(numbersToSum ...[]int)(sums []int){
//	length:=len(numbersToSum)
//	sums=make([]int,length)//make 创建指定长度的切片
//	for i,numbers := range numbersToSum{
//		sums[i]=Sum(numbers)
//	}
//	return
//}
func SumAll(numberToSum ...[]int)[]int{
	var sums []int
	for _,numbers:=range numberToSum{
		sums = append(sums, Sum(numbers))
	}
	return sums
}
func SumAllTails(numberToSum ...[]int)[]int{
	var sums[]int
	for _,numbers:=range numberToSum{
		if len(numbers)!=0{
			tail:=numbers[1:]
			sums=append(sums,Sum(tail))
		} else{
			sums=append(sums,0)
		}

	}
	return sums
}
func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers:=[]int{1,2,3}
		got:=Sum(numbers)
		want:=6
		if got !=want{
			t.Errorf("got %d want %d given,%v", got, want, numbers)
		}
	})

}
func TestSumAll(t *testing.T){
	got:=SumAll([]int{1,2},[]int{0,9})
	want:=[]int{3,9}
	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v want %v",got,want)
	}
}
func TestSumAllTails(t *testing.T){
	checkSums:= func(t *testing.T,got,want []int) {
		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v want %v",got,want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got:=SumAllTails([]int{1,2},[]int{0,9})
		want:=[]int{2,9}
		checkSums(t,got,want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got:=SumAllTails([]int{},[]int{1,7,8})
		want:=[]int{0,15}
		checkSums(t,got,want)
	})

}


keypoints:
	1:append: s=append(s,ele)
	2:reflect.DeepEqual() :can compare different type of ele, it' not type safe
	3:len()