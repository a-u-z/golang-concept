package as

import "log"

func diffArrSlice() {
	arr()
	slice()
}

// 這是 array
func arr() {
	arr := [3]int{0, 1, 2}
	f := func(v [3]int) {
		v[0] = 100
	}
	f(arr)
	log.Printf("here is arr:%+v", arr)
}

// 這是 slice
func slice() {
	sl := []int{0, 1, 2}
	f := func(v []int) {
		v[0] = 100
	}

	f(sl)
	log.Printf("here is sl:%+v", sl)
}
