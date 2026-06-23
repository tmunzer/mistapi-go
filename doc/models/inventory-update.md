
# Inventory Update

Request to assign, unassign, delete, or change management mode for inventory devices

*This model accepts additional fields of type interface{}.*

## Structure

`InventoryUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableAutoConfig` | `*bool` | Optional | If `op`==`assign`, this disables the default behavior of a cloud-ready switch/gateway being managed/configured by Mist. Setting this to `true` means you want to disable the default behavior and do not want the device to be Mist-managed.<br><br>**Default**: `false` |
| `Macs` | `[]string` | Optional | If `op`==`assign`, `op`==`unassign`, `op`==`upgrade_to_mist`, or `op`==`downgrade_to_jsi`, device MAC addresses to update |
| `Managed` | `*bool` | Optional | If `op`==`assign`. An adopted switch/gateway will not be managed/configured by Mist by default. Setting this parameter to `true` enables the adopted switch/gateway to be managed/configured by Mist.<br><br>**Default**: `false` |
| `MistConfigured` | `*bool` | Optional | Whether the device can be configured by Mist. Replaces `managed` for adopted devices and `disable_auto_config` for claimed devices |
| `NoReassign` | `*bool` | Optional | If `op`==`assign`, if true, treat site assignment against an already assigned AP as error<br><br>**Default**: `false` |
| `Op` | [`models.InventoryUpdateOperationEnum`](../../doc/models/inventory-update-operation-enum.md) | Required | enum:<br><br>* `upgrade_to_mist`: Upgrade to mist-managed<br>* `downgrade_to_jsi`: Downgrade to basic monitoring. When downgrading a VC member to jsi, we will move the cloud connection of the VC to jsi-terminator and keep all VC device/inventories intact for pain-free upgrading back to mist.<br>* `assign`: Assign inventory to a site<br>* `unassign`: Unassign inventory from a site<br>* `delete`: Delete multiple inventory from org. If the device is already assigned to a site, it will be unassigned |
| `Serials` | `[]string` | Optional | If `op`==`delete`, device serial numbers to delete from inventory |
| `SiteId` | `*uuid.UUID` | Optional | If `op`==`assign`, target site ID for the inventory operation |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    inventoryUpdate := models.InventoryUpdate{
        DisableAutoConfig:    models.ToPointer(false),
        Macs:                 []string{
            "macs3",
            "macs2",
            "macs1",
        },
        Managed:              models.ToPointer(false),
        MistConfigured:       models.ToPointer(false),
        NoReassign:           models.ToPointer(false),
        Op:                   models.InventoryUpdateOperationEnum_DELETE,
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

