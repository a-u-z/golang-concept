package de

func De() {
	// recover()
	// panic(1)

	// defer recover()
	// panic(1)

	defer func() { // recover 要放在 defer 且一定要有 func 來處理
		recover()
	}()
	panic(1)
}
