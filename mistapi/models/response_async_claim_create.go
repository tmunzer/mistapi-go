// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ResponseAsyncClaimCreate represents a ResponseAsyncClaimCreate struct.
// Response to an async inventory claim request
type ResponseAsyncClaimCreate struct {
	// Unique identifier for the async claim job, used to poll status
	ClaimId *uuid.UUID `json:"claim_id,omitempty"`
	// Inventory devices pending asynchronous claim processing
	InventoryPending     []ResponseClaimLicenseInventoryPendingItem `json:"inventory_pending,omitempty"`
	AdditionalProperties map[string]interface{}                     `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAsyncClaimCreate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAsyncClaimCreate) String() string {
	return fmt.Sprintf(
		"ResponseAsyncClaimCreate[ClaimId=%v, InventoryPending=%v, AdditionalProperties=%v]",
		r.ClaimId, r.InventoryPending, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAsyncClaimCreate.
// It customizes the JSON marshaling process for ResponseAsyncClaimCreate objects.
func (r ResponseAsyncClaimCreate) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"claim_id", "inventory_pending"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAsyncClaimCreate object to a map representation for JSON marshaling.
func (r ResponseAsyncClaimCreate) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.ClaimId != nil {
		structMap["claim_id"] = r.ClaimId
	}
	if r.InventoryPending != nil {
		structMap["inventory_pending"] = r.InventoryPending
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAsyncClaimCreate.
// It customizes the JSON unmarshaling process for ResponseAsyncClaimCreate objects.
func (r *ResponseAsyncClaimCreate) UnmarshalJSON(input []byte) error {
	var temp tempResponseAsyncClaimCreate
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "claim_id", "inventory_pending")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.ClaimId = temp.ClaimId
	r.InventoryPending = temp.InventoryPending
	return nil
}

// tempResponseAsyncClaimCreate is a temporary struct used for validating the fields of ResponseAsyncClaimCreate.
type tempResponseAsyncClaimCreate struct {
	ClaimId          *uuid.UUID                                 `json:"claim_id,omitempty"`
	InventoryPending []ResponseClaimLicenseInventoryPendingItem `json:"inventory_pending,omitempty"`
}
