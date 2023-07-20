package helpers

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/assert"
	api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	api_http "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/client"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/client/models"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/config_http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func CreateDeviceHTTP(platform string, userId string) models.CreateDeviceRequest {
	device := models.CreateDeviceRequest{
		Platform: platform,
		UserID:   userId,
	}
	return device
}

func CreateDeviceHTTPWithAssertions(t *testing.T, platform string, userId string) int {
	client := api_http.NewHTTPClient(5, 1*time.Second)
	device := models.CreateDeviceRequest{
		Platform: platform,
		UserID:   userId,
	}
	createDeviceResponse, _, err := client.CreateDevice(context.Background(), device)
	idCreateDeviceResponse := createDeviceResponse.DeviceID
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, idCreateDeviceResponse, 1)
	return idCreateDeviceResponse
}

func CreateDeviceGRPC(platform string, userId uint64) *api.CreateDeviceV1Request {
	device := &api.CreateDeviceV1Request{
		Platform: platform,
		UserId:   userId,
	}
	return device
}

//func CreateDeviceWithAssertions(t *testing.T, platform string, userId uint64, conn *grpc.ClientConn) *api.CreateDeviceV1Response {
//	client := api.NewActDeviceApiServiceClient(conn)
//	device := api.CreateDeviceV1Request{Platform: platform, UserId: userId}
//	createDeviceResponse, err := client.CreateDeviceV1(context.Background(), &device)
//
//	assert.NoError(t, err)
//	assert.Equal(t, codes.OK, status.Code(err))
//	assert.GreaterOrEqual(t, createDeviceResponse.DeviceId, amountDevices)
//	return createDeviceResponse
//}

func CreateDeviceWithAssertionsGRPC(t *testing.T, platform string, userId uint64, conn *grpc.ClientConn) uint64 {
	clientDevice := api.NewActDeviceApiServiceClient(conn)

	newDevice := api.CreateDeviceV1Request{Platform: platform, UserId: userId}
	createDeviceResponse, err := clientDevice.CreateDeviceV1(context.Background(), &newDevice)
	idCreatedDevice := createDeviceResponse.DeviceId

	assert.NoError(t, err)
	assert.Equal(t, codes.OK, status.Code(err))
	return idCreatedDevice
}

func CreateDeviceWithAssertionsGRPCAllure(t provider.T, platform string, userId uint64, conn *grpc.ClientConn) uint64 {
	clientDevice := api.NewActDeviceApiServiceClient(conn)

	newDevice := api.CreateDeviceV1Request{Platform: platform, UserId: userId}
	createDeviceResponse, err := clientDevice.CreateDeviceV1(context.Background(), &newDevice)
	idCreatedDevice := createDeviceResponse.DeviceId

	assert.NoError(t, err)
	assert.Equal(t, codes.OK, status.Code(err))
	return idCreatedDevice
}

func CreateDeviceGRPCAllure(t provider.T, values ValuesForDevice, conn *grpc.ClientConn) (api.ActDeviceApiServiceClient, *api.CreateDeviceV1Request) {
	deviceApiClient := api.NewActDeviceApiServiceClient(conn)
	request := &api.CreateDeviceV1Request{
		Platform: values.Platform,
		UserId:   values.UserId,
	}
	t.Logf("Platform: %v; userId: %v", values.Platform, values.UserId)
	return deviceApiClient, request
}

//func CreateDeviceWithAssertionsGRPCProvider(sCtx, platform string, userId uint64, conn *grpc.ClientConn) uint64 {
//	clientDevice := api.NewActDeviceApiServiceClient(conn)
//
//	newDevice := api.CreateDeviceV1Request{Platform: platform, UserId: userId}
//	createDeviceResponse, err := clientDevice.CreateDeviceV1(context.Background(), &newDevice)
//	idCreatedDevice := createDeviceResponse.DeviceId
//
//	assert.NoError(t, err)
//	assert.Equal(t, codes.OK, status.Code(err))
//	return idCreatedDevice
//}

func UpdateDeviceWithAssertionsGRPC(t *testing.T, idCreatedDevice uint64, platform string, userId uint64, conn *grpc.ClientConn) (bool, error) {
	client := api.NewActDeviceApiServiceClient(conn)

	updateDeviceRequest := api.UpdateDeviceV1Request{DeviceId: idCreatedDevice, Platform: platform, UserId: userId}
	updateDeviceResponse, err := client.UpdateDeviceV1(context.Background(), &updateDeviceRequest)
	successResponse := updateDeviceResponse.Success

	assert.NoError(t, err)
	assert.Equal(t, codes.OK, status.Code(err))
	return successResponse, err
}

func UpdateDeviceWithAssertionsGRPCAllure(t provider.T, idCreatedDevice uint64, platform string, userId uint64, conn *grpc.ClientConn) (bool, error) {
	client := api.NewActDeviceApiServiceClient(conn)

	updateDeviceRequest := api.UpdateDeviceV1Request{DeviceId: idCreatedDevice, Platform: platform, UserId: userId}
	updateDeviceResponse, err := client.UpdateDeviceV1(context.Background(), &updateDeviceRequest)
	successResponse := updateDeviceResponse.Success

	assert.NoError(t, err)
	assert.Equal(t, codes.OK, status.Code(err))
	return successResponse, err
}

func RemoveDeviceWithAssertionsGRPC(t *testing.T, deviceId uint64, conn *grpc.ClientConn) error {
	client := api.NewActDeviceApiServiceClient(conn)

	removeDeviceRequest := &api.RemoveDeviceV1Request{DeviceId: deviceId}
	removeDeviceResponse, err := client.RemoveDeviceV1(context.Background(), removeDeviceRequest)
	successRemoveDeviceResponse := removeDeviceResponse.Found

	assert.NoError(t, err)
	assert.Equal(t, true, successRemoveDeviceResponse)
	//assert.Equal(t, codes.NotFound, status.Code(err))
	return err
}

func DescribeDeviceWithAssertionsGRPC(t *testing.T, deviceId uint64, conn *grpc.ClientConn) (*api.DescribeDeviceV1Response, error) {
	client := api.NewActDeviceApiServiceClient(conn)
	var amountDevices uint64 = 1
	describeDeviceRequest := api.DescribeDeviceV1Request{DeviceId: deviceId}
	describeDeviceResponse, err := client.DescribeDeviceV1(context.Background(), &describeDeviceRequest)
	valueDescribeDeviceResponse := describeDeviceResponse.Value.Id

	assert.GreaterOrEqual(t, valueDescribeDeviceResponse, amountDevices)
	assert.NoError(t, err)
	assert.Equal(t, codes.OK, status.Code(err))

	return describeDeviceResponse, err
}

func DescribeDeviceWithAssertionsGRPCAllure(t provider.T, deviceId uint64, conn *grpc.ClientConn) (*api.DescribeDeviceV1Response, error) {
	client := api.NewActDeviceApiServiceClient(conn)
	var amountDevices uint64 = 1
	describeDeviceRequest := api.DescribeDeviceV1Request{DeviceId: deviceId}
	describeDeviceResponse, err := client.DescribeDeviceV1(context.Background(), &describeDeviceRequest)
	valueDescribeDeviceResponse := describeDeviceResponse.Value.Id

	assert.GreaterOrEqual(t, valueDescribeDeviceResponse, amountDevices)
	assert.NoError(t, err)
	assert.Equal(t, codes.OK, status.Code(err))

	return describeDeviceResponse, err
}

//func tt(t *testing.T, deviceId uint64, conn *grpc.ClientConn) (*api.DescribeDeviceV1Response, error) {
//	client := api.NewActDeviceApiServiceClient(conn)
//	var amountDevices uint64 = 1
//	describeDeviceRequest := api.DescribeDeviceV1Request{DeviceId: deviceId}
//	describeDeviceResponse, err := client.DescribeDeviceV1(context.Background(), &describeDeviceRequest)
//	valueDescribeDeviceResponse := describeDeviceResponse.Value
//
//	assert.GreaterOrEqual(t, valueDescribeDeviceResponse, amountDevices)
//	assert.NoError(t, err)
//	assert.Equal(t, codes.OK, status.Code(err))
//
//	return valueDescribeDeviceResponse, err
//}

func SendNotificationWithAssertionsGRPC(t *testing.T, deviceId uint64, message string, lang api.Language, conn *grpc.ClientConn) (uint64, error) {
	// arrange
	clientNot := api.NewActNotificationApiServiceClient(conn)

	notification := api.Notification{DeviceId: deviceId, Message: message, Lang: lang}
	sendNotificationV1Request := api.SendNotificationV1Request{Notification: &notification}
	// act
	sendNotificationV1Response, err := clientNot.SendNotificationV1(context.Background(), &sendNotificationV1Request)
	idSendNotificationV1Response := sendNotificationV1Response.NotificationId
	// assert
	assert.NoError(t, err)
	assert.Equal(t, codes.OK, status.Code(err))
	return idSendNotificationV1Response, err
}

//func GetNotificationWithAssertionsGRPC(t *testing.T, deviceId uint64, conn *grpc.ClientConn) (*api.GetNotificationV1Response, error) {
//	clientNot := api.NewActNotificationApiServiceClient(conn)
//
//	notification := api.Notification{DeviceId: deviceId, Message: message, Lang: lang}
//	getNotificationV1Request := api.GetNotificationV1Request{Notification: &notification}
//	sendNotificationV1Response, err := clientNot.GetNotification(context.Background(), &getNotificationV1Request)
//
//	assert.NoError(t, err)
//	assert.Equal(t, codes.OK, status.Code(err))
//	return sendNotificationV1Response, err
//}

func GetNotificationWithAssertionsGRPC(t *testing.T, deviceId uint64, conn *grpc.ClientConn) (*api.GetNotificationV1Response, error) {
	clientNot := api.NewActNotificationApiServiceClient(conn)
	getNotificationV1Request := api.GetNotificationV1Request{DeviceId: deviceId}
	getNotificationV1Response, err := clientNot.GetNotification(context.Background(), &getNotificationV1Request)
	assert.Equal(t, codes.OK, status.Code(err))
	return getNotificationV1Response, err
}

func AckNotificationWithAssertionsGRPC(t *testing.T, notificationId uint64, conn *grpc.ClientConn) (*api.AckNotificationV1Response, error) {
	client := api.NewActNotificationApiServiceClient(conn)
	ackNotificationRequest := api.AckNotificationV1Request{NotificationId: notificationId}
	ackNotificationResponse, err := client.AckNotification(context.Background(), &ackNotificationRequest)
	assert.Equal(t, codes.OK, status.Code(err))
	return ackNotificationResponse, err
}

func OpenConnection() *grpc.ClientConn {
	host := config_http.GetBasePathGrpc()
	conn, _ := grpc.Dial(host, grpc.WithInsecure())
	return conn
}

func CloseConnection(t *testing.T, conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		t.Log("connection error")
	}
}

func CloseConnectionProvider(t provider.T, conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		t.Log("connection error")
	}
}

func HelpersDescribeDevice(platform string, userId string) models.DescribeDeviceResponse {
	device := models.DescribeDeviceResponse{
		Value: models.Item{
			ID:       "1",
			Platform: "1",
			UserID:   "1",
		},
	}
	return device
}

func HelpersUpdateDevice(platform string, userId string) models.UpdateDeviceRequest {
	device := models.UpdateDeviceRequest{
		Platform: platform,
		UserID:   userId,
	}
	return device
}

func HelpersUpdateDeviceGRPC(deviceId uint64, platform string, userId uint64) *api.UpdateDeviceV1Request {
	device := &api.UpdateDeviceV1Request{
		DeviceId: deviceId,
		Platform: platform,
		UserId:   userId,
	}
	return device
}

func HelpersDescribeDeviceGRPC(deviceId uint64) *api.DescribeDeviceV1Request {
	device := &api.DescribeDeviceV1Request{
		DeviceId: deviceId,
	}
	return device
}

func CompareTimeInterval(h int, lang api.Language) (message string) {
	message = "ошибка"
	if h >= 6 && h < 11 {
		messages := [4]string{"Good morning", "Доброе утро", "Buenos dias", "Buon giorno"}
		return messages[lang]
	} else if h >= 11 && h < 15 {
		messages := [4]string{"Good afternoon", "Добрый день", "Buenas tardes", "Buon pomeriggio"}
		return messages[lang]
	} else if h >= 15 && h < 21 {
		messages := [4]string{"Good evening", "Добрый вечер", "Buenas noches", "Buona serata"}
		return messages[lang]
	} else if (h >= 21 && h < 24) || (h >= 0 && h < 6) {
		messages := [4]string{"Good night", "Доброй ночи", "Buenas noches", "Buona notte"}
		return messages[lang]
	}
	return message
}
