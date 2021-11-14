package ipwhois

type IPDetails struct {
	IP                string  `json:"ip,omitempty"`                 // IP address used for the query
	Success           bool    `json:"success,omitempty"`            // Indicates if the query was successful. true or false
	Message           string  `json:"message,omitempty"`            // Error message if the query was not successful
	Type              string  `json:"type,omitempty"`               // The type of the query. IPv4 or IPv6
	Continent         string  `json:"continent,omitempty"`          // The continent of the IP address. I.e North America
	ContinentCode     string  `json:"continent_code,omitempty"`     // The continent code of the IP address. I.e NA
	Country           string  `json:"country,omitempty"`            // The country of the IP address. I.e United States
	CountryCode       string  `json:"country_code,omitempty"`       // The country code of the IP address. I.e US
	CountryFlag       string  `json:"country_flag,omitempty"`       // URL to an image of the country's flag.
	CountryCapital    string  `json:"country_capital,omitempty"`    // The capital city of the country. I.e Washington
	CountryPhone      string  `json:"country_phone,omitempty"`      // The country's phone code. I.e +1
	CountryNeighbours string  `json:"country_neighbours,omitempty"` // The country's neighbours. I.e CA, MX
	Region            string  `json:"region,omitempty"`             // The region/state of the IP address. I.e Virginia
	City              string  `json:"city,omitempty"`               // The city of the IP address. I.e Ashburn
	Latitude          float64 `json:"latitude,omitempty"`           // The latitude of the IP address. I.e 39.0437567
	Longitude         float64 `json:"longitude,omitempty"`          // The longitude of the IP address. I.e -77.4874416
	Asn               string  `json:"asn,omitempty"`                // The Autonomous System Number of the IP address. I.e AS15169
	Org               string  `json:"org,omitempty"`                // The organization of the IP address. I.e Google, Inc.
	Isp               string  `json:"isp,omitempty"`                // The ISP of the IP address. I.e Comcast Cable Communications, Inc.
	Timezone          string  `json:"timezone,omitempty"`           // The timezone of the IP address's Host Location. I.e America/New_York
	TimezoneName      string  `json:"timezone_name,omitempty"`      // The timezone name of the IP address's Host Location. I.e Eastern Stanard Time
	TimezoneDstOffset float64 `json:"timezone_dstOffset,omitempty"` // The timezone DST offset of the IP address's Host Location. I.e 0
	TimezoneGmtOffset float64 `json:"timezone_gmtOffset,omitempty"` // The timezone GMT offset of the IP address's Host Location. I.e -18000
	TimezoneGmt       string  `json:"timezone_gmt,omitempty"`       // The timezone GMT of the IP address's Host Location. I.e GMT -5:00
	Currency          string  `json:"currency,omitempty"`           // The currency of the IP address's Host Country. I.e US Dollar
	CurrencyCode      string  `json:"currency_code,omitempty"`      // The currency code of the IP address's Host Country. I.e USD
	CurrencySymbol    string  `json:"currency_symbol,omitempty"`    // The currency symbol of the IP address's Host Country. I.e $
	CurrencyRates     float64 `json:"currency_rates,omitempty"`     // The current exchange rate against the US dollar. I.e 1
	CurrencyPlural    string  `json:"currency_plural,omitempty"`    // The currency plural of the IP address. I.e US Dollars
	CompletedRequests int     `json:"completed_requests,omitempty"` // Number of API calls for the current month (Updated every 2 minutes).
}

// Options represents the options for the GetIPDetails request.
type Options struct {
	Objects  []string `url:"objects,comma"` // Objects is a list of objects to get details for. I.e ["country","city","timezone"] this would only return those fields. It is recommended to always add "success" and "message" when filtering so these messages are returned as well
	Language string   `url:"lang"`          // Language localizes city, region, country and continent names. Defaults to English. Other Options: "en", "ru", "de", "es", "pt-br", "fr", "zh-cn", "ja"
}

func (c *Client) GetIPDetails(ip *string, opts *Options) (*IPDetails, error) {
	raw, err := c.get(ip, &opts)
	if err != nil {
		return nil, err
	}

	var data IPDetails
	err = c.unmarshal(&raw, &data)
	if err != nil {
		return nil, err
	}

	return &data, err
}
