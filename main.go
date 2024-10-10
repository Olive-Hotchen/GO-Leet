package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {

	//TWO SUM
	//var testCase = []int{2, 7, 11, 15}
	//fmt.Println(TwoSum(testCase, 26))

	// WRITE STACK WITH QUEUE
	//queue := Queue{
	//	Elements: []int{1, 2, 0, 4, 5, 6},
	//	Size:     6,
	//}
	//
	//fmt.Println("Queue values: ", queue.Elements)
	//
	//stack := Stack{Queue: Queue{Elements: []int{}, Size: 6}, Size: 6}
	//
	//
	//for i := 0; i < stack.Size; i++ {
	//	stack.Push(queue.Elements[i])
	//}
	//
	//for i := 0; i < stack.Size; i++ {
	//	pop, _ := stack.Pop()
	//	fmt.Print(strconv.Itoa(pop) + ", ")
	//}

	// VALID ANAGRAM
	//fmt.Println(ValidAnagram("abc", "cba"))
	//fmt.Println(ValidAnagram("abc", "cca"))

	// REVERSE INTEGER
	//fmt.Println(ReverseInteger(12345))
	//fmt.Println(ReverseInteger(-32472))

	//REVERSE LINKED LIST
	//head := ListNode{
	//	1,
	//	&ListNode{
	//		2,
	//		nil,
	//	},
	//}
	//
	//list := LinkedList{Head: &head}
	//fmt.Println(list.AsArray())
	//list.Reverse()
	//fmt.Println(list.AsArray())

	// LONGEST COMMON PREFIX
	//fmt.Println(LongestCommonPrefix([]string{"flower", "fly", "flame"}))
	//fmt.Println(LongestCommonPrefix([]string{"space", "spy", "suffer"}))
	println(ValidParentheses("{}[]{()"))
	println(ValidParentheses("{}()"))
	println(ValidParentheses("{(}()"))
	println(ValidParentheses("{{{([[()]])}}}"))

}

func ValidParentheses(s string) bool {
	var brackets = Stack{}
	var openingList = "{(["

	var bracketMap = make(map[rune]rune)
	bracketMap['{'] = '}'
	bracketMap['('] = ')'
	bracketMap['['] = ']'

	openingRunes := []rune(s)

	for _, r := range openingRunes {
		found := strings.ContainsRune(openingList, r)
		if found {
			brackets.Push(r)
		}
	}

	for _, r := range s {
		if bracketMap[brackets.Peek().(rune)] == r {
			brackets.Pop()
		}
	}

	return brackets.Empty()
}

func LongestCommonPrefix(strings []string) string {
	longest := ""

	if len(strings) == 0 {
		return longest
	}

	broken := false
	pointer := 0

	for !broken {
		charsAtPointer := make([]string, 0)
		isMatching := true
		for _, str := range strings {
			charsAtPointer = append(charsAtPointer, string(str[pointer]))
		}

		for i := 0; i < len(charsAtPointer); i++ {
			if charsAtPointer[i] != charsAtPointer[0] {
				isMatching = false
			}
		}

		if isMatching {
			pointer++
			longest += charsAtPointer[0]
		} else {
			broken = true
		}
	}

	return longest
}

func (list *LinkedList) Reverse() {
	var prevNode *ListNode = nil

	for list.Head != nil {
		nextNode := list.Head.Next
		list.Head.Next = prevNode
		prevNode = list.Head
		list.Head = nextNode
	}

	list.Head = prevNode
}

func (list *LinkedList) AsArray() []int {
	arr := make([]int, 1)

	if list.Head == nil {
		return []int{}
	}

	arr[0] = list.Head.Data
	pointer := list.Head.Next
	for pointer != nil {
		arr = append(arr, pointer.Data)
		pointer = pointer.Next
	}

	return arr
}

type ListNode struct {
	Data int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func ReverseInteger(num int) int {
	negative := num < 0
	reversed := 0
	if negative {
		num *= -1
	}

	for num > 0 {
		reversed = (reversed * 10) + (num % 10)
		num = int(math.Floor(float64(num / 10)))
	}

	if reversed > int(math.Pow(2, 31-1)) {
		return 0
	}

	if negative {
		return reversed * -1
	}

	return reversed
}

func ValidAnagram(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	var counts = make(map[string]int, len(s1))
	for i := 0; i < len(s1); i++ {
		counts[string(s1[i])]++
	}

	for i := 0; i < len(s2); i++ {
		counts[string(s2[i])]--
	}

	finalDiffCount := 0
	for _, t := range counts {
		if t < 0 {
			continue
		}

		finalDiffCount += t
	}

	return finalDiffCount == 0
}

func TwoSum(nums []int, target int) []int {
	var indices = make(map[int]int)
	var res = make([]int, 0)

	for i := 0; i < len(nums); i++ {
		var element = nums[i]
		complement := target - element

		if val, ok := indices[complement]; ok {
			res = append(res, val)
			res = append(res, i)
			return res
		}

		indices[element] = i
	}

	return res
}

type Stack struct {
	elements []interface{}
}

func (stack *Stack) Push(element interface{}) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack) Pop() interface{} {
	if len(stack.elements) == 0 {
		return nil
	}

	element := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return element
}

func (stack *Stack) Empty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}

func (stack *Stack) Peek() interface{} {
	return stack.elements[len(stack.elements)-1]
}

type QueueStack struct {
	Queue Queue
	Size  int
}

func (stack *QueueStack) Push(element int) {
	rotation := stack.Queue.Size
	stack.Queue.Enqueue(element)
	for rotation > 0 {
		stack.Queue.Enqueue(stack.Queue.Dequeue())
		rotation--
	}
}

func (stack *QueueStack) Pop() (int, error) {
	return stack.Queue.Dequeue(), nil
}

func (stack *QueueStack) Top() (int, error) {
	return stack.Queue.Peek()
}

func (stack *QueueStack) Empty() bool {
	return stack.Queue.IsEmpty()
}

type Queue struct {
	Elements []int
	Size     int
}

func (q *Queue) Enqueue(elem int) {
	if q.GetLength() == q.Size {
		fmt.Println("Overflow")
		return
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return 0
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element // Slice off the element once it is dequeued.
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}
