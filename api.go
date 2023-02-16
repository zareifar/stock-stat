package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
)

func FetchStockData(symbol string, startDate time.Time, endDate time.Time) []byte {
	// Fetches historical stock data for a given symbol between the specified start and end dates from the NASDAQ API.
	// Parameters:
	//   symbol (string): The stock symbol to retrieve data for.
	//   startDate (time.Time): The start date of the data range, inclusive.
	//   endDate (time.Time): The end date of the data range, inclusive.
	// Returns:
	//   []byte: The historical stock data in JSON format as bytes.

	config := GetConfig()

	reqUrl, err := url.Parse(config.Api.Nasdaq.EndpointUrl)
	if err != nil {
		panic(err)
	}

	urlParams := reqUrl.Query()

	urlParams.Add("api_key", config.Api.Nasdaq.ApiKey)
	urlParams.Add("date.gte", startDate.Format(config.Api.Nasdaq.DateFormat))
	urlParams.Add("date.lte", endDate.Format(config.Api.Nasdaq.DateFormat))
	urlParams.Add("qopts.columns", config.Api.Nasdaq.QueryColumns)
	urlParams.Add("ticker", symbol)

	reqUrl.RawQuery = urlParams.Encode()

	resp, err := http.Get(reqUrl.String())
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func PostMessageToTelegramChannel(message string) {
	// Sends a message to a specified Telegram channel using the Telegram API.
	//
	// Parameters:
	// 		message (string): Message to be sent to Telegram channel
	config := GetConfig()

	if config.Debug {
		return
	}

	reqUrl, err := url.Parse(config.Api.Telegram.EndpointURL + config.Api.Telegram.ApiKey)
	if err != nil {
		panic(err)
	}
	reqUrl.Path = path.Join(reqUrl.Path, config.Api.Telegram.SendMessageSlug)

	data := url.Values{
		"chat_id":{config.Api.Telegram.ChannelId},
		"text": {message},
	}

	resp, err := http.PostForm(reqUrl.String(), data)
	if err != nil {
		log.Println(err)
	}
	
	if resp.StatusCode != 200 {
		log.Println("There was a problem with posting information to Telegram:", resp.Status, resp.Body)
	} else {
		fmt.Println("The results were published to the Telegram channel 'stockstatz'")
	}
}
