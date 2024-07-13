package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// RfDiagInfoItem represents a RfDiagInfoItem struct.
type RfDiagInfoItem struct {
    // if `type`==`asset`, id of the asset
    AssetId              *uuid.UUID       `json:"asset_id,omitempty"`
    // if `type`==`asset`, name of the asset
    AssetName            *string          `json:"asset_name,omitempty"`
    // if `type`==`client`, hostname of the client
    ClientName           *string          `json:"client_name,omitempty"`
    // recording length in seconds, max is 120
    Duration             int              `json:"duration"`
    // timestamp of end of recording
    EndTime              int              `json:"end_time"`
    // Number of frames in the output
    FrameCount           int              `json:"frame_count"`
    Id                   *string          `json:"id,omitempty"`
    // if `type`==`client` or `asset`, mac of the device
    Mac                  *string          `json:"mac,omitempty"`
    MapId                uuid.UUID        `json:"map_id"`
    Name                 string           `json:"name"`
    // Optional. id of the next recoding if present. Only valid for site survey.
    Next                 *string          `json:"next,omitempty"`
    // URL to a JSON file that contains array of raw location diag events
    RawEvents            string           `json:"raw_events"`
    // whether it’s ready for playback
    Ready                bool             `json:"ready"`
    // if `type`==`sdkclient`, sdkclient_id of this recording
    SdkclientId          *uuid.UUID       `json:"sdkclient_id,omitempty"`
    // if `type`==`sdkclient`, name of the sdkclient
    SdkclientName        *string          `json:"sdkclient_name,omitempty"`
    // if `type`==`sdkclient`, device_id of sdkclient
    SdkclientUuid        *uuid.UUID       `json:"sdkclient_uuid,omitempty"`
    // timestamp of the recording (the start)
    StartTime            int              `json:"start_time"`
    Type                 RfClientTypeEnum `json:"type"`
    // URL to a JSON file that contains an array of frames, each frame is the same format
    Url                  string           `json:"url"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RfDiagInfoItem.
// It customizes the JSON marshaling process for RfDiagInfoItem objects.
func (r RfDiagInfoItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RfDiagInfoItem object to a map representation for JSON marshaling.
func (r RfDiagInfoItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.AssetId != nil {
        structMap["asset_id"] = r.AssetId
    }
    if r.AssetName != nil {
        structMap["asset_name"] = r.AssetName
    }
    if r.ClientName != nil {
        structMap["client_name"] = r.ClientName
    }
    structMap["duration"] = r.Duration
    structMap["end_time"] = r.EndTime
    structMap["frame_count"] = r.FrameCount
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.Mac != nil {
        structMap["mac"] = r.Mac
    }
    structMap["map_id"] = r.MapId
    structMap["name"] = r.Name
    if r.Next != nil {
        structMap["next"] = r.Next
    }
    structMap["raw_events"] = r.RawEvents
    structMap["ready"] = r.Ready
    if r.SdkclientId != nil {
        structMap["sdkclient_id"] = r.SdkclientId
    }
    if r.SdkclientName != nil {
        structMap["sdkclient_name"] = r.SdkclientName
    }
    if r.SdkclientUuid != nil {
        structMap["sdkclient_uuid"] = r.SdkclientUuid
    }
    structMap["start_time"] = r.StartTime
    structMap["type"] = r.Type
    structMap["url"] = r.Url
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RfDiagInfoItem.
// It customizes the JSON unmarshaling process for RfDiagInfoItem objects.
func (r *RfDiagInfoItem) UnmarshalJSON(input []byte) error {
    var temp rfDiagInfoItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "asset_id", "asset_name", "client_name", "duration", "end_time", "frame_count", "id", "mac", "map_id", "name", "next", "raw_events", "ready", "sdkclient_id", "sdkclient_name", "sdkclient_uuid", "start_time", "type", "url")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.AssetId = temp.AssetId
    r.AssetName = temp.AssetName
    r.ClientName = temp.ClientName
    r.Duration = *temp.Duration
    r.EndTime = *temp.EndTime
    r.FrameCount = *temp.FrameCount
    r.Id = temp.Id
    r.Mac = temp.Mac
    r.MapId = *temp.MapId
    r.Name = *temp.Name
    r.Next = temp.Next
    r.RawEvents = *temp.RawEvents
    r.Ready = *temp.Ready
    r.SdkclientId = temp.SdkclientId
    r.SdkclientName = temp.SdkclientName
    r.SdkclientUuid = temp.SdkclientUuid
    r.StartTime = *temp.StartTime
    r.Type = *temp.Type
    r.Url = *temp.Url
    return nil
}

// rfDiagInfoItem is a temporary struct used for validating the fields of RfDiagInfoItem.
type rfDiagInfoItem  struct {
    AssetId       *uuid.UUID        `json:"asset_id,omitempty"`
    AssetName     *string           `json:"asset_name,omitempty"`
    ClientName    *string           `json:"client_name,omitempty"`
    Duration      *int              `json:"duration"`
    EndTime       *int              `json:"end_time"`
    FrameCount    *int              `json:"frame_count"`
    Id            *string           `json:"id,omitempty"`
    Mac           *string           `json:"mac,omitempty"`
    MapId         *uuid.UUID        `json:"map_id"`
    Name          *string           `json:"name"`
    Next          *string           `json:"next,omitempty"`
    RawEvents     *string           `json:"raw_events"`
    Ready         *bool             `json:"ready"`
    SdkclientId   *uuid.UUID        `json:"sdkclient_id,omitempty"`
    SdkclientName *string           `json:"sdkclient_name,omitempty"`
    SdkclientUuid *uuid.UUID        `json:"sdkclient_uuid,omitempty"`
    StartTime     *int              `json:"start_time"`
    Type          *RfClientTypeEnum `json:"type"`
    Url           *string           `json:"url"`
}

func (r *rfDiagInfoItem) validate() error {
    var errs []string
    if r.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.EndTime == nil {
        errs = append(errs, "required field `end_time` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.FrameCount == nil {
        errs = append(errs, "required field `frame_count` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.RawEvents == nil {
        errs = append(errs, "required field `raw_events` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.Ready == nil {
        errs = append(errs, "required field `ready` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.StartTime == nil {
        errs = append(errs, "required field `start_time` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Rf_Diag_Info_Item`")
    }
    if r.Url == nil {
        errs = append(errs, "required field `url` is missing for type `Rf_Diag_Info_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}