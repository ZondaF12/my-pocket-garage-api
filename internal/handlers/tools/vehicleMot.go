package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ZondaF12/my-pocket-garage/config"
)

type MotData []struct {
	Registration      string     `json:"registration"`
	Make              string     `json:"make"`
	Model             string     `json:"model"`
	FirstUsedDate     string     `json:"firstUsedDate"`
	FuelType          string     `json:"fuelType"`
	PrimaryColour     string     `json:"primaryColour"`
	MotTestExpiryDate string     `json:"MotTestExpiryDate"`
	MotTests          []MotTests `json:"motTests"`
}

type MotTests struct {
	CompletedDate  string `json:"completedDate"`
	TestResult     string `json:"testResult"`
	ExpiryDate     string `json:"expiryDate"`
	OdometerValue  string `json:"odometerValue"`
	OdometerUnit   string `json:"odometerUnit"`
	MotTestNumber  string `json:"motTestNumber"`
	RfrAndComments []any  `json:"rfrAndComments"`
}

func DoVehicleMotRequest(registration string) (MotData, error) {
	env, err := config.LoadConfig()
	if err != nil {
		return MotData{}, err
	}

	requestURL := "https://beta.check-mot.service.gov.uk/trade/vehicles/mot-tests/?registration=" + registration
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return MotData{}, fmt.Errorf("could not create request %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", env.DVSA_API)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return MotData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return MotData{}, errors.New("invalid registration number")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MotData{}, err
	}

	var vehicleResponse MotData
	if err = json.Unmarshal(body, &vehicleResponse); err != nil {
		fmt.Printf("Error: %v", err)
	}

	return vehicleResponse, nil
}
