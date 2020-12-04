package main

import "fmt"

type Iface interface {
	DoSomething()
}

type BaseStruct struct {
}

func NewBaseStruct() Iface {
	return &BaseStruct{}
}

func (b BaseStruct) DoSomething() {
	fmt.Print("DoSomething() From Base Struct")
}

type AbstractIface interface {
	Iface
	DoSomeOtherThing()
}

type AbstractStruct struct {
	Iface         // super
	AbstractIface //derived
}

func NewAbstractStruct(derived AbstractIface) AbstractIface {
	return AbstractStruct{NewBaseStruct(), derived}
}

//override
func (a AbstractStruct) DoSomething() {
	fmt.Println("DoSomething() From Abstract Struct")
	a.DoSomeOtherThing()
}

type DerivedStruct struct {
	AbstractIface
}

func NewDerivedStruct() AbstractIface {
	return NewAbstractStruct(DerivedStruct{})
}

func (d DerivedStruct) DoSomeOtherThing() {
	fmt.Println("SomeOtherThing() From Derived")
}

func main() {
	d := NewDerivedStruct()
	d.DoSomething()
}
