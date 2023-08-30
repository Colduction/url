# url
 Fast, optimized and pool-based `url` package inspired from `"net/url"` written in Go

## Installation

Use go get.

    go get github.com/colduction/url
    
## :warning: CAUTION
Be aware of the `(*url).Reset(put bool)` function, set the `put` parameter to **true** when you are done with that `*url` instance and if you want to use it multiple times, you must set the `put` parameter to **false**, otherwise you will increase the pool size with unused `*url` structures.