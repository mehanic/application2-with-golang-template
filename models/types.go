package models

import "time"

// Структура для передачи данных в шаблон
type PageData struct {
	Title      string
	Heading    string
	Message    string
	Form       map[string]string
	StatusLine string
}
type Information struct {
	PageData
	Items map[string]string
	Sage  Sage
}
type Sage struct {
	Name  string
	Motto string
}

type Sagen struct {
	Name1  string
	Motto1 string
}
type Car struct {
	Manufacturer string
	Model        string
	Doors        int
}
type AboutData struct {
	Information
	Wisdom    []Sagen
	Transport []Car
}
type SecondSaga struct {
	Name2  string
	Motto2 string
}

type SecondCar struct {
	Manufacturer string
	Model        string
	Doors        int
}

type NewData struct {
	PageData
	Thinkers []SecondSaga
	Vehicles []SecondCar
	SomeDate time.Time
	Number   int
	FloatNum float64
}

type User struct {
	Name  string
	Motto string
	Admin bool
}

type Course struct {
	Number string
	Name   string
	Units  string
}

type Semester struct {
	Term    string
	Courses []Course
}

type Year struct {
	AcaYear string
	Fall    Semester
	Spring  Semester
	Summer  Semester
}

type EducationData struct {
	PageData
	Years []Year
}

type Hotel struct {
	Name, Address, City, Zip, Region string
}

type Region struct {
	Region string
	Hotels []Hotel
}

type RelaxData struct {
	PageData
	Regions []Region
	Menu    Items
}

type Item struct {
	Name, Descrip, Meal string
	Price               float64
}

type Items []Item

type Meal struct {
	Meal string
	Item []Item
}

type Menu []Meal

type Record struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int64
	AdjClose float64
}

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}
