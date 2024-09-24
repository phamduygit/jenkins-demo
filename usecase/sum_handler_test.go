package usecase

import (
	"errors"
	"testing"

	"github.com/phamduygit/jenkins-demo/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking SumService to isolate and test SumHandler independently
type MockSumService struct {
	mock.Mock
}

func (m *MockSumService) Sum(number1, number2 int) (int, error) {
	args := m.Called(number1, number2)
	return args.Int(0), args.Error(1)
}

func TestSumHandler_Sum_Success(t *testing.T) {
	mockSumService := new(MockSumService)
	handler := NewSumHandler(mockSumService)

	// Test data
	req := &domain.SumNumberRequest{Number1: 5, Number2: 10}
	expectedResponse := &domain.SumNumberResponse{Result: 15}

	// Setting up mock behavior
	mockSumService.On("Sum", 5, 10).Return(15, nil)

	// Execute the function
	resp, err := handler.Sum(req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedResponse.Result, resp.Result)

	// Verifying that the mock method was called
	mockSumService.AssertCalled(t, "Sum", 5, 10)
}

func TestSumHandler_Sum_Error(t *testing.T) {
	mockSumService := new(MockSumService)
	handler := NewSumHandler(mockSumService)

	// Test data
	req := &domain.SumNumberRequest{Number1: 5, Number2: 10}

	// Simulate an error returned by SumService
	mockSumService.On("Sum", 5, 10).Return(0, errors.New("numbers must be positive"))

	// Execute the function
	resp, err := handler.Sum(req)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.EqualError(t, err, "failed: numbers must be positive")

	// Verifying that the mock method was called
	mockSumService.AssertCalled(t, "Sum", 5, 10)
}

func TestNewSumHandler(t *testing.T) {
	mockSumService := new(MockSumService)

	// Test the constructor
	handler := NewSumHandler(mockSumService)

	// Assertions
	assert.NotNil(t, handler)
	assert.Equal(t, mockSumService, handler.SumService)
}
