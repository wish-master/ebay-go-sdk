package ebay_go_sdk

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const path = "/sell/fulfillment/v1/order"

type IFulfilmentAPI interface {
	GetOrders(ctx context.Context, ids []string, filter []TFilterField, limit, offset string, fieldGroups []string) (*TGetOrdersResponse, error)
}

func (s ebaySDK) GetOrders(ctx context.Context, ids []string, filter []TFilterField, limit, offset string, fieldGroups []string) (*TGetOrdersResponse, error) {
	req, err := http.NewRequest("GET", s.apiHost+path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get orders")
	}

	req = req.WithContext(ctx)
	q := req.URL.Query()

	if len(ids) != 0 {
		q.Add("ids", strings.Join(ids, ","))
	}

	if len(filter) != 0 {
		tmp := make([]string, len(filter))
		for i, field := range filter {
			b, err := json.Marshal(field)
			if err != nil {
				return nil, errors.Wrap(err, "failed to marshal TFilterField")
			}

			tmp[i] = string(b)
		}

		q.Add("filter", strings.Join(ids, ","))
	}

	if limit != "" {
		q.Add("limit", limit)
	}

	if offset != "" {
		q.Add("offset", offset)
	}

	req.URL.RawQuery = q.Encode()
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send GetOrders request")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)

	var result TGetOrdersResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetOrders response")
	}

	return &result, nil
}
