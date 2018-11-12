package wex

import (
	"testing"

	"github.com/thrasher-/gocryptotrader/config"
	"github.com/thrasher-/gocryptotrader/currency/symbol"
	exchange "github.com/thrasher-/gocryptotrader/exchanges"
)

var w WEX

// Please supply your own keys for better unit testing
const (
	apiKey    = ""
	apiSecret = ""
)

func TestSetDefaults(t *testing.T) {
	w.SetDefaults()
}

func TestSetup(t *testing.T) {
	wexConfig := config.GetConfig()
	wexConfig.LoadConfig("../../testdata/configtest.json")
	conf, err := wexConfig.GetExchangeConfig("WEX")
	if err != nil {
		t.Error("Test Failed - WEX init error")
	}
	conf.APIKey = apiKey
	conf.APISecret = apiSecret
	conf.AuthenticatedAPISupport = true

	w.Setup(conf)
}

func TestGetTradablePairs(t *testing.T) {
	t.Parallel()
	_, err := w.GetTradablePairs()
	if err != nil {
		t.Errorf("Test failed. GetTradablePairs err: %s", err)
	}
}

func TestGetInfo(t *testing.T) {
	t.Parallel()
	_, err := w.GetInfo()
	if err != nil {
		t.Error("Test Failed - GetInfo() error")
	}
}

func TestGetTicker(t *testing.T) {
	t.Parallel()
	_, err := w.GetTicker("btc_usd")
	if err != nil {
		t.Error("Test Failed - GetTicker() error", err)
	}
}

func TestGetDepth(t *testing.T) {
	t.Parallel()
	_, err := w.GetDepth("btc_usd")
	if err != nil {
		t.Error("Test Failed - GetDepth() error", err)
	}
}

func TestGetTrades(t *testing.T) {
	t.Parallel()
	_, err := w.GetTrades("btc_usd")
	if err != nil {
		t.Error("Test Failed - GetTrades() error", err)
	}
}

func TestGetAccountInfo(t *testing.T) {
	t.Parallel()
	_, err := w.GetAccountInfo()
	if err == nil {
		t.Error("Test Failed - GetAccountInfo() error", err)
	}
}

func TestGetActiveOrders(t *testing.T) {
	t.Parallel()
	_, err := w.GetActiveOrders("")
	if err == nil {
		t.Error("Test Failed - GetActiveOrders() error", err)
	}
}

func TestGetOrderInfo(t *testing.T) {
	t.Parallel()
	_, err := w.GetOrderInfo(6196974)
	if err == nil {
		t.Error("Test Failed - GetOrderInfo() error", err)
	}
}

func TestCancelOrder(t *testing.T) {
	t.Parallel()
	_, err := w.CancelOrder(1337)
	if err == nil {
		t.Error("Test Failed - CancelOrder() error", err)
	}
}

func TestTrade(t *testing.T) {
	t.Parallel()
	_, err := w.Trade("", "buy", 0, 0)
	if err == nil {
		t.Error("Test Failed - Trade() error", err)
	}
}

func TestGetTransactionHistory(t *testing.T) {
	t.Parallel()
	_, err := w.GetTransactionHistory(0, 0, 0, "", "", "")
	if err == nil {
		t.Error("Test Failed - GetTransactionHistory() error", err)
	}
}

func TestGetTradeHistory(t *testing.T) {
	t.Parallel()
	_, err := w.GetTradeHistory(0, 0, 0, "", "", "", "")
	if err == nil {
		t.Error("Test Failed - GetTradeHistory() error", err)
	}
}

func TestWithdrawCoins(t *testing.T) {
	t.Parallel()
	_, err := w.WithdrawCoins("", 0, "")
	if err == nil {
		t.Error("Test Failed - WithdrawCoins() error", err)
	}
}

func TestCoinDepositAddress(t *testing.T) {
	t.Parallel()
	_, err := w.CoinDepositAddress("btc")
	if err == nil {
		t.Error("Test Failed - WithdrawCoins() error", err)
	}
}

func TestCreateCoupon(t *testing.T) {
	t.Parallel()
	_, err := w.CreateCoupon("bla", 0)
	if err == nil {
		t.Error("Test Failed - CreateCoupon() error", err)
	}
}

func TestRedeemCoupon(t *testing.T) {
	t.Parallel()
	_, err := w.RedeemCoupon("bla")
	if err == nil {
		t.Error("Test Failed - RedeemCoupon() error", err)
	}
}

func setFeeBuilder() exchange.FeeBuilder {
	return exchange.FeeBuilder{
		Amount:              1,
		Delimiter:           "_",
		FeeType:             exchange.CryptocurrencyTradeFee,
		FirstCurrency:       symbol.LTC,
		SecondCurrency:      symbol.BTC,
		IsMaker:             false,
		PurchasePrice:       1,
		CurrencyItem:        symbol.USD,
		BankTransactionType: exchange.WireTransfer,
	}
}

func TestGetFee(t *testing.T) {
	w.SetDefaults()
	TestSetup(t)
	var feeBuilder = setFeeBuilder()

	// CryptocurrencyTradeFee Basic
	if resp, err := w.GetFee(feeBuilder); resp != float64(0.002) || err != nil {
		t.Error(err)
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0.002), resp)
	}

	// CryptocurrencyTradeFee High quantity
	feeBuilder = setFeeBuilder()
	feeBuilder.Amount = 1000
	feeBuilder.PurchasePrice = 1000
	if resp, err := w.GetFee(feeBuilder); resp != float64(2000) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(2000), resp)
		t.Error(err)
	}

	// CryptocurrencyTradeFee IsMaker
	feeBuilder = setFeeBuilder()
	feeBuilder.IsMaker = true
	if resp, err := w.GetFee(feeBuilder); resp != float64(0.002) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0.002), resp)
		t.Error(err)
	}

	// CryptocurrencyTradeFee Negative purchase price
	feeBuilder = setFeeBuilder()
	feeBuilder.PurchasePrice = -1000
	if resp, err := w.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
		t.Error(err)
	}
	// CryptocurrencyWithdrawalFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.CryptocurrencyWithdrawalFee
	if resp, err := w.GetFee(feeBuilder); resp != float64(0.001) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0.001), resp)
		t.Error(err)
	}

	// CryptocurrencyWithdrawalFee Invalid currency
	feeBuilder = setFeeBuilder()
	feeBuilder.FirstCurrency = "hello"
	feeBuilder.FeeType = exchange.CryptocurrencyWithdrawalFee
	if resp, err := w.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
		t.Error(err)
	}

	// CyptocurrencyDepositFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.CyptocurrencyDepositFee
	if resp, err := w.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
		t.Error(err)
	}

	// InternationalBankDepositFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.InternationalBankDepositFee
	if resp, err := w.GetFee(feeBuilder); resp != float64(0.065) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0.065), resp)
		t.Error(err)
	}

	// InternationalBankWithdrawalFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.InternationalBankWithdrawalFee
	feeBuilder.CurrencyItem = symbol.USD
	if resp, err := w.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
		t.Error(err)
	}
}

func TestFormatWithdrawPermissions(t *testing.T) {
	// Arrange
	w.SetDefaults()
	expectedResult := exchange.AutoWithdrawCryptoWithAPIPermissionText
	// Act
	withdrawPermissions := w.FormatWithdrawPermissions()
	// Assert
	if withdrawPermissions != expectedResult {
		t.Errorf("Expected: %s, Recieved: %s", expectedResult, withdrawPermissions)
	}
}
