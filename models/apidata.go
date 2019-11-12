package models

/*
Exchange ...
*/
type Exchange struct {
	Name, URL, Method, Type, Currency, Endpoint string
}

/*
GetPricesData ...
*/
type GetPricesData struct {
	Type     string  `json:"type"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

/*
GetPricesResponse ...
*/
type GetPricesResponse struct {
	Type          string         `json:"type"`
	Currency      string         `json:"currency"`
	BestRate      float64        `json:"bestrate"`
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
Satos ... POST https://api.satos.nl/v1/public/getDepthRate
*/
type Satos struct {
	Data []struct {
		Btc struct {
			Rateid      int     `json:"rateid"`
			AskPrice    int     `json:"ask_price"`
			AskVolume   float64 `json:"ask_volume"`
			AskCoin     float64 `json:"ask_coin"`
			AskEurBank  float64 `json:"ask_eur_bank"`
			AskEurCash  float64 `json:"ask_eur_cash"`
			AskRateBank float64 `json:"ask_rate_bank"`
			AskRateCash float64 `json:"ask_rate_cash"`
			BidPrice    int     `json:"bid_price"`
			BidVolume   float64 `json:"bid_volume"`
			BidCoin     float64 `json:"bid_coin"`
			BidEurBank  float64 `json:"bid_eur_bank"`
			BidEurCash  float64 `json:"bid_eur_cash"`
			BidRateBank float64 `json:"bid_rate_bank"`
			BidRateCash float64 `json:"bid_rate_cash"`
			UpdatedAt   int     `json:"updated_at"`
		} `json:"btc"`
		Eth struct {
			Rateid      int     `json:"rateid"`
			AskPrice    float64 `json:"ask_price"`
			AskVolume   float64 `json:"ask_volume"`
			AskCoin     float64 `json:"ask_coin"`
			AskEurBank  float64 `json:"ask_eur_bank"`
			AskEurCash  float64 `json:"ask_eur_cash"`
			AskRateBank float64 `json:"ask_rate_bank"`
			AskRateCash float64 `json:"ask_rate_cash"`
			BidPrice    int     `json:"bid_price"`
			BidVolume   float64 `json:"bid_volume"`
			BidCoin     float64 `json:"bid_coin"`
			BidEurBank  float64 `json:"bid_eur_bank"`
			BidEurCash  float64 `json:"bid_eur_cash"`
			BidRateBank float64 `json:"bid_rate_bank"`
			BidRateCash float64 `json:"bid_rate_cash"`
			UpdatedAt   int     `json:"updated_at"`
		} `json:"eth"`
		Xrp struct {
			Rateid      int     `json:"rateid"`
			AskPrice    float64 `json:"ask_price"`
			AskVolume   float64 `json:"ask_volume"`
			AskCoin     float64 `json:"ask_coin"`
			AskEurBank  float64 `json:"ask_eur_bank"`
			AskEurCash  float64 `json:"ask_eur_cash"`
			AskRateBank float64 `json:"ask_rate_bank"`
			AskRateCash float64 `json:"ask_rate_cash"`
			BidPrice    float64 `json:"bid_price"`
			BidVolume   float64 `json:"bid_volume"`
			BidCoin     float64 `json:"bid_coin"`
			BidEurBank  float64 `json:"bid_eur_bank"`
			BidEurCash  float64 `json:"bid_eur_cash"`
			BidRateBank float64 `json:"bid_rate_bank"`
			BidRateCash float64 `json:"bid_rate_cash"`
			UpdatedAt   int     `json:"updated_at"`
		} `json:"xrp"`
		Bch struct {
			Rateid      int     `json:"rateid"`
			AskPrice    float64 `json:"ask_price"`
			AskVolume   float64 `json:"ask_volume"`
			AskCoin     float64 `json:"ask_coin"`
			AskEurBank  float64 `json:"ask_eur_bank"`
			AskEurCash  float64 `json:"ask_eur_cash"`
			AskRateBank float64 `json:"ask_rate_bank"`
			AskRateCash float64 `json:"ask_rate_cash"`
			BidPrice    float64 `json:"bid_price"`
			BidVolume   float64 `json:"bid_volume"`
			BidCoin     float64 `json:"bid_coin"`
			BidEurBank  float64 `json:"bid_eur_bank"`
			BidEurCash  float64 `json:"bid_eur_cash"`
			BidRateBank float64 `json:"bid_rate_bank"`
			BidRateCash float64 `json:"bid_rate_cash"`
			UpdatedAt   int     `json:"updated_at"`
		} `json:"bch"`
		Ltc struct {
			Rateid      int     `json:"rateid"`
			AskPrice    float64 `json:"ask_price"`
			AskVolume   float64 `json:"ask_volume"`
			AskCoin     float64 `json:"ask_coin"`
			AskEurBank  float64 `json:"ask_eur_bank"`
			AskEurCash  float64 `json:"ask_eur_cash"`
			AskRateBank float64 `json:"ask_rate_bank"`
			AskRateCash float64 `json:"ask_rate_cash"`
			BidPrice    float64 `json:"bid_price"`
			BidVolume   float64 `json:"bid_volume"`
			BidCoin     float64 `json:"bid_coin"`
			BidEurBank  float64 `json:"bid_eur_bank"`
			BidEurCash  float64 `json:"bid_eur_cash"`
			BidRateBank float64 `json:"bid_rate_bank"`
			BidRateCash float64 `json:"bid_rate_cash"`
			UpdatedAt   int     `json:"updated_at"`
		} `json:"ltc"`
	} `json:"data"`
}

/*
Bitmymoney ... GET https://www.bitmymoney.com/market/
*/
type Bitmymoney struct {
	Markets struct {
		LTC struct {
			DailyChange string `json:"daily_change"`
			BuyPrice    string `json:"buy_price"`
			SellPrice   string `json:"sell_price"`
			MidPrice    string `json:"mid_price"`
			Balance     struct {
				LTC struct {
					Highest string `json:"highest"`
					Total   string `json:"total"`
				} `json:"LTC"`
				EUR struct {
					Highest string `json:"highest"`
					Total   string `json:"total"`
				} `json:"EUR"`
			} `json:"balance"`
		} `json:"LTC"`
		ETH struct {
			DailyChange string `json:"daily_change"`
			BuyPrice    string `json:"buy_price"`
			SellPrice   string `json:"sell_price"`
			MidPrice    string `json:"mid_price"`
			Balance     struct {
				ETH struct {
					Highest string `json:"highest"`
					Total   string `json:"total"`
				} `json:"ETH"`
				EUR struct {
					Highest string `json:"highest"`
					Total   string `json:"total"`
				} `json:"EUR"`
			} `json:"balance"`
		} `json:"ETH"`
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
		LTC struct {
			IconAlt         string `json:"icon-alt"`
			Name            string `json:"name"`
			PriceDataMerged string `json:"price_data_merged"`
			Sign            string `json:"sign"`
			PriceData       string `json:"price_data"`
			DisplayDecimals int    `json:"display_decimals"`
			Decimals        int    `json:"decimals"`
			EventData       string `json:"event_data"`
			Icon            string `json:"icon"`
		} `json:"LTC"`
		ETH struct {
			IconAlt         string `json:"icon-alt"`
			Name            string `json:"name"`
			PriceDataMerged string `json:"price_data_merged"`
			Sign            string `json:"sign"`
			PriceData       string `json:"price_data"`
			DisplayDecimals int    `json:"display_decimals"`
			Decimals        int    `json:"decimals"`
			EventData       string `json:"event_data"`
			Icon            string `json:"icon"`
		} `json:"ETH"`
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
	BuyMinLtc                 float64 `json:"buyMinLtc"`
	BuyMaxLtc                 int     `json:"buyMaxLtc"`
	BuyMinEth                 float64 `json:"buyMinEth"`
	BuyMaxEth                 int     `json:"buyMaxEth"`
	BuyMinBch                 float64 `json:"buyMinBch"`
	BuyMaxBch                 int     `json:"buyMaxBch"`
	BuyMinXrp                 int     `json:"buyMinXrp"`
	BuyMaxXrp                 int     `json:"buyMaxXrp"`
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
Knaken ...
*/
type Knaken struct {
	DASH struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DASH"`
	BTC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BTC"`
	ETH struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ETH"`
	BCH struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BCH"`
	ETC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ETC"`
	LTC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"LTC"`
	REP struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"REP"`
	XMR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XMR"`
	XRP struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XRP"`
	ZEC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ZEC"`
	DOGE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DOGE"`
	VTC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VTC"`
	PPC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PPC"`
	FTC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"FTC"`
	RDD struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"RDD"`
	NXT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"NXT"`
	POT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"POT"`
	BLK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BLK"`
	EMC2 struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"EMC2"`
	XMY struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XMY"`
	GRS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"GRS"`
	NLG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"NLG"`
	MONA struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MONA"`
	VRC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VRC"`
	CURE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CURE"`
	XST struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XST"`
	VIA struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VIA"`
	PINK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PINK"`
	IOC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"IOC"`
	SYS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SYS"`
	DGB struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DGB"`
	BURST struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BURST"`
	EXCL struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"EXCL"`
	BLOCK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BLOCK"`
	BTS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BTS"`
	GAME struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"GAME"`
	NXS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"NXS"`
	GEO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"GEO"`
	FLO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"FLO"`
	MUE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MUE"`
	XEM struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XEM"`
	SPHR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SPHR"`
	OK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"OK"`
	AEON struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"AEON"`
	EXP struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"EXP"`
	FCT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"FCT"`
	SLS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SLS"`
	RADS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"RADS"`
	DCR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DCR"`
	XVG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XVG"`
	PIVX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PIVX"`
	MEME struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MEME"`
	STEEM struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"STEEM"`
	LSK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"LSK"`
	WAVES struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"WAVES"`
	LBC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"LBC"`
	SBD struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SBD"`
	STRAT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"STRAT"`
	ARDR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ARDR"`
	XZC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XZC"`
	NEO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"NEO"`
	UBQ struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"UBQ"`
	KMD struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"KMD"`
	SIB struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SIB"`
	CRW struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CRW"`
	ARK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ARK"`
	INCNT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"INCNT"`
	GBYTE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"GBYTE"`
	GNT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"GNT"`
	EDG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"EDG"`
	MORE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MORE"`
	RLC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"RLC"`
	HMQ struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"HMQ"`
	ANT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ANT"`
	SC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SC"`
	BAT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BAT"`
	ZEN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ZEN"`
	QRL struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"QRL"`
	PTOY struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PTOY"`
	BNT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BNT"`
	SNT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SNT"`
	DCT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DCT"`
	XEL struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XEL"`
	MCO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MCO"`
	ADT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ADT"`
	PAY struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PAY"`
	MTL struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MTL"`
	STORJ struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"STORJ"`
	ADX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ADX"`
	OMG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"OMG"`
	CVC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CVC"`
	PART struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PART"`
	DNT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DNT"`
	MANA struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MANA"`
	RCN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"RCN"`
	VIB struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VIB"`
	MER struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MER"`
	ENG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ENG"`
	UKG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"UKG"`
	IGNIS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"IGNIS"`
	SRN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SRN"`
	ZRX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ZRX"`
	VEE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VEE"`
	TRX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"TRX"`
	TUSD struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"TUSD"`
	LRC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"LRC"`
	DMT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DMT"`
	STORM struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"STORM"`
	GTO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"GTO"`
	TUBE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"TUBE"`
	CMCT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CMCT"`
	NLC2 struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"NLC2"`
	BKX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BKX"`
	MFT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MFT"`
	LOOM struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"LOOM"`
	RFR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"RFR"`
	RVN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"RVN"`
	BFT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BFT"`
	GO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"GO"`
	HYDRO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"HYDRO"`
	UPP struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"UPP"`
	ENJ struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ENJ"`
	MET struct {
		Price     int    `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MET"`
	DTA struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DTA"`
	EDR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"EDR"`
	IHT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"IHT"`
	XHV struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"XHV"`
	NPXS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"NPXS"`
	PAL struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PAL"`
	PAX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PAX"`
	ZIL struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ZIL"`
	MOC struct {
		Price     int    `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MOC"`
	OST struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"OST"`
	SPC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SPC"`
	IOST struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"IOST"`
	SOLVE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SOLVE"`
	USDS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"USDS"`
	JNT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"JNT"`
	LBA struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"LBA"`
	DENT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DENT"`
	DRGN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DRGN"`
	VITE struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VITE"`
	IOTX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"IOTX"`
	BTM struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BTM"`
	ELF struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ELF"`
	QNT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"QNT"`
	BTU struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BTU"`
	SPND struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SPND"`
	BTT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BTT"`
	NKN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"NKN"`
	GRIN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"GRIN"`
	CTXC struct {
		Price     int    `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CTXC"`
	HXRO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"HXRO"`
	META struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"META"`
	ANKR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ANKR"`
	TRAC struct {
		Price     int    `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"TRAC"`
	CRO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CRO"`
	ONT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ONT"`
	ONG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ONG"`
	AERGO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"AERGO"`
	TTC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"TTC"`
	SLT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SLT"`
	PTON struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PTON"`
	PI struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PI"`
	PLA struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PLA"`
	ART struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ART"`
	ORBS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ORBS"`
	VBK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VBK"`
	BORA struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BORA"`
	CND struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CND"`
	FX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"FX"`
	OCEAN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"OCEAN"`
	BWX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BWX"`
	VDX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VDX"`
	MAID struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MAID"`
	COSM struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"COSM"`
	LAMB struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"LAMB"`
	STPT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"STPT"`
	DAI struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DAI"`
	CPT struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CPT"`
	PROM struct {
		Price     int    `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PROM"`
	ABYSS struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"ABYSS"`
	FXC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"FXC"`
	DUSK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"DUSK"`
	URAC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"URAC"`
	OneST struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"1ST"`
	TEMCO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"TEMCO"`
	SPIN struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SPIN"`
	LUNA struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"LUNA"`
	BLOC struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"BLOC"`
	CHR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"CHR"`
	TUDA struct {
		Price     int    `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"TUDA"`
	UTK struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"UTK"`
	PXL struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PXL"`
	AKRO struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"AKRO"`
	TSHP struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"TSHP"`
	HEDG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"HEDG"`
	WAXP struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"WAXP"`
	MRPH struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MRPH"`
	HBAR struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"HBAR"`
	PLG struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"PLG"`
	VET struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"VET"`
	SIX struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"SIX"`
	MED struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"MED"`
	FNB struct {
		Price     string `json:"price"`
		SellPrice string `json:"sellPrice"`
	} `json:"FNB"`
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
