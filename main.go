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
	Name         string  `validate:"required"`
	Age          uint    `validate:"required,gte=18,lte=100"`
	Salary       float64 `validate:"required,gt=0"`
	Email        string  `validate:"required,email"`
	PAN_Number   string  `validate:"required,len=10,pan"`
	Phone_Number string  `validate:"required,len=13,phone"`
}

// Making the structure of Department having attributes as Dept Name and Employee struct
type Department struct {
	DeptName  string
	Employees []Employee
}

// Method to calculte the Vaerage Salary of the Employee
func (d Department) averageSalary() (float64, error) {
	if len(d.Employees) == 0 {
		return 0, errors.New("the department is empty so cannot calculate the salary")
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
			emp.Name, emp.Age, emp.Salary, emp.Email, emp.PAN_Number, emp.Phone_Number)
	}
}

// Method to add a Employee by user
func (d *Department) addEmployeeByUser() error {
	var name string
	var age uint
	var salary float64
	var email string
	var pan_number string
	var phone_number string

	fmt.Print("Enter the Name of Employee: ")
	if _, err := fmt.Scanln(&name); err != nil {
		return fmt.Errorf("invalid age: %w", err)
	}
	fmt.Print("Enter the Age of Employee: ")
	if _, err := fmt.Scanln(&age); err != nil {
		return fmt.Errorf("invalid age: %w", err)
	}

	fmt.Print("Enter the Salary of Employee: ")
	if _, err := fmt.Scanln(&salary); err != nil {
		return fmt.Errorf("invalid salary: %w", err)
	}

	fmt.Print("Enter the Email of Employee: ")
	if _, err := fmt.Scanln(&email); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}

	fmt.Print("Enter the PAN Number (CapitalCase) of Employee: ")
	if _, err := fmt.Scanln(&pan_number); err != nil {
		return fmt.Errorf("invalid PAN Number: %w", err)
	}

	fmt.Print("Enter the Phone Number of Employee: ")
	if _, err := fmt.Scanln(&phone_number); err != nil {
		return fmt.Errorf("invalid Phone Number: %w", err)
	}

	emp := Employee{Name: name,
		Age:          age,
		Salary:       salary,
		Email:        email,
		PAN_Number:   pan_number,
		Phone_Number: phone_number,
	}

	if err := validate.Struct(emp); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				fmt.Printf("Validation error: Field '%s' failed on '%s'\n", e.Field(), e.Tag())
			}
			return errors.New("validation failed")
		}
		fmt.Println("Validation failed with:", err)
		return err
	}

	d.Employees = append(d.Employees, emp)
	fmt.Println("Employee successfully added.")
	return nil
}

// Method to remove an Employee
func (d Department) removeEmployee(name string) (Department, error) {
	for i, emp := range d.Employees {
		if emp.Name == name {
			d.Employees = append(d.Employees[:i], d.Employees[i+1:]...)
			fmt.Printf("Employee '%s' has been removed.\n", name)
			return d, nil
		}
	}
	return d, fmt.Errorf("Employee with name %s does not exist", name)
}

// Method to give a raise to an Employee
func (d Department) toGiveRaise(name string, amount float64) (Department, error) {
	fmt.Println("After Upraisal")
	for i, emp := range d.Employees {
		if emp.Name == name {
			emp.Salary += amount
			d.Employees[i] = emp
			fmt.Println("The Raise is given to", emp.Name)
			return d, nil
		}
	}
	return d, fmt.Errorf("the salary of %s has been raised", name)
}

// Custom function for validating the PAN_Number
func validatePAN(f1 validator.FieldLevel) bool {
	pan := f1.Field().String()
	fmt.Println("Checking PAN:", pan)
	if len(pan) != 10 {
		return false
	}
	for index := 0; index < 5; index++ {
		if pan[index] < 'A' || pan[index] > 'Z' {
			fmt.Println("First")
			return false

		}
	}
	for index := 5; index < 9; index++ {
		if pan[index] < '0' || pan[index] > '9' {
			fmt.Println("2nd")
			return false
		}
	}
	if pan[9] < 'A' || pan[9] > 'Z' {
		fmt.Println("last")
		return false
	}
	return true
}

// Custom function to validate phone number
func validatePhone(f1 validator.FieldLevel) bool {
	phone := f1.Field().String()
	if len(phone) == 10 {
		for _, ph := range phone {
			if ph < '0' || ph > '9' {
				return false
			}
		}
		return true
	} else if len(phone) == 13 {
		if phone[0] != '+' || phone[1] != '9' || phone[2] != '1' {
			return false
		}
		for index := 3; index < 13; index++ {
			if phone[index] < '0' || phone[index] > '9' {
				return false
			}
		}
	}
	fmt.Println("Enter phone number including the country code (+1,+91,etc)")
	return true
}
func main() {

	// Validating things goes this way
	validate = validator.New() //New validator instance

	// Registering the custom validator functions
	validate.RegisterValidation("pan", validatePAN)
	validate.RegisterValidation("phone", validatePhone)

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
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("\nAverage salary in %s: ₹%.2f\n", dept.DeptName, avgSalary)
			}
		case 3: // Create and add a new employee from user input
			err := dept.addEmployeeByUser()
			if err != nil && err.Error() != "no new employee to add" {
				fmt.Println("Error adding employee:", err)
			}
		case 4: // Giving Raise to the employee
			var name string
			var raise float64

			fmt.Print("Enter employee name to give raise: ")
			fmt.Scanln(&name)
			fmt.Print("Enter raise amount: ")
			fmt.Scanln(&raise)

			updatedDept, err := dept.toGiveRaise(name, raise)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				dept = updatedDept
			}
		case 5: // To remove a employee
			var name string
			fmt.Print("Enter employee name to remove: ")
			fmt.Scanln(&name)

			updatedDept, err := dept.removeEmployee(name)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				dept = updatedDept
			}
		case 6: // Exit the code
			fmt.Println("Exiting program. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option (1-6).")
		}
	}
}