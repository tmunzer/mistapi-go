package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// MapJibestream represents a MapJibestream struct.
type MapJibestream struct {
    // the client id
    ClientId             uuid.UUID              `json:"client_id"`
    // the client secret
    ClientSecret         string                 `json:"client_secret"`
    // the jibestream customer record id
    CustomerId           int                    `json:"customer_id"`
    // the map contents endpoint host
    EndpointUrl          string                 `json:"endpoint_url"`
    // the jibestream map id
    MapId                uuid.UUID              `json:"map_id"`
    // millimeter per pixel
    Mmpp                 int                    `json:"mmpp"`
    // pixel per meter, same as the map JSON value.
    Ppm                  float64                `json:"ppm"`
    // the vendor ‘jibestream’. enum: `jibestream`
    VendorName           string                 `json:"vendor_name"`
    // the venue or organization id
    VenueId              int                    `json:"venue_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapJibestream.
// It customizes the JSON marshaling process for MapJibestream objects.
func (m MapJibestream) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "client_id", "client_secret", "customer_id", "endpoint_url", "map_id", "mmpp", "ppm", "vendor_name", "venue_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapJibestream object to a map representation for JSON marshaling.
func (m MapJibestream) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["client_id"] = m.ClientId
    structMap["client_secret"] = m.ClientSecret
    structMap["customer_id"] = m.CustomerId
    structMap["endpoint_url"] = m.EndpointUrl
    structMap["map_id"] = m.MapId
    structMap["mmpp"] = m.Mmpp
    structMap["ppm"] = m.Ppm
    structMap["vendor_name"] = m.VendorName
    structMap["venue_id"] = m.VenueId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapJibestream.
// It customizes the JSON unmarshaling process for MapJibestream objects.
func (m *MapJibestream) UnmarshalJSON(input []byte) error {
    var temp tempMapJibestream
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "client_id", "client_secret", "customer_id", "endpoint_url", "map_id", "mmpp", "ppm", "vendor_name", "venue_id")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.ClientId = *temp.ClientId
    m.ClientSecret = *temp.ClientSecret
    m.CustomerId = *temp.CustomerId
    m.EndpointUrl = *temp.EndpointUrl
    m.MapId = *temp.MapId
    m.Mmpp = *temp.Mmpp
    m.Ppm = *temp.Ppm
    m.VendorName = *temp.VendorName
    m.VenueId = *temp.VenueId
    return nil
}

// tempMapJibestream is a temporary struct used for validating the fields of MapJibestream.
type tempMapJibestream  struct {
    ClientId     *uuid.UUID `json:"client_id"`
    ClientSecret *string    `json:"client_secret"`
    CustomerId   *int       `json:"customer_id"`
    EndpointUrl  *string    `json:"endpoint_url"`
    MapId        *uuid.UUID `json:"map_id"`
    Mmpp         *int       `json:"mmpp"`
    Ppm          *float64   `json:"ppm"`
    VendorName   *string    `json:"vendor_name"`
    VenueId      *int       `json:"venue_id"`
}

func (m *tempMapJibestream) validate() error {
    var errs []string
    if m.ClientId == nil {
        errs = append(errs, "required field `client_id` is missing for type `map_jibestream`")
    }
    if m.ClientSecret == nil {
        errs = append(errs, "required field `client_secret` is missing for type `map_jibestream`")
    }
    if m.CustomerId == nil {
        errs = append(errs, "required field `customer_id` is missing for type `map_jibestream`")
    }
    if m.EndpointUrl == nil {
        errs = append(errs, "required field `endpoint_url` is missing for type `map_jibestream`")
    }
    if m.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `map_jibestream`")
    }
    if m.Mmpp == nil {
        errs = append(errs, "required field `mmpp` is missing for type `map_jibestream`")
    }
    if m.Ppm == nil {
        errs = append(errs, "required field `ppm` is missing for type `map_jibestream`")
    }
    if m.VendorName == nil {
        errs = append(errs, "required field `vendor_name` is missing for type `map_jibestream`")
    }
    if m.VenueId == nil {
        errs = append(errs, "required field `venue_id` is missing for type `map_jibestream`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
