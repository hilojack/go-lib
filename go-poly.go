/**

作者：痕无落 链接：https://www.jianshu.com/p/b333c5f34ef6
*/
package main

import (
    "fmt"
)

type Animal interface {
    Sleep()
    Age() int
    Type() string
}

type Cat struct {
    MaxAge int
}

func (this *Cat) Sleep() {
    fmt.Println("Cat need sleep")
}
func (this *Cat) Age() int {
    return this.MaxAge
}
func (this *Cat) Type() string {
    return "Cat"
}

type Dog struct {
    MaxAge int
}

func (this *Dog) Sleep() {
    fmt.Println("Dog need sleep")
}
func (this *Dog) Age() int {
    return this.MaxAge
}
func (this *Dog) Type() string {
    return "Dog"
}


func Factory(name string) Animal {
    switch name {
    case "dog":
        return &Dog{MaxAge: 20}
    case "cat":
        return &Cat{MaxAge: 10}
    default:
        panic("No such animal")
    }
}


func main() {
    animal := Factory("dog")
    animal.Sleep()
    fmt.Printf("%s max age is: %d", animal.Type(), animal.Age())
}

