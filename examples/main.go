package main

import (
	"fmt"

	"github.com/rocketgate/rocketgate-go-sdk/request"
	"github.com/rocketgate/rocketgate-go-sdk/response"
	"github.com/rocketgate/rocketgate-go-sdk/service"
)

func main() {
	// Create gateway objects
	gatewayRequest := request.NewGatewayRequest()
	gatewayResponse := response.NewGatewayResponse()
	gatewayService := service.NewGatewayService()

	gatewayRequest.Set(request.MERCHANT_ID, "1")
	gatewayRequest.Set(request.MERCHANT_PASSWORD, "testpassword")
	// TODO Add time prefix
	gatewayRequest.Set(request.MERCHANT_CUSTOMER_ID, "Go Test")
	gatewayRequest.Set(request.MERCHANT_INVOICE_ID, "Sale Test")
	//
	gatewayRequest.Set(request.AMOUNT, "9.99")
	gatewayRequest.Set(request.CARDNO, "4111111111111111")
	gatewayRequest.Set(request.EXPIRE_MONTH, "02")
	gatewayRequest.Set(request.EXPIRE_YEAR, "2023")
	gatewayRequest.Set(request.CVV2, "999")
	gatewayRequest.Set(request.CUSTOMER_FIRSTNAME, "Joe")
	gatewayRequest.Set(request.CUSTOMER_LASTNAME, "GO Tester")
	gatewayRequest.Set(request.EMAIL, "gotest@fakedomain.com")
	gatewayRequest.Set(request.IPADDRESS, "192.168.1.1")
	//
	gatewayRequest.Set(request.BILLING_ADDRESS, "123 Main St")
	gatewayRequest.Set(request.BILLING_CITY, "Las Vegas")
	gatewayRequest.Set(request.BILLING_STATE, "NV")
	gatewayRequest.Set(request.BILLING_ZIPCODE, "89141")
	gatewayRequest.Set(request.BILLING_COUNTRY, "US")
	//
	gatewayRequest.Set(request.SCRUB, "IGNORE")
	gatewayRequest.Set(request.CVV2_CHECK, "IGNORE")
	gatewayRequest.Set(request.AVS_CHECK, "IGNORE")

	gatewayService.SetTestMode(true)

	// Optional manual gateway
	gatewayRequest.Set(request.GATEWAY_SERVER, "local.rocketgate.com")
	gatewayRequest.Set(request.GATEWAY_PORTNO, "8443")
	gatewayRequest.Set(request.GATEWAY_PROTOCOL, "https")

	if gatewayService.PerformPurchase(gatewayRequest, gatewayResponse) {
		fmt.Println("Purchase succeeded")
		fmt.Println("Response Code: " + gatewayResponse.Get(response.RESPONSE_CODE))
		fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
	} else {
		fmt.Println("Purchase failed")
		fmt.Println("Response Code: " + gatewayResponse.Get(response.RESPONSE_CODE))
		fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
	}
}
