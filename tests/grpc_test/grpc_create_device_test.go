//go:build e2e_grpc
// +build e2e_grpc

package grpc_test_test

func TestCreateDevice(t *testing.T) {
	conn := helpers.OpenConnection()
	defer helpers.CloseConnection(t, conn)
	testTable := test_data.DataCreateDeviceGRPC()
	for _, tc := range testTable {
		t.Run("grpc test create device", func(t *testing.T) {
			// arrange
			var amountDevices uint64 = 1
			client := api.NewActDeviceApiServiceClient(conn)
			createDeviceRequest := api.CreateDeviceV1Request{Platform: tc.Platform, UserId: tc.UserId}

			// act
			createDeviceResponse, err := client.CreateDeviceV1(context.Background(), &createDeviceRequest)

			// assert
			assert.NoError(t, err)
			assert.Equal(t, codes.OK, status.Code(err))
			assert.GreaterOrEqual(t, createDeviceResponse.DeviceId, amountDevices)
		})
	}
}
