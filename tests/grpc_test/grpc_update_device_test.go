//go:build e2e_grpc
// +build e2e_grpc

package grpc_test_test

func TestUpdateDevice(t *testing.T) {
	testTable := test_data.DataCreateDeviceGRPC()
	conn := helpers.OpenConnection()
	defer helpers.CloseConnection(t, conn)

	for _, tc := range testTable {
		t.Run("update device", func(t *testing.T) {
			// arrange
			//client := api.NewActDeviceApiServiceClient(conn)
			var userId uint64 = 999
			var newPlatform = "new platform"

			// создаем устройство
			idCreatedDevice := helpers.CreateDeviceWithAssertionsGRPC(t, tc.Platform, tc.UserId, conn)

			// обновляем устройство
			_, err := helpers.UpdateDeviceWithAssertionsGRPC(t, idCreatedDevice, newPlatform, userId, conn)
			require.NoError(t, err)

			// делаем запрос на получение инфы об устройстве
			describedDeviceResponse, err := helpers.DescribeDeviceWithAssertionsGRPC(t, idCreatedDevice, conn)

			deviceIdResponse := describedDeviceResponse.Value.Id
			platformResponse := describedDeviceResponse.Value.Platform
			userIdResponse := describedDeviceResponse.Value.UserId

			// assert
			assert.NoError(t, err)
			assert.Equal(t, idCreatedDevice, deviceIdResponse)
			assert.Equal(t, newPlatform, platformResponse)
			assert.Equal(t, userId, userIdResponse)
		})

	}
}
