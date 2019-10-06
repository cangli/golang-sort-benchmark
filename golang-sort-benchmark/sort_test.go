package golang_sort_benchmark

import "testing"

var arrLength int = 50000

//func TestA(t *testing.T){
//	arr := randomArr(50)
//	fmt.Println(arr)
//	fmt.Println(bubbleSort(shuffle(arr)))
//	fmt.Println(insertionSort(shuffle(arr)))
//	fmt.Println(selectionSort(shuffle(arr)))
//	fmt.Println(mergeSortMain(shuffle(arr)))
//	fmt.Println(quickSortMain(shuffle(arr)))
//	fmt.Println(heapSort(shuffle(arr)))
//}
func BenchmarkBubbleSort(b *testing.B) {
	for i:=0;i<b.N;i++{
		b.StopTimer()
		arr := randomArr(arrLength)
		b.StartTimer()
		bubbleSort(arr)
	}
}

func BenchmarkInsertionSort(b *testing.B){
	for i:=0;i<b.N;i++{
		b.StopTimer()
		arr := randomArr(arrLength)
		b.StartTimer()
		insertionSort(arr)
	}
}

func BenchmarkSelectionSort(b *testing.B){
	for i:=0;i<b.N;i++{
		b.StopTimer()
		arr := randomArr(arrLength)
		b.StartTimer()
		selectionSort(arr)
	}
}

func BenchmarkMergeSort(b *testing.B){
	for i:=0;i<b.N;i++{
		b.StopTimer()
		arr := randomArr(arrLength)
		b.StartTimer()
		mergeSortMain(arr)
	}
}

func BenchmarkQuickSort(b *testing.B){
	for i:=0;i<b.N;i++{
		b.StopTimer()
		arr := randomArr(arrLength)
		b.StartTimer()
		quickSortMain(arr)
	}
}

func BenchmarkHeapSort(b *testing.B){
	for i:=0;i<b.N;i++{
		b.StopTimer()
		arr := randomArr(arrLength)
		b.StartTimer()
		heapSort(arr)
	}
}
