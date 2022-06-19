package errorh

func Handle(e error) {
	panic(e.Error())
}
