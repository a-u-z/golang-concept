package as

import "log"

func eff() {
	sl := []int{1, 2, 3}

	// 這種 v 會複製一份出來，所以耗效能，如果是 sl 裡面是 struct 那就更耗效能
	for _, v := range sl {
		log.Printf("here is v:%+v", v)
	}

	// 用這種方式，效能比較好，不過是會改到原本的 sl
	for i := 0; i < len(sl); i++ {

	}

}
