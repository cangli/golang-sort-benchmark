package golang_sort_benchmark

import (
	"testing"
)

var arrLength int = 50000

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
