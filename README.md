# `money`

this project is currently intended to enable a few things:

  - formatting agnostic (ie 1000 vs 10.00 vs $100), w.r.t. strings mostly

  - primitive agnostic, but consolidated.

  - provide a simple way for financial applications to be created with a higher degree of semantic value.

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
### rev: [#357692f](https://github.com/coip/money/commit/357692f) date: **3/31/2019**
    Benchmarks were ran on a i7/16Gb SurfaceBook (1st Gen) w/ dGPU, in GOOS=Windows, while plugged in and on the Performance power plan.

# `goos`: linux
```
ethan@vanitas:/mnt/c/Users/ethan/go/src/github.com/coip/money$ go test -benchmem -bench .
goos: linux
goarch: amd64
pkg: github.com/coip/money
BenchmarkFromString10-4                                  5000000               332 ns/op             128 B/op          4 allocs/op
BenchmarkFromString10DollarSign-4                        5000000               293 ns/op             128 B/op          4 allocs/op
BenchmarkFromString10DollarSignNeg-4                     3000000               377 ns/op             128 B/op          4 allocs/op
BenchmarkFromString1000-4                                5000000               365 ns/op             128 B/op          4 allocs/op
BenchmarkFromString10Bucks-4                             3000000               440 ns/op             176 B/op          5 allocs/op
BenchmarkFromStringInvalid-4                             3000000               555 ns/op             192 B/op          8 allocs/op
BenchmarkFromi100-4                                     20000000                83.1 ns/op            80 B/op          2 allocs/op
BenchmarkFromi1000-4                                    20000000                83.8 ns/op            80 B/op          2 allocs/op
BenchmarkFromiMaxInt32-4                                20000000                85.3 ns/op            80 B/op          2 allocs/op
BenchmarkFromiMinInt32-4                                20000000                85.3 ns/op            80 B/op          2 allocs/op
BenchmarkFromiMaxInt64-4                                20000000                85.9 ns/op            80 B/op          2 allocs/op
BenchmarkFromiMinInt64-4                                20000000                83.8 ns/op            80 B/op          2 allocs/op
BenchmarkFromi10-4                                      20000000                83.9 ns/op            80 B/op          2 allocs/op
BenchmarkFromi64_10-4                                   20000000                84.4 ns/op            80 B/op          2 allocs/op
BenchmarkFromf32_10-4                                   10000000               183 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_10-4                                    5000000               232 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_Max-4                                     50000             27463 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_Max95-4                                   50000             23733 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_Max80-4                                  100000             17017 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_Max65-4                                  100000             11307 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_Max50-4                                  200000              6197 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_Max35-4                                  500000              2773 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_Max10-4                                 2000000               802 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_MaxNeg-4                                  50000             26374 ns/op              80 B/op          2 allocs/op
BenchmarkFromf64_Min-4                                     30000             42256 ns/op              80 B/op          2 allocs/op
BenchmarkString-4                                        1000000              1832 ns/op            1200 B/op         33 allocs/op
BenchmarkFormatter-4                                     1000000              1878 ns/op            1200 B/op         33 allocs/op
BenchmarkStringNeg-4                                     1000000              1954 ns/op            1280 B/op         35 allocs/op
BenchmarkFormatterNeg-4                                   500000              2376 ns/op            1280 B/op         35 allocs/op
BenchmarkMoneyEqMoney-4                                  2000000               631 ns/op             416 B/op         10 allocs/op
BenchmarkAddingPennyToPennies-4                         2000000000               0.43 ns/op            0 B/op          0 allocs/op
BenchmarkAdding1MMPennyToPennies-4                      2000000000               0.77 ns/op            0 B/op          0 allocs/op
BenchmarkAddingPenniesToMoney-4                          3000000               555 ns/op             480 B/op         12 allocs/op
BenchmarkAddingPennyToMoney-4                            3000000               559 ns/op             480 B/op         12 allocs/op
BenchmarkAddingMoneyToPenniesAndPenny-4                  1000000              2028 ns/op            1296 B/op         32 allocs/op
BenchmarkCastingAddingPenniesNative-4                   2000000000               0.89 ns/op            0 B/op          0 allocs/op
BenchmarkCastingAddingPenniesNativeExplicit-4           2000000000               0.51 ns/op            0 B/op          0 allocs/op
BenchmarkCastingPenniesToInt64-4                        2000000000               0.80 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/coip/money   68.887s
```
# `goos`: windows
```
Running tool: C:\Go\bin\go.exe test -benchmem -run=^$ github.com\coip\money -bench . -coverprofile=C:\Users\ethan\AppData\Local\Temp\vscode-goqaqmJr\go-code-cover

goos: windows
goarch: amd64
pkg: github.com/coip/money
BenchmarkFromString10-4                         	 5000000	       352 ns/op	     128 B/op	       4 allocs/op
BenchmarkFromString10DollarSign-4               	 5000000	       378 ns/op	     128 B/op	       4 allocs/op
BenchmarkFromString10DollarSignNeg-4            	 5000000	       329 ns/op	     128 B/op	       4 allocs/op
BenchmarkFromString1000-4                       	 5000000	       338 ns/op	     128 B/op	       4 allocs/op
BenchmarkFromString10Bucks-4                    	 3000000	       456 ns/op	     176 B/op	       5 allocs/op
BenchmarkFromStringInvalid-4                    	 3000000	       626 ns/op	     192 B/op	       8 allocs/op
BenchmarkFromi100-4                             	20000000	        99.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromi1000-4                            	20000000	        81.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromiMaxInt32-4                        	20000000	        81.5 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromiMinInt32-4                        	20000000	        79.1 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromiMaxInt64-4                        	20000000	        79.5 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromiMinInt64-4                        	20000000	        89.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromi10-4                              	20000000	        92.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromi64_10-4                           	20000000	        89.5 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf32_10-4                           	10000000	       205 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_10-4                           	 5000000	       256 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_Max-4                          	   50000	     28660 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_Max95-4                        	  100000	     26730 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_Max80-4                        	  100000	     17439 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_Max65-4                        	  200000	     11207 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_Max50-4                        	  300000	      5916 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_Max35-4                        	 1000000	      2629 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_Max10-4                        	 2000000	       829 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_MaxNeg-4                       	   50000	     26213 ns/op	      80 B/op	       2 allocs/op
BenchmarkFromf64_Min-4                          	   30000	     43066 ns/op	      80 B/op	       2 allocs/op
BenchmarkString-4                               	 1000000	      1802 ns/op	    1200 B/op	      33 allocs/op
BenchmarkFormatter-4                            	 1000000	      1850 ns/op	    1200 B/op	      33 allocs/op
BenchmarkStringNeg-4                            	 1000000	      1901 ns/op	    1280 B/op	      35 allocs/op
BenchmarkFormatterNeg-4                         	 1000000	      1942 ns/op	    1280 B/op	      35 allocs/op
BenchmarkMoneyEqMoney-4                         	 3000000	       517 ns/op	     416 B/op	      10 allocs/op
BenchmarkAddingPennyToPennies-4                 	2000000000	         0.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkAdding1MMPennyToPennies-4              	2000000000	         0.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddingPenniesToMoney-4                 	 3000000	       549 ns/op	     479 B/op	      12 allocs/op
BenchmarkAddingPennyToMoney-4                   	 3000000	       536 ns/op	     479 B/op	      12 allocs/op
BenchmarkAddingMoneyToPenniesAndPenny-4         	  500000	      2218 ns/op	    1328 B/op	      33 allocs/op
BenchmarkCastingAddingPenniesNative-4           	2000000000	         0.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastingAddingPenniesNativeExplicit-4   	2000000000	         1.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastingPenniesToInt64-4                	2000000000	         0.56 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	github.com/coip/money	72.130s
Success: Benchmarks passed.
```