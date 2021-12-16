package main

import (
	"fmt"
	"time"

	"github.com/rocketgate/rocketgate-go-sdk/request"
	"github.com/rocketgate/rocketgate-go-sdk/response"
	"github.com/rocketgate/rocketgate-go-sdk/service"
)

func main() {
	// Create gateway objects
	gatewayRequest := request.NewGatewayRequest()
	gatewayResponse := response.NewGatewayResponse()
	gatewayService := service.NewGatewayService()
	// Setup the Auth-Only request.
	gatewayRequest.Set(request.MERCHANT_ID, "1")
	gatewayRequest.Set(request.MERCHANT_PASSWORD, "testpassword")
	// For example/testing, we set the order id and customer as the unix timestamp as a convienent sequencing value
	// appending a test name to the order id to facilitate some clarity when reviewing the tests
	time := time.Now().Unix()
	gatewayRequest.Set(request.MERCHANT_CUSTOMER_ID, fmt.Sprint(time)+".GoTest")
	gatewayRequest.Set(request.MERCHANT_INVOICE_ID, fmt.Sprint(time)+".AuthOnlyTest")
	//
	gatewayRequest.Set(request.AMOUNT, "9.99")
	gatewayRequest.Set(request.CURRENCY, "USD")

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
	// Risk/Scrub Request Setting
	gatewayRequest.Set(request.SCRUB, "IGNORE")
	gatewayRequest.Set(request.CVV2_CHECK, "IGNORE")
	gatewayRequest.Set(request.AVS_CHECK, "IGNORE")

	// Setup test parameters in the service and request.
	gatewayService.SetTestMode(true)

	// Optional manual gateway
	gatewayRequest.Set(request.GATEWAY_SERVER, "local.rocketgate.com")
	gatewayRequest.Set(request.GATEWAY_PORTNO, "8443")
	gatewayRequest.Set(request.GATEWAY_PROTOCOL, "https")

	//	Perform the Auth-Only transaction.
	if gatewayService.PerformAuthOnly(gatewayRequest, gatewayResponse) {
		fmt.Println("Auth-Only succeeded")
		fmt.Println("Response Code: " + gatewayResponse.Get(response.RESPONSE_CODE))
		fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
		fmt.Println("Auth No: " + gatewayResponse.Get(response.AUTH_NO))
		fmt.Println("AVS: " + gatewayResponse.Get(response.AVS_RESPONSE))
		fmt.Println("CVV2: " + gatewayResponse.Get(response.CVV2_CODE))
		fmt.Println("GUID: " + gatewayResponse.Get(response.TRANSACT_ID))
		fmt.Println("Transaction time: " + gatewayResponse.Get(response.TRANSACTION_TIME))
		fmt.Println("Account: " + gatewayResponse.Get(response.MERCHANT_ACCOUNT))
		fmt.Println("Scrub: " + gatewayResponse.Get(response.SCRUB_RESULTS))
	} else {
		fmt.Println("Auth-Only failed")
		fmt.Println("GUID: " + gatewayResponse.Get(response.TRANSACT_ID))
		fmt.Println("Transaction time: " + gatewayResponse.Get(response.TRANSACTION_TIME))
		fmt.Println("Response Code: " + gatewayResponse.Get(response.RESPONSE_CODE))
		fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
		fmt.Println("Exception: " + gatewayResponse.Get(response.EXCEPTION))
		fmt.Println("Scrub: " + gatewayResponse.Get(response.SCRUB_RESULTS))
	}
}
