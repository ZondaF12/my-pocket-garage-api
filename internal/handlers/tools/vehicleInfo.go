package tools

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ZondaF12/my-pocket-garage/config"
)

type VehicleData struct {
	RegistrationNumber       string `json:"registrationNumber"`
	TaxStatus                string `json:"taxStatus"`
	TaxDueDate               string `json:"taxDueDate"`
	ArtEndDate               string `json:"artEndDate"`
	MotStatus                string `json:"motStatus"`
	Make                     string `json:"make"`
	YearOfManufacture        int    `json:"yearOfManufacture"`
	EngineCapacity           int    `json:"engineCapacity"`
	Co2Emissions             int    `json:"co2Emissions"`
	FuelType                 string `json:"fuelType"`
	MarkedForExport          bool   `json:"markedForExport"`
	Colour                   string `json:"colour"`
	TypeApproval             string `json:"typeApproval"`
	RevenueWeight            int    `json:"revenueWeight"`
	EuroStatus               string `json:"euroStatus"`
	DateOfLastV5CIssued      string `json:"dateOfLastV5CIssued"`
	MotExpiryDate            string `json:"motExpiryDate"`
	Wheelplan                string `json:"wheelplan"`
	MonthOfFirstRegistration string `json:"monthOfFirstRegistration"`
}

func DoVehicleInfoRequest(vehicleReg string) (VehicleData, error) {
	env, err := config.LoadConfig()
	if err != nil {
		return VehicleData{}, err
	}

	jsonBody := []byte(fmt.Sprintf(`{"registrationNumber": "%s"}`, vehicleReg))
	bodyReader := bytes.NewBuffer(jsonBody)

	requestURL := "https://driver-vehicle-licensing.api.gov.uk/vehicle-enquiry/v1/vehicles"
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return VehicleData{}, fmt.Errorf("could not create request %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", env.DVLA_API)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return VehicleData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return VehicleData{}, errors.New("invalid registration number")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return VehicleData{}, err
	}

	var vehicleResponse VehicleData
	if err = json.Unmarshal(body, &vehicleResponse); err != nil {
		fmt.Printf("Error: %v", err)
	}

	return vehicleResponse, nil
}