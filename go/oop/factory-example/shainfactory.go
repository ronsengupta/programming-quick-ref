package main

type shainfactory interface {
	factory(string) Shain
}

type fact struct {
}

func (this *fact) factory(shaintype string) Shain {
	return &tanto{}
}
