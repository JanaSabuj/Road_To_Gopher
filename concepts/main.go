package main

import (
	"fmt"
	"time"
)

type Book struct {
	title string
	author string

	isSaved bool
	savedAt time.Time
}

// Classic Go Style function
func (book *Book) saveBook() {
	book.isSaved = true
	book.savedAt = time.Now()
}

// call by ref C style - same utility as the above the func 
// func saveBook(book *Book) {
// 	book.isSaved = true
// 	book.savedAt = time.Now()
// }

// call by value - no modification
// func saveBook(book Book) {
// 	book.isSaved = true
// 	book.savedAt = time.Now()
// }

func main() {
	b := Book{
		title: "Harry Potter",
		author: "J K Rowling",
	}

	fmt.Println(b)
	// saveBook(&b)
	b.saveBook() // same as the above statement - b is implicitly converted to pointer when calling saveBook()
	fmt.Println(b)

	// Anything in Go is either Data/Struct or Functions
	// Struct only stores Data, and no functions
}



