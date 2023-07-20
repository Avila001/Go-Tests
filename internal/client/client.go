package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/client/models"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/config_http"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/internal/pkg/logger"
)

// Client is a client for the Act Device API
type Client interface {
	Do(req *http.Request) (*http.Response, error)
	CreateDevice(ctx context.Context, body models.CreateDeviceRequest) (models.CreateDeviceResponse, *http.Response, error)
	ListDevices(ctx context.Context, opts url.Values) (models.ListItemsResponse, *http.Response, error)
	DescribeDevice(ctx context.Context, deviceID string) (models.DescribeDeviceResponse, *http.Response, error)
	RemoveDevice(ctx context.Context, deviceID string) (models.RemovedDevice, *http.Response, error)
	UpdateDevice(ctx context.Context, deviceID string, body models.UpdateDeviceRequest) (models.AckNotificationResponse, *http.Response, error)
	CreateNotification(ctx context.Context, body models.CreateNotificationRequest) (models.CreateNotificationResponse, *http.Response, error)
	GetNotification(ctx context.Context, opts url.Values) (models.GetNotificationResponse, *http.Response, error)
	AckNotification(ctx context.Context, deviceID string) (models.AckNotificationResponse, *http.Response, error)
}

type client struct {
	client   *retryablehttp.Client
	BasePath string
}

// NewHTTPClient creates a new HTTP client.
func NewHTTPClient(retryMax int, timeout time.Duration) Client {
	c := &retryablehttp.Client{
		HTTPClient:   &http.Client{Timeout: timeout},
		RetryMax:     retryMax,
		RetryWaitMin: 1 * time.Second,
		RetryWaitMax: 10 * time.Second,
		CheckRetry:   retryablehttp.DefaultRetryPolicy,
		Backoff:      retryablehttp.DefaultBackoff,
		//RequestLogHook:  requestHook,
		//ResponseLogHook: responseHook,
	}
	basePath := config_http.GetBasePath()
	client := &client{client: c, BasePath: basePath}
	return client
}

func requestHook(_ retryablehttp.Logger, req *http.Request, retry int) {
	dump, err := httputil.DumpRequest(req, true) // better way
	if err != nil {
		logger.ErrorKV(req.Context(), "can't dump request")
	}

	logger.InfoKV(
		req.Context(),
		fmt.Sprintf("Retry request %d", retry),
		"request", string(dump),
		"url", req.URL.String(),
	)
}

func responseHook(_ retryablehttp.Logger, res *http.Response) {
	dump, err := httputil.DumpResponse(res, true) // better way
	if err != nil {
		logger.ErrorKV(res.Request.Context(), "can't dump response")
	}

	logger.InfoKV(
		res.Request.Context(),
		"Responded",
		"response", dump,
		"url", res.Request.URL.String(),
		"status_code", res.StatusCode,
	)
}

func (c *client) Do(request *http.Request) (*http.Response, error) {
	req, err := retryablehttp.FromRequest(request)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *client) CreateDevice(ctx context.Context, body models.CreateDeviceRequest) (models.CreateDeviceResponse, *http.Response, error) {
	var localResponse models.CreateDeviceResponse

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.BasePath+"/api/v1/devices", b)
	if err != nil {
		return localResponse, nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.ErrorKV(ctx, "Error on Body reading", err)
		}
	}(res.Body)

	device := new(models.CreateDeviceResponse)
	err = json.NewDecoder(res.Body).Decode(device)
	if err != nil {
		return localResponse, res, err
	}
	return *device, res, nil
}

func (c *client) ListDevices(ctx context.Context, opts url.Values) (models.ListItemsResponse, *http.Response, error) {
	var localResponse models.ListItemsResponse

	apiURL, err := url.Parse(c.BasePath + "/api/v1/devices")
	if err != nil {
		return localResponse, nil, err
	}

	query := apiURL.Query()
	for k, v := range opts {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	apiURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	devices := new(models.ListItemsResponse)
	err = json.NewDecoder(res.Body).Decode(devices)
	if err != nil {
		return localResponse, res, err
	}
	return *devices, res, nil
}

func (c *client) DescribeDevice(ctx context.Context, deviceID string) (models.DescribeDeviceResponse, *http.Response, error) {
	var localResponse models.DescribeDeviceResponse

	apiURLString := c.BasePath + "/api/v1/devices/{deviceId}"
	apiURLString = strings.Replace(apiURLString, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	apiURL, err := url.Parse(apiURLString)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	device := new(models.DescribeDeviceResponse)
	err = json.NewDecoder(res.Body).Decode(device)
	if err != nil {
		return localResponse, res, err
	}
	return *device, res, nil
}

func (c *client) RemoveDevice(ctx context.Context, deviceID string) (models.RemovedDevice, *http.Response, error) {
	var localResponse models.RemovedDevice

	apiURLString := c.BasePath + "/api/v1/devices/{deviceId}"
	apiURLString = strings.Replace(apiURLString, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	apiURL, err := url.Parse(apiURLString)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, apiURL.String(), nil)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	removedDevice := new(models.RemovedDevice)

	err = json.NewDecoder(res.Body).Decode(removedDevice)
	if err != nil {
		return localResponse, res, err
	}
	return *removedDevice, res, nil
}

func (c *client) UpdateDevice(ctx context.Context, deviceID string, body models.UpdateDeviceRequest) (models.AckNotificationResponse, *http.Response, error) {
	var localResponse models.AckNotificationResponse

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return localResponse, nil, err
	}

	apiURLString := c.BasePath + "/api/v1/devices/{deviceId}"
	apiURLString = strings.Replace(apiURLString, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	apiURL, err := url.Parse(apiURLString)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodPut, apiURL.String(), b)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.ErrorKV(ctx, "Error on Body reading", err)
		}
	}(res.Body)

	updatedDevice := new(models.AckNotificationResponse)

	err = json.NewDecoder(res.Body).Decode(updatedDevice)
	if err != nil {
		return localResponse, res, err
	}
	return *updatedDevice, res, nil

}

func (c *client) CreateNotification(ctx context.Context, body models.CreateNotificationRequest) (models.CreateNotificationResponse, *http.Response, error) {
	var localResponse models.CreateNotificationResponse

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.BasePath+"/api/v1/notification", b)
	if err != nil {
		return localResponse, nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.ErrorKV(ctx, "Error on Body reading", err)
		}
	}(res.Body)

	device := new(models.CreateNotificationResponse)
	err = json.NewDecoder(res.Body).Decode(device)
	if err != nil {
		return localResponse, res, err
	}
	return *device, res, nil
}

func (c *client) AckNotification(ctx context.Context, deviceID string) (models.AckNotificationResponse, *http.Response, error) {
	var localResponse models.AckNotificationResponse

	apiURLString := c.BasePath + "/api/v1/notification/{deviceId}"
	apiURLString = strings.Replace(apiURLString, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	apiURL, err := url.Parse(apiURLString)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, apiURL.String(), nil)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	ackNotification := new(models.AckNotificationResponse)

	err = json.NewDecoder(res.Body).Decode(ackNotification)
	if err != nil {
		return localResponse, res, err
	}
	return *ackNotification, res, nil
}

func (c *client) GetNotification(ctx context.Context, opts url.Values) (models.GetNotificationResponse, *http.Response, error) {
	var localResponse models.GetNotificationResponse

	apiURL, err := url.Parse(c.BasePath + "/api/v1/notification")
	if err != nil {
		return localResponse, nil, err
	}

	query := apiURL.Query()

	//for k, v := range opts {
	//	for _, iv := range v {
	//		query.Add(k, iv)
	//	}
	//}
	//

	query.Add("deviceId", "1")

	apiURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodDelete, apiURL.String(), nil)
	//req, err := http.NewRequest(http.MethodDelete, query.Get("deviceId"), nil)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	if res.StatusCode != http.StatusOK {
		logger.ErrorKV(ctx, "Bad status code", res.StatusCode)
	}

	receivedNotification := new(models.GetNotificationResponse)

	err = json.NewDecoder(res.Body).Decode(receivedNotification)
	if err != nil {
		return localResponse, res, err
	}
	return *receivedNotification, res, nil
}
