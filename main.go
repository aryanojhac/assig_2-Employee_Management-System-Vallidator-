// Assignment 2:- Develop a system to manage employees and their departments. Each employee should have a name, age, and salary. Each department should have a name, a list of employees, and a method to calculate the average salary of its employees. Additionally, implement methods to add and remove employees from departments and to give a raise to an employee.

package main

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Making the Structure of Employee having attibutes as Name, Age and Salary
type Employee struct {
	Name        string  `validate:"required"`
	Age         uint    `validate:"required,gte=18,lte=100"`
	Salary      float64 `validate:"required,gt=0"`
	Email       string  `validate:"required,email"`
	PANNumber   string  `validate:"required,len=10,PANNumber"`
	PhoneNumber string  `validate:"required,PhoneNumber"`
}

// Making the structure of Department having attributes as Dept Name and Employee struct
type Department struct {
	DeptName  string
	Employees []Employee
}

// Method to calculte the Vaerage Salary of the Employee
func (d *Department) averageSalary() (float64, error) {
	if len(d.Employees) == 0 {
		return 0, errors.New("The department is empty so cannot calculate the salary")
	}

	total := 0.0
	for _, emp := range d.Employees {
		total += emp.Salary
	}
	// Returning the float value for the average salary
	return float64(total) / float64(len(d.Employees)), nil
}

// Method to display all the Employees
func (d Department) listAll() {
	for _, emp := range d.Employees { // Iterating on all employees to display the details
		fmt.Printf("%-10s | %2d |  ₹%.2f | %s | %s | %s \n",
			emp.Name, emp.Age, emp.Salary, emp.Email, emp.PANNumber, emp.PhoneNumber)
	}
}

// Method to add a Employee by user
func (d *Department) addEmployeeByUser() error {
	var name string
	var age uint
	var salary float64
	var email string
	var pannumber string
	var phonenumber string

	fmt.Print("Enter the Name of Employee: ")
	if _, err := fmt.Scanln(&name); err != nil {
		return fmt.Errorf("The employee name must conatin only Alphabets. %w", err)
	}
	fmt.Print("Enter the Age of Employee: ")
	if _, err := fmt.Scanln(&age); err != nil {
		return fmt.Errorf("Age must be between 18 to 100 years. %w", err)
	}

	fmt.Print("Enter the Salary of Employee: ")
	if _, err := fmt.Scanln(&salary); err != nil {
		return fmt.Errorf("The Salary must be numeric and more than Rs0. %w", err)
	}

	fmt.Print("Enter the Email of Employee: ")
	if _, err := fmt.Scanln(&email); err != nil {
		return fmt.Errorf("The Email must contain @ and .com. %w", err)
	}

	fmt.Print("Enter the PAN Number (CapitalCase) of Employee: ")
	if _, err := fmt.Scanln(&pannumber); err != nil {
		return fmt.Errorf("PAN number is invalid, it must be Alphanumeric %w", err)
	}

	fmt.Print("Enter the Phone Number of Employee: ")
	if _, err := fmt.Scanln(&phonenumber); err != nil {
		return fmt.Errorf("The Phone Number must  %w", err)
	}

	emp := Employee{Name: name,
		Age:         age,
		Salary:      salary,
		Email:       email,
		PANNumber:   pannumber,
		PhoneNumber: phonenumber,
	}

	if err := validate.Struct(emp); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				fmt.Printf("Please give an Valid Input", e)
			}
			return errors.New("validation failed")
		}
		fmt.Println("Validation failed with:", err)
		return err
	}

	d.Employees = append(d.Employees, emp)
	fmt.Println("Congratulations! Employee successfully added.")
	return nil
}

// Method to remove an Employee
func (d *Department) removeEmployee(name string) error {
	for i, emp := range d.Employees {
		if emp.Name == name {
			d.Employees = append(d.Employees[:i], d.Employees[i+1:]...)
			fmt.Printf("Employee '%s' has been removed.\n", name)
			return nil
		}
	}
	return fmt.Errorf("Employee with name %s does not exist", name)
}

// Method to give a raise to an Employee
func (d *Department) toGiveRaise(name string, amount float64) error {
	fmt.Println("After Upraisal")
	for i, emp := range d.Employees {
		if emp.Name == name {
			emp.Salary += amount
			d.Employees[i] = emp
			fmt.Println("The Raise is given to", emp.Name)
			return nil
		}
	}
	return fmt.Errorf("the salary of %s has been raised", name)
}

// Custom function for validating the PAN_Number
func validatePAN(f1 validator.FieldLevel) bool {
	pan := f1.Field().String()
	fmt.Println("Checking PAN:", pan)
	if len(pan) != 10 {
		return false
	}
	for index := 0; index < 5; index++ { //Checking 1st 5 entries are Alphabets or not
		if pan[index] < 'A' || pan[index] > 'Z' {
			return false
		}
	}
	for index := 5; index < 9; index++ { //Checking middle 4 entries are are numbers or not
		if pan[index] < '0' || pan[index] > '9' {
			return false
		}
	}
	if pan[9] < 'A' || pan[9] > 'Z' { //Checking last entry is Alphabets or not
		return false
	}
	return true
}

// Custom function to validate phone number
func validatePhone(f1 validator.FieldLevel) bool {
	phone := f1.Field().String()
	// Allowing 10-digit phone number
	if len(phone) == 10 {
		for _, ch := range phone {
			if ch < '0' || ch > '9' {
				return false
			}
		}
		return true
	}

	if len(phone) == 13 && phone[0] == '+' {
		for i := 1; i < 13; i++ {
			if phone[i] < '0' || phone[i] > '9' {
				return false
			}
		}
		return true
	}
	return false
}
func main() {

	// Validating things goes this way
	validate = validator.New() //New validator instance

	// Registering the custom validator functions
	validate.RegisterValidation("PANNumber", validatePAN)
	validate.RegisterValidation("PhoneNumber", validatePhone)

	dept := Department{
		DeptName: "IT",

		// Adding employees manually (Hardcoded)
		Employees: []Employee{
			{"Aryan", 21, 12500, "aryan@gmail.com", "AFLPO9601H", "7058242415"},
			{"Atharv", 23, 11000, "atharv@gmail.com", "WASTG9876J", "1234567892"},
			{"Purva", 25, 75000, "purva@gmail.com", "OKJUY6543Y", "9876542108"},
		},
	}

	// Choices for the admin
	for {
		fmt.Println("Menu for the Admin:")
		fmt.Println("1. See the List of all Employees | 2. See the Average salary | 3. Add a New Employee | 4. Give Raise to an Employee | 5. Remove an Employee | 6. Exit")
		var choice uint
		fmt.Println("Please Enter your Choice (1 to 6)")
		fmt.Scanln(&choice)

		switch choice {
		case 1: // Listing all the Employees
			fmt.Println("The first list of Employees:")
			dept.listAll()

		case 2: // Printing average salary of Employees
			avgSalary, err := dept.averageSalary()
			if err != nil {
				fmt.Println("Cannot calculate average salary.", err)
			} else {
				fmt.Printf("\nAverage salary in %s: ₹%.2f\n", dept.DeptName, avgSalary)
			}

		case 3: // Create and add a new employee from user input
			err := dept.addEmployeeByUser()
			if err != nil && err.Error() != "no new employee to add" {
				fmt.Println("Failed to add employee. Please try again.", err)
			}

		case 4: // Giving Raise to the employee
			var name string
			var raise float64

			fmt.Print("Enter employee name to give raise: ")
			fmt.Scanln(&name)
			fmt.Print("Enter raise amount: ")
			fmt.Scanln(&raise)

			err := dept.toGiveRaise(name, raise)
			if err != nil {
				fmt.Println("The employee whom you want to give raise is not in the organization.", err)
			}

		case 5: // To remove a employee
			var name string
			fmt.Print("Enter employee name to remove: ")
			fmt.Scanln(&name)
			err := dept.removeEmployee(name)
			if err != nil {
				fmt.Println("The employee you want to remove is not in the organization.", err)
			}

		case 6: // Exit the code
			fmt.Println("Exiting program. Thank You!")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option (1-6).")
		}
	}
}
