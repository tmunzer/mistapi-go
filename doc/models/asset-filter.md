
# Asset Filter

BLE asset filter definition; all specified criteria must match

*This model accepts additional fields of type interface{}.*

## Structure

`AssetFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | Access point MAC address that must observe the BLE asset |
| `Beam` | `*int` | Optional | BLE beam number used to filter asset observations |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Disabled` | `*bool` | Optional | Whether the asset filter is disabled<br><br>**Default**: `false` |
| `EddystoneUidNamespace` | `*string` | Optional | Eddystone uid namespace used to filter assets |
| `EddystoneUrl` | `*string` | Optional | Eddystone url used to filter assets |
| `ForSite` | `*bool` | Optional, Read-only | Whether this asset filter is scoped directly to a site |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | iBeacon UUID value, or null when no iBeacon UUID is configured |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `MfgCompanyId` | `*int` | Optional | BLE manufacturing-specific company-id used to filter assets |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MqttTopic` | `*string` | Optional | If set, matching BLE advertisements are forwarded to this MQTT topic when MQTT publishing is enabled |
| `Name` | `string` | Required | Display name for this BLE asset filter |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Rssi` | `*int` | Optional | Signal strength threshold used to filter BLE asset observations |
| `ServiceUuid` | `*uuid.UUID` | Optional | BLE service data uuid used to filter assets |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    assetFilter := models.AssetFilter{
        ApMac:                 models.ToPointer("ap_mac8"),
        Beam:                  models.ToPointer(206),
        CreatedTime:           models.ToPointer(float64(39.96)),
        Disabled:              models.ToPointer(false),
        EddystoneUidNamespace: models.ToPointer("2818e3868dec25629ede"),
        EddystoneUrl:          models.ToPointer("https://www.abc.com"),
        IbeaconMajor:          models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:           models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
        Id:                    models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MfgCompanyId:          models.ToPointer(935),
        Name:                  "Visitor Tags",
        OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        ServiceUuid:           models.ToPointer(uuid.MustParse("0000fe6a-0000-1000-8000-0030459b3cfb")),
        SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties:  map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

