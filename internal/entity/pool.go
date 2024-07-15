package entity

type Pool struct {
	Type      string `json:"type"`
	ProgramID string `json:"programId"`
	ID        string `json:"id"`
	MintA     struct {
		ChainID    int      `json:"chainId"`
		Address    string   `json:"address"`
		ProgramID  string   `json:"programId"`
		LogoURI    string   `json:"logoURI"`
		Symbol     string   `json:"symbol"`
		Name       string   `json:"name"`
		Decimals   int      `json:"decimals"`
		Tags       []string `json:"tags"`
		Extensions struct {
		} `json:"extensions"`
	} `json:"mintA"`
	MintB struct {
		ChainID    int           `json:"chainId"`
		Address    string        `json:"address"`
		ProgramID  string        `json:"programId"`
		LogoURI    string        `json:"logoURI"`
		Symbol     string        `json:"symbol"`
		Name       string        `json:"name"`
		Decimals   int           `json:"decimals"`
		Tags       []interface{} `json:"tags"`
		Extensions struct {
		} `json:"extensions"`
	} `json:"mintB"`
	Price       float64 `json:"price"`
	MintAmountA float64 `json:"mintAmountA"`
	MintAmountB float64 `json:"mintAmountB"`
	FeeRate     float64 `json:"feeRate"`
	OpenTime    string  `json:"openTime"`
	Tvl         float64 `json:"tvl"`
	Day         struct {
		Volume      float64       `json:"volume"`
		VolumeQuote float64       `json:"volumeQuote"`
		VolumeFee   float64       `json:"volumeFee"`
		Apr         float64       `json:"apr"`
		FeeApr      float64       `json:"feeApr"`
		PriceMin    float64       `json:"priceMin"`
		PriceMax    float64       `json:"priceMax"`
		RewardApr   []interface{} `json:"rewardApr"`
	} `json:"day"`
	Week struct {
		Volume      float64       `json:"volume"`
		VolumeQuote float64       `json:"volumeQuote"`
		VolumeFee   float64       `json:"volumeFee"`
		Apr         float64       `json:"apr"`
		FeeApr      float64       `json:"feeApr"`
		PriceMin    float64       `json:"priceMin"`
		PriceMax    float64       `json:"priceMax"`
		RewardApr   []interface{} `json:"rewardApr"`
	} `json:"week"`
	Month struct {
		Volume      float64       `json:"volume"`
		VolumeQuote float64       `json:"volumeQuote"`
		VolumeFee   float64       `json:"volumeFee"`
		Apr         float64       `json:"apr"`
		FeeApr      float64       `json:"feeApr"`
		PriceMin    float64       `json:"priceMin"`
		PriceMax    float64       `json:"priceMax"`
		RewardApr   []interface{} `json:"rewardApr"`
	} `json:"month"`
	Pooltype           []string      `json:"pooltype"`
	RewardDefaultInfos []interface{} `json:"rewardDefaultInfos"`
	FarmUpcomingCount  int           `json:"farmUpcomingCount"`
	FarmOngoingCount   int           `json:"farmOngoingCount"`
	FarmFinishedCount  int           `json:"farmFinishedCount"`
	MarketID           string        `json:"marketId,omitempty"`
	LpMint             struct {
		ChainID    int           `json:"chainId"`
		Address    string        `json:"address"`
		ProgramID  string        `json:"programId"`
		LogoURI    string        `json:"logoURI"`
		Symbol     string        `json:"symbol"`
		Name       string        `json:"name"`
		Decimals   int           `json:"decimals"`
		Tags       []interface{} `json:"tags"`
		Extensions struct {
		} `json:"extensions"`
	} `json:"lpMint,omitempty"`
	LpPrice                float64 `json:"lpPrice,omitempty"`
	LpAmount               float64 `json:"lpAmount,omitempty"`
	RewardDefaultPoolInfos string  `json:"rewardDefaultPoolInfos,omitempty"`
	Config                 struct {
		ID                string    `json:"id"`
		Index             int       `json:"index"`
		ProtocolFeeRate   int       `json:"protocolFeeRate"`
		TradeFeeRate      int       `json:"tradeFeeRate"`
		TickSpacing       int       `json:"tickSpacing"`
		FundFeeRate       int       `json:"fundFeeRate"`
		Description       string    `json:"description"`
		DefaultRange      float64   `json:"defaultRange"`
		DefaultRangePoint []float64 `json:"defaultRangePoint"`
	} `json:"config,omitempty"`
}
