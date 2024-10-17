package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResponseAssignSuccess represents a ResponseAssignSuccess struct.
type ResponseAssignSuccess struct {
	Success              []string       `json:"success"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseAssignSuccess.
// It customizes the JSON marshaling process for ResponseAssignSuccess objects.
func (r ResponseAssignSuccess) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAssignSuccess object to a map representation for JSON marshaling.
func (r ResponseAssignSuccess) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["success"] = r.Success
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAssignSuccess.
// It customizes the JSON unmarshaling process for ResponseAssignSuccess objects.
func (r *ResponseAssignSuccess) UnmarshalJSON(input []byte) error {
	var temp tempResponseAssignSuccess
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "success")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Success = *temp.Success
	return nil
}

// tempResponseAssignSuccess is a temporary struct used for validating the fields of ResponseAssignSuccess.
type tempResponseAssignSuccess struct {
	Success *[]string `json:"success"`
}

func (r *tempResponseAssignSuccess) validate() error {
	var errs []string
	if r.Success == nil {
		errs = append(errs, "required field `success` is missing for type `response_assign_success`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
