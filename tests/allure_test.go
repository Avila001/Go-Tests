package tests

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	apiClient "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/client"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/helpers"
	sql "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/sql_test"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/test_data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"testing"
	"time"
)

type Device struct {
	platform string
	userId   uint64
}

func (e *Device) String() string {
	return fmt.Sprintf("%s#%d", e.platform, e.userId)
}

type CreateDeviceSuite struct {
	suite.Suite
	ParamMyTest []*Device
}

func (s *CreateDeviceSuite) BeforeAll(t provider.T) {
	var params []*allure.Parameter
	testTable := test_data.DataCreateDeviceGRPC()
	for _, tc := range testTable {
		param := &Device{
			platform: tc.Platform,
			userId:   tc.UserId,
		}
		//params = append(params, allure.NewParameter(fmt.Sprintf("Ex %d", i), param))
		s.ParamMyTest = append(s.ParamMyTest, param)
	}
	t.NewStep("BeforeAllStep", params...)
}

func (s *CreateDeviceSuite) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("BeforeAfter")
	t.NewStep("This Step will be before Each")
}

func (s *CreateDeviceSuite) AfterEach(t provider.T) {
	t.NewStep("AfterEach Step")
}

func (s *CreateDeviceSuite) AfterAll(t provider.T) {
	t.NewStep("AfterAll Step")
}

func (s *CreateDeviceSuite) TableTestMyTest(t provider.T, device *Device) {
	conn := helpers.OpenConnection()
	defer helpers.CloseConnectionProvider(t, conn)
	//testTable := test_data.DataCreateDeviceGRPC()
	t.Title("gRPC Create Device Test")
	t.Description(`
				This Test will have all labels from SetupTest function
				Unique labels:
			ID = "example5"
			Severity = "trivial"
			Unique tag = "Example5"`)

	// arrange
	t.Parallel()
	var (
		platform string
		userId   uint64
		ctx      context.Context
	)
	//var amountDevices uint64 = 1
	//client := api.NewActDeviceApiServiceClient(conn)

	// act
	t.WithTestSetup(func(t provider.T) {
		t.WithNewStep("init platform", func(sCtx provider.StepCtx) {
			platform = device.platform
			sCtx.WithNewParameters("platform", platform)
		})
		t.WithNewStep("init userId", func(sCtx provider.StepCtx) {
			userId = device.userId
			sCtx.WithNewParameters("userId", userId)
		})
		t.WithNewStep("init ctx", func(sCtx provider.StepCtx) {
			ctx = context.Background()
			sCtx.WithNewParameters("ctx", ctx)
		})
	})
	//createDeviceRequest := api.CreateDeviceV1Request{Platform: device.platform, UserId: device.userId}
	//createDeviceResponse, err := client.CreateDeviceV1(context.Background(), &createDeviceRequest)
	// assert
	t.Require().NotEqual(1, userId)
	//t.Require().NoError(err)
	//t.Require().Equal(codes.OK, status.Code(err))
	//t.Require().GreaterOrEqual(createDeviceResponse.DeviceId, amountDevices)

	//require.NoError(t, err)
	//require.Equal(t, codes.OK, status.Code(err))
	//require.GreaterOrEqual(t, createDeviceResponse.DeviceId, amountDevices)
}

func TestCreateDevice(t *testing.T) {
	conn := helpers.OpenConnection()
	defer helpers.CloseConnection(t, conn)
	testTable := test_data.DataCreateDeviceGRPC()
	for _, tc := range testTable {
		runner.Run(t, "name", func(t provider.T) {
			t.Title("gRPC Create Device Test")
			t.Description(`Создание девайс протокол gRPC`)
			client := api.NewActDeviceApiServiceClient(conn)
			// arrange
			//values := allure.NewParameters("platform", tc.Platform, "userId", tc.UserId)
			t.WithNewStep("arrange", func(sCtx provider.StepCtx) {
				var amountDevices uint64 = 1
				uintAmountDevices := strconv.FormatUint(amountDevices, 10)

				createDeviceRequest := &api.CreateDeviceV1Request{Platform: tc.Platform, UserId: tc.UserId}
				sCtx.WithNewAttachment("Test data", allure.Text, []byte(fmt.Sprintf("%+v", createDeviceRequest)))
				// act

				createDeviceResponse, err := client.CreateDeviceV1(context.Background(), createDeviceRequest)
				id := strconv.FormatUint(createDeviceResponse.DeviceId, 10)
				sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", createDeviceResponse)))

				// assert
				sCtx.Assert().NoError(err)
				sCtx.Assert().Equal(codes.OK.String(), status.Code(err).String())
				sCtx.Assert().GreaterOrEqual(id, uintAmountDevices, "Id девайса >= 1")
			})

		})
	}
}

func TestCreateDeviceAllure(t *testing.T) {
	conn := helpers.OpenConnection()
	defer helpers.CloseConnection(t, conn)
	testTable := test_data.DataCreateDeviceGRPC()
	for _, tc := range testTable {
		runner.Run(t, "name", func(t provider.T) {
			t.Title("gRPC Create Device Test")
			t.Description(`Создание девайс протокол gRPC`)
			//client := api.NewActDeviceApiServiceClient(conn)
			// arrange
			values := allure.NewParameters("platform", tc.Platform, "userId", tc.UserId)
			var amountDevices uint64 = 1
			uintAmountDevices := strconv.FormatUint(amountDevices, 10)
			t.WithNewStep("arrange", func(sCtx provider.StepCtx) {
				// act

				client, createDeviceRequest := helpers.CreateDeviceGRPCAllure(t, tc, conn)
				sCtx.WithNewAttachment("Test data", allure.Text, []byte(fmt.Sprintf("%+v", createDeviceRequest)))

				createDeviceResponse, err := client.CreateDeviceV1(context.Background(), createDeviceRequest)

				id := strconv.FormatUint(createDeviceResponse.DeviceId, 10)
				t.Log("Сreated a new device", id)
				sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", createDeviceResponse)))

				// assert
				sCtx.Assert().NoError(err)
				sCtx.Assert().Equal(codes.OK.String(), status.Code(err).String())
				sCtx.Assert().GreaterOrEqual(id, uintAmountDevices, "Id девайса >= 1")
			}, values...)

		})
	}
}

func TestUpdateDevice(t *testing.T) {
	testTable := test_data.DataCreateDeviceGRPC()
	conn := helpers.OpenConnection()
	defer helpers.CloseConnection(t, conn)
	//db := helpers.DBConnection(t)
	for _, tc := range testTable {

		runner.Run(t, "update device", func(t provider.T) {
			values := allure.NewParameters("platform", tc.Platform, "userId", tc.UserId)
			t.Title("gRPC Create Device Test")
			t.Description(`Создание девайс протокол gRPC`)
			// arrange
			//client := api.NewActDeviceApiServiceClient(conn)
			//var userId uint64 = 999
			//var newPlatform = "new platform"
			t.WithNewStep("Update Test", func(sCtx provider.StepCtx) {
				// создаем устройство
				idCreatedDevice := helpers.CreateDeviceWithAssertionsGRPCAllure(t, tc.Platform, tc.UserId, conn)

				// обновляем устройство
				_, err := helpers.UpdateDeviceWithAssertionsGRPCAllure(t, idCreatedDevice, tc.Platform, tc.UserId, conn)
				//require.NoError(t, err)

				// делаем запрос на получение инфы об устройстве
				describedDeviceResponse, err := helpers.DescribeDeviceWithAssertionsGRPCAllure(t, idCreatedDevice, conn)

				deviceIdResponse := describedDeviceResponse.Value.Id
				platformResponse := describedDeviceResponse.Value.Platform
				userIdResponse := describedDeviceResponse.Value.UserId

				// assert
				assert.NoError(t, err)
				assert.Equal(t, idCreatedDevice, deviceIdResponse)
				assert.Equal(t, tc.Platform, platformResponse)
				assert.Equal(t, tc.UserId, userIdResponse)
			}, values...)
		})

	}
}

func TestDB(t *testing.T) {

}

func TestRemoveDevice(t *testing.T) {
	testTable := test_data.DataCreateDevice()
	ctx := context.Background()
	db := sql.GetDBConnection(t)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			t.Fatal("Failed to close DB connection")
		}
	}(db.DB)
	for _, tc := range testTable {
		t.Run("create, delete and describe this device", func(t *testing.T) {
			client := apiClient.NewHTTPClient(5, 1*time.Second)
			// создаем девайс
			idCreatedDevice := helpers.CreateDeviceHTTPWithAssertions(t, tc.Platform, tc.UserId)
			// удаляем девайс
			removeDeviceResponse, _, err := client.RemoveDevice(ctx, strconv.Itoa(idCreatedDevice))
			bdRequest, err := db.RemoveDevice(ctx, uint64(idCreatedDevice))

			assert.Equal(t, bdRequest, true)
			assert.NoError(t, err)
			assert.Equal(t, true, removeDeviceResponse.Found)
		})
	}
}

//func TestAckNotificationGRPC(t *testing.T) {
//	conn := helpers.OpenConnection()
//	defer helpers.CloseConnection(t, conn)
//
//	runner.Run(t, "test ack notification", func(t provider.T) {
//		// создаем новый девайс для получения на него уведомлений
//		idCreatedDevice := helpers.CreateDeviceWithAssertionsGRPC(t, "Ios", 100, conn)
//
//		// создаем новое уведомление на девайс
//		_, err := helpers.SendNotificationWithAssertionsGRPC(t, idCreatedDevice, "mes", 0, conn)
//
//		// отправляем уведомление на устройство
//		getNotificationResponse, _ := helpers.GetNotificationWithAssertionsGRPC(t, idCreatedDevice, conn)
//		idGetNotificationResponse := getNotificationResponse.Notification[0].NotificationId
//
//		// подтверждаем получение уведомления на устройстве
//		ackNotificationResponse, err := helpers.AckNotificationWithAssertionsGRPC(t, idGetNotificationResponse, conn)
//		res := ackNotificationResponse.Success
//
//		//assert
//		assert.NoError(t, err)
//		assert.Equal(t, codes.OK, status.Code(err))
//		assert.Equal(t, true, res)
//	})
//}

func TestRunner(t *testing.T) {
	suite.RunSuite(t, new(CreateDeviceSuite))
}
