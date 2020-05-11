package responses

type GeneralResponse struct {
	Success		bool		`json:"success"`
	Data		interface{}	`json:"data"`
	Message		interface{}	`json:"message"`
	Error		string		`json:"error"`
}