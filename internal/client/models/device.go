package models

// CreateDeviceResponse is the response body from the CreateDevice endpoint
type CreateDeviceResponse struct {
	DeviceID int `json:"deviceId,string"`
}

// DescribeDeviceResponse is the response of DescribeDeviceRequest
type DescribeDeviceResponse struct {
	Value Item `json:"value"`
}

// CreateDeviceRequest is the request struct for api CreateDevice
type CreateDeviceRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

// RemovedDevice Response is the response of DeletedDeviceRequest
type RemovedDevice struct {
	Found bool `json:"found"`
}

type UpdateDeviceRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

type UpdateDeviceResponse struct {
	Success bool `json:"success"`
}

type CreateNotificationRequest struct {
	Value NotificationRequestItem `json:"notification"`
}

type CreateNotificationResponse struct {
	NotificationId int `json:"notificationId,string"`
}

type GetNotificationResponse struct {
	Notification []Notification `json:"notification"`
}
type Notification struct {
	NotificationID string `json:"notificationId"`
	Message        string `json:"message"`
}

type SubscribeNotification struct {
	Result struct {
		NotificationID string `json:"notificationId"`
		Message        string `json:"message"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []struct {
			TypeURL string `json:"typeUrl"`
			Value   string `json:"value"`
		} `json:"details"`
	} `json:"error"`
}

type AckNotificationResponse struct {
	Success bool `json:"success"`
}
