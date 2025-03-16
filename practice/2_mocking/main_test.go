package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// smsServiceMock implements the MessageService interface
type smsServiceMock struct {
	mock.Mock
}

// Implement the MessageService interface method
func (m *smsServiceMock) SendChargeNotification(value int) error {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	args := m.Called(value)
	// Return error (nil or actual error) instead of bool
	return args.Error(0)
}

// TestChargeCustomer tests the ChargeCustomer functionality
func TestChargeCustomer(t *testing.T) {
	// Create new mock instance
	smsService := new(smsServiceMock)

	// Set up expectation: when SendChargeNotification is called with 100,
	// return nil error to simulate successful notification
	smsService.On("SendChargeNotification", 100).Return(nil)

	// Create MyService instance with our mock
	myService := MyService{
		messageService: smsService,
	}

	// Call the method we want to test
	err := myService.ChargeCustomer(100)
	if err != nil {
		t.Errorf("ChargeCustomer returned an error: %v", err)
	}

	// Verify that all expected mock calls were made
	smsService.AssertExpectations(t)
}
