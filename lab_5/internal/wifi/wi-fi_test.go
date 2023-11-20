package wifi_test

import (
	"errors"
	myWifi "example_mock/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestGetAddresses(t *testing.T) {
	mockWiFi := NewWiFi(t)
	wifiService := myWifi.WiFiService{WiFi: mockWiFi}

	testTable := []struct {
		hardwareAddrs []net.HardwareAddr
		errExpected   error
	}{
		{hardwareAddrs: []net.HardwareAddr{}, errExpected: nil},
		{hardwareAddrs: nil, errExpected: nil},
		{hardwareAddrs: nil, errExpected: errors.New("some error")},
	}

	for i, row := range testTable {
		mockIfaces := []*wifi.Interface{}
		for _, addr := range row.hardwareAddrs {
			mockIfaces = append(mockIfaces, &wifi.Interface{HardwareAddr: addr})
		}

		mockWiFi.On("Interfaces").Return(mockIfaces, row.errExpected)
		actualAddrs, err := wifiService.GetAddresses()

		if row.errExpected != nil {
			require.Error(t, err, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, actualAddrs)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)

		if len(row.hardwareAddrs) == 0 {
			require.Nil(t, actualAddrs, "row: %d, actual addrs must be nil for empty result", i)
		} else {
			require.Equal(t, row.hardwareAddrs, actualAddrs, "row: %d, expected addrs: %v, actual addrs: %v", i, row.hardwareAddrs, actualAddrs)
		}

		mockWiFi.On("Interfaces").Unset()
	}

	mockWiFi.AssertExpectations(t)
}

func TestGetNames(t *testing.T) {
	mockWifi := NewWiFi(t)
	wifiService := myWifi.WiFiService{WiFi: mockWifi}

	testTable := []struct {
		names       []string
		errExpected error
	}{
		{names: []string{"wlan0", "wlan1"}, errExpected: nil},
		{names: nil, errExpected: nil},
		{names: nil, errExpected: errors.New("some error")},
	}

	for i, row := range testTable {
		mockIfaces := []*wifi.Interface{}
		for _, name := range row.names {
			mockIfaces = append(mockIfaces, &wifi.Interface{Name: name})
		}

		mockWifi.On("Interfaces").Return(mockIfaces, row.errExpected)
		actualNames, err := wifiService.GetNames()

		if row.errExpected != nil {
			require.Error(t, err, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, actualNames, "row: %d, expected names: %v, actual names: %v", i, row.names, actualNames)

		mockWifi.On("Interfaces").Unset()
	}

	mockWifi.AssertExpectations(t)
}
