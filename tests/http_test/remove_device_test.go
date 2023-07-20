//go:build e2e_http
// +build e2e_http

package http_test

func TestRemoveDevice(t *testing.T) {
	testTable := test_data.DataCreateDevice()
	for _, tc := range testTable {
		t.Run("create, delete and describe this device", func(t *testing.T) {
			client := apiClient.NewHTTPClient(5, 1*time.Second)
			// создаем девайс
			idCreatedDevice := helpers.CreateDeviceHTTPWithAssertions(t, tc.Platform, tc.UserId)
			// удаляем девайс
			removeDeviceResponse, _, err := client.RemoveDevice(context.Background(), strconv.Itoa(idCreatedDevice))
			assert.NoError(t, err)
			assert.Equal(t, true, removeDeviceResponse.Found)
		})
	}
}
