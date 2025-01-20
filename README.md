# Location Managment Api's in Go

## Api Endpoints

```
**POST**

**BaseURL/locations**

RequestBody : {
  "name": string,
  "address": string,
  "latitude": float,
  "lonitude": float,
  "category": string
}
```

```
**GET**

**BaseURL/locations/{category}**
```

```
**POST**

**BaseURL/search**

RequestBody : {
  "latitude": float,
  "lonitude": float,
  "category": string,
  "radius": int
}
```

```
**POST**

**BaseURL/trip-cost/{location_id}**

RequestBody : {
  "latitude": float,
  "lonitude": float
}
```
