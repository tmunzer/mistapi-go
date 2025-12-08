
# Map Jibestream

## Structure

`MapJibestream`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `uuid.UUID` | Required | Client id |
| `ClientSecret` | `string` | Required | Client secret |
| `CustomerId` | `int` | Required | Jibestream customer record id |
| `EndpointUrl` | `string` | Required | Map contents endpoint host |
| `MapId` | `uuid.UUID` | Required | Jibestream map id |
| `Mmpp` | `int` | Required | Millimeter per pixel |
| `Ppm` | `float64` | Required | Pixel per meter, same as the map JSON value. |
| `VendorName` | `string` | Required, Constant | The vendor ‘jibestream’. enum: `jibestream`<br><br>**Value**: `"jibestream"` |
| `VenueId` | `int` | Required | Venue or organization id |

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

