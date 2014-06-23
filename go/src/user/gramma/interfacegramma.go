package gramma

import (
	"fmt"
	"reflect"
)

func BaseInterface() {
	fmt.Println("#########BaseInterface###########")
	followHumans()

	reflection()
	fmt.Println("#########BaseInterface###########")
}

type Human interface {
	Name() string
	Age() int
	Tell()
}

type HumanX interface {
	Human
	Answer(q Question)
}

type Question interface {
	Title() string
	Body() string
}

type Person2 struct {
	MName string
	mAge  int
	Sex   string
}

type Student2 struct {
	Person2
	mSchool string
}

func (p Person2) Name() string {
	return p.MName
}

func (p *Person2) PName() string {
	return p.MName
}

func (p Person2) Age() int {
	return p.mAge
}

func (p Person2) Tell() {
	fmt.Printf("%s is %d years old.\n", p.MName, p.mAge)
}

func (s Student2) Tell() {
	fmt.Printf("%s is %d years old, studying in %s.\n", s.MName, s.mAge, s.mSchool)
}

func (p Person2) String() string {
	return "This is person2 " + p.Name()
}

func likeHuman(h Human) {
	fmt.Printf("[%s %d]: ", h.Name(), h.Age())
	h.Tell()

}

func followHumans() {
	p1 := Person2{MName: "da wang", mAge: 33}
	p2 := Student2{Person2: Person2{MName: "xiao wang", mAge: 6}, mSchool: "yi fu zhong"}

	(&p1).PName()
	p1.PName()
	fmt.Println(p1)
	likeHuman(p1)
	fmt.Println(realHumanType(p1))
	fmt.Println(p2)
	likeHuman(p2)
	fmt.Println(realHumanType(p2))
}

type TypeContainer interface{}

// judge comma-ok
func realHumanType(h Human) string {
	var t TypeContainer = h

	// method 1
	switch t.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case Person2:
		return "Person2"
	case Student2:
		return "Student2"
	}

	// method 2
	if _, ok := t.(float64); ok {
		return "float"
	}

	return "nothing"
}

// reflection
func reflection() {
	x := Person2{MName: "da wang", mAge: 33}
	t := reflect.TypeOf(&x)
	v := reflect.ValueOf(&x)

	fmt.Println(t) // *gramma.Person2
	fmt.Println(v) // <*gramma.Person2 Value>

	// for pointer(*T), use Elem(), means addressable
	tag := t.Elem().Field(0).Name
	name := v.Elem().Field(0).String()
	canSet := v.Elem().Field(0).CanSet()
	canAddr := v.Elem().Field(0).CanAddr()
	canInterface := v.CanInterface()

	x.MName = "da zhang"
	// because the first letter is lower, it can not be set out of this package(be privated)
	// v.Elem().FieldByName("mAge").SetInt(31)
	v.Elem().FieldByName("MName").SetString("da wang")
	v.Elem().Field(0).SetString("da wang")
	v.Elem().FieldByName("Sex").SetString("male")

	fmt.Println(tag)
	fmt.Println(name)
	fmt.Println(v.Kind(), canSet, " ", canAddr, " ", canInterface)

	// this reflection is readonly
	t2 := reflect.TypeOf(x)
	v2 := reflect.ValueOf(x)

	tag2 := t2.Field(0).Name
	name2 := v2.Field(0).String()
	canSet = v2.Field(0).CanSet()
	canAddr = v2.Field(0).CanAddr()
	canInterface = v2.CanInterface()

	// this will be fail, bacause x is not addressable, but *x is.
	// v2.FieldByName("MName").SetString("da wang")
	fmt.Println(t2)
	fmt.Println(v2)
	fmt.Println(tag2)
	fmt.Println(name2)
	fmt.Println(v2.Kind(), canSet, " ", canAddr, " ", canInterface)

	it := v2.Interface()
	if _, ok := it.(Person2); ok {
		fmt.Println("It's Person2")
	}

	var h Human = x
	//var h interface{} = x
	t3 := reflect.TypeOf(h)  // gramma.Person2
	v3 := reflect.ValueOf(h) // <gramma.Person2 Value>

	tag3 := t3.Field(0).Name
	name3 := v3.Field(0).String()
	canSet = v3.Field(0).CanSet()
	canAddr = v3.Field(0).CanAddr()
	canInterface = v3.CanInterface()

	fmt.Println(t3)
	fmt.Println(v3)
	fmt.Println(tag3)
	fmt.Println(name3)
	fmt.Println(v3.Kind(), canSet, " ", canAddr, " ", canInterface)

}
