package gramma

import "fmt"

func BaseStruct() {
	usePerson()
	useStudent()
	useObjectMethod()
}

type Person struct {
	name string
	age  int
}

func usePerson() {
	var p Person
	p.name = "wang da nian"
	p.age = 30

	fmt.Println(p)

	var pp1 *Person = new(Person) // poniter init with new
	pp1.name = "wang xiao jun"
	pp1.age = 5

	fmt.Println(pp1)

	p2 := Person{name: "li xiao mei", age: 28}
	fmt.Println(p2)
}

type Student struct {
	Person     // anonymous field,  means Inherit from Person
	SchoolName string
	Score      int
	int        // anonymous field, use like student_x.int
}

type CollegeStudent struct {
	Student
	Score int
}

func useStudent() {
	// how to use the anonymous field
	xiaoming := Student{SchoolName: "er zhong"}
	xiaoming.name = "xiaoming"
	xiaoming.Person.age = 12
	xiaoming.Score = 80
	xiaoming.int = 2

	fmt.Println(xiaoming)

	xiaohong := Student{Person: Person{name: "xiaohong", age: 12}, SchoolName: "er zhong", Score: 90, int: 3}
	fmt.Println(xiaohong)

	daming := CollegeStudent{Student: Student{Person: Person{name: "daming", age: 22}, SchoolName: "song da", Score: 90, int: 3}}
	daming.Score = 90
	daming.Student.Score = 80
	fmt.Println(daming)
}

type Point struct {
	X int64
	Y int64
}

type Rect struct {
	TopLeft Point
	Width   float64
	Height  float64
}

type CirCle struct {
	Center Point
	Radius float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

const pi = 3.1415926

func (c CirCle) Area() float64 {
	return 2 * pi * c.Radius
}

func useObjectMethod() {
	rect1 := Rect{Width: 100, Height: 200}
	circle1 := CirCle{Radius: 300, Center: Point{X: 1, Y: 2}}
	rect1.TopLeft.X = 1
	rect1.TopLeft.Y = 2

	fmt.Println(rect1.Area())
	fmt.Println(circle1.Area())

	typeofwhat()

	employeeTell()
}

type Rects []Rect
type Los []int

func (rs Rects) HowMany() int {
	return len(rs)
}

func (rs Los) HowMany() int {
	return len(rs)
}

func typeofwhat() {
	var rects Rects
	rects = make(Rects, 10)
	fmt.Println(rects.HowMany())

	var loss Los = []int{1, 2, 3}
	fmt.Println(loss.HowMany())
}

type Employee struct {
	Person
	Position string
	Salary   int
}

type FirstLine struct {
	Employee
}

type Manager struct {
	Employee
	ManLevel uint
}

func (e Employee) Tell() {
	fmt.Printf("My name is %v. I am in %s position\n", e.name, e.Position)
}

func (e Manager) Tell() {
	fmt.Printf("My name is %s. I am in %s position, level %d\n", e.name, e.Position, e.ManLevel)
}

func employeeTell() {
	emp1 := Employee{Person: Person{name: "Steve Mona", age: 40}, Position: "Market Senior Acer", Salary: 500}
	emp2 := FirstLine{Employee: Employee{Person: Person{name: "Steve Mona", age: 40}, Position: "Market Senior Acer", Salary: 500}}
	emp3 := Manager{Employee: Employee{Person: Person{name: "Jerry Fonda", age: 40}, Position: "IT Support Manager", Salary: 400}, ManLevel: 3}

	emp1.Tell()
	emp2.Tell()
	emp3.Tell()
}
