# seoul-impact-hackathon

## Server

Build the server:

```sh
cd ./server
go build
```

Run the server:

```sh
./youth-korea
```

### API endpoints

The server is started on port `:8000`

- `/shelters`: Returns all shelters irrespective of status
  A shelter adheres to the following data model:

```go
type Shelter struct {
	Id            int           `json:"id"`
	OpeningTime   string        `json:"openingTime"`
	Facilities    *[]string     `json:"facilities,omitempty"`
	Requirements  *string       `json:"requirements,omitempty"`
	Rules         string        `json:"rules"`
	MaxCapacity   int           `json:"maxCapacity"`
	Capacity      int           `json:"capacity"`
	Rating        *int          `json:"rating,omitempty"`
	Desc          *string       `json:"desc,omitempty"`
	Address       string        `json:"address"`
	Name          string        `json:"name"`
	CurrentStatus ShelterStatus `json:"currentStatus"`
	Announcement  *string       `json:"announcement,omitempty"`
	Contact       *string       `json:"contact,omitempty"`
	Website       *string       `json:"website,omitempty"`
}

The endpoint accepts the following query parameters for filtering the returned shelters:

    - `hasCapacity`: boolean
    - `status`: string
```

- `shelter/:id`: Returns a single shelter with the passed ID. IDs should be numeric. If no shelters are found, the API
  will return a `404`. In case of non-numeric IDs, the error code will be `400`.

- `/food`: Returns all facilities irrespective of status
  A facility adheres to the following data model:

```go
type Facility struct {
	Id            int              `json:"id"`
	OpeningTime   string           `json:"openingTime"`
	Rating        *int             `json:"rating,omitempty"`
	Desc          *string          `json:"desc,omitempty"`
	Address       string           `json:"address"`
	Name          string           `json:"name"`
	CurrentStatus RestaurantStatus `json:"currentStatus"`
	Announcement  *string          `json:"announcement,omitempty"`
	Contact       *string          `json:"contact,omitempty"`
	Website       *string          `json:"website,omitempty"`
	FoodTypes     []FoodOptions    `json:"foodTypes"`
}
```

The endpoint accepts the following query parameters for filtering the returned facilities:

    - `status`: string
    - `foodOptions`: comma-separated list

- `food/:id`: Returns a single facility with the passed ID. IDs should be numeric. If no facilities are found, the API
  will return a `404`. In case of non-numeric IDs, the error code will be `400`.
