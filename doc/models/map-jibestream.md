
# Map Jibestream

## Structure

`MapJibestream`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `uuid.UUID` | Required | the client id |
| `ClientSecret` | `string` | Required | the client secret |
| `CustomerId` | `int` | Required | the jibestream customer record id |
| `EndpointUrl` | `string` | Required | the map contents endpoint host |
| `MapId` | `uuid.UUID` | Required | the jibestream map id |
| `Mmpp` | `int` | Required | millimeter per pixel |
| `Ppm` | `float64` | Required | pixel per meter, same as the map JSON value. |
| `VendorName` | `string` | Required, Constant | the vendor ‘jibestream’<br>**Default**: `"jibestream"` |
| `VenueId` | `int` | Required | the venue or organization id |

## Example (as JSON)

```json
{
  "client_id": "199d6770-0f6f-407a-9bd5-fc33c7840194",
  "client_secret": "/9Nog3yDzcYj0bY91XJZQLCt+m9DXaIVhx+Ghk3ddd",
  "customer_id": 123,
  "endpoint_url": "https://api.jibestream.com",
  "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "mmpp": 223,
  "ppm": 4.0,
  "vendor_name": "jibestream",
  "venue_id": 123
}
```

