
# Inventory Update Operation Enum

enum:

* `upgrade_to_mist`: Upgrade to mist-managed
* `downgrade_to_jsi`: Downgrade to basic monitoring. When downgrading a VC member to jsi, we will move the cloud connection of the VC to jsi-terminator and keep all VC device/inventories intact for pain-free upgrading back to mist.
* `assign`: Assign inventory to a site
* `unassign`: Unassign inventory from a site
* `delete`: Delete multiple inventory from org. If the device is already assigned to a site, it will be unassigned

## Enumeration

`InventoryUpdateOperationEnum`

## Fields

| Name |
|  --- |
| `assign` |
| `delete` |
| `downgrade_to_jsi` |
| `unassign` |
| `upgrade_to_mist` |

