
# Jsi Inventory Count Distinct Enum

Distinct field used when counting Juniper Support Insight inventory records. enum: `account_id`, `claimed`, `has_support`, `end_of_sale_time`, `eos_time`, `version_time`, `model`, `sku`, `status`, `type`, `version`, `warranty_type`

## Enumeration

`JsiInventoryCountDistinctEnum`

## Fields

| Name |
|  --- |
| `account_id` |
| `claimed` |
| `has_support` |
| `end_of_sale_time` |
| `eos_time` |
| `version_time` |
| `model` |
| `sku` |
| `status` |
| `type` |
| `version` |
| `warranty_type` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    jsiInventoryCountDistinct := models.JsiInventoryCountDistinctEnum_VERSION

}
```

