package web

func getBlog(id int) (Model, error) {
	return Do(id)
}
