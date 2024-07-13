package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ZoneStatsDetails represents a ZoneStatsDetails struct.
// Zone details statistics
type ZoneStatsDetails struct {
    // list of ble assets currently in the zone and when they entered
    Assets               []string                    `json:"assets,omitempty"`
    // client wait time right now
    ClientWaits          ZoneStatsDetailsClientWaits `json:"client_waits"`
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

// MarshalJSON implements the json.Marshaler interface for ZoneStatsDetails.
// It customizes the JSON marshaling process for ZoneStatsDetails objects.
func (z ZoneStatsDetails) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneStatsDetails object to a map representation for JSON marshaling.
func (z ZoneStatsDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, z.AdditionalProperties)
    if z.Assets != nil {
        structMap["assets"] = z.Assets
    }
    structMap["client_waits"] = z.ClientWaits.toMap()
    if z.Clients != nil {
        structMap["clients"] = z.Clients
    }
    structMap["id"] = z.Id
    structMap["map_id"] = z.MapId
    structMap["name"] = z.Name
    structMap["num_clients"] = z.NumClients
    structMap["num_sdkclients"] = z.NumSdkclients
    if z.Sdkclients != nil {
        structMap["sdkclients"] = z.Sdkclients
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneStatsDetails.
// It customizes the JSON unmarshaling process for ZoneStatsDetails objects.
func (z *ZoneStatsDetails) UnmarshalJSON(input []byte) error {
    var temp zoneStatsDetails
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
    
    z.AdditionalProperties = additionalProperties
    z.Assets = temp.Assets
    z.ClientWaits = *temp.ClientWaits
    z.Clients = temp.Clients
    z.Id = *temp.Id
    z.MapId = *temp.MapId
    z.Name = *temp.Name
    z.NumClients = *temp.NumClients
    z.NumSdkclients = *temp.NumSdkclients
    z.Sdkclients = temp.Sdkclients
    return nil
}

// zoneStatsDetails is a temporary struct used for validating the fields of ZoneStatsDetails.
type zoneStatsDetails  struct {
    Assets        []string                     `json:"assets,omitempty"`
    ClientWaits   *ZoneStatsDetailsClientWaits `json:"client_waits"`
    Clients       []string                     `json:"clients,omitempty"`
    Id            *uuid.UUID                   `json:"id"`
    MapId         *uuid.UUID                   `json:"map_id"`
    Name          *string                      `json:"name"`
    NumClients    *int                         `json:"num_clients"`
    NumSdkclients *int                         `json:"num_sdkclients"`
    Sdkclients    []string                     `json:"sdkclients,omitempty"`
}

func (z *zoneStatsDetails) validate() error {
    var errs []string
    if z.ClientWaits == nil {
        errs = append(errs, "required field `client_waits` is missing for type `Zone_Stats_Details`")
    }
    if z.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Zone_Stats_Details`")
    }
    if z.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Zone_Stats_Details`")
    }
    if z.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Zone_Stats_Details`")
    }
    if z.NumClients == nil {
        errs = append(errs, "required field `num_clients` is missing for type `Zone_Stats_Details`")
    }
    if z.NumSdkclients == nil {
        errs = append(errs, "required field `num_sdkclients` is missing for type `Zone_Stats_Details`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
