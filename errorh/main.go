package errorh

func Handle(e error) {
	if e != nil {
		panic(e.Error())
	}
}
