package main

import (
	"errors"
)

type Entitlement struct {
	Teams struct {
		Enabled bool    `json:"enabled,omitempty"`
		Size    Counter `json:"size,omitempty"`
	} `json:"teams,omitempty"`
}

var ErrCounterExhausted = errors.New("counter exhausted")

type Counter int64

func NewCounter(i int64) *Counter {
	c := Counter(i)
	return &c
}

func (c *Counter) Value() int64 {
	if *c == 0 {
		return 0
	}

	return int64(*c)
}

func (c *Counter) Take() error { return c.TakeN(1) }

func (c *Counter) TakeN(i int64) error {
	if *c <= 0 {
		return ErrCounterExhausted
	}

	*c -= Counter(i)

	return nil
}

// meu código personalizado

//var age int = 20 // não dá pra criar métodos em um tipo primitivo

// type age int // cria um tipo para poder criar métodos nele

// func (myage age) Show() age {
// 	return myage * 4
// }

// func main() {
// 	juniorage := age(50)
// 	fmt.Println(juniorage.Show())
// }
