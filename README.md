# money

this typedef is currently intended to enable a few things:

  - formatting agnostic (ie 1000 vs 10.00 vs $100), w.r.t. strings mostly

  - primitive agnostic, but consolidated.

Doesnt handle currently: (feel free to open a pr :-) )
 - -$100
 - -$100.00
 - ($100)
 - (100)

----

to elaborate on ~agnostic: [money.Money](./money.go) *is* impl'd via [decimal](https://github.com/shopspring/decimal) under the hood at the moment,

but provides a means of homologating various input types while consolidating into a single underlying type.

so if something else strikes your fancy &#8212; maybe `type pennies int` &#8212; you may lift `type Money decimal.Decimal` instead via `type Money pennies`, without the need to refactor client code which might depends on `money.Money`.