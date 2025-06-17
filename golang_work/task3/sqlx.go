package task3

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// Sqlx入门
/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

// Employee 结构体对应 employees 表
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func sqlx1() {
	// 使用 Sqlx 连接数据库
	db, err := sqlx.Connect("mysql", "user:password@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		log.Fatalln(err)
	}

	// 查询部门为 "技术部" 的所有员工
	var employees []Employee
	err = db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("技术部员工:")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", emp.ID, emp.Name, emp.Salary)
	}

	// 查询工资最高的员工
	var highestSalaryEmp Employee
	err = db.Get(&highestSalaryEmp, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n工资最高的员工:\nID: %d, Name: %s, Salary: %.2f\n", highestSalaryEmp.ID, highestSalaryEmp.Name, highestSalaryEmp.Salary)
}

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

// Book 结构体对应 books 表
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func sqlx2() {
	// 使用 Sqlx 连接数据库
	db, err := sqlx.Connect("mysql", "user:password@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		log.Fatalln(err)
	}

	// 查询价格大于 50 元的书籍
	var books []Book
	err = db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("价格大于 50 元的书籍:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}
