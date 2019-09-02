package golang_sort_benchmark

import (
	"fmt"
	"math/rand"
	//"sort"
)


func bubbleSort(arr []int) []int {
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-1-i;j++{
			if arr[j+1] < arr[j]{
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i:=1;i<len(arr);i++{
		for j:=i-1;j>=0 && arr[j] > arr[j+1];j--{
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return arr
}

func selectionSort(arr []int)[]int{
	for i:=0;i<len(arr)-1;i++{
		minIndex := i
		for j:=i+1;j<len(arr);j++{
			if arr[j] < arr[minIndex]{
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func randomArr(n int)[]int{
	out := make([]int, 0, n)
	for i:=0;i<n;i++{
		out = append(out, rand.Intn(50))
	}
	return out
}
func randomArrArr(m, n int)[][]int{
	out := make([][]int, 0, m)
	for i:=0;i<m;i++{
		out = append(out, randomArr(n))
	}
	return out
}

func mergeSortMerge(arr1, arr2 []int)[]int{
	out := make([]int, 0, len(arr1)+len(arr2))
	var i, j = 0, 0
	for i<len(arr1)&&j<len(arr2){
		if arr1[i] <= arr2[j]{
			out = append(out, arr1[i])
			i++
		} else {
			out = append(out, arr2[j])
			j++
		}
	}
	if len(arr1) > i{
		out = append(out, arr1[i:]...)
	} else if len(arr2) > j{
		out = append(out, arr2[j:]...)
	}
	return out
}
func mergeSort(arr []int, start, end int)[]int{
	if start >= end {
		return arr[start:end+1]
	}
	middle := (start + end) / 2
	arr1, arr2 := mergeSort(arr, start, middle), mergeSort(arr, middle+1, end)
	return mergeSortMerge(arr1, arr2)
}

func mergeSortMain(arr []int)[]int{
	return mergeSort(arr, 0, len(arr)-1)
}

func partition(arr []int, start, end int)int{
	middle := (start + end) / 2
	max, min := start, end
	if arr[max] < arr[min]{
		max, min = min, max
	}
	if arr[max] < arr[middle]{
		max, middle = middle, max
	}
	if arr[middle] < arr[min]{
		middle, min = min, middle
	}
	//pivot := arr[middle]
	arr[middle], arr[end] = arr[end], arr[middle]
	i, j:= start, start
	for ;j < end;j++{
		if arr[j] < arr[end]{
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func quickSort(arr []int, start, end int){
	if start >= end {
		return
	}
	p := partition(arr, start, end)
	quickSort(arr, start, p - 1)
	quickSort(arr, p + 1, end)
}

func quickSortMain(arr []int)[]int{
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func shuffle(arr []int)[]int{
	rand.Shuffle(len(arr), func(i, j int){
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func main(){
	arr := randomArr(50)
	fmt.Println(arr)
	fmt.Println(bubbleSort(shuffle(arr)))
	fmt.Println(insertionSort(shuffle(arr)))
	fmt.Println(selectionSort(shuffle(arr)))
	fmt.Println(mergeSortMain(shuffle(arr)))
	fmt.Println(quickSortMain(shuffle(arr)))
}