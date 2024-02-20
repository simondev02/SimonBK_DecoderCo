package service

import (
	"SimonBK_DecoderCo/db"
	"fmt"
)

type VehicleInfo struct {
	Plate       string
	Imei        string
	VehicleType string
	IdCompany   int
	Company     string
	IdCustomer  int
	Customer    string
}

func GetVehiclesInfo(imei string) (VehicleInfo, error) {
	// Crear una conexión a la base de datos
	db, err := db.ConnectDirectly()
	if err != nil {
		return VehicleInfo{}, fmt.Errorf("[GetVehiclesInfo] - error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Preparar la consulta SQL
	query := `
		SELECT 
			COALESCE(vehicles.plate, '')      as Plate, 
			COALESCE(avl_devices.imei, '')    as Imei,
			COALESCE(companies.id, 0)         as Id_company,
			COALESCE(companies.name, '')      as Company,
			COALESCE(customers.id, 0)         as Id_customer,
			COALESCE(customers.name, '')      as Customer
		FROM avl_devices
		FULL JOIN vehicles         on vehicles.fk_avl_device   = avl_devices.id
		FULL JOIN companies        on companies.id             = vehicles.fk_company
		FULL JOIN customers        on customers.id             = vehicles.fk_customer
		WHERE avl_devices.imei = $1
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return VehicleInfo{}, fmt.Errorf("[GetVehiclesInfo] - error al preparar la consulta: %v", err)
	}
	defer stmt.Close()

	// Ejecutar la consulta y obtener el resultado
	var info VehicleInfo
	err = stmt.QueryRow(imei).Scan(&info.Plate, &info.Imei, &info.IdCompany, &info.Company, &info.IdCustomer, &info.Customer)
	if err != nil {
		return VehicleInfo{}, fmt.Errorf("[GetVehiclesInfo] - error al ejecutar la consulta: %v", err)
	}

	// Devolver la información del vehículo
	return info, nil
}
