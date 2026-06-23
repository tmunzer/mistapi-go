
# Response Client Nac Search

Paginated NAC client search response

## Structure

`ResponseClientNacSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Search window end timestamp for NAC clients, in epoch seconds |
| `Limit` | `*int` | Optional | Maximum number of NAC client results requested |
| `Next` | `*string` | Optional | URL for the next page of NAC client results |
| `Results` | [`[]models.ClientNac`](../../doc/models/client-nac.md) | Optional | NAC client records returned by a search response |
| `Start` | `*int` | Optional | Search window start timestamp for NAC clients, in epoch seconds |
| `Total` | `*int` | Optional | Number of NAC client records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseClientNacSearch := models.ResponseClientNacSearch{
        End:                  models.ToPointer(1513362753),
        Limit:                models.ToPointer(3),
        Next:                 models.ToPointer("next4"),
        Results:              []models.ClientNac{
            models.ClientNac{
                Ap:                   []string{
                    "ap6",
                    "ap7",
                    "ap8",
                },
                AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTEAP),
                CertCn:               []string{
                    "cert_cn9",
                    "cert_cn8",
                },
                CertIssuer:           []string{
                    "cert_issuer2",
                },
                CertSerial:           []string{
                    "cert_serial8",
                },
            },
            models.ClientNac{
                Ap:                   []string{
                    "ap6",
                    "ap7",
                    "ap8",
                },
                AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTEAP),
                CertCn:               []string{
                    "cert_cn9",
                    "cert_cn8",
                },
                CertIssuer:           []string{
                    "cert_issuer2",
                },
                CertSerial:           []string{
                    "cert_serial8",
                },
            },
        },
        Start:                models.ToPointer(1513276353),
        Total:                models.ToPointer(2),
    }

}
```

