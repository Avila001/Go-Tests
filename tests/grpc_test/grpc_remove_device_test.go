//go:build e2e_grpc
// +build e2e_grpc

package grpc_test_test

func TestRemoveDevice(t *testing.T) {
	testTable := test_data.DataCreateDeviceGRPC()
	conn := helpers.OpenConnection()
	defer helpers.CloseConnection(t, conn)
	ctx := context.Background()
	client := api.NewActDeviceApiServiceClient(conn)
	for _, tc := range testTable {
		t.Run("grpc remove device", func(t *testing.T) {
			// создаем новый девайс
			idCreatedDevice := helpers.CreateDeviceWithAssertionsGRPC(t, tc.Platform, tc.UserId, conn)

			// удаляем девайс
			err := helpers.RemoveDeviceWithAssertionsGRPC(t, idCreatedDevice, conn)
			require.NoError(t, err)

			// делаем запрос на получение инфы об удаленном устройстве
			_, err = client.DescribeDeviceV1(ctx, &api.DescribeDeviceV1Request{DeviceId: idCreatedDevice})

			// assert
			assert.Equal(t, codes.NotFound, status.Code(err))
		})
	}
}
