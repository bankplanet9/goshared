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


func TestClientAssignLimit_Assigned(t *testing.T) {
	limit := 1000.00
	c := &Client{
	}
	c.AssignLimit(limit)

	if *c.Limit != limit {
		t.Fatal("User has been assigned a limit, but the final limit differs from the assigned one")
	}
}


func TestClientAssignQuota_Assigned(t *testing.T) {
	limit := 1000.00
	c := &Client{
	}
	c.AssignQuota(limit)

	if *c.Limit != limit {
		t.Fatal("User has been assigned a quota, but the final limit differs from the assigned one")
	}
}