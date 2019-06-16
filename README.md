# Avida

[![CircleCI](https://circleci.com/gh/jkielbaey/avida.svg?style=svg)](https://circleci.com/gh/jkielbaey/avida)

This tool allows you to get a total overview of your crypto assets. These
assets can either be on an exchange (currently only binance) or in another
wallet.

## Getting started

### Installation

Download the relevant binary from [releases](https://github.com/jkielbaey/avida/releases).

### Configuration

All configuration is done using a `.avida.toml` file in your user's home directory.

To add an exchange you should add a section as below. (Currently only binance is supported).

```toml
[[exchanges]]
exchange = "binance"
apikey = "<API_KEY>"
apisecret = "<API_SECRET>"
```

You can also add fixed positions/coins (not on an exchange).

```toml
[[positions]]
symbol = "XMR"
amount = 199.34
```

### Usage

Just run `avida`

```sh
$ avida
  XMR :  199.34 => $ 199.34
  LTC :    0.00 => $   0.01
                 ------------
 Total ==>         $ 199.35
```

## Known issues

Unfortunately Binance and CoinMarketCap don't use the same symbols for all cryptocurrencies. This can lead to an incorrect total values of your portfolio.

## Contributing

1. Fork it!
1. Create your feature branch: `git checkout -b my-new-feature`
1. Commit your changes: `git commit -am 'Add some feature'`
1. Push to the branch: `git push origin my-new-feature`
1. Submit a pull request :D.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
