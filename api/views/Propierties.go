package views

type Properties struct {
	AcceleratorPedalPosition int `json:"Accelerator Pedal Position"`
	ActiveGSMOperator        int `json:"Active GSM Operator"`
	AnalogInput1             int `json:"Analog Input 1"`
	BTStatus                 int `json:"BT Status"`
	BatteryLevel             int `json:"Battery Level"`
	BatteryVoltage           int `json:"Battery Voltage"`
	DigitalOutput1           int `json:"Digital Output 1"`
	EngineRPM                int `json:"Engine RPM"`
	EngineTemperature        int `json:"Engine Temperature"`
	ExternalVoltage          int `json:"External Voltage"`
	FuelConsumedCounted      int `json:"Fuel Consumed (counted)"`
	FuelLevel                int `json:"Fuel Level"`
	GNSSHDOP                 int `json:"GNSS HDOP"`
	GNSSPDOP                 int `json:"GNSS PDOP"`
	GNSSStatus               int `json:"GNSS Status"`
	GSMSignal                int `json:"GSM Signal"`
	Ignition                 int `json:"Ignition"`
	Movement                 int `json:"Movement"`
	SleepMode                int `json:"Sleep Mode"`
	Speed                    int `json:"Speed"`
	TotalMileage             int `json:"Total Mileage"`
	TotalMileageCounted      int `json:"Total Mileage (counted)"`
	TotalOdometer            int `json:"Total Odometer"`
	VehicleSpeed             int `json:"Vehicle Speed"`
}
