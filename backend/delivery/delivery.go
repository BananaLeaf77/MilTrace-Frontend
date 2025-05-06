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
}

func (h *deviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var deviceDataPayload domain.Device
	err := json.NewDecoder(r.Body).Decode(&deviceDataPayload)
	if err != nil {
		config.NetHTTPBadRequest(w, "Invalid request payload", err.Error())
		return
	}

	// Update device
	err = h.deviceService.UpdateDevice(&deviceDataPayload, r.Context())
	if err != nil {
		config.NetHTTPInternalServerError(w, "Failed to update device", err.Error())
		return
	}

	// Respond with success
	config.NetHTTPStatusOK(w, "Device updated successfully", nil)
}

func (h *deviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var deviceUUID domain.Device
	err := json.NewDecoder(r.Body).Decode(&deviceUUID)
	if err != nil {
		config.NetHTTPBadRequest(w, "Invalid request payload", err.Error())
		return
	}

	// Delete device
	err = h.deviceService.DeleteDevice(&deviceUUID.ID, r.Context())
	if err != nil {
		config.NetHTTPInternalServerError(w, "Failed to delete device", err.Error())
		return
	}

	// Respond with success
	config.NetHTTPStatusNoContent(w, "Device deleted successfully")
}

func (h *deviceHandler) GetAllDeviceData(w http.ResponseWriter, r *http.Request) {
	// Get all device data
	devices, err := h.deviceService.GetAllDeviceData(r.Context())
	if err != nil {
		config.NetHTTPInternalServerError(w, "Failed to retrieve devices", err.Error())
		return
	}

	// Respond with success
	config.NetHTTPStatusOK(w, "Devices retrieved successfully", &devices)
}

func (h *deviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var deviceUUID domain.Device
	err := json.NewDecoder(r.Body).Decode(&deviceUUID)
	if err != nil {
		config.NetHTTPBadRequest(w, "Invalid request payload", err.Error())
		return
	}

	// Get device data
	device, err := h.deviceService.GetDevice(&deviceUUID.ID)
	if err != nil {
		config.NetHTTPInternalServerError(w, "Failed to retrieve device", err.Error())
		return
	}

	// Respond with success
	config.NetHTTPStatusOK(w, "Device retrieved successfully", &device)
}
