package sort

import (
	"log"
	"math/rand"
	"testing"
)

// Bubble sort
func BubbleSort(arr []int) {

	count := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				count = count + 1
			}
		}
	}
	log.Print("count:", count)
	//log.Print(arr)

}

func BubbleSort2(arrayzor []int) []int {

	count := 0
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor)-1; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				swap(arrayzor, i, i+1)
				//arrayzor[i],arrayzor[i+1] = arrayzor[i+1],arrayzor[i]
				swapped = true
				count = count + 1
			}
		}
	}

	log.Print("count:", count)
	return arrayzor
}

//0.150 ns/op
func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func TestBubbleSort(t *testing.T) {

	arr := randInt(10)
	BubbleSort(arr)

}

func TestBubbleSort2(t *testing.T) {

	arr := randInt(10)
	BubbleSort2(arr)

}

//109333425 ns/op
func BenchmarkBubbleSort(b *testing.B) {
	rand.Seed(213534765)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := randInt(10000)
		b.StartTimer()
		BubbleSort(arr)
	}
	b.StopTimer()

}

//153117341 ns/op
func BenchmarkBubbleSort2(b *testing.B) {
	rand.Seed(213534765)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := randInt(10000)
		b.StartTimer()
		BubbleSort2(arr)
	}
	b.StopTimer()
}


//64941607 ns/op
func BenchmarkBubbleOrderSort(b *testing.B) {
	arr := orderInt(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
	b.StopTimer()

}

//12772 ns/op
//24933921
func BenchmarkBubbleOrderSort2(b *testing.B) {
	arr := orderInt(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort2(arr)
	}
	b.StopTimer()
}

func randInt(len int) []int {
	//rand.Seed(time.Now().UnixNano())
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

func orderInt(len int) []int {
	ints := make([]int, len)

	for i := 1; i < len; i++ {
		ints[i] = i + 100
	}
	return ints
}

func TestRandInt(t *testing.T) {

	log.Println(randInt(10))
}
