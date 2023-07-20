package test_data

type GetDevicesFields []struct {
	Platform string
	UserId   string
}

type GetDevicesFieldsGRPC []struct {
	Platform string
	UserId   uint64
}

func DataCreateDevice() GetDevicesFields {
	test := GetDevicesFields{
		{
			Platform: "ios",
			UserId:   "7",
		},
		{
			Platform: "android",
			UserId:   "14",
		},
	}
	return test
}

func DataCreateDeviceGRPC() GetDevicesFieldsGRPC {
	test := GetDevicesFieldsGRPC{
		{
			Platform: "ios",
			UserId:   14,
		},
		{
			Platform: "android",
			UserId:   8,
		},
	}
	return test
}

func DataCreateDeviceInvalidBody() GetDevicesFields {
	test := GetDevicesFields{
		{
			Platform: "",
			UserId:   "",
		},
	}
	return test
}

func DataCreateDeviceInvalidUserId() GetDevicesFields {
	test := GetDevicesFields{
		{
			Platform: "test",
			UserId:   "f",
		},
	}
	return test
}

func DataDescribeDevice() GetDevicesFields {
	test := GetDevicesFields{
		{
			Platform: "test",
			UserId:   "14",
		},
		{
			Platform: "test",
			UserId:   "14",
		},
	}
	return test
}

func DataUpdateDevice() GetDevicesFields {
	test := GetDevicesFields{
		{
			Platform: "new_updated platform",
			UserId:   "999",
		},
	}
	return test
}
