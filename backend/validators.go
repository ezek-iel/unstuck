package main

type Validator struct {
	errors map[string]string
}

func (v *Validator) Check(key, message string, result bool) {
	if !result {
		if _, ok := v.errors[key]; ok {
			return
		}
		v.errors[key] = message
	}
}

func NewValidator() *Validator {
	errors := make(map[string]string)
	return &Validator{errors: errors}
}
