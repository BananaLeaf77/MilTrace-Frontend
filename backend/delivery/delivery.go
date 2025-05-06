package delivery

import (
	"MilTrace/domain"
	"net/http"
)

type deviceHandler struct {
	deviceService domain.DeviceService
}

func NewUserAuthHandler(netHttp *http.ServeMux, deviceService domain.DeviceService) {
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
	// Handle device registration logic here

}
