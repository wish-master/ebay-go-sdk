package ebay_go_sdk

type TRangeValue struct {
	End            string `json:"end"`
	ExclusiveEnd   bool   `json:"exclusiveEnd"`
	ExclusiveStart bool   `json:"exclusiveStart"`
	Range          bool   `json:"range"`
	Start          string `json:"start"`
}

type TFilterField struct {
	Field string      `json:"field"`
	Range TRangeValue `json:"range"`
	Set   []string    `json:"set"`
	Value string      `json:"value"`
}

type TGetOrdersResponse struct {
	Href   string `json:"href"`
	Limit  string `json:"limit"`
	Next   string `json:"next"`
	Offset string `json:"offset"`
	Orders []struct {
		Buyer struct {
			BuyerRegistrationAddress struct {
				CompanyName    string `json:"companyName"`
				ContactAddress struct {
					AddressLine1    string `json:"addressLine1"`
					AddressLine2    string `json:"addressLine2"`
					City            string `json:"city"`
					CountryCode     string `json:"countryCode"`
					County          string `json:"county"`
					PostalCode      string `json:"postalCode"`
					StateOrProvince string `json:"stateOrProvince"`
				} `json:"contactAddress"`
				Email        string `json:"email"`
				FullName     string `json:"fullName"`
				PrimaryPhone struct {
					PhoneNumber string `json:"phoneNumber"`
				} `json:"primaryPhone"`
			} `json:"buyerRegistrationAddress"`
			TaxAddress struct {
				City            string `json:"city"`
				CountryCode     string `json:"countryCode"`
				PostalCode      string `json:"postalCode"`
				StateOrProvince string `json:"stateOrProvince"`
			} `json:"taxAddress"`
			TaxIdentifier struct {
				TaxpayerId        string `json:"taxpayerId"`
				TaxIdentifierType string `json:"taxIdentifierType"`
				IssuingCountry    string `json:"issuingCountry"`
			} `json:"taxIdentifier"`
			Username string `json:"username"`
		} `json:"buyer"`
		BuyerCheckoutNotes string `json:"buyerCheckoutNotes"`
		CancelStatus       struct {
			CancelledDate  string `json:"cancelledDate"`
			CancelRequests []struct {
				CancelCompletedDate string `json:"cancelCompletedDate"`
				CancelInitiator     string `json:"cancelInitiator"`
				CancelReason        string `json:"cancelReason"`
				CancelRequestedDate string `json:"cancelRequestedDate"`
				CancelRequestId     string `json:"cancelRequestId"`
				CancelRequestState  string `json:"cancelRequestState"`
			} `json:"cancelRequests"`
			CancelState string `json:"cancelState"`
		} `json:"cancelStatus"`
		CreationDate                 string   `json:"creationDate"`
		EbayCollectAndRemitTax       string   `json:"ebayCollectAndRemitTax"`
		FulfillmentHrefs             []string `json:"fulfillmentHrefs"`
		FulfillmentStartInstructions []struct {
			DestinationTimeZone      string `json:"destinationTimeZone"`
			EbaySupportedFulfillment string `json:"ebaySupportedFulfillment"`
			FinalDestinationAddress  struct {
				AddressLine1    string `json:"addressLine1"`
				AddressLine2    string `json:"addressLine2"`
				City            string `json:"city"`
				CountryCode     string `json:"countryCode"`
				County          string `json:"county"`
				PostalCode      string `json:"postalCode"`
				StateOrProvince string `json:"stateOrProvince"`
			} `json:"finalDestinationAddress"`
			FulfillmentInstructionsType string `json:"fulfillmentInstructionsType"`
			MaxEstimatedDeliveryDate    string `json:"maxEstimatedDeliveryDate"`
			MinEstimatedDeliveryDate    string `json:"minEstimatedDeliveryDate"`
			PickupStep                  struct {
				MerchantLocationKey string `json:"merchantLocationKey"`
			} `json:"pickupStep"`
			ShippingStep struct {
				ShippingCarrierCode string `json:"shippingCarrierCode"`
				ShippingServiceCode string `json:"shippingServiceCode"`
				ShipTo              struct {
					CompanyName    string `json:"companyName"`
					ContactAddress struct {
						AddressLine1    string `json:"addressLine1"`
						AddressLine2    string `json:"addressLine2"`
						City            string `json:"city"`
						CountryCode     string `json:"countryCode"`
						County          string `json:"county"`
						PostalCode      string `json:"postalCode"`
						StateOrProvince string `json:"stateOrProvince"`
					} `json:"contactAddress"`
					Email        string `json:"email"`
					FullName     string `json:"fullName"`
					PrimaryPhone struct {
						PhoneNumber string `json:"phoneNumber"`
					} `json:"primaryPhone"`
				} `json:"shipTo"`
				ShipToReferenceId string `json:"shipToReferenceId"`
			} `json:"shippingStep"`
		} `json:"fulfillmentStartInstructions"`
		LastModifiedDate string `json:"lastModifiedDate"`
		LegacyOrderId    string `json:"legacyOrderId"`
		LineItems        []struct {
			AppliedPromotions []struct {
				Description    string `json:"description"`
				DiscountAmount struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"discountAmount"`
				PromotionId string `json:"promotionId"`
			} `json:"appliedPromotions"`
			DeliveryCost struct {
				DiscountAmount string `json:"discountAmount"`
				HandlingCost   string `json:"handlingCost"`
				ImportCharges  struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"importCharges"`
				ShippingCost struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"shippingCost"`
				ShippingIntermediationFee struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"shippingIntermediationFee"`
			} `json:"deliveryCost"`
			DiscountedLineItemCost struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"discountedLineItemCost"`
			EbayCollectAndRemitTaxes []struct {
				Amount struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"amount"`
				EbayReference struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"ebayReference"`
				TaxType          string `json:"taxType"`
				CollectionMethod string `json:"collectionMethod"`
			} `json:"ebayCollectAndRemitTaxes"`
			EbayCollectedCharges struct {
				EbayShipping struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"ebayShipping"`
			} `json:"ebayCollectedCharges"`
			GiftDetails struct {
				Message        string `json:"message"`
				RecipientEmail string `json:"recipientEmail"`
				SenderName     string `json:"senderName"`
			} `json:"giftDetails"`
			ItemLocation struct {
				CountryCode string `json:"countryCode"`
				Location    string `json:"location"`
				PostalCode  string `json:"postalCode"`
			} `json:"itemLocation"`
			LegacyItemId      string `json:"legacyItemId"`
			LegacyVariationId string `json:"legacyVariationId"`
			LineItemCost      struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"lineItemCost"`
			LineItemFulfillmentInstructions struct {
				DestinationTimeZone      string `json:"destinationTimeZone"`
				GuaranteedDelivery       string `json:"guaranteedDelivery"`
				MaxEstimatedDeliveryDate string `json:"maxEstimatedDeliveryDate"`
				MinEstimatedDeliveryDate string `json:"minEstimatedDeliveryDate"`
				ShipByDate               string `json:"shipByDate"`
				SourceTimeZone           string `json:"sourceTimeZone"`
			} `json:"lineItemFulfillmentInstructions"`
			LineItemFulfillmentStatus string `json:"lineItemFulfillmentStatus"`
			LineItemId                string `json:"lineItemId"`
			LinkedOrderLineItems      []struct {
				LineItemAspects []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"lineItemAspects"`
				LineItemId               string `json:"lineItemId"`
				MaxEstimatedDeliveryDate string `json:"maxEstimatedDeliveryDate"`
				MinEstimatedDeliveryDate string `json:"minEstimatedDeliveryDate"`
				OrderId                  string `json:"orderId"`
				SellerId                 string `json:"sellerId"`
				Shipments                []struct {
					ShipmentTrackingNumber string `json:"shipmentTrackingNumber"`
					ShippingCarrierCode    string `json:"shippingCarrierCode"`
				} `json:"shipments"`
				Title string `json:"title"`
			} `json:"linkedOrderLineItems"`
			ListingMarketplaceId string `json:"listingMarketplaceId"`
			Properties           struct {
				BuyerProtection   string `json:"buyerProtection"`
				FromBestOffer     string `json:"fromBestOffer"`
				SoldViaAdCampaign string `json:"soldViaAdCampaign"`
			} `json:"properties"`
			PurchaseMarketplaceId string `json:"purchaseMarketplaceId"`
			Quantity              string `json:"quantity"`
			Refunds               []struct {
				Amount struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"amount"`
				RefundDate        string `json:"refundDate"`
				RefundId          string `json:"refundId"`
				RefundReferenceId string `json:"refundReferenceId"`
			} `json:"refunds"`
			Sku        string `json:"sku"`
			SoldFormat string `json:"soldFormat"`
			Taxes      []struct {
				Amount struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"amount"`
				TaxType string `json:"taxType"`
			} `json:"taxes"`
			Title string `json:"title"`
			Total struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"total"`
			VariationAspects []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"variationAspects"`
		} `json:"lineItems"`
		OrderFulfillmentStatus string `json:"orderFulfillmentStatus"`
		OrderId                string `json:"orderId"`
		OrderPaymentStatus     string `json:"orderPaymentStatus"`
		PaymentSummary         struct {
			Payments []struct {
				Amount struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"amount"`
				PaymentDate  string `json:"paymentDate"`
				PaymentHolds []struct {
					ExpectedReleaseDate string `json:"expectedReleaseDate"`
					HoldAmount          struct {
						ConvertedFromCurrency string `json:"convertedFromCurrency"`
						ConvertedFromValue    string `json:"convertedFromValue"`
						Currency              string `json:"currency"`
						Value                 string `json:"value"`
					} `json:"holdAmount"`
					HoldReason             string `json:"holdReason"`
					HoldState              string `json:"holdState"`
					ReleaseDate            string `json:"releaseDate"`
					SellerActionsToRelease []struct {
						SellerActionToRelease string `json:"sellerActionToRelease"`
					} `json:"sellerActionsToRelease"`
				} `json:"paymentHolds"`
				PaymentMethod      string `json:"paymentMethod"`
				PaymentReferenceId string `json:"paymentReferenceId"`
				PaymentStatus      string `json:"paymentStatus"`
			} `json:"payments"`
			Refunds []struct {
				Amount struct {
					ConvertedFromCurrency string `json:"convertedFromCurrency"`
					ConvertedFromValue    string `json:"convertedFromValue"`
					Currency              string `json:"currency"`
					Value                 string `json:"value"`
				} `json:"amount"`
				RefundDate        string `json:"refundDate"`
				RefundId          string `json:"refundId"`
				RefundReferenceId string `json:"refundReferenceId"`
				RefundStatus      string `json:"refundStatus"`
			} `json:"refunds"`
			TotalDueSeller struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"totalDueSeller"`
		} `json:"paymentSummary"`
		PricingSummary struct {
			Adjustment struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"adjustment"`
			DeliveryCost struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"deliveryCost"`
			DeliveryDiscount struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"deliveryDiscount"`
			Fee struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"fee"`
			PriceDiscount struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"priceDiscount"`
			PriceSubtotal struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"priceSubtotal"`
			Tax struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"tax"`
			Total struct {
				ConvertedFromCurrency string `json:"convertedFromCurrency"`
				ConvertedFromValue    string `json:"convertedFromValue"`
				Currency              string `json:"currency"`
				Value                 string `json:"value"`
			} `json:"total"`
		} `json:"pricingSummary"`
		Program struct {
			AuthenticityVerification struct {
				OutcomeReason string `json:"outcomeReason"`
				Status        string `json:"status"`
			} `json:"authenticityVerification"`
			EbayShipping struct {
				ShippingLabelProvidedBy string `json:"shippingLabelProvidedBy"`
			} `json:"ebayShipping"`
			EbayVault struct {
				FulfillmentType string `json:"fulfillmentType"`
			} `json:"ebayVault"`
			EbayInternationalShipping struct {
				ReturnsManagedBy string `json:"returnsManagedBy"`
			} `json:"ebayInternationalShipping"`
			FulfillmentProgram struct {
				FulfilledBy string `json:"fulfilledBy"`
			} `json:"fulfillmentProgram"`
		} `json:"program"`
		SalesRecordReference string `json:"salesRecordReference"`
		SellerId             string `json:"sellerId"`
		TotalFeeBasisAmount  struct {
			ConvertedFromCurrency string `json:"convertedFromCurrency"`
			ConvertedFromValue    string `json:"convertedFromValue"`
			Currency              string `json:"currency"`
			Value                 string `json:"value"`
		} `json:"totalFeeBasisAmount"`
		TotalMarketplaceFee struct {
			ConvertedFromCurrency string `json:"convertedFromCurrency"`
			ConvertedFromValue    string `json:"convertedFromValue"`
			Currency              string `json:"currency"`
			Value                 string `json:"value"`
		} `json:"totalMarketplaceFee"`
	} `json:"orders"`
	Prev     string `json:"prev"`
	Total    string `json:"total"`
	Warnings []struct {
		Category     string   `json:"category"`
		Domain       string   `json:"domain"`
		ErrorId      string   `json:"errorId"`
		InputRefIds  []string `json:"inputRefIds"`
		LongMessage  string   `json:"longMessage"`
		Message      string   `json:"message"`
		OutputRefIds []string `json:"outputRefIds"`
		Parameters   []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"parameters"`
		Subdomain string `json:"subdomain"`
	} `json:"warnings"`
}
