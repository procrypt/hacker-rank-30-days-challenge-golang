package main

import "fmt"

type Stack struct {
	container []interface{}
}

type Queue struct {
	container []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func(s *Stack) PushCharacter(item interface{}) {
	s.container = append(s.container, item)
}

func(s *Stack) PopCharacter() interface{} {
	element := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return element
}

func(q *Queue) EnqueCharacter(item interface{}) {
	q.container = append(q.container, item)
}

func(q *Queue) DequeCharacter() interface{} {
	element := q.container[0]
	q.container = q.container[1:]
	return element
}

func main() {
	  s := NewStack()
	  q := NewQueue()
	  a := "abhishek"

	  for i := range a {
	  	s.PushCharacter(string(a[i]))
	  }

	  fmt.Println("+++Pop+++")
	  for i:=0; i<len(a); i++ {
	  	fmt.Println(s.PopCharacter())
	  }

	  for i := range a {
	  	q.EnqueCharacter(string(a[i]))
	  }

	  fmt.Println("+++Deque+++")
	  for i:=0; i<len(a); i++ {
	  	fmt.Println(q.DequeCharacter())
	  }
}
