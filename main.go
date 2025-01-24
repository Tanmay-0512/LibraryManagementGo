package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Book Struct
type Book struct {
    Title     string
    Author    string
    ISBN      string
    Available bool
}

// Initialize a new book
func NewBook(title, author, isbn string, available bool) Book {
    return Book{Title: title, Author: author, ISBN: isbn, Available: available}
}

// Display the book details
func (b Book) DisplayDetails() {
    fmt.Printf("Title: %s\nAuthor: %s\nISBN: %s\nAvailable: %v\n\n", b.Title, b.Author, b.ISBN, b.Available)
}

// EBook Struct (inherits from Book via embedding)
type EBook struct {
    Book
    FileSize int // File size in MB
}

// Override the DisplayDetails method for EBook
func (e EBook) DisplayDetails() {
    fmt.Printf("Title: %s\nAuthor: %s\nISBN: %s\nAvailable: %v\nFileSize: %dMB\n\n", e.Title, e.Author, e.ISBN, e.Available, e.FileSize)
}

// Book Interface
type BookInterface interface {
    DisplayDetails()
}

// Library Struct
type Library struct {
    Books []interface{} // Use interface{} to handle both Book and EBook types
}

// Add a book to the library
func (l *Library) AddBook(book interface{}) {
    l.Books = append(l.Books, book)
}

// Remove a book from the library using ISBN
func (l *Library) RemoveBook(isbn string) error {
    for i, item := range l.Books {
        switch book := item.(type) {
        case Book:
            if book.ISBN == isbn {
                l.Books = append(l.Books[:i], l.Books[i+1:]...)
                return nil
            }
        case EBook:
            if book.ISBN == isbn {
                l.Books = append(l.Books[:i], l.Books[i+1:]...)
                return nil
            }
        }
    }
    return fmt.Errorf("book with ISBN %s not found", isbn)
}

// Search books by title
func (l *Library) SearchBookByTitle(title string) []interface{} {
    var results []interface{}
    for _, item := range l.Books {
        switch book := item.(type) {
        case Book:
            if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
                results = append(results, book)
            }
        case EBook:
            if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
                results = append(results, book)
            }
        }
    }
    return results
}

// List all books and ebooks
func (l *Library) ListBooks() {
    if len(l.Books) == 0 {
        fmt.Println("No books in the library.")
        return
    }
    for _, item := range l.Books {
        if bookItem, ok := item.(BookInterface); ok {
            bookItem.DisplayDetails()
        }
    }
}

func main() {
    // Initialize library
    library := &Library{}

    // Menu loop
    for {
        fmt.Println("\nLibrary Management System")
        fmt.Println("1. Add a Book/EBook")
        fmt.Println("2. Remove a Book/EBook")
        fmt.Println("3. Search for Books by Title")
        fmt.Println("4. List all Books/EBooks")
        fmt.Println("5. Exit")
        fmt.Print("Choose an option: ")

        var option string
        // Use bufio.NewReader to capture the full input line
        reader := bufio.NewReader(os.Stdin)
        option, _ = reader.ReadString('\n')
        option = strings.TrimSpace(option) // Remove trailing newlines

        switch option {
        case "1":
            // Add Book/EBook
            fmt.Print("Enter the type of book (Book/EBook): ")
            bookType, _ := reader.ReadString('\n')
            bookType = strings.TrimSpace(bookType)

            fmt.Print("Enter Title: ")
            title, _ := reader.ReadString('\n')
            title = strings.TrimSpace(title)

            fmt.Print("Enter Author: ")
            author, _ := reader.ReadString('\n')
            author = strings.TrimSpace(author)

            fmt.Print("Enter ISBN: ")
            isbn, _ := reader.ReadString('\n')
            isbn = strings.TrimSpace(isbn)

            var available bool
            // Handle boolean input validation
            for {
                fmt.Print("Is the book available (true/false): ")
                availableInput, _ := reader.ReadString('\n')
                availableInput = strings.TrimSpace(availableInput)

                // Convert input to boolean
                if availableInput == "true" {
                    available = true
                    break
                } else if availableInput == "false" {
                    available = false
                    break
                } else {
                    fmt.Println("Invalid input, please enter 'true' or 'false'.")
                }
            }

            if bookType == "Book" {
                library.AddBook(NewBook(title, author, isbn, available))
            } else if bookType == "EBook" {
                var fileSize int
                fmt.Print("Enter File Size (MB): ")
                _, err := fmt.Scanf("%d", &fileSize)
                if err != nil {
                    fmt.Println("Invalid input for file size.")
                    break
                }
                ebook := EBook{Book: NewBook(title, author, isbn, available), FileSize: fileSize}
                library.AddBook(ebook)
            } else {
                fmt.Println("Invalid book type entered.")
            }

        case "2":
            // Remove Book/EBook
            fmt.Print("Enter ISBN of the book to remove: ")
            isbn, _ := reader.ReadString('\n')
            isbn = strings.TrimSpace(isbn)
            err := library.RemoveBook(isbn)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Book removed successfully.")
            }

        case "3":
            // Search for Books by Title
            fmt.Print("Enter Title to search: ")
            title, _ := reader.ReadString('\n')
            title = strings.TrimSpace(title)
            searchResults := library.SearchBookByTitle(title)
            if len(searchResults) == 0 {
                fmt.Println("No books found.")
            } else {
                for _, item := range searchResults {
                    if bookItem, ok := item.(BookInterface); ok {
                        bookItem.DisplayDetails()
                    }
                }
            }

        case "4":
            // List all Books/EBooks
            library.ListBooks()

        case "5":
            // Exit
            fmt.Println("Exiting...")
            return

        default:
            fmt.Println("Invalid option. Try again.")
        }
    }
}
