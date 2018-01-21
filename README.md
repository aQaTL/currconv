# mmc [![Build Status](https://travis-ci.org/NARKOZ/mmc.svg)](https://travis-ci.org/NARKOZ/mmc)

currconv is a simple commandline currency converter.
It's a fork of [mmc](https://github.com/narkoz/mmc), but I had to change the name, 
because it conflicts with Microsoft Management Console (`mmc`) on Windows.

# Installation

```sh
go get -u github.com/aQaTL/currconv
```

# Usage

The usage is similar to Google's Currency converter.

Run from your terminal (case insensitive):

```
currconv 100 USD AUD
```

or:

```
currconv 12.5 btc usd
```

The first argument is an amount for conversion, the second and last arguments
are currency codes.

If you want to get just the converted value, add `-short` before arguments:

```
currconv -short 100 usd pln
```

For a list of supported 150+ currencies and their respective codes see
[`data/currencies.json`](https://github.com/aQaTL/currconv/blob/master/data/currencies.json)

# License

Released under the BSD 2-clause license. See LICENSE.txt for details.
