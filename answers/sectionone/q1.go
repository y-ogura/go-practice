package sectionone

import "fmt"

// Q1 map[string]{"a": "a"} と map[string]string{"b": "b"}をcにマージしてください
func Q1() map[string]string {
	a := map[string]string{"a": "a"}
	b := map[string]string{"b": "b"}

	fmt.Println(a, b)

	c := map[string]string{}

	return c
}
