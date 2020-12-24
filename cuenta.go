package main

type Cuenta struct {
	Name     string `json:"name"`
	Rfc      string `json:"rfc"`
	NoCuenta int    `json:"no_cuenta"`
}

type Cuentas []Cuenta
