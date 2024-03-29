package dcollect

import (
	"encoding/json"
	"fmt"
	"graduatework/internal/config"
	"graduatework/internal/model"
	"io"
	"net/http"
)

func (m *MicroServiceStr) ReadIncidentData() (outputData []model.IncidentData, respStatusCode int) {

	outputData = make([]model.IncidentData, 0)

	response, err := http.Get(config.GlobalConfig.IncidentAddr)
	if err != nil {
		fmt.Print("can't GET incident-data: ", err)
	}

	if response.StatusCode != 200 {
		return outputData, response.StatusCode
	}

	responseInBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print("can't read incident-data: ", err)
	}

	err = json.Unmarshal(responseInBytes, &outputData)
	if err != nil {
		return outputData, response.StatusCode
	}

	reserv := make([]model.IncidentData, 0)
	for i := range outputData {
		reserv = outputData

		reserv = append(reserv[:i], reserv[i+1:]...)
	}
	outputData = reserv

	return outputData, response.StatusCode
}
