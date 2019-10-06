# golang-sort-benchmark
go语言实现排序算法的压力测试

# 简介
学习了极客时间的数据结构与算法，自己动手实现了以下几个排序算法，并做了压力测试，对算法的复杂度有了更直观的理解。
- 冒泡排序
- 插入排序
- 选择排序
- 归并排序
- 快速排序

# 压力测试
go test -bench=. -benchmem
## arrLength=500
    goos: windows
    goarch: amd64
    pkg: gopro
    BenchmarkBubbleSort-8      	   10000	    217112 ns/op	       1 B/op	       0 allocs/op
    BenchmarkInsertionSort-8   	   20000	     96505 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSelectionSort-8   	   10000	    159309 ns/op	       1 B/op	       0 allocs/op
    BenchmarkMergeSort-8       	   30000	     45269 ns/op	   36672 B/op	     499 allocs/op
    BenchmarkQuickSort-8       	  100000	     20461 ns/op	       0 B/op	       0 allocs/op
    PASS
    ok  	gopro	15.583s
## arrLength=5000
    goos: windows
    goarch: amd64
    pkg: gopro/golang-sort-benchmark
    BenchmarkBubbleSort-8      	     100	  22461283 ns/op	     182 B/op	       0 allocs/op
    BenchmarkInsertionSort-8   	     200	   9255530 ns/op	       8 B/op	       0 allocs/op
    BenchmarkSelectionSort-8   	     100	  13380763 ns/op	       1 B/op	       0 allocs/op
    BenchmarkMergeSort-8       	    3000	    518695 ns/op	  522642 B/op	    4999 allocs/op
    BenchmarkQuickSort-8       	    5000	    338819 ns/op	      14 B/op	       0 allocs/op
    PASS
    ok  	gopro/golang-sort-benchmark	11.479s
## arrLength=50000
    goos: windows
    goarch: amd64
    pkg: gopro/golang-sort-benchmark
    BenchmarkBubbleSort-8      	       1	3482199200 ns/op	   17984 B/op	      10 allocs/op
    BenchmarkInsertionSort-8   	       2	 920552650 ns/op	     268 B/op	       1 allocs/op
    BenchmarkSelectionSort-8   	       1	1332076200 ns/op	    1152 B/op	       1 allocs/op
    BenchmarkMergeSort-8       	     300	   5833670 ns/op	 6751855 B/op	   49999 allocs/op
    BenchmarkQuickSort-8       	     100	  14160808 ns/op	      83 B/op	       0 allocs/op
    PASS
    ok  	gopro/golang-sort-benchmark	12.503s

以上三个benchmark结果分别是在数组长度500、5000、50000的前提下得出的。
可以看出冒泡、插入、选择排序时间都是以n^2规模增长。而在这三个算法中，以插入排序表现最佳，这就是插入排序会作为小数据排序算法的原因吧。
归并排序和快速排序的结果有点出乎我意料，500和5000的数组长度下，快速排序更快，而在50000的数组长度下，则是归并排序反超。
快速排序的时间是以略大于nlogn的规模在增长，这个结果与快排的时间复杂度一致。
然而从测试结果看，归并排序时间是以接近n的规模增长，这令我有点想不通。留待思考出结果后再行补充。
 