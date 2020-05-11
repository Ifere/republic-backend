package errors
import "fmt"

//This package  defines custom app errors
type NotCreated struct {
	Resource string
}

type ErrorGetting struct {
	Resource string
}

type ErrorUpdating struct {
	Resource string
}

type ErrorDeleting struct {
	Resource string
}

type DuplicateError struct {
	Resource string
}

type InvalidLoginCredentials struct {
}

type NotFound struct {
	Resource string
}

type ErrorVerifying struct {
	Resource string
}

type ErrorConverting struct {
	Resource1 string
	Resource2 string
}

type ErrorTransferringBudget struct {
}

type ErrorSaving struct {
	Resource string
}
type ErrorCharging struct {
	Resource string
}

type ErrorFreezing struct {
	Resource string
}

type ErrorFunding struct {
	Resource string
}

type InsufficientBudgetFunds struct{}

type ErrorAllocatingFunds struct{}

type ErrorCreditingBudget struct{}

type DuplicatedTransactionError struct{}

type InsufficientTransferFunds struct{}

type ErrorTransferringFunds struct {
}

func (e NotCreated) Error() string {
	err := fmt.Sprintf("unable to create %s at this time", e.Resource)
	return err
}

func (e ErrorGetting) Error() string {
	return fmt.Sprintf("unable to get %s at this time", e.Resource)
}

func (e ErrorUpdating) Error() string {
	return fmt.Sprintf("unable to update %s at this time", e.Resource)
}

func (e ErrorDeleting) Error() string {
	return fmt.Sprintf("unable to delete %s at this time", e.Resource)
}

func (e DuplicateError) Error() string {
	return fmt.Sprintf("%s already exsist", e.Resource)
}

func (e InvalidLoginCredentials) Error() string {
	return fmt.Sprintf("invalid email or password")
}

func (e NotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

func (e ErrorVerifying) Error() string {
	return fmt.Sprintf("unable to verify %v", e.Resource)
}

func (e ErrorConverting) Error() string {
	return fmt.Sprintf("unable to convert %v to %v", e.Resource1, e.Resource2)
}

func (e ErrorTransferringBudget) Error() string {
	return fmt.Sprintf("unable to transfer budget at this time")
}

func (e ErrorAllocatingFunds) Error() string {
	return "error allocation funds at this time"
}

func (e InsufficientBudgetFunds) Error() string {
	return "budget does not have enough funds for this allocation"
}

func (e ErrorCreditingBudget) Error() string {
	return "unable to credit budget at this time"
}

func (e DuplicatedTransactionError) Error() string {
	return "transaction already occurred"
}

func (e ErrorCharging) Error() string {
	return fmt.Sprintf("unable to charge %v at this time", e.Resource)
}

func (e ErrorSaving) Error() string {
	return fmt.Sprintf("unable to save %v at this time", e.Resource)
}

func (e InsufficientTransferFunds) Error() string {
	return fmt.Sprintf("insufficeint funds")
}

func (e ErrorTransferringFunds) Error() string {
	return fmt.Sprintf("unable to transfer funds at this time ")
}

func (e ErrorFreezing) Error() string {
	return fmt.Sprintf("unable to freeze %v at this time", e.Resource)
}

func (e ErrorFunding) Error() string {
	return fmt.Sprintf("unable to fund %v at this time", e.Resource)
}
