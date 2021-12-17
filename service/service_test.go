package service

import (
	"fmt"
	"github.com/rocketgate/rocketgate-go-sdk/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGatewayService(t *testing.T) {
	service := NewGatewayService()
	fmt.Println(service)
	assert.Equal(t, "/gateway/servlet/ServiceDispatcherAccess", service._ROCKETGATE_SERVLET)
	assert.Equal(t, "gateway.rocketgate.com", service._ROCKETGATE_HOST)
	assert.Equal(t, "https", service._ROCKETGATE_PROTOCOL)
	assert.Equal(t, 443, service._ROCKETGATE_PORTNO)
	assert.Equal(t, 10, service._ROCKETGATE_CONNECT_TIMEOUT)
	assert.Equal(t, 90, service._ROCKETGATE_READ_TIMEOUT)
}

func TestGetServiceUrlDefault(t *testing.T) {
	gatewayService := NewGatewayService()
	gatewayRequest := request.NewGatewayRequest()
	serviceUrl := gatewayService.getServiceUrl("gateway.rocketgate.com", gatewayRequest)
	assert.Equal(t, "https://gateway.rocketgate.com:443/gateway/servlet/ServiceDispatcherAccess", serviceUrl)
}

func TestGetServiceUrlSetTestMode1(t *testing.T) {
	gatewayService := NewGatewayService()
	gatewayService.SetTestMode(true)
	gatewayRequest := request.NewGatewayRequest()
	serviceUrl := gatewayService.getServiceUrl(gatewayService._ROCKETGATE_HOST, gatewayRequest)
	assert.Equal(t, "https://dev-gateway.rocketgate.com:443/gateway/servlet/ServiceDispatcherAccess", serviceUrl)
}

func TestGetServiceUrlSetTestMode2(t *testing.T) {
	gatewayService := NewGatewayService()
	gatewayService.SetTestMode(true)
	gatewayService.SetTestMode(false)
	gatewayRequest := request.NewGatewayRequest()
	serviceUrl := gatewayService.getServiceUrl(gatewayService._ROCKETGATE_HOST, gatewayRequest)
	assert.Equal(t, "https://gateway.rocketgate.com:443/gateway/servlet/ServiceDispatcherAccess", serviceUrl)
}

func TestGetServiceSets(t *testing.T) {
	gatewayService := NewGatewayService()

	gatewayService.SetServlet("/servlet")
	gatewayService.SetTestMode(true)

	assert.Equal(t, gatewayService._ROCKETGATE_HOST, "dev-gateway.rocketgate.com")
	assert.Equal(t, gatewayService._ROCKETGATE_PROTOCOL, "https")
	assert.Equal(t, gatewayService._ROCKETGATE_PORTNO, 443)

	gatewayService.SetHost("test.rocketgate.com")
	gatewayService.SetPortNo(8443)
	gatewayService.SetProtocol("http")
	gatewayService.SetConnectTimeout(5555)
	gatewayService.SetReadTimeout(7777)

	assert.Equal(t, gatewayService._ROCKETGATE_SERVLET, "/servlet")
	assert.Equal(t, gatewayService._ROCKETGATE_HOST, "test.rocketgate.com")
	assert.Equal(t, gatewayService._ROCKETGATE_PROTOCOL, "http")
	assert.Equal(t, gatewayService._ROCKETGATE_PORTNO, 8443)
	assert.Equal(t, gatewayService._ROCKETGATE_CONNECT_TIMEOUT, 5555)
	assert.Equal(t, gatewayService._ROCKETGATE_READ_TIMEOUT, 7777)

}
