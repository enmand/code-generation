package generated

import "errors"

type Pet interface {
	Speak() string
	Walk() error
}
type Dog struct {
	speaks   string
	walkable bool
}

func NewDog() *Dog {
	return &Dog{
		speaks:   "bark",
		walkable: true,
	}
}
func (p *Dog) Speak() string {
	return p.speaks
}
func (p *Dog) Walk() error {
	if p.walkable == true {
		return errors.New("unable to walk pet")
	}
	return nil
}

type Cat struct {
	speaks   string
	walkable bool
}

func NewCat() *Cat {
	return &Cat{
		speaks:   "meow",
		walkable: false,
	}
}
func (p *Cat) Speak() string {
	return p.speaks
}
func (p *Cat) Walk() error {
	if p.walkable == true {
		return errors.New("unable to walk pet")
	}
	return nil
}

type Horse struct {
	speaks   string
	walkable bool
}

func NewHorse() *Horse {
	return &Horse{
		speaks:   "neigh",
		walkable: true,
	}
}
func (p *Horse) Speak() string {
	return p.speaks
}
func (p *Horse) Walk() error {
	if p.walkable == true {
		return errors.New("unable to walk pet")
	}
	return nil
}

type Parrot struct {
	speaks   string
	walkable bool
}

func NewParrot() *Parrot {
	return &Parrot{
		speaks:   "english?",
		walkable: false,
	}
}
func (p *Parrot) Speak() string {
	return p.speaks
}
func (p *Parrot) Walk() error {
	if p.walkable == true {
		return errors.New("unable to walk pet")
	}
	return nil
}
