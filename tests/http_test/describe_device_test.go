//go:build e2e_http
// +build e2e_http

package http_test

func TestDescribeDevice(t *testing.T) {
	testTable := test_data.DataCreateDevice()
	for _, tc := range testTable {
		t.Run("create device and describe device", func(t *testing.T) {
			// arrange
			client := api.NewHTTPClient(5, 1*time.Second)
			ctx := context.Background()

			// создание девайса
			idCreatedDevice := helpers.CreateDeviceHTTPWithAssertions(t, tc.Platform, tc.UserId)

			// act
			respDescribeDevice, resp, err := client.DescribeDevice(ctx, strconv.Itoa(idCreatedDevice))
			respDescribeDeviceId, _ := strconv.Atoi(respDescribeDevice.Value.ID)

			// assert
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			assert.GreaterOrEqual(t, respDescribeDeviceId, 1)
			assert.Equal(t, tc.UserId, respDescribeDevice.Value.UserID)
			assert.Equal(t, tc.Platform, respDescribeDevice.Value.Platform)
		})
	}
}
