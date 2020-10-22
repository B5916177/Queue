// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/B5916177/app/ent/dentist"
	"github.com/B5916177/app/ent/employee"
	"github.com/B5916177/app/ent/patient"
	"github.com/B5916177/app/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	dentistFields := schema.Dentist{}.Fields()
	_ = dentistFields
	// dentistDescDentistName is the schema descriptor for dentist_name field.
	dentistDescDentistName := dentistFields[0].Descriptor()
	// dentist.DentistNameValidator is a validator for the "dentist_name" field. It is called by the builders before save.
	dentist.DentistNameValidator = dentistDescDentistName.Validators[0].(func(string) error)
	// dentistDescDentistEmail is the schema descriptor for dentist_email field.
	dentistDescDentistEmail := dentistFields[1].Descriptor()
	// dentist.DentistEmailValidator is a validator for the "dentist_email" field. It is called by the builders before save.
	dentist.DentistEmailValidator = dentistDescDentistEmail.Validators[0].(func(string) error)
	// dentistDescDentistPhone is the schema descriptor for dentist_phone field.
	dentistDescDentistPhone := dentistFields[2].Descriptor()
	// dentist.DentistPhoneValidator is a validator for the "dentist_phone" field. It is called by the builders before save.
	dentist.DentistPhoneValidator = dentistDescDentistPhone.Validators[0].(func(int) error)
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescEmployeeName is the schema descriptor for employee_name field.
	employeeDescEmployeeName := employeeFields[0].Descriptor()
	// employee.EmployeeNameValidator is a validator for the "employee_name" field. It is called by the builders before save.
	employee.EmployeeNameValidator = employeeDescEmployeeName.Validators[0].(func(string) error)
	// employeeDescEmployeePhone is the schema descriptor for employee_phone field.
	employeeDescEmployeePhone := employeeFields[3].Descriptor()
	// employee.EmployeePhoneValidator is a validator for the "employee_phone" field. It is called by the builders before save.
	employee.EmployeePhoneValidator = employeeDescEmployeePhone.Validators[0].(func(int) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientName is the schema descriptor for patient_name field.
	patientDescPatientName := patientFields[0].Descriptor()
	// patient.PatientNameValidator is a validator for the "patient_name" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	// patientDescPatientGender is the schema descriptor for patient_gender field.
	patientDescPatientGender := patientFields[1].Descriptor()
	// patient.PatientGenderValidator is a validator for the "patient_gender" field. It is called by the builders before save.
	patient.PatientGenderValidator = patientDescPatientGender.Validators[0].(func(string) error)
	// patientDescPatientPhone is the schema descriptor for patient_phone field.
	patientDescPatientPhone := patientFields[3].Descriptor()
	// patient.PatientPhoneValidator is a validator for the "patient_phone" field. It is called by the builders before save.
	patient.PatientPhoneValidator = patientDescPatientPhone.Validators[0].(func(int) error)
}