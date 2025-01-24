# Library Management System in Go

## Overview
This program is a console-based Library Management System implemented in Go. It allows users to manage books and eBooks in a library through operations like adding, removing, searching, and listing books.

## Features
- Manage both physical books and eBooks with relevant details.
- Perform CRUD operations such as adding, removing, and searching books.
- Display all available books and eBooks in the library.

## Installation and Running

### Prerequisites
- Install [Go](https://golang.org/doc/install) on your system.

### Steps to Run
1. Clone the repository or save the code in a file named `main.go`.
2. Open a terminal in the project directory.
3. Run the following command to execute the program:
   ```bash
   go run main.go
   ```

## Usage Instructions
1. Choose an option from the menu displayed.
2. Follow prompts to enter book details, search terms, or ISBNs.
3. View results or perform further operations.

## Example Commands

### Adding a Book
```
1
Enter the type of book (Book/EBook): Book
Enter Title: The Go Programming Language
Enter Author: Alan A. A. Donovan
Enter ISBN: 9780134190440
Is the book available (true/false): true
```

### Listing All Books
```
4
```

### Removing a Book
```
2
Enter ISBN of the book to remove: 9780134190440
```

