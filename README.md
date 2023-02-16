# Stockstatz CLI Tool

This is a simple CLI tool written in Go that takes in a ticker symbol, start date, and end date and calculates the simple return and maximum drawdown for the given period. It then publishes the results to the 'stockstatz' Telegram channel.

## Installation

To use this tool, you need to have Go installed on your machine. You can download it from the official Go website.

Once you have Go installed, you can install the tool using the following command:

```go install github.com/zareifar/stock-stat```

## Usage

To use the tool, you can either run the program using go by running the following command in the project directory:

```go run . 2017-01-01 to 2018-01-01```

or build the project and use it independantly, by running the following command in the project directory:

```go build```

and then run the executable followed by the ticker symbol, start date, and end date. For example:

```stock-stat AAPL 2017-01-01 to 2018-01-01```

This will calculate the simple return and maximum drawdown for AAPL stock between January 1, 2017 and January 1, 2018 and publish the results to the 'stockstatz' Telegram channel.
Please note that This program is using the Nasdaq API in "none-subscriber" mode, and as a result you can view a subset of the data.
The sample includes data from 2017-09-01 to 2017-10-31 for the following tickers:

**MMM, AXP, AAPL, BA, CAT, CVX, CSCO, KO, DIS, XOM, GE, GS, HD, IBM, INTC, JNJ, JPM, MCD, MRK, MSFT, NKE, PFE, PG, TRV, UTX, UNH, VZ, V, WMT**

## License

This tool is released under the MIT License. See LICENSE for more information.
