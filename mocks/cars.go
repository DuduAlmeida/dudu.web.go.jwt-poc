package mocks

import "github.com/DuduAlmeida/dudu.web.go.jwt-poc/domain/car"

var CarsMocked = car.CarList{
	{Modelo: "SuperGT", Marca: "Velocitas", Ano: 2025},
	{Modelo: "EcoDrive", Marca: "Sustentia", Ano: 2025},
	{Modelo: "WorkVan", Marca: "Utilitario", Ano: 2025},
}
