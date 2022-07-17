

//已有函数round13
func rand5() int {
	for {
		res := rand13()
		if res <= 5 {
			break
		}
	}

	return res
}

//已有函数round5
func rand13() int {
	for {
		t1 := rand5()
		t2 := rand5()

		n := (t1-1)*6 + t2
		if n <= 26 {
			break
		}
	}

	return 1 + n%13

}