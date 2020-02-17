package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker interface {
	printPosition()
	printNameOfWork()
	printNameOfWorker()
}
type Boss interface {
	printInformationAboutBoss()
}

//person

type person struct {
	age         int
	gender      string
	sitizenShip string
	lastName    string
	firstName   string
}

func (p person) printPosition() {

}
func (p person) printNameOfWork() {

}
func (p person) printNameOfWorker() {

}
func newPerson(age int, gender, sitizenShip, lastName, firstName string) person {
	return person{age: age, gender: gender, sitizenShip: sitizenShip, lastName: lastName, firstName: firstName}
}

//person

//programming
type programmerBoss struct {
	person
}

func newProgrammerBoss(age int, gender, sitizenShip, lastName, firstName string) programmerBoss {
	return programmerBoss{newPerson(age, gender, sitizenShip, lastName, firstName)}
}
func (pb programmerBoss) printInformationAboutBoss() {
	fmt.Println("Programmer boss. My name is: ", pb.firstName, " ", pb.lastName)
}

type programmer struct {
	person
	position                    string
	nameOfWork                  string
	quantityOfProgrammsPerMonth int
}

func (p programmer) printPosition() {
	fmt.Println("Worker position: ", p.position)
}
func (p programmer) printNameOfWork() {
	fmt.Println("Name of work: ", p.nameOfWork)
}
func (p programmer) printNameOfWorker() {
	fmt.Println("Worker name: ", p.firstName, " ", p.lastName)
}
func newProgrammer(age, quantityOfProgrammsPerMonth int, gender, sitizenShip, lastName, firstName, nameOfWork, position string) programmer {
	return programmer{
		person:                      newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork:                  nameOfWork,
		position:                    position,
		quantityOfProgrammsPerMonth: quantityOfProgrammsPerMonth,
	}
}

//programming

//bank
type bankerBoss struct {
	person
}

func newBankerBoss(age int, gender, sitizenShip, lastName, firstName string) bankerBoss {
	return bankerBoss{newPerson(age, gender, sitizenShip, lastName, firstName)}
}
func (bb bankerBoss) printInformationAboutBoss() {
	fmt.Println("Banker boss. My name is: ", bb.firstName, " ", bb.lastName)
}

type banker struct {
	person
	position          string
	nameOfWork        string
	hoursToWorkPerDay int
}

func (b banker) printPosition() {
	fmt.Println("Worker position: ", b.position)
}
func (b banker) printNameOfWork() {
	fmt.Println("Name of work: ", b.nameOfWork)
}
func (b banker) printNameOfWorker() {
	fmt.Println("Worker name: ", b.firstName, " ", b.lastName)
}
func newBanker(age, hoursToWorkPerDay int, gender, sitizenShip, lastName, firstName, nameOfWork, position string) banker {
	return banker{
		person:            newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork:        nameOfWork,
		position:          position,
		hoursToWorkPerDay: hoursToWorkPerDay,
	}
}

//bank

//taxi
type taxiDriverBoss struct {
	person
}

func newTaxiDriverBoss(age int, gender, sitizenShip, lastName, firstName string) taxiDriverBoss {
	return taxiDriverBoss{newPerson(age, gender, sitizenShip, lastName, firstName)}
}
func (tb taxiDriverBoss) printInformationAboutBoss() {
	fmt.Println("Taxi driver boss. My name is: ", tb.firstName, " ", tb.lastName)
}

type taxiDriver struct {
	person
	position              string
	nameOfWork            string
	quantityOfRidesPerDay int
}

func (t taxiDriver) printPosition() {
	fmt.Println("Worker position: ", t.position)
}
func (t taxiDriver) printNameOfWork() {
	fmt.Println("Name of work: ", t.nameOfWork)
}
func (t taxiDriver) printNameOfWorker() {
	fmt.Println("Worker name: ", t.firstName, " ", t.lastName)
}
func newTaxiDriver(age, quantityOfRidesPerDay int, gender, sitizenShip, lastName, firstName string, nameOfWork, position string) taxiDriver {
	return taxiDriver{
		person:                newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork:            nameOfWork,
		position:              position,
		quantityOfRidesPerDay: quantityOfRidesPerDay,
	}
}

//taxi

//school
type teacherBoss struct {
	person
}

func newTeacherBoss(age int, gender, sitizenShip, lastName, firstName string) teacherBoss {
	return teacherBoss{newPerson(age, gender, sitizenShip, lastName, firstName)}
}
func (tb teacherBoss) printInformationAboutBoss() {
	fmt.Println("Teacher boss. My name is: ", tb.firstName, " ", tb.lastName)
}

type teacher struct {
	person
	position               string
	nameOfWork             string
	quantityOfPupilsPerDay int
}

func (t teacher) printPosition() {
	fmt.Println("Worker position: ", t.position)
}
func (t teacher) printNameOfWork() {
	fmt.Println("Name of work: ", t.nameOfWork)
}
func (t teacher) printNameOfWorker() {
	fmt.Println("Worker name: ", t.firstName, " ", t.lastName)
}
func newTeacher(age, quantityOfPupilsPerDay int, gender, sitizenShip, lastName, firstName, nameOfWork, position string) teacher {
	return teacher{
		person:                 newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork:             nameOfWork,
		position:               position,
		quantityOfPupilsPerDay: quantityOfPupilsPerDay,
	}
}

//school

//factory
type workerOfFactoryBoss struct {
	person
}

func newWorkerOfFactoryBoss(age int, gender, sitizenShip, lastName, firstName string) workerOfFactoryBoss {
	return workerOfFactoryBoss{newPerson(age, gender, sitizenShip, lastName, firstName)}
}
func (wb workerOfFactoryBoss) printInformationAboutBoss() {
	fmt.Println("Worker boss. My name is: ", wb.firstName, " ", wb.lastName)
}

type workerOfFactory struct {
	person
	position              string
	nameOfWork            string
	quantityOfWorkPerHour int
}

func (w workerOfFactory) printPosition() {
	fmt.Println("Worker position: ", w.position)
}
func (w workerOfFactory) printNameOfWork() {
	fmt.Println("Name of work: ", w.nameOfWork)
}
func (w workerOfFactory) printNameOfWorker() {
	fmt.Println("Worker name: ", w.firstName, " ", w.lastName)
}
func newWorkerOfFactory(age, quantityOfWorkPerHour int, gender, sitizenShip, lastName, firstName, nameOfWork, position string) workerOfFactory {
	return workerOfFactory{
		person:                newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork:            nameOfWork,
		position:              position,
		quantityOfWorkPerHour: quantityOfWorkPerHour,
	}
}

//factory

//other
var mt sync.Mutex
var wg sync.WaitGroup

func printInformationAboutAllBoses(boses []Boss) {
	mt.Lock()
	for _, value := range boses {
		value.printInformationAboutBoss()
	}
	mt.Unlock()

}
func printInformationAboutAllWorkers(workers map[string]Worker) {
	mt.Lock()
	for _, value := range workers {
		value.printNameOfWorker()
		value.printNameOfWork()
		value.printPosition()
	}
	mt.Unlock()
}

// func convertWorkerToHuman(w Worker) person {
// 	switch p := w.(type) {
// 	case programmer:
// 		return p.(person)
// 	}
// }

//other

func main() {

	workers := make(map[string]Worker, 5)
	teacher := newTeacher(
		30,
		15,
		"Male",
		"Ukrainian",
		"Sydorenko",
		"Denis",
		"School",
		"Teacher",
	)
	workers[teacher.lastName+" "+teacher.firstName] = teacher
	taxiDriver := newTaxiDriver(
		45,
		10,
		"Male",
		"American",
		"Elton",
		"John",
		"Taxi",
		"Driver",
	)
	workers[taxiDriver.lastName+" "+taxiDriver.firstName] = taxiDriver
	programmer := newProgrammer(
		25,
		4,
		"Female",
		"Indian",
		"Ghas",
		"Matum",
		"Programming",
		"Programmer",
	)
	workers[programmer.lastName+" "+programmer.firstName] = programmer
	banker := newBanker(
		32,
		4,
		"Female",
		"Canadian",
		"Pokidko",
		"Danya",
		"Bank",
		"Banker",
	)
	workers[banker.lastName+" "+banker.firstName] = banker
	factoryWorker := newWorkerOfFactory(
		15,
		7,
		"Male",
		"Chinese",
		"Kolbaca",
		"Vasilek",
		"Factory",
		"Worker",
	)
	workers[factoryWorker.lastName+" "+factoryWorker.firstName] = factoryWorker

	bosses := make([]Boss, 0)
	var programmerBoss programmerBoss = newProgrammerBoss(44, "Female", "Polish", "Kirz", "kobain")
	bosses = append(bosses, programmerBoss)
	var bankerBoss bankerBoss = newBankerBoss(65, "Female", "British", "Gontaeva", "Valeriy")
	bosses = append(bosses, bankerBoss)
	var taxiDriverBoss taxiDriverBoss = newTaxiDriverBoss(30, "Male", "Moldova", "Igif", "Rate")
	bosses = append(bosses, taxiDriverBoss)
	var workerOfFactoryBoss workerOfFactoryBoss = newWorkerOfFactoryBoss(50, "Male", "Belarus", "Goroh", "Vanya")
	bosses = append(bosses, workerOfFactoryBoss)
	var teacherBoss teacherBoss = newTeacherBoss(40, "Female", "Ukrainian", "Metuh", "Natasha")
	bosses = append(bosses, teacherBoss)

	wg.Add(1)
	for i := 0; i < 5; i++ {
		go printInformationAboutAllBoses(bosses)
		time.Sleep(time.Second)
		fmt.Println()
		fmt.Println()
	}
	wg.Done()
	for i := 0; i < 5; i++ {
		go printInformationAboutAllWorkers(workers)
		time.Sleep(time.Second)
		fmt.Println()
		fmt.Println()
	}

	fmt.Scanln()
}
