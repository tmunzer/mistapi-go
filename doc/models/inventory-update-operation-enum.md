
# Inventory Update Operation Enum

* if `op`== `upgrade_to_mist`: Upgrade to mist-managed
* if `op`== `downgrade_to_jsi`: Downgrade to basic monitoring. When downgrading a VC member to jsi, we will move the cloud connection of the VC to jsi-terminator and keep all VC device/inventories intact for pain-free upgrading back to mist.
* if `op`== `assign`: Assign inventory to a site
* if `op`== `unassign`: Unassign inventory from a site
* if `op`== `delete`: Delete multiple inventory from org. If the device is already assigned to a site, it will be unassigned.

## Enumeration

`InventoryUpdateOperationEnum`

## Fields

| Name |
|  --- |
| `assign` |
| `unassign` |
| `delete` |
| `upgrade_to_mist` |
| `downgrade_to_jsi` |
