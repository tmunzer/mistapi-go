
# Response Inventory Error Exception

Result of adding device claim codes to organization inventory

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseInventoryErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Added` | `[]string` | Optional | Claim codes accepted into organization inventory |
| `Duplicated` | `[]string` | Optional | Claim codes already present in organization inventory |
| `Error` | `[]string` | Optional | Claim codes rejected by the inventory add operation |
| `InventoryAdded` | [`[]models.ResponseInventoryInventoryAddedItems`](../../doc/models/response-inventory-inventory-added-items.md) | Optional | Detailed inventory records added by the claim operation<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `InventoryDuplicated` | [`[]models.ResponseInventoryInventoryDuplicatedItems`](../../doc/models/response-inventory-inventory-duplicated-items.md) | Optional | Detailed inventory records already present during the claim operation<br><br>**Constraints**: *Unique Items Required* |
| `Reason` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseInventoryErrorException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

