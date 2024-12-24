package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// Site represents a Site struct.
// Site
type Site struct {
    // full address of the site
    Address              *string                `json:"address,omitempty"`
    // Alarm Template ID, this takes precedence over the Org-level alarmtemplate_id
    AlarmtemplateId      Optional[uuid.UUID]    `json:"alarmtemplate_id"`
    // AP Template ID, used by APs
    AptemplateId         Optional[uuid.UUID]    `json:"aptemplate_id"`
    // country code for the site (for AP config generation), in two-character
    CountryCode          *string                `json:"country_code,omitempty"`
    // when the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // Gateway Template ID, used by gateways
    GatewaytemplateId    Optional[uuid.UUID]    `json:"gatewaytemplate_id"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Latlng               *LatLng                `json:"latlng,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Name                 string                 `json:"name"`
    // Network Template ID, this takes precedence over Site Settings
    NetworktemplateId    Optional[uuid.UUID]    `json:"networktemplate_id"`
    // optional, any notes about the site
    Notes                *string                `json:"notes,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // RF Template ID, this takes precedence over Site Settings
    RftemplateId         Optional[uuid.UUID]    `json:"rftemplate_id"`
    // SecPolicy ID
    SecpolicyId          Optional[uuid.UUID]    `json:"secpolicy_id"`
    // sitegroups this site belongs to
    SitegroupIds         []uuid.UUID            `json:"sitegroup_ids,omitempty"`
    // Site Template ID
    SitetemplateId       Optional[uuid.UUID]    `json:"sitetemplate_id"`
    // Timezone the site is at
    Timezone             *string                `json:"timezone,omitempty"`
    Tzoffset             *int                   `json:"tzoffset,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Site,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Site) String() string {
    return fmt.Sprintf(
    	"Site[Address=%v, AlarmtemplateId=%v, AptemplateId=%v, CountryCode=%v, CreatedTime=%v, GatewaytemplateId=%v, Id=%v, Latlng=%v, ModifiedTime=%v, Name=%v, NetworktemplateId=%v, Notes=%v, OrgId=%v, RftemplateId=%v, SecpolicyId=%v, SitegroupIds=%v, SitetemplateId=%v, Timezone=%v, Tzoffset=%v, AdditionalProperties=%v]",
    	s.Address, s.AlarmtemplateId, s.AptemplateId, s.CountryCode, s.CreatedTime, s.GatewaytemplateId, s.Id, s.Latlng, s.ModifiedTime, s.Name, s.NetworktemplateId, s.Notes, s.OrgId, s.RftemplateId, s.SecpolicyId, s.SitegroupIds, s.SitetemplateId, s.Timezone, s.Tzoffset, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Site.
// It customizes the JSON marshaling process for Site objects.
func (s Site) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "address", "alarmtemplate_id", "aptemplate_id", "country_code", "created_time", "gatewaytemplate_id", "id", "latlng", "modified_time", "name", "networktemplate_id", "notes", "org_id", "rftemplate_id", "secpolicy_id", "sitegroup_ids", "sitetemplate_id", "timezone", "tzoffset"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Site object to a map representation for JSON marshaling.
func (s Site) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Address != nil {
        structMap["address"] = s.Address
    }
    if s.AlarmtemplateId.IsValueSet() {
        if s.AlarmtemplateId.Value() != nil {
            structMap["alarmtemplate_id"] = s.AlarmtemplateId.Value()
        } else {
            structMap["alarmtemplate_id"] = nil
        }
    }
    if s.AptemplateId.IsValueSet() {
        if s.AptemplateId.Value() != nil {
            structMap["aptemplate_id"] = s.AptemplateId.Value()
        } else {
            structMap["aptemplate_id"] = nil
        }
    }
    if s.CountryCode != nil {
        structMap["country_code"] = s.CountryCode
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.GatewaytemplateId.IsValueSet() {
        if s.GatewaytemplateId.Value() != nil {
            structMap["gatewaytemplate_id"] = s.GatewaytemplateId.Value()
        } else {
            structMap["gatewaytemplate_id"] = nil
        }
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Latlng != nil {
        structMap["latlng"] = s.Latlng.toMap()
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    structMap["name"] = s.Name
    if s.NetworktemplateId.IsValueSet() {
        if s.NetworktemplateId.Value() != nil {
            structMap["networktemplate_id"] = s.NetworktemplateId.Value()
        } else {
            structMap["networktemplate_id"] = nil
        }
    }
    if s.Notes != nil {
        structMap["notes"] = s.Notes
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.RftemplateId.IsValueSet() {
        if s.RftemplateId.Value() != nil {
            structMap["rftemplate_id"] = s.RftemplateId.Value()
        } else {
            structMap["rftemplate_id"] = nil
        }
    }
    if s.SecpolicyId.IsValueSet() {
        if s.SecpolicyId.Value() != nil {
            structMap["secpolicy_id"] = s.SecpolicyId.Value()
        } else {
            structMap["secpolicy_id"] = nil
        }
    }
    if s.SitegroupIds != nil {
        structMap["sitegroup_ids"] = s.SitegroupIds
    }
    if s.SitetemplateId.IsValueSet() {
        if s.SitetemplateId.Value() != nil {
            structMap["sitetemplate_id"] = s.SitetemplateId.Value()
        } else {
            structMap["sitetemplate_id"] = nil
        }
    }
    if s.Timezone != nil {
        structMap["timezone"] = s.Timezone
    }
    if s.Tzoffset != nil {
        structMap["tzoffset"] = s.Tzoffset
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Site.
// It customizes the JSON unmarshaling process for Site objects.
func (s *Site) UnmarshalJSON(input []byte) error {
    var temp tempSite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "address", "alarmtemplate_id", "aptemplate_id", "country_code", "created_time", "gatewaytemplate_id", "id", "latlng", "modified_time", "name", "networktemplate_id", "notes", "org_id", "rftemplate_id", "secpolicy_id", "sitegroup_ids", "sitetemplate_id", "timezone", "tzoffset")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Address = temp.Address
    s.AlarmtemplateId = temp.AlarmtemplateId
    s.AptemplateId = temp.AptemplateId
    s.CountryCode = temp.CountryCode
    s.CreatedTime = temp.CreatedTime
    s.GatewaytemplateId = temp.GatewaytemplateId
    s.Id = temp.Id
    s.Latlng = temp.Latlng
    s.ModifiedTime = temp.ModifiedTime
    s.Name = *temp.Name
    s.NetworktemplateId = temp.NetworktemplateId
    s.Notes = temp.Notes
    s.OrgId = temp.OrgId
    s.RftemplateId = temp.RftemplateId
    s.SecpolicyId = temp.SecpolicyId
    s.SitegroupIds = temp.SitegroupIds
    s.SitetemplateId = temp.SitetemplateId
    s.Timezone = temp.Timezone
    s.Tzoffset = temp.Tzoffset
    return nil
}

// tempSite is a temporary struct used for validating the fields of Site.
type tempSite  struct {
    Address           *string             `json:"address,omitempty"`
    AlarmtemplateId   Optional[uuid.UUID] `json:"alarmtemplate_id"`
    AptemplateId      Optional[uuid.UUID] `json:"aptemplate_id"`
    CountryCode       *string             `json:"country_code,omitempty"`
    CreatedTime       *float64            `json:"created_time,omitempty"`
    GatewaytemplateId Optional[uuid.UUID] `json:"gatewaytemplate_id"`
    Id                *uuid.UUID          `json:"id,omitempty"`
    Latlng            *LatLng             `json:"latlng,omitempty"`
    ModifiedTime      *float64            `json:"modified_time,omitempty"`
    Name              *string             `json:"name"`
    NetworktemplateId Optional[uuid.UUID] `json:"networktemplate_id"`
    Notes             *string             `json:"notes,omitempty"`
    OrgId             *uuid.UUID          `json:"org_id,omitempty"`
    RftemplateId      Optional[uuid.UUID] `json:"rftemplate_id"`
    SecpolicyId       Optional[uuid.UUID] `json:"secpolicy_id"`
    SitegroupIds      []uuid.UUID         `json:"sitegroup_ids,omitempty"`
    SitetemplateId    Optional[uuid.UUID] `json:"sitetemplate_id"`
    Timezone          *string             `json:"timezone,omitempty"`
    Tzoffset          *int                `json:"tzoffset,omitempty"`
}

func (s *tempSite) validate() error {
    var errs []string
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `site`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
