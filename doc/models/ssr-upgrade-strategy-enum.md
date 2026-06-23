
# Ssr Upgrade Strategy Enum

enum:

* `big_bang`: upgrade all at once
* `serial`: one at a time

## Enumeration

`SsrUpgradeStrategyEnum`

## Fields

| Name |
|  --- |
| `big_bang` |
| `serial` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssrUpgradeStrategy := models.SsrUpgradeStrategyEnum_BIGBANG

}
```

