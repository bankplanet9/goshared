package dto

import (
	"testing"
)

func TestClientIsPreApproved_GoodCHAndLimit(t *testing.T) {
	hasGoodCreditHistory := true
	limit := 2000.00
	c := &Client{
		IsIdentified:         true,
		HasGoodCreditHistory: &hasGoodCreditHistory,
		Limit:                &limit,
	}

	if !c.IsPreApproved() {
		t.Fatal("User has good credit history & limit assigned but has been not pre-approved")
	}
}

func TestClientIsPreApproved_BadClientButQuotaAssigned(t *testing.T) {
	hasGoodCreditHistory := false
	HasQuotaAssigned := true
	limit := 1000.00
	c := &Client{
		IsIdentified:         true,
		HasGoodCreditHistory: &hasGoodCreditHistory,
		Limit:                &limit,
		HasQuotaAssigned:     &HasQuotaAssigned,
	}

	if !c.IsPreApproved() {
		t.Fatal("User has bad credit history and quota assigned but has been not pre-approved")
	}
}

func TestClientIsPreApproved_BadClientNotQuotaAssigned(t *testing.T) {
	hasGoodCreditHistory := false
	HasQuotaAssigned := false
	limit := 1000.00
	c := &Client{
		IsIdentified:         true,
		HasGoodCreditHistory: &hasGoodCreditHistory,
		Limit:                &limit,
		HasQuotaAssigned:     &HasQuotaAssigned,
	}

	if c.IsPreApproved() {
		t.Fatal("User has bad credit history and has not quota assigned but has been pre-approved")
	}
}

func TestClientIsPreApproved_NotIdentified(t *testing.T) {
	limit := 1000.00
	c := &Client{
		IsIdentified:         false,
		HasGoodCreditHistory: nil,
		Limit:                &limit,
	}

	if !c.IsPreApproved() {
		t.Fatal("User is not identified, has a limit assigned but has not been pre-approved")
	}
}
