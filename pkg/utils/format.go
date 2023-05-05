package utils

import "errors"

func FormatCPF(cpf string) (string, error) {
	cpf = GetOnlyNumbers(cpf)
	if !ValidateCPF(cpf) {
		return "", errors.New("CPF inválido")
	}
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:], nil
}
