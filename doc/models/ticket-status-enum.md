
# Ticket Status Enum

Ticket status. enum: `closed`, `open`, `pending`, `solved`. `open` means Mist is working on it, `pending` means requester attention is needed, `solved` means Mist considers it resolved but it can still be updated or rated, and `closed` means it is archived

## Enumeration

`TicketStatusEnum`

## Fields

| Name |
|  --- |
| `closed` |
| `open` |
| `pending` |
| `solved` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ticketStatus := models.TicketStatusEnum_CLOSED

}
```

