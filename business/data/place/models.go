package place

// Place contains the place data points captured from the API.
type Place struct {
	ID               string   `json:"id,omitempty"`
	Category         string   `json:"category"`
	CityID           CityID   `json:"city"`
	PlaceID          string   `json:"place_id"`
	CityName         string   `json:"city_name"`
	Name             string   `json:"name"`
	Address          string   `json:"address"`
	Lat              float64  `json:"lat"`
	Lng              float64  `json:"lng"`
	LocationType     []string `json:"location_type"`
	AvgUserRating    float32  `json:"avg_user_rating"`
	NumberOfRatings  int      `json:"no_user_rating"`
	GmapsURL         string   `json:"gmaps_url"`
	PhotoReferenceID string   `json:"photo_id"`
}

// CityID is used to capture the city id in relationships.
type CityID struct {
	ID string `json:"id"`
}

type addResult struct {
	AddPlace struct {
		Place []struct {
			ID string `json:"id"`
		} `json:"place"`
	} `json:"addPlace"`
}

func (addResult) document() string {
	return `{
		place {
			id
		}
	}`
}

type updateCityResult struct {
	UpdateCity struct {
		City []struct {
			ID string `json:"id"`
		} `json:"city"`
	} `json:"updateCity"`
}

func (updateCityResult) document() string {
	return `{
		city {
			id
		}
	}`
}
