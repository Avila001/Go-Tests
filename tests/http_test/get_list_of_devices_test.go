//go:build e2e_http
// +build e2e_http

package http_test

func TestListOfDevices(t *testing.T) {
	t.Run("get devices list of devices", func(t *testing.T) {
		// Arrange
		client := apiClient.NewHTTPClient(5, 1*time.Second)

		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "100")

		// Act
		items, resp, err := client.ListDevices(context.Background(), opts)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.GreaterOrEqual(t, len(items.Items), 1)
	})
}
