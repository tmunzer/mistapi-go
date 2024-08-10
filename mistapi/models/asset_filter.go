package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// AssetFilter represents a AssetFilter struct.
// Asset Filter
type AssetFilter struct {
    ApMac                 *string        `json:"ap_mac,omitempty"`
    Beam                  *int           `json:"beam,omitempty"`
    CreatedTime           *float64       `json:"created_time,omitempty"`
    Disabled              *bool          `json:"disabled,omitempty"`
    // whether the asset filter is disabled
    Disasbled             *bool          `json:"disasbled,omitempty"`
    // eddystone uid namespace used to filter assets
    EddystoneUidNamespace *string        `json:"eddystone_uid_namespace,omitempty"`
    // eddystone url used to filter assets
    EddystoneUrl          *string        `json:"eddystone_url,omitempty"`
    ForSite               *bool          `json:"for_site,omitempty"`
    // ibeacon major value used to filter assets
    IbeaconMajor          *int           `json:"ibeacon_major,omitempty"`
    // ibeacon uuid used to filter assets
    IbeaconUuid           *uuid.UUID     `json:"ibeacon_uuid,omitempty"`
    Id                    *uuid.UUID     `json:"id,omitempty"`
    // ble manufacturing-specific company-id used to filter assets
    MfgCompanyId          *int           `json:"mfg_company_id,omitempty"`
    ModifiedTime          *float64       `json:"modified_time,omitempty"`
    Name                  string         `json:"name"`
    OrgId                 *uuid.UUID     `json:"org_id,omitempty"`
    Rssi                  *int           `json:"rssi,omitempty"`
    // ble service data uuid used to filter assets
    ServiceUuid           *uuid.UUID     `json:"service_uuid,omitempty"`
    SiteId                *uuid.UUID     `json:"site_id,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AssetFilter.
// It customizes the JSON marshaling process for AssetFilter objects.
func (a AssetFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AssetFilter object to a map representation for JSON marshaling.
func (a AssetFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ApMac != nil {
        structMap["ap_mac"] = a.ApMac
    }
    if a.Beam != nil {
        structMap["beam"] = a.Beam
    }
    if a.CreatedTime != nil {
        structMap["created_time"] = a.CreatedTime
    }
    if a.Disabled != nil {
        structMap["disabled"] = a.Disabled
    }
    if a.Disasbled != nil {
        structMap["disasbled"] = a.Disasbled
    }
    if a.EddystoneUidNamespace != nil {
        structMap["eddystone_uid_namespace"] = a.EddystoneUidNamespace
    }
    if a.EddystoneUrl != nil {
        structMap["eddystone_url"] = a.EddystoneUrl
    }
    if a.ForSite != nil {
        structMap["for_site"] = a.ForSite
    }
    if a.IbeaconMajor != nil {
        structMap["ibeacon_major"] = a.IbeaconMajor
    }
    if a.IbeaconUuid != nil {
        structMap["ibeacon_uuid"] = a.IbeaconUuid
    }
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.MfgCompanyId != nil {
        structMap["mfg_company_id"] = a.MfgCompanyId
    }
    if a.ModifiedTime != nil {
        structMap["modified_time"] = a.ModifiedTime
    }
    structMap["name"] = a.Name
    if a.OrgId != nil {
        structMap["org_id"] = a.OrgId
    }
    if a.Rssi != nil {
        structMap["rssi"] = a.Rssi
    }
    if a.ServiceUuid != nil {
        structMap["service_uuid"] = a.ServiceUuid
    }
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AssetFilter.
// It customizes the JSON unmarshaling process for AssetFilter objects.
func (a *AssetFilter) UnmarshalJSON(input []byte) error {
    var temp tempAssetFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "beam", "created_time", "disabled", "disasbled", "eddystone_uid_namespace", "eddystone_url", "for_site", "ibeacon_major", "ibeacon_uuid", "id", "mfg_company_id", "modified_time", "name", "org_id", "rssi", "service_uuid", "site_id")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.ApMac = temp.ApMac
    a.Beam = temp.Beam
    a.CreatedTime = temp.CreatedTime
    a.Disabled = temp.Disabled
    a.Disasbled = temp.Disasbled
    a.EddystoneUidNamespace = temp.EddystoneUidNamespace
    a.EddystoneUrl = temp.EddystoneUrl
    a.ForSite = temp.ForSite
    a.IbeaconMajor = temp.IbeaconMajor
    a.IbeaconUuid = temp.IbeaconUuid
    a.Id = temp.Id
    a.MfgCompanyId = temp.MfgCompanyId
    a.ModifiedTime = temp.ModifiedTime
    a.Name = *temp.Name
    a.OrgId = temp.OrgId
    a.Rssi = temp.Rssi
    a.ServiceUuid = temp.ServiceUuid
    a.SiteId = temp.SiteId
    return nil
}

// tempAssetFilter is a temporary struct used for validating the fields of AssetFilter.
type tempAssetFilter  struct {
    ApMac                 *string    `json:"ap_mac,omitempty"`
    Beam                  *int       `json:"beam,omitempty"`
    CreatedTime           *float64   `json:"created_time,omitempty"`
    Disabled              *bool      `json:"disabled,omitempty"`
    Disasbled             *bool      `json:"disasbled,omitempty"`
    EddystoneUidNamespace *string    `json:"eddystone_uid_namespace,omitempty"`
    EddystoneUrl          *string    `json:"eddystone_url,omitempty"`
    ForSite               *bool      `json:"for_site,omitempty"`
    IbeaconMajor          *int       `json:"ibeacon_major,omitempty"`
    IbeaconUuid           *uuid.UUID `json:"ibeacon_uuid,omitempty"`
    Id                    *uuid.UUID `json:"id,omitempty"`
    MfgCompanyId          *int       `json:"mfg_company_id,omitempty"`
    ModifiedTime          *float64   `json:"modified_time,omitempty"`
    Name                  *string    `json:"name"`
    OrgId                 *uuid.UUID `json:"org_id,omitempty"`
    Rssi                  *int       `json:"rssi,omitempty"`
    ServiceUuid           *uuid.UUID `json:"service_uuid,omitempty"`
    SiteId                *uuid.UUID `json:"site_id,omitempty"`
}

func (a *tempAssetFilter) validate() error {
    var errs []string
    if a.Name == nil {
        errs = append(errs, "required field `name` is missing for type `asset_filter`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
