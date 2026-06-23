
# Wayfinding Import Json

Vendor wayfinding map metadata imported from Jibestream or Micello

## Class Name

`WayfindingImportJson`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.MapJibestream`](../../../doc/models/map-jibestream.md) | models.WayfindingImportJsonContainer.FromMapJibestream(models.MapJibestream mapJibestream) |
| [`models.MapMicello`](../../../doc/models/map-micello.md) | models.WayfindingImportJsonContainer.FromMapMicello(models.MapMicello mapMicello) |

## models.MapJibestream

### Initialization Code

#### Example

```go
value := models.WayfindingImportJsonContainer.FromMapJibestream(models.MapJibestream{
    ClientId:             uuid.MustParse("199d6770-0f6f-407a-9bd5-fc33c7840194"),
    ClientSecret:         "/9Nog3yDzcYj0bY91XJZQLCt+m9DXaIVhx+Ghk3ddd",
    CustomerId:           123,
    EndpointUrl:          "https://api.jibestream.com",
    MapId:                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
    Mmpp:                 223,
    Ppm:                  float64(4),
    VendorName:           "jibestream",
    VenueId:              123,
})
```

## models.MapMicello

### Initialization Code

#### Example

```go
value := models.WayfindingImportJsonContainer.FromMapMicello(models.MapMicello{
    AccountKey:           "",
    DefaultLevelId:       5,
    MapId:                uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"),
    VendorName:           "micello",
})
```

