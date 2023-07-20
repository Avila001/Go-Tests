//go:build e2e_http
// +build e2e_http

package http_test

func TestUpdateDevice(t *testing.T) {
	testTable := test_data.DataCreateDevice()
	for _, tc := range testTable {
		t.Run("create, update and describe device", func(t *testing.T) {
			client := apiClient.NewHTTPClient(5, 1*time.Second)
			ctx := context.Background()
			// создаем девайс
			idCreatedDevice := helpers.CreateDeviceHTTPWithAssertions(t, tc.Platform, tc.UserId)

			// обновляем девайс
			newUserID := strconv.Itoa(999)
			newPlatform := "updated platform"
			updatedDevice := helpers.HelpersUpdateDevice(newPlatform, newUserID)
			respUpdateDevice, _, err := client.UpdateDevice(ctx, strconv.Itoa(idCreatedDevice), updatedDevice)
			require.NoError(t, err)
			require.Equal(t, true, respUpdateDevice.Success)

			// получение ин-фии по девайсу!
			respDescribeDevice, _, _ := client.DescribeDevice(ctx, strconv.Itoa(idCreatedDevice))
			// assert
			assert.Equal(t, newUserID, respDescribeDevice.Value.UserID)
			assert.Equal(t, newPlatform, respDescribeDevice.Value.Platform)
		})
	}
}
