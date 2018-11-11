package mock

type Retriver struct {
	Contents string
}

func (r Retriver) Get(url string) string {
	r.Contents = "this is from mock.retriver.Get"
	return url
}

