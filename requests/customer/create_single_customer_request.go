package requests

type CreateSingleCustomerRequest struct {
	Firstname string
	Lastname  string
	Email     string
	Age       uint8
	Gender    string
}
