//go:build e2e_grpc
// +build e2e_grpc

package grpc_test_test

func TestListDevice(t *testing.T) {
	conn := helpers.OpenConnection()
	defer helpers.CloseConnection(t, conn)
	client := api.NewActDeviceApiServiceClient(conn)

	testTable := []struct {
		name    string
		page    uint64
		perPage uint64
	}{
		{
			name:    "test 1",
			page:    1,
			perPage: 10,
		},
		{
			name:    "test 2",
			page:    2,
			perPage: 3,
		},
	}

	for _, tc := range testTable {
		t.Run("grpc list of devices", func(t *testing.T) {
			// arrange
			opts := &api.ListDevicesV1Request{Page: tc.page, PerPage: tc.perPage}
			// act
			listDevicesRequest, err := client.ListDevicesV1(context.Background(), opts)
			listDevicesResponse := listDevicesRequest.Items

			//assert
			assert.NoError(t, err)
			assert.Equal(t, codes.OK, status.Code(err))
			assert.GreaterOrEqual(t, len(listDevicesResponse), 1)
		})
	}

}
