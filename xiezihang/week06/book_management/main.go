package main

import (
	"fmt"
)

// 定义 Book 结构体
type Book struct {
	ID          int
	Title       string
	Author      string
	IsAvailable bool
}

// 定义 Magazine 结构体
type Magazine struct {
	ID          int
	Title       string
	Issue       int
	IsAvailable bool
}

// 定义 Manageable 接口
type Manageable interface {
	Borrow() bool
	Return() bool
	GetInfo() string
}

// 为 Book 实现 Manageable 接口
func (b *Book) Borrow() bool {
	if b.IsAvailable {
		b.IsAvailable = false
		return true
	}
	return false
}

func (b *Book) Return() bool {
	if !b.IsAvailable {
		b.IsAvailable = true
		return true
	}
	return false
}

func (b *Book) GetInfo() string {
	return fmt.Sprintf("Book ID: %d, Title: %s, Author: %s, Available: %t", b.ID, b.Title, b.Author, b.IsAvailable)
}

// 为 Magazine 实现 Manageable 接口
func (m *Magazine) Borrow() bool {
	if m.IsAvailable {
		m.IsAvailable = false
		return true
	}
	return false
}

func (m *Magazine) Return() bool {
	if !m.IsAvailable {
		m.IsAvailable = true
		return true
	}
	return false
}

func (m *Magazine) GetInfo() string {
	return fmt.Sprintf("Magazine ID: %d, Title: %s, Issue: %d, Available: %t", m.ID, m.Title, m.Issue, m.IsAvailable)
}

// 定义 Library 结构体
type Library struct {
	Books     []*Book
	Magazines []*Magazine
	Name      string
}

// 添加一本书到图书馆
func (l *Library) AddBook(book *Book) {
	l.Books = append(l.Books, book)
}

// 根据 ID 查找书籍
func (l *Library) FindBookByID(id int) *Book {
	for _, book := range l.Books {
		if book.ID == id {
			return book
		}
	}
	return nil
}

// 获取所有可借阅的书籍
func (l *Library) GetAvailableBooks() []*Book {
	var availableBooks []*Book
	for _, book := range l.Books {
		if book.IsAvailable {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// 添加一本杂志到图书馆
func (l *Library) AddMagazine(magazine *Magazine) {
	l.Magazines = append(l.Magazines, magazine)
}

// 根据 ID 查找杂志
func (l *Library) FindMagazineByID(id int) *Magazine {
	for _, magazine := range l.Magazines {
		if magazine.ID == id {
			return magazine
		}
	}
	return nil
}

// 获取所有可借阅的杂志
func (l *Library) GetAvailableMagazines() []*Magazine {
	var availableMagazines []*Magazine
	for _, magazine := range l.Magazines {
		if magazine.IsAvailable {
			availableMagazines = append(availableMagazines, magazine)
		}
	}
	return availableMagazines
}

// 打印所有可借阅的书籍和杂志
func (l *Library) PrintAvailableItems() {
	fmt.Println("Available Books:")
	for _, book := range l.GetAvailableBooks() {
		fmt.Println(book.GetInfo())
	}
	fmt.Println("Available Magazines:")
	for _, magazine := range l.GetAvailableMagazines() {
		fmt.Println(magazine.GetInfo())
	}
}

func main() {
	lib := Library{Name: "City Library"}

	// 添加书籍和杂志
	lib.AddBook(&Book{ID: 1, Title: "Go Programming", Author: "Alan", IsAvailable: true})
	lib.AddBook(&Book{ID: 2, Title: "Learning AI", Author: "Lina", IsAvailable: true})
	lib.AddBook(&Book{ID: 3, Title: "Data Structures", Author: "Max", IsAvailable: true})
	lib.AddMagazine(&Magazine{ID: 101, Title: "Science Weekly", Issue: 45, IsAvailable: true})
	lib.AddMagazine(&Magazine{ID: 102, Title: "Tech Today", Issue: 12, IsAvailable: true})

	// 打印初始状态
	fmt.Println("初始可用书籍和杂志列表:")
	lib.PrintAvailableItems()

	// 借书操作
	book := lib.FindBookByID(1)
	if book != nil && book.Borrow() {
		fmt.Println("成功借出书籍:", book.GetInfo())
	}

	// 借杂志操作
	magazine := lib.FindMagazineByID(101)
	if magazine != nil && magazine.Borrow() {
		fmt.Println("成功借出杂志:", magazine.GetInfo())
	}

	// 打印借阅后的状态
	fmt.Println("\n借阅后可用书籍和杂志列表:")
	lib.PrintAvailableItems()

	// 还书操作
	if book != nil && book.Return() {
		fmt.Println("成功归还书籍:", book.GetInfo())
	}

	// 还杂志操作
	if magazine != nil && magazine.Return() {
		fmt.Println("成功归还杂志:", magazine.GetInfo())
	}

	// 打印归还后的状态
	fmt.Println("\n归还后可用书籍和杂志列表:")
	lib.PrintAvailableItems()
}
