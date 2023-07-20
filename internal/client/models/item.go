package models

import "time"

// Item is a struct
type Item struct {
	ID        string     `json:"id"`
	Platform  string     `json:"platform"`
	UserID    string     `json:"userId"`
	EnteredAt *time.Time `json:"enteredAt"`
}

// ItemRequest is a struct
type ItemRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

// ListItemsResponse is a struct
type ListItemsResponse struct {
	Items []struct {
		Item
	} `json:"items"`
}

// Act notification api, Create Notification

type ListNotificationsRequest struct {
	Items []struct {
		NotificationRequestItem
	} `json:"notification"`
}

type NotificationRequestItem struct {
	NotificationID     int    `json:"notificationId"`
	DeviceID           int    `json:"deviceId"`
	Username           string `json:"username"`
	Message            string `json:"message"`
	Lang               string `json:"lang"`
	NotificationStatus string `json:"notificationStatus"`
}
