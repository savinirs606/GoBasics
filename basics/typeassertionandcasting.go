package main

import "fmt"

type Lion struct{ Name string }

// value receiver: both Dog and *Dog implement Speaker
func (d Lion) Speak() string { return "roar: " + d.Name }

func assertExample(s Speaker) {
	// safe assertion forms
	if dPtr, ok := s.(*Lion); ok {
		fmt.Println("asserted to *Lion:", dPtr.Name)
	} else if dVal, ok := s.(Lion); ok {
		fmt.Println("asserted to Lion:", dVal.Name)
	} else {
		fmt.Println("not a Dog")
	}

	// type switch
	switch v := s.(type) {
	case *Lion:
		fmt.Println("type switch: *Lion", v.Name)
	case Lion:
		fmt.Println("type switch: Lion", v.Name)
	default:
		fmt.Println("unknown")
	}
}

func typeassertion() {
	var s1 Speaker = Lion{Name: "Fido"} // dynamic type: Dog
	var s2 Speaker = &Lion{Name: "Rex"} // dynamic type: *Dog

	assertExample(s1)
	assertExample(s2)
}
