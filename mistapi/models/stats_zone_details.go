package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsZoneDetails represents a StatsZoneDetails struct.
// Zone details statistics
type StatsZoneDetails struct {
    // list of ble assets currently in the zone and when they entered
    Assets               []string                    `json:"assets,omitempty"`
    // client wait time right now
    ClientWaits          StatsZoneDetailsClientWaits `json:"client_waits"`
    // list of clients currently in the zone and when they entered
    Clients              []string                    `json:"clients,omitempty"`
    // id of the zone
    Id                   uuid.UUID                   `json:"id"`
    // map_id of the zone
    MapId                uuid.UUID                   `json:"map_id"`
    // name of the zone
    Name                 string                      `json:"name"`
    NumClients           int                         `json:"num_clients"`
    // sdkclient wait time right now
    NumSdkclients        int                         `json:"num_sdkclients"`
    // list of sdkclients currently in the zone and when they entered
    Sdkclients           []string                    `json:"sdkclients,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsZoneDetails.
// It customizes the JSON marshaling process for StatsZoneDetails objects.
func (s StatsZoneDetails) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsZoneDetails object to a map representation for JSON marshaling.
func (s StatsZoneDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Assets != nil {
        structMap["assets"] = s.Assets
    }
    structMap["client_waits"] = s.ClientWaits.toMap()
    if s.Clients != nil {
        structMap["clients"] = s.Clients
    }
    structMap["id"] = s.Id
    structMap["map_id"] = s.MapId
    structMap["name"] = s.Name
    structMap["num_clients"] = s.NumClients
    structMap["num_sdkclients"] = s.NumSdkclients
    if s.Sdkclients != nil {
        structMap["sdkclients"] = s.Sdkclients
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsZoneDetails.
// It customizes the JSON unmarshaling process for StatsZoneDetails objects.
func (s *StatsZoneDetails) UnmarshalJSON(input []byte) error {
    var temp tempStatsZoneDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "assets", "client_waits", "clients", "id", "map_id", "name", "num_clients", "num_sdkclients", "sdkclients")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Assets = temp.Assets
    s.ClientWaits = *temp.ClientWaits
    s.Clients = temp.Clients
    s.Id = *temp.Id
    s.MapId = *temp.MapId
    s.Name = *temp.Name
    s.NumClients = *temp.NumClients
    s.NumSdkclients = *temp.NumSdkclients
    s.Sdkclients = temp.Sdkclients
    return nil
}

// tempStatsZoneDetails is a temporary struct used for validating the fields of StatsZoneDetails.
type tempStatsZoneDetails  struct {
    Assets        []string                     `json:"assets,omitempty"`
    ClientWaits   *StatsZoneDetailsClientWaits `json:"client_waits"`
    Clients       []string                     `json:"clients,omitempty"`
    Id            *uuid.UUID                   `json:"id"`
    MapId         *uuid.UUID                   `json:"map_id"`
    Name          *string                      `json:"name"`
    NumClients    *int                         `json:"num_clients"`
    NumSdkclients *int                         `json:"num_sdkclients"`
    Sdkclients    []string                     `json:"sdkclients,omitempty"`
}

func (s *tempStatsZoneDetails) validate() error {
    var errs []string
    if s.ClientWaits == nil {
        errs = append(errs, "required field `client_waits` is missing for type `stats_zone_details`")
    }
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `stats_zone_details`")
    }
    if s.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `stats_zone_details`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `stats_zone_details`")
    }
    if s.NumClients == nil {
        errs = append(errs, "required field `num_clients` is missing for type `stats_zone_details`")
    }
    if s.NumSdkclients == nil {
        errs = append(errs, "required field `num_sdkclients` is missing for type `stats_zone_details`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
