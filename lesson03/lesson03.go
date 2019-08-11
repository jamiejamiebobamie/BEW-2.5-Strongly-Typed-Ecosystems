// NOUNS IN CODE:
//
// ANIMATE
// INANIMATE
// MAMMAL
// HUMAN
// DOG
// CAT
// MALE
// FEMALE
// GIRL
// WOMAN

// VERBS IN CODE:
//
// DANCE
// [SAY] HELLO
// SNEEZE

package main

import (
    "fmt"
    "strconv"
)

// strconv.Itoa([int to convert])
// ^^^there's another method,
// but it seems more involved than the one above

type thing struct {
    living bool
}

type human struct {
    thing
    gender string
    age int
    name string
}

type Dancer interface {
    dance()
    hello()
    }

func (S *human) dance() {
    for i := 0; i < 10; i++{
        fmt.Println(S.name + " does the cha-cha.")
    }
}

func sneeze(person human) {
    fmt.Println(person.name + " sneezes.")
    if person.age > 60 {
        if person.gender == "male" {
            fmt.Println("Must be his age!")
        } else {
            fmt.Println("Must be her age!")
        }
    } else {
        fmt.Println("God bless you!")
    }
}

func (dan *human) hello() (ret string) {
    ret = "Hello, my name is " + dan.name + ". I am " + strconv.Itoa(dan.age) + " years old."
    return
}

func main() {
    S := human{gender:"female", age:52, name:"Sally"}
    S.living = true // how do i set this at initialization?
    dan := human{gender:"male", age:62, name:"Dan"}
    fmt.Println(S)
    S.dance()
    fmt.Println(dan.hello())
    sneeze(dan)
    sneeze(S)
}

// FURTHER READING:
// https://travix.io/type-embedding-in-go-ba40dd4264df
// https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a
