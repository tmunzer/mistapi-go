
# Snmp Config

SNMP configuration for managed network devices

## Structure

`SnmpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientList` | [`[]models.SnmpConfigClientList`](../../doc/models/snmp-config-client-list.md) | Optional | SNMP client allowlists that can be referenced by communities |
| `Contact` | `*string` | Optional | Administrative contact string advertised through SNMP |
| `Description` | `*string` | Optional | Device description string advertised through SNMP |
| `Enabled` | `*bool` | Optional | Whether SNMP is enabled<br><br>**Default**: `true` |
| `EngineId` | `*string` | Optional | SNMP engine ID value used for SNMPv3<br><br>**Constraints**: *Maximum Length*: `27` |
| `EngineIdType` | [`*models.SnmpConfigEngineIdTypeEnum`](../../doc/models/snmp-config-engine-id-type-enum.md) | Optional | Method used to derive the SNMP engine ID. enum: `local`, `use_mac_address`<br><br>**Default**: `"local"` |
| `Location` | `*string` | Optional | Physical location string advertised through SNMP |
| `Name` | `*string` | Optional | System name advertised through SNMP |
| `Network` | `*string` | Optional | Management network used for SNMP traffic<br><br>**Default**: `"default"` |
| `TrapGroups` | [`[]models.SnmpConfigTrapGroup`](../../doc/models/snmp-config-trap-group.md) | Optional | SNMP trap group definitions |
| `V2cConfig` | [`[]models.SnmpConfigV2cConfig`](../../doc/models/snmp-config-v2-c-config.md) | Optional | SNMPv2c community configuration entries for this SNMP profile |
| `V3Config` | [`*models.Snmpv3Config`](../../doc/models/snmpv-3-config.md) | Optional | SNMPv3 notification, target, USM, and VACM configuration |
| `Views` | [`[]models.SnmpConfigView`](../../doc/models/snmp-config-view.md) | Optional | SNMP MIB view definitions |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpConfig := models.SnmpConfig{
        ClientList:           []models.SnmpConfigClientList{
            models.SnmpConfigClientList{
                ClientListName:       models.ToPointer("client_list_name2"),
                Clients:              []string{
                    "clients4",
                },
            },
            models.SnmpConfigClientList{
                ClientListName:       models.ToPointer("client_list_name2"),
                Clients:              []string{
                    "clients4",
                },
            },
        },
        Contact:              models.ToPointer("cns@juniper.net"),
        Description:          models.ToPointer("Juniper QFX Series Switch - 1K_5LA"),
        Enabled:              models.ToPointer(true),
        EngineId:             models.ToPointer("engine_id4"),
        EngineIdType:         models.ToPointer(models.SnmpConfigEngineIdTypeEnum_LOCAL),
        Location:             models.ToPointer("Las Vegas, NV"),
        Name:                 models.ToPointer("TGH-1K-QFX10K"),
        Network:              models.ToPointer("default"),
    }

}
```

