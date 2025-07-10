package delivery

import (
	"MilTrace/config"
	"MilTrace/domain"

	"github.com/gin-gonic/gin"
)

type deviceHandler struct {
	deviceService domain.DeviceService
}

func NewDeviceHandler(router *gin.Engine, deviceService domain.DeviceService) {
	deviceHandler := &deviceHandler{
		deviceService: deviceService,
	}

	router.GET("/ping", func(c *gin.Context) {
		config.GinStatusOK(c, "Ping Received", nil)
	})
	router.POST("/device", deviceHandler.RegisterNewDevice)
	router.PUT("/device/update", deviceHandler.UpdateDevice)
	router.DELETE("/device/delete/:deviceid", deviceHandler.DeleteDevice)
	router.GET("/device/all", deviceHandler.GetAllDeviceData)
	router.GET("/device/get/:deviceid", deviceHandler.GetDevice)
	router.POST("/device/receiveLocation", deviceHandler.ReceiveLocationData)
}

func (h *deviceHandler) RegisterNewDevice(c *gin.Context) {
	var device domain.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		config.GinBadRequest(c, "Invalid input", err)
		return
	}

	if err := h.deviceService.RegisterNewDevice(c, &device); err != nil {
		config.GinInternalServerError(c, "Failed to register device", err)
		return
	}

	config.GinCreated(c, "Device registered successfully")
}

func (h *deviceHandler) UpdateDevice(c *gin.Context) {
	var device domain.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		config.GinBadRequest(c, "Invalid input", err)
		return
	}

	if err := h.deviceService.UpdateDevice(c, &device); err != nil {
		config.GinInternalServerError(c, "Failed to update device", err)
		return
	}

	config.GinStatusOK(c, "Device updated successfully", nil)
}

func (h *deviceHandler) DeleteDevice(c *gin.Context) {
	deviceID := c.Param("deviceid")
	if deviceID == "" {
		config.GinBadRequest(c, "Device ID is required", nil)
		return
	}

	if err := h.deviceService.DeleteDevice(c, deviceID); err != nil {
		config.GinInternalServerError(c, "Failed to delete device", err)
		return
	}

	config.GinStatusOK(c, "Device deleted successfully", nil)
}

func (h *deviceHandler) GetAllDeviceData(c *gin.Context) {
	data, err := h.deviceService.GetAllDeviceData(c)
	if err != nil {
		config.GinInternalServerError(c, "Failed to get all devices", err)
		return
	}

	config.GinStatusOK(c, "All devices retrieved successfully", data)
}

func (h *deviceHandler) GetDevice(c *gin.Context) {
	deviceID := c.Param("deviceid")
	if deviceID == "" {
		config.GinBadRequest(c, "Device ID is required", nil)
		return
	}

	data, err := h.deviceService.GetDevice(c, deviceID)
	if err != nil {
		config.GinInternalServerError(c, "Failed to get device", err)
		return
	}

	config.GinStatusOK(c, "Device retrieved successfully", data)
}

func (h *deviceHandler) ReceiveLocationData(c *gin.Context) {
	var payload domain.Device
	if err := c.ShouldBindJSON(&payload); err != nil {
		config.GinBadRequest(c, "Invalid input", err)
		return
	}

	if err := h.deviceService.ReceiveLocationData(c, &payload); err != nil {
		config.GinInternalServerError(c, "Failed to receive location data", err)
		return
	}

	config.GinStatusOK(c, "Location data received successfully", nil)
}
