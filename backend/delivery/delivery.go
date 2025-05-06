package delivery

import (
	"MilTrace/config"
	"MilTrace/domain"
	"encoding/json"
	"net/http"
)

type deviceHandler struct {
	deviceService domain.DeviceService
}

func NewDeviceHandler(netHttp *http.ServeMux, deviceService domain.DeviceService) {
	deviceHandler := &deviceHandler{
		deviceService: deviceService,
	}

	netHttp.HandleFunc("/device", deviceHandler.RegisterNewDevice)
	netHttp.HandleFunc("/device/update", deviceHandler.UpdateDevice)
	netHttp.HandleFunc("/device/delete", deviceHandler.DeleteDevice)
	netHttp.HandleFunc("/device/all", deviceHandler.GetAllDeviceData)
	netHttp.HandleFunc("/device/get", deviceHandler.GetDevice)

}

func (h *deviceHandler) RegisterNewDevice(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var deviceDataPayload domain.Device
	err := json.NewDecoder(r.Body).Decode(&deviceDataPayload)
	if err != nil {
		config.NetHTTPBadRequest(w, "Invalid request payload", err.Error())
		return
	}

	// Register device
	err = h.deviceService.RegisterNewDevice(&deviceDataPayload, r.Context())
	if err != nil {
		config.NetHTTPInternalServerError(w, "Failed to register device", err.Error())
		return
	}

	// Respond with success
	config.NetHTTPStatusCreated(w, "Device registered successfully")
	return
}

func (h *deviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {

}

func (h *deviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {

}

func (h *deviceHandler) GetAllDeviceData(w http.ResponseWriter, r *http.Request) {

}

func (h *deviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {

}
