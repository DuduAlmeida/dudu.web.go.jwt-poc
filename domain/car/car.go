package car

type Car struct {
	Modelo string `json:"modelo"`
	Marca  string `json:"marca"`
	Ano    int    `json:"ano"`
}

type CarList []Car
