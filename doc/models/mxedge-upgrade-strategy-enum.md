
# Mxedge Upgrade Strategy Enum

enum:

* `big_bang`: upgrade all at once, no orchestration
* `serial`: one at a time'
* `canary`: upgrade in phases

## Enumeration

`MxedgeUpgradeStrategyEnum`

## Fields

| Name |
|  --- |
| `canary` |
| `big_bang` |
| `serial` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeUpgradeStrategy := models.MxedgeUpgradeStrategyEnum_CANARY

}
```

