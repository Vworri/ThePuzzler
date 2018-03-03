package puzzler

import "time"

var prinable = string("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \n\t")
var algorithm string

var emptyHashes emptyHashStruct
var unhashingTime time.Duration
