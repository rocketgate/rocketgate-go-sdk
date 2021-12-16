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
	// Setup the Purchase request.
	gatewayRequest.Set(request.MERCHANT_ID, "1")
	gatewayRequest.Set(request.MERCHANT_PASSWORD, "testpassword")
	// For example/testing, we set the order id and customer as the unix timestamp as a convienent sequencing value
	// appending a test name to the order id to facilitate some clarity when reviewing the tests
	time := time.Now().Unix()
	cust_id := fmt.Sprint(time) + ".GoTest"
	inv_id := fmt.Sprint(time) + ".3DSTest"
	gatewayRequest.Set(request.MERCHANT_CUSTOMER_ID, cust_id)
	gatewayRequest.Set(request.MERCHANT_INVOICE_ID, inv_id)
	//
	gatewayRequest.Set(request.CURRENCY, "USD")
	gatewayRequest.Set(request.AMOUNT, "9.99")

	gatewayRequest.Set(request.CARDNO, "4000000000001091") // This card will trigger a 3DS 2.0 stepUp in the TestProcessor
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
	// Request 3DS
	gatewayRequest.Set(request.USE_3D_SECURE, "TRUE")
	gatewayRequest.Set(request.BROWSER_USER_AGENT, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36")
	gatewayRequest.Set(request.BROWSER_ACCEPT_HEADER, "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")

	// Setup test parameters in the service and request.
	gatewayService.SetTestMode(true)

	// Optional manual gateway
	gatewayRequest.Set(request.GATEWAY_SERVER, "local.rocketgate.com")
	gatewayRequest.Set(request.GATEWAY_PORTNO, "8443")
	gatewayRequest.Set(request.GATEWAY_PROTOCOL, "https")

	//	Step 1: Perform the BIN intelligence transaction.
	gatewayService.PerformPurchase(gatewayRequest, gatewayResponse)
	response_code := gatewayResponse.Get(response.RESPONSE_CODE)
	reason_code := gatewayResponse.Get(response.REASON_CODE)
	if response_code != fmt.Sprint(2) && reason_code != fmt.Sprint(225) { // RESPONSE_RISK_FAIL, REASON_3DSECURE_INITIATION
		fmt.Println("Response Code: " + response_code)
		fmt.Println("Reason Code: " + reason_code)
		return
	}

	fmt.Println("3DS 2.0 Device Fingerprinting Succeeded!")
	fmt.Println("Response Code: " + response_code)
	fmt.Println("Reason Code: " + reason_code)
	fmt.Println("Device Fingerprinting URL: " + gatewayResponse.Get(response.V_3DSECURE_DEVICE_COLLECTION_URL))
	fmt.Println("Device Fingerprinting JWT: " + gatewayResponse.Get(response.V_3DSECURE_DEVICE_COLLECTION_JWT))
	fmt.Println("Exception: " + gatewayResponse.Get(response.EXCEPTION))

	// Recycle the first request and add two new fields
	gatewayRequest.Set(request.V_3DSECURE_DF_REFERENCE_ID, "fake")
	gatewayRequest.Set(request.V_3DSECURE_REDIRECT_URL, "fake")

	gatewayRequest.Set(request.BROWSER_JAVA_ENABLED, "TRUE")
	gatewayRequest.Set(request.BROWSER_LANGUAGE, "en-CA")
	gatewayRequest.Set(request.BROWSER_COLOR_DEPTH, "32")
	gatewayRequest.Set(request.BROWSER_SCREEN_HEIGHT, "1080")
	gatewayRequest.Set(request.BROWSER_SCREEN_WIDTH, "1920")
	gatewayRequest.Set(request.BROWSER_TIME_ZONE, "-240")

	//	Step 2: Perform the Lookup transaction.
	if gatewayService.PerformPurchase(gatewayRequest, gatewayResponse) {
		fmt.Println("Purchase succeeded")
		fmt.Println("Response Code: " + gatewayResponse.Get(response.RESPONSE_CODE))
		fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
		fmt.Println("GUID: " + gatewayResponse.Get(response.TRANSACT_ID))
		fmt.Println("Card Issuer: " + gatewayResponse.Get(response.CARD_ISSUER_NAME))
		fmt.Println("Account: " + gatewayResponse.Get(response.MERCHANT_ACCOUNT))
		fmt.Println("Exception: " + gatewayResponse.Get(response.EXCEPTION))
	} else if gatewayResponse.Get(response.REASON_CODE) == fmt.Sprint(202) {
		fmt.Println("3DS Lookup succeeded")
		fmt.Println("GUID: " + gatewayResponse.Get(response.TRANSACT_ID))
		fmt.Println("3DS Version: " + gatewayResponse.Get(response.V_3DSECURE_VERSION))
		fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
		fmt.Println("PAREQ: " + gatewayResponse.Get(response.PAREQ))
		fmt.Println("ACS URL: " + gatewayResponse.Get(response.ACS_URL))
		fmt.Println("STEP-UP URL: " + gatewayResponse.Get(response.V_3DSECURE_STEP_UP_URL))
		fmt.Println("STEP-UP JWT: " + gatewayResponse.Get(response.V_3DSECURE_STEP_UP_JWT))
		//	Setup the 3rd request.
		gatewayRequest := request.NewGatewayRequest()

		gatewayRequest.Set(request.MERCHANT_ID, "1")
		gatewayRequest.Set(request.MERCHANT_PASSWORD, "testpassword")

		gatewayRequest.Set(request.CVV2, "999")
		gatewayRequest.Set(request.REFERENCE_GUID, gatewayResponse.Get(response.TRANSACT_ID))
		// In a real transaction this would include the PARES returned from the Authentication
		// On dev we send through the SimulatedPARES + TRANSACT_ID
		pares := "SimulatedPARES" + gatewayResponse.Get(response.TRANSACT_ID)
		gatewayRequest.Set(request.PARES, pares)
		// Risk/Scrub Request Setting
		gatewayRequest.Set(request.SCRUB, "IGNORE")
		gatewayRequest.Set(request.CVV2_CHECK, "IGNORE")
		gatewayRequest.Set(request.AVS_CHECK, "IGNORE")

		gatewayRequest.Set(request.MERCHANT_CUSTOMER_ID, cust_id)
		//
		gatewayRequest.Set(request.CURRENCY, "USD")
		gatewayRequest.Set(request.AMOUNT, "9.99")
		// Optional manual gateway
		gatewayRequest.Set(request.GATEWAY_SERVER, "local.rocketgate.com")
		gatewayRequest.Set(request.GATEWAY_PORTNO, "8443")
		gatewayRequest.Set(request.GATEWAY_PROTOCOL, "https")

		// Step 3: Perform the Purchase transaction.
		if gatewayService.PerformPurchase(gatewayRequest, gatewayResponse) {
			fmt.Println("Purchase succeeded")
			fmt.Println("Response Code: " + gatewayResponse.Get(response.RESPONSE_CODE))
			fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
			fmt.Println("GUID: " + gatewayResponse.Get(response.TRANSACT_ID))
		} else {
			fmt.Println("Purchase failed")
			fmt.Println("Response Code: " + gatewayResponse.Get(response.RESPONSE_CODE))
			fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
			fmt.Println("GUID: " + gatewayResponse.Get(response.TRANSACT_ID))
			fmt.Println("Exception: " + gatewayResponse.Get(response.EXCEPTION))
		}
	} else {
		fmt.Println("Purchase failed")
		fmt.Println("Response Code: " + gatewayResponse.Get(response.RESPONSE_CODE))
		fmt.Println("Reason Code: " + gatewayResponse.Get(response.REASON_CODE))
		fmt.Println("GUID: " + gatewayResponse.Get(response.TRANSACT_ID))
		fmt.Println("Exception: " + gatewayResponse.Get(response.EXCEPTION))
	}
}
