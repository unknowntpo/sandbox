package main

import (
	"fmt"
	"time"
)

// Dog has a name and a list of people caring for him and watching
// what he does
type Dog struct {
	name    string
	sitters map[string][]chan string
}

// AddSitter adds an event listener to the Dog struct instance
func (b *Dog) AddSitter(e string, ch chan string) {
	if b.sitters == nil {
		b.sitters = make(map[string][]chan string)
	}
	// if sitter exist
	if _, ok := b.sitters[e]; ok {
		b.sitters[e] = append(b.sitters[e], ch)
	} else {
		// What does it means ?
		b.sitters[e] = []chan string{ch}
	}
}

// RemoveSitter removes an event listener from the Dog struct instance
func (b *Dog) RemoveSitter(e string, ch chan string) {
	if _, ok := b.sitters[e]; ok {
		for i := range b.sitters[e] {
			if b.sitters[e][i] == ch {
				b.sitters[e] = append(b.sitters[e][:i], b.sitters[e][i+1:]...)
				break
			}
		}
	}
}

// Emit emits an event on the Dog struct instance
func (b *Dog) Emit(e string, response string) {
	// if sitters for that event exist
	if _, ok := b.sitters[e]; ok {
		for _, handler := range b.sitters[e] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}

func main() {
	balto := Dog{"Balto", nil}

	flavio := make(chan string)

	balto.AddSitter("bark", flavio)
	balto.AddSitter("poop", flavio)
	balto.AddSitter("hungry", flavio)

	go func() {
		for {
			// Wait for message from channel flavio sent by (*Dog).Emit()
			msg := <-flavio
			fmt.Println("Flavio: I " + msg)
		}
	}()

	fmt.Println("The dog barked!")
	balto.Emit("bark", "told it not to bark!")

	fmt.Println("The dog pooped!")
	balto.Emit("poop", "picked up poop!")

	fmt.Println("The dog is hungry!")
	balto.Emit("hungry", "fed the dog!")

	time.Sleep(3 * time.Second)

	fmt.Printf("\n.\n.\n.\n")
	balto.RemoveSitter("poop", flavio)

	// Hired a dog sitter to pick up poop
	dogsitter := make(chan string)
	balto.AddSitter("poop", dogsitter)
	fmt.Println("Flavio hired a dogsitter to pick up poop, won't pick it up again")

	go func() {
		for {
			msg := <-dogsitter
			fmt.Println("Dogsitter: I " + msg)
		}
	}()

	fmt.Println("Dog barked!")
	balto.Emit("bark", "told it not to bark!")

	fmt.Println("Dog has pooped!")
	balto.Emit("poop", "picked up poop!")
	fmt.Println("Dog has pooped again!")
	balto.Emit("poop", "picked up poop, again!")
	fmt.Println("The dog is hungry!")
	balto.Emit("hungry", "fed the dog!")

	fmt.Scanln()
}
