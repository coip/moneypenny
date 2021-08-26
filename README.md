# `moneypenny` ![badge](https://action-badges.now.sh/coip/moneypenny)

this project is currently intended to enable a few things:

  - formatting agnostic (ie 1000 vs 10.00 vs $100), w.r.t. strings mostly

  - primitive agnostic, but consolidated.

  - provide a simple way for financial applications to be created with a higher degree of semantic value.

  - provide a straightforward reference for actions

Doesnt handle currently: (feel free to open a pr :-) )
 - -$100
 - -$100.00
 - ($100)
 - (100)

----

to elaborate on ~agnostic: [money.Money](./money.go) *is* impl'd via [decimal](https://github.com/shopspring/decimal) under the hood at the moment,

but provides a means of homologating various input types while consolidating into a single underlying type that plays well with both presentation-layer and persistence-layer.

----

# latest benchmarks
Visit the Go Build task under the Actions tab to find `unit test` and `benchmarking` results..
