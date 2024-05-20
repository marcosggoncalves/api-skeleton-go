package request

type Login struct {
	CPF   string `json:"cpf"`
	Senha string `json:"senha"`
}
