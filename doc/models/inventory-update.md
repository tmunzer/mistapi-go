
# Inventory Update

## Structure

`InventoryUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableAutoConfig` | `*bool` | Optional | if `op`==`assign`, a **cloud-ready** switch/gateway will be managed/configured by Mist by default, this disabled the behavior<br>**Default**: `false` |
| `Macs` | `[]string` | Optional | if `op`==`assign`, `op`==`unassign`, `op`==`upgrade_to_mist`or `op`==`downgrade_to_jsi` , list of MAC, e.g. ["5c5b350e0001"] |
| `Managed` | `*bool` | Optional | if `op`==`assign`, an **adopted** switch/gateway will not be managed/configured by Mist by default, this enables the behavior<br>**Default**: `false` |
| `NoReassign` | `*bool` | Optional | if `op`==`assign`, if true, treat site assignment against an already assigned AP as error |
| `Op` | [`models.InventoryUpdateOperationEnum`](../../doc/models/inventory-update-operation-enum.md) | Required | enum:<br><br>* `upgrade_to_mist`: Upgrade to mist-managed<br>* `downgrade_to_jsi`: Downgrade to basic monitoring. When downgrading a VC member to jsi, we will move the cloud connection of the VC to jsi-terminator and keep all VC device/inventories intact for pain-free upgrading back to mist.<br>* `assign`: Assign inventory to a site<br>* `unassign`: Unassign inventory from a site<br>* `delete`: Delete multiple inventory from org. If the device is already assigned to a site, it will be unassigned |
| `Serials` | `[]string` | Optional | if `op`==`delete`, list of serial numbers, e.g. ["FXLH2015150025"] |
| `SiteId` | `*uuid.UUID` | Optional | if `op`==`assign`, target site id |

## Example (as JSON)

```json
{
  "disable_auto_config": false,
  "managed": false,
  "op": "unassign",
  "macs": [
    "macs9"
  ],
  "no_reassign": false,
  "serials": [
    "serials0",
    "serials1"
  ]
}
```

