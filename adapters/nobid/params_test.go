package nobid

import (
	"encoding/json"
	"testing"

	"github.com/prebid/prebid-server/v3/openrtb_ext"
)

func TestValidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, validParam := range validParams {
		if err := validator.Validate(openrtb_ext.BidderNoBid, json.RawMessage(validParam)); err != nil {
			t.Errorf("Schema rejected nobid params: %s", validParam)
		}
	}
}

func TestInvalidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, invalidParam := range invalidParams {
		if err := validator.Validate(openrtb_ext.BidderNoBid, json.RawMessage(invalidParam)); err == nil {
			t.Errorf("Schema allowed unexpected params: %s", invalidParam)
		}
	}
}

var validParams = []string{
	`{"siteId": 25251, "bidfloor": 0.01}`,
}

var invalidParams = []string{
	`null`,
	`nil`,
	``,
	`[]`,
	`true`,
	`{}`,
	`{"tagid":12345678}`,
	`{"tagId":"12345678"}`,
	`{"tagid":"25251", "bidfloor": 0.01}`,
	`{"tagId": 12345678}`,
	`{"placementId": "1"}`,
	`{"placementId": 1}`,
}
