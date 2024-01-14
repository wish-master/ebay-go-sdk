package ebay_go_sdk

import (
	"net/http"
)

const (
	productionAPIHost = "https://api.ebay.com"
	sandboxAPIHost    = "https://api.sandbox.ebay.com"
)

type IEbaySDK interface {
	IFulfilmentAPI
}

type ebaySDK struct {
	client       *http.Client
	apiHost      string
	refreshToken string
}

func NewEbaySDK(clientID, clientSecret, ruName, refreshToken string, isSandboxMode bool) IEbaySDK {
	sdk := ebaySDK{
		client:       nil,
		apiHost:      sandboxAPIHost,
		refreshToken: refreshToken,
	}

	if !isSandboxMode {
		sdk.apiHost = productionAPIHost
	}

	return &sdk
}
