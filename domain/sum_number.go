package domain

type SumNumberRequest struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type SumNumberResponse struct {
    Result int `json:"result"`
}

type SumService interface {
	Sum(num1, num2 int) (int, error)
}