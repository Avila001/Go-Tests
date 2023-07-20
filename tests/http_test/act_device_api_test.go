//go:build e2e_http
// +build e2e_http

package http_test

func TestCreateDevice(t *testing.T) {
	testTable := test_data.DataCreateDevice()
	for _, tc := range testTable {
		t.Run("test create device", func(t *testing.T) {
			// Arrange
			client := api.NewHTTPClient(5, 1*time.Second)

			// act
			createDeviceRequest := helpers.CreateDeviceHTTP(tc.Platform, tc.UserId)
			createDeviceResponse, resp, err := client.CreateDevice(context.Background(), createDeviceRequest)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.GreaterOrEqual(t, createDeviceResponse.DeviceID, 1)
		})
	}
}

func TestCreateDeviceDescribeNewDevice(t *testing.T) {
	testTable := test_data.DataCreateDevice()
	for _, tc := range testTable {
		t.Run("create device and describe device", func(t *testing.T) {
			client := api.NewHTTPClient(5, 1*time.Second)

			// создаем новый девайс
			idCreatedDevice := helpers.CreateDeviceHTTPWithAssertions(t, tc.Platform, tc.UserId)
			respDescribeDevice, resp, _ := client.DescribeDevice(context.Background(), strconv.Itoa(idCreatedDevice))

			// assert
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tc.Platform, respDescribeDevice.Value.Platform)
			assert.Equal(t, tc.UserId, respDescribeDevice.Value.UserID)
		})
	}
}

func TestCreateDeviceWithInvalidBody(t *testing.T) {
	testTable := test_data.DataCreateDeviceInvalidBody()
	for _, tc := range testTable {
		t.Run("test with invalid body", func(t *testing.T) {
			client := api.NewHTTPClient(5, 1*time.Second)

			createDeviceRequest := helpers.CreateDeviceHTTP(tc.Platform, tc.UserId)
			responseDeviceRequest, resp, err := client.CreateDevice(context.Background(), createDeviceRequest)

			assert.NotNil(t, responseDeviceRequest)
			assert.NoError(t, err)
			assert.Equal(t, 400, resp.StatusCode)
		})
	}
}

func TestCreateDeviceWithInvalidUserId(t *testing.T) {
	testTable := test_data.DataCreateDeviceInvalidUserId()
	for _, tc := range testTable {
		t.Run("test with invalid userId", func(t *testing.T) {
			client := api.NewHTTPClient(5, 1*time.Second)

			createDeviceRequest := helpers.CreateDeviceHTTP(tc.Platform, tc.UserId)
			_, resp, err := client.CreateDevice(context.Background(), createDeviceRequest)

			assert.NoError(t, err)
			assert.Equal(t, 400, resp.StatusCode)
		})
	}
}
