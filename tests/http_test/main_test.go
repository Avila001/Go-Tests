//go:build e2e_http
// +build e2e_http

package http_test

func TestMain(m *testing.M) {
	cfg, err := config_http.GetConfig()
	checkURL := url.URL{
		Scheme: "HTTP",
		Host:   cfg.Host,
		Path:   cfg.Livecheck,
	}

	// livecheck
	res, err := http.Get(checkURL.String())
	if err != nil {
		log.Fatalf("Service is not running: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Status is not OK: %v", res.StatusCode)
	}
	m.Run()
}
