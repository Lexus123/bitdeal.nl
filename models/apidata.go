package models

/*
Exchange ...
*/
type Exchange struct {
	Name, URL, Method, Type, Currency, Endpoint string
}

/*
ExchangeInfo ...
*/
type ExchangeInfo struct {
	Name, Link, TrustLink string
	Reviews, Stars        int
	Broker                bool
}

/*
GetPricesData ...
*/
type GetPricesData struct {
	Type     string `json:"type"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

/*
GetPricesResponse ...
*/
type GetPricesResponse struct {
	Type          string         `json:"type"`
	Currency      string         `json:"currency"`
	BestRate      float64        `json:"bestrate"`
	BestAmount    float64        `json:"bestamount"`
	MostReviews   int            `json:"mostreviews"`
	ExchangeRates []ExchangeRate `json:"exchangerates"`
}

/*
ExchangeRate ...
*/
type ExchangeRate struct {
	Exchange string  `json:"exchange"`
	Rate     float64 `json:"rate"`
	Amount   float64 `json:"amount"`
	Link     string  `json:"link"`
	Reviews  int     `json:"reviews"`
	Stars    int     `json:"stars"`
	Status   Status  `json:"status"`
	Broker   bool    `json:"broker"`
}

/*
Status ...
*/
type Status int

const (
	// OK means the broker/exchange replied with a status 200
	OK Status = 200

	// RequestTimeout means bitdeal cutoff the request (after 500ms)
	RequestTimeout Status = 408

	// GeneralError means something went wrong
	GeneralError Status = 777
)

/*
GetPricesError ...
*/
type GetPricesError struct {
	Exchange string `json:"exchange"`
	Error    Status `json:"error"`
}

/*
Bitvavo ... GET https://api.bitvavo.com/v2/BTC-EUR/book
*/
type Bitvavo struct {
	Market string     `json:"market"`
	Nonce  int        `json:"nonce"`
	Bids   [][]string `json:"bids"`
	Asks   [][]string `json:"asks"`
}

/*
Bitmymoney ... GET https://www.bitmymoney.com/market/
*/
type Bitmymoney struct {
	Markets struct {
		BTC struct {
			DailyChange string `json:"daily_change"`
			BuyPrice    string `json:"buy_price"`
			SellPrice   string `json:"sell_price"`
			MidPrice    string `json:"mid_price"`
			Balance     struct {
				BTC struct {
					Highest string `json:"highest"`
					Total   string `json:"total"`
				} `json:"BTC"`
				EUR struct {
					Highest string `json:"highest"`
					Total   string `json:"total"`
				} `json:"EUR"`
			} `json:"balance"`
		} `json:"BTC"`
	} `json:"markets"`
	CurrencyInfo struct {
		BTC struct {
			IconAlt         string `json:"icon-alt"`
			Name            string `json:"name"`
			PriceDataMerged string `json:"price_data_merged"`
			Sign            string `json:"sign"`
			PriceData       string `json:"price_data"`
			DisplayDecimals int    `json:"display_decimals"`
			Decimals        int    `json:"decimals"`
			EventData       string `json:"event_data"`
			Icon            string `json:"icon"`
		} `json:"BTC"`
	} `json:"currency_info"`
	Plugins struct {
		Sell struct {
			Sepa struct {
				Fee    string `json:"fee"`
				Margin string `json:"margin"`
				Name   string `json:"name"`
				View   string `json:"view"`
			} `json:"sepa"`
		} `json:"sell"`
		Buy struct {
			Ideal struct {
				Fee    string `json:"fee"`
				Margin string `json:"margin"`
				Name   string `json:"name"`
				View   string `json:"view"`
			} `json:"ideal"`
			Bankdeposit struct {
				Fee    string `json:"fee"`
				Margin string `json:"margin"`
				Name   string `json:"name"`
			} `json:"bankdeposit"`
			Mrcash struct {
				Fee                  string `json:"fee"`
				Name                 string `json:"name"`
				MaxVerifiedAmount    int    `json:"max_verified_amount"`
				MaxFirstNoHoldAmount int    `json:"max_first_no_hold_amount"`
				Margin               string `json:"margin"`
				MinAmount            string `json:"min_amount"`
				MaxUnverifiedAmount  int    `json:"max_unverified_amount"`
				View                 string `json:"view"`
			} `json:"mrcash"`
		} `json:"buy"`
		Deposit struct {
			Bankdeposit struct {
				RequiresDepositAddress bool   `json:"requires_deposit_address"`
				Fee                    string `json:"fee"`
				Name                   string `json:"name"`
				MaxVerifiedAmount      int    `json:"max_verified_amount"`
				MaxFirstNoHoldAmount   int    `json:"max_first_no_hold_amount"`
				Margin                 string `json:"margin"`
				MaxUnverifiedAmount    int    `json:"max_unverified_amount"`
				View                   string `json:"view"`
			} `json:"bankdeposit"`
			Deposit struct {
				Fee                    string `json:"fee"`
				Margin                 string `json:"margin"`
				Name                   string `json:"name"`
				RequiresDepositAddress bool   `json:"requires_deposit_address"`
				View                   string `json:"view"`
			} `json:"deposit"`
			Payment struct {
				Fee                    string `json:"fee"`
				Name                   string `json:"name"`
				RestOnly               bool   `json:"rest_only"`
				RequiresDepositAddress bool   `json:"requires_deposit_address"`
				Margin                 string `json:"margin"`
				View                   string `json:"view"`
			} `json:"payment"`
		} `json:"deposit"`
		Send struct {
			Account struct {
				Fee                    string `json:"fee"`
				Margin                 string `json:"margin"`
				Name                   string `json:"name"`
				RequiresDepositAddress bool   `json:"requires_deposit_address"`
				View                   string `json:"view"`
			} `json:"account"`
			Address struct {
				Fee                    string `json:"fee"`
				Margin                 string `json:"margin"`
				Name                   string `json:"name"`
				RequiresDepositAddress bool   `json:"requires_deposit_address"`
				View                   string `json:"view"`
			} `json:"address"`
		} `json:"send"`
	} `json:"plugins"`
}

/*
Litebit ... GET https://api.litebit.eu/market/btc
*/
type Litebit struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Result  struct {
		Name      string `json:"name"`
		Abbr      string `json:"abbr"`
		Available string `json:"available"`
		Volume    string `json:"volume"`
		Buy       string `json:"buy"`
		Sell      string `json:"sell"`
	} `json:"result"`
}

/*
Btcdirect ... GET https://my.btcdirect.eu/config
*/
type Btcdirect struct {
	SellMinEur                int     `json:"sellMinEur"`
	SellMaxEur                int     `json:"sellMaxEur"`
	SellMinEurLtc             int     `json:"sellMinEurLtc"`
	SellMaxEurLtc             int     `json:"sellMaxEurLtc"`
	SellMinEurEth             int     `json:"sellMinEurEth"`
	SellMaxEurEth             int     `json:"sellMaxEurEth"`
	SellMinEurBch             int     `json:"sellMinEurBch"`
	SellMaxEurBch             int     `json:"sellMaxEurBch"`
	SellMinEurXrp             int     `json:"sellMinEurXrp"`
	SellMaxEurXrp             int     `json:"sellMaxEurXrp"`
	BuyMinBtc                 float64 `json:"buyMinBtc"`
	BuyMaxBtc                 int     `json:"buyMaxBtc"`
	SellEnabled               int     `json:"sellEnabled"`
	SellEnabledGeneral        int     `json:"sellEnabledGeneral"`
	BuyEnabled                int     `json:"buyEnabled"`
	BuyEnabledGeneral         int     `json:"buyEnabledGeneral"`
	ShopEnabled               int     `json:"shopEnabled"`
	SmsThreshold              int     `json:"smsThreshold"`
	IdealTransactionCostEur   float64 `json:"idealTransactionCostEur"`
	MaxSofortOpenAmount       int     `json:"maxSofortOpenAmount"`
	MaxGiropayOpenAmount      int     `json:"maxGiropayOpenAmount"`
	APICommunicationErrorBuy  string  `json:"apiCommunicationErrorBuy"`
	APICommunicationErrorSell string  `json:"apiCommunicationErrorSell"`
	ExchangeRates             struct {
		Sell struct {
			Btc string `json:"btc"`
			Ltc string `json:"ltc"`
			Eth string `json:"eth"`
			Bch string `json:"bch"`
			Xrp string `json:"xrp"`
		} `json:"sell"`
		Buy struct {
			Btc string `json:"btc"`
			Ltc string `json:"ltc"`
			Eth string `json:"eth"`
			Bch string `json:"bch"`
			Xrp string `json:"xrp"`
		} `json:"buy"`
	} `json:"exchangeRates"`
}

/*
KnakenBuy ...
*/
type KnakenBuy struct {
	Ok    bool   `json:"ok"`
	Price string `json:"price"`
}

/*
KnakenSell ...
*/
type KnakenSell struct {
	BTC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BTC"`
}

/*
BitonicBuy ...
*/
type BitonicBuy struct {
	Success bool    `json:"success"`
	Btc     float64 `json:"btc"`
	Eur     float64 `json:"eur"`
	Price   float64 `json:"price"`
	Method  string  `json:"method"`
}

/*
BitonicSell ...
*/
type BitonicSell struct {
	Success bool    `json:"success"`
	Btc     float64 `json:"btc"`
	Eur     float64 `json:"eur"`
	Price   float64 `json:"price"`
}

/*
BitrushBuy ...
*/
type BitrushBuy struct {
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	FiatQuantity float64 `json:"fiat_quantity"`
	CoinQuantity float64 `json:"coin_quantity"`
}

/*
BitrushSell ...
*/
type BitrushSell struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

/*
Coinmerce ...
*/
type Coinmerce struct {
	Status int `json:"status"`
	Prices struct {
		BTC struct {
			Buy  string `json:"buy"`
			Sell string `json:"sell"`
		} `json:"BTC"`
	} `json:"prices"`
	Version int `json:"version"`
}

/*
Bitqist ...
*/
type Bitqist struct {
	BaseCurrency  int         `json:"base_currency"`
	QuoteCurrency int         `json:"quote_currency"`
	Time          interface{} `json:"time"`
	VolumeBaseNow interface{} `json:"volume_base_now"`
	AskNow        float64     `json:"ask_now"`
	AskNowInclFee float64     `json:"ask_now_incl_fee"`
	BidNow        float64     `json:"bid_now"`
	BidNowInclFee float64     `json:"bid_now_incl_fee"`
	Ask1H         interface{} `json:"ask_1h"`
	Ask1D         interface{} `json:"ask_1d"`
	Ask7D         interface{} `json:"ask_7d"`
}

/*
Satos ...
*/
type Satos struct {
	Pair      string `json:"pair"`
	Side      string `json:"side"`
	Value     string `json:"value"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
	Timestamp int    `json:"timestamp"`
}
