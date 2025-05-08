package delivery

import (
	"MilTrace/config"
	"MilTrace/domain"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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
	netHttp.HandleFunc("/device/sendLocation", deviceHandler.GetDevice)

}

func (h *deviceHandler) RegisterNewDevice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var deviceDataPayload domain.Device
		err := json.NewDecoder(r.Body).Decode(&deviceDataPayload)
		if err != nil {
			config.NetHTTPBadRequest(w, "Invalid request payload", err.Error())
			return
		}

		err = h.deviceService.RegisterNewDevice(&deviceDataPayload, r.Context())
		if err != nil {
			config.NetHTTPInternalServerError(w, "Failed to register device", err.Error())
			return
		}

		config.NetHTTPStatusCreated(w, "Device registered successfully")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *deviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	switch r.Method {
	case http.MethodPut:
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

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *deviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		// Decode request body
		var deviceUUID domain.Device
		err := json.NewDecoder(r.Body).Decode(&deviceUUID)
		if err != nil {
			config.NetHTTPBadRequest(w, "Invalid request payload", err.Error())
			return
		}

		// Delete device
		err = h.deviceService.DeleteDevice(&deviceUUID.DeviceID, r.Context())
		if err != nil {
			config.NetHTTPInternalServerError(w, "Failed to delete device", err.Error())
			return
		}

		// Respond with success
		config.NetHTTPStatusNoContent(w, "Device deleted successfully")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *deviceHandler) GetAllDeviceData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Get all device data
		devices, err := h.deviceService.GetAllDeviceData(r.Context())
		if err != nil {
			config.NetHTTPInternalServerError(w, "Failed to retrieve devices", err.Error())
			return
		}

		// Respond with success
		config.NetHTTPStatusOK(w, "Devices retrieved successfully", &devices)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *deviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Decode request body
		var deviceUUID domain.Device
		err := json.NewDecoder(r.Body).Decode(&deviceUUID)
		if err != nil {
			config.NetHTTPBadRequest(w, "Invalid request payload", err.Error())
			return
		}

		// Get device data
		device, err := h.deviceService.GetDevice(&deviceUUID.DeviceID, r.Context())
		if err != nil {
			config.NetHTTPInternalServerError(w, "Failed to retrieve device", err.Error())
			return
		}

		// Respond with success
		config.NetHTTPStatusOK(w, "Device retrieved successfully", &device)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *deviceHandler) ReceiveLocationData(w http.ResponseWriter, r *http.Request) {
	type payloadReceiver struct {
		UUID     uuid.UUID       `json:"device_uuid"`
		Location domain.Location `json:"location"`
	}

	switch r.Method {
	case http.MethodPost:
		var payload payloadReceiver
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			config.NetHTTPBadRequest(w, "Failed to parse location data", err.Error())
			return
		}

		err = h.deviceService.ReceiveLocationData(&payload.UUID, &payload.Location, r.Context())
		if err != nil {
			config.NetHTTPInternalServerError(w, "Failed to send device location", err.Error())
		}

		config.NetHTTPStatusOK(w, "Location data received successfully", nil)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
