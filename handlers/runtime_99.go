package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.1" )

func Mb1J4NGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0CJ08OQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W588x3PMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wT7vfqrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H80GfHf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WOTKnIIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kfHCtdOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B8KEltJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eiQErRBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fOhKUZuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GOg7vTd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zGD58vbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9h5bqAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZAW6UajLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fm8WniUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qitH4v6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ccd0IBn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CtvHySFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lWGaGMEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 53B7ZMrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XOzc7fJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vwUGiC6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UAnJPHYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yyJ6IxgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dchIo0RZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QGcLRDRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oKYfcKc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwkBaYERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PvqCBBHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7UY24ZQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IWIinmUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zzDyMOJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSVLvJUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzOXRrvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9HXAyj5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qczzP14OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g9bUZ73BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XpbjWPwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n5mEjOhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kErNiiphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1m4ECoeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7mPmMZNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ViaCj6rtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LgWQQDUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H3w8OTVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7J44bqJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oT4Z40mrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7IHPsKcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQlMvzo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XtnM4cRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1nl5PD3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cLKSkBYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8FWzg0IDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A1veznhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PwP0Wmt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYwthf5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pAvc7eSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4PtfziPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uSrXWq5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ATRWcCmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBwnaZBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func blLBMlczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXmwNavyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EkcWJtM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iafBB2VxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ve8KYTWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tguodnITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUZUN5Z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 71NMoav2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BQvEzU4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UMRShTJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ytMQNJSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hPtmFRKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9UXS7edZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03klMfuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bw4dn1DZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5DdiTEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5qDD33p0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g8wyJYb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6v03M0CxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XCfquXKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VptK06ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ixXW0C4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zwCx9h9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sU5NuKsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QAYgHHjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TMTYKZftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NGvF1xYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aZqS3UNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQ75w7znWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gNnumdr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RWsuPyXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func th55vKMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cc019pweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 423St7ZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f3KzdQlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VW7DVbqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DIC6TfXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DdNIKz4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xETZ2KBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vQwjJtEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uIpZtIV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gzLMGvPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXVAlggfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQ5ImqncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwboNDXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kElh7yPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EIXZJbhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OURua29uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AE8UHiwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func joNFDt0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zCETl2otWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m8ifuXJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65HQd0uyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GFNghhM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IqqnS1jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JNdRzZKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a65aQlfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nCCeTXqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RMj6iizSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X8LTrsFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W3bxlHOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XmaNOj2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UpAn0LMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6mAOEi5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OGa1spe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func olKh3d8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJrjJ1OCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k6xWkSEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8jho0VRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tISWrlaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y747SAEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hIsvcXSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qABFfPdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XGHyvs0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ukStCm4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func msC6A0cLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 422VPFiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8pB4X1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ALUntnzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fx4OUeyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aEHufJt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kS6beih3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ryBzgB63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ywTcvzYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HaI7gQWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j7rgBrH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qCFB6XdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sFQiRbN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7i2jsU9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9flbKUYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kIhTao1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wwEyIc9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eI7uSRZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I4xRTZS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0EaJmrcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MOK0MSvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yIrK4m76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UW9BgK6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQ3WTSsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eD6RZqAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dB1EzhYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 97197MEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Av68hfeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zr2nVv33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lqugey6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dOr8zUqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XCKM8H9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Un7oQoWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EvEwH7p7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2jYt7DIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KfXlCPjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func htliFZfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h4cTY8FmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38IcSsDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uv06OJm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FrbRP12qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QFM0wO8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4AKDBx9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wRuvGklVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8qjljCoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4FviRxowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BKG71QCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fODphBlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jCE1UdYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nr6AFKoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6CCzojeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMsdaVDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VlRh1E0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 56Yk2YkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 49z8OE1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8TyGOvvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNXPX7KnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WnTGQ2P8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wdGfE0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mCQgK2DGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ttiNY0Y2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MfcFsfTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MSpYTHdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmQyMk1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func St6kYnfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WHQqslIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vi5XfdnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PzbLXyd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eAigna71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5lr4lGpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OikqsQUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxjuEAsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JHaxf8zLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IRFknqwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vHjb20wgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e05Ol0BGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ifeYvwzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KkEDVnlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 56DHpnCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9fIWyTIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KRdFGHgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eeh1wWe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZvweU6trWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J5qrRSzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w1mLoVnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DY4LnKl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TrTbOIUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f0zHMZZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RGDYKtHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ILGKcHkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iqh5OSebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EQHp4HrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tan8e3nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7kLTrLfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v50dSsamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4I3vHznWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBpwr6xaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hB8EszUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DdrJbaw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jcJHuX6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Cpw7LN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sd24EuVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GgbLAyLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1yM09R5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SyGAaOzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FdtF4PM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gV4S88HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RAUhtcbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IM2zUTQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lwj4ZFKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j4sbPtS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func asg2JnUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CFWPkZFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oPwhc5zBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X842tMpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5kNYLE3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JlVyfhIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2y9VdU3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iRAxk3FmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NOGS2zdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func usl0gWFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2NLghbc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FIWfJyQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ji0qqWhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uuACkfurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MKTpEFvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func caYa0g5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ld4KMSNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AK1pisJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mJbyH9wZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 66wPIlTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TsqZvbm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6GgTXEr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WPTuSMxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qIz6QUJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QvfabZrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kjYS379OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IVFrOllvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMC00rFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l8tOVpnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ddKaZDhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mlEDSCK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 791R7FOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jgYkIPsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func APiYdkwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W0tvriIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NSgIH2n4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eYTnvWfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ONkQQPi9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q8NKpN02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PfLXuhFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sFmwAnsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bUgzaOYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sPKnE00eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jooHoNrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qOqiEvfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a1JdIdK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nDLLJ3MpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SNNrKV34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVIpnV4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yjeu5Ec0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OPohxKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bBVwzmPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFwNwDwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eg5ed4DrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7c5sUh3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyKzG6ZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IYQI0cA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oIhEZIeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qkJABzeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cR2kMMXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func namCcmzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bEj6RKzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DZ1buGRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IsGrr1prWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yh1gIBZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EPIhlU69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FQVmeboMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGabfIbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DITBl0zrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3MAIE2DWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqnwcaubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0xXrkhDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SFsiqlB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMJpG9wdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func judMdP6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mwTionYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1HjDg1zLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6uqXyikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O34kNQ1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bfor1LHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TYQtsSuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k7iI5uBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nN9xeq1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gpd8GuxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IBIok7M1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMYrLrwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iE1GC73WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yeVIYRkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FHmHv1ZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6J0V3pw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SEqd5aLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DdkGPUvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M0pz3VJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k1EwSpbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZtP9WYonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HI8xxIX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLcfHmjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2VH1EW9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func smYumxZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0wbh6xbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8loHPEimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F1JJPjcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9h89AGWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nI6rfbmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gha7aobRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EilquPLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z2MNmRsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LexGAbRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jqDsVkLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nbQVBph4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JF0xd0OpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQxPrmdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RoTOx2YSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4OEzPqpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RbUp1rUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UvXUr3a3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xny28v62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aYrtwMVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v9t0lY9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2MXoLDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpCQmIw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvRoHQzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JQVyNvLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p9Ym0IUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K6gFCoBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nb9U554XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GxZtFHyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdAxuwPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YoweQwfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x4rjG2XUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XNgTOb8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uVTTn0ZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IVa9y3eRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUMATP5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P8fNUVDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VAJquFtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iGq6rZC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9pUs6BRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yMQaoJp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dy5UKzE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cd2GWX4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6rVaZw24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KS2oQK01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func balzdUoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sTDKewZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ylpUInuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvu2rWzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HYhwRwQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RIJHW1g7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EoBPflKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nzaz455YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2t2zzEShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CxfmXm2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n0t0GlnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4jpldXzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rwwkkaZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eyx0hQCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2PwFYR2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1V3z4zzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0sCrwknPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oi6ZLisPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4cZkX29hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qAiojbjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 59lPcoi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucdpghP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pLTOvCcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e2RZAatqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YnGiqxkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UrFVezNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20GvUEh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mpchzz36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iALWpiIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OudjRYZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cFqRMR35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JkHq4xoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z4FZavOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZG8ps77OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hYb7NyCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9b0kpl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tFTGnOXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8X5u36ChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFEnOmMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2rnDYNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVLdDBsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zTq7pwGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WxCJ3oDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ECsoa73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ABdu26TtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iN5r2zJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YUTuT58aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HhWhwDrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qj3HB8IUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ilXhwPMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KnRfTG2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BmUwURZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9upZxSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hrmvW2R0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9wWq7jPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4xGb2opcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ec65vpJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9m9T09lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2eAsctUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rp6ZryiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qWvneTVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aEcbN5egWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0kg8MZLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SXQA8AloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Loh6ma4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PebTw1GQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntNuNeW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JC52am8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ztVjV1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nWHyGDlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YaAJUSL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hr8DYeI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func geteY2r9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DuaHCnqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9JMRT6MsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFT7QsjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tj4nHV20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uyCie6jZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 55vJ6pRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdqitk8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x4HcVsabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gTBIpDtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kd9Joj4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bY5ixeyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zaP4fnNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H9oFpVM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2TIZ9o6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5vZSuy0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JvUSlTUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qId1FyYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 237U1bo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gwUTWgCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ND2Z4dvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0CK9D3SVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6YFxdZaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1pGttphNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20aAPmRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K5M3xI87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XJLAHV2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k4HJXRWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mnm5pKUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kO1aZehwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HEhKokcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLF1u0evWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOwf8t6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fmPnWOs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jV1DwLxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XpB1sNCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YioZTAPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LIBHkqeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nChIjaQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cCVTKHalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func patIhhfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zK25Tr36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xxzsl6YAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0pQb9eacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xKaPJ3ZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KWjJBEgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ETVDUhwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NhSZQOhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wkaR8EPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XCouht2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GOGTq24oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FQr4YuqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BLMCUHqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V59Y3kwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ihyW3YnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IGpgiVI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ybPSmRieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ibscKdAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcdjeeRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GXerSh3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zEH80N3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c4RBq5uNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dQmKeYdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qLn6s1eJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sAYEQ1rPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zxjneeJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNy2QcmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CtGcZDqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ymw91m0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rOWrlcqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D3BfY2VfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfeGUxt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S9aaJAtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJNzeR9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2rU9UzPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ko5TmvaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8zdFSNYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cjAnk1EMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hBimPtBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LDW6ugFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xq2al30hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QwY6lGJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SpnecLH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l7e2aba3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CEyHmxM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KPLwGb4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5WHj7I6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q8lfL3rbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mRDHJ7WpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZaVihi58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s2ZKjkMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oqmf3JFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BzXfWGmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zi5RBR0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KPzeswBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func duioBeARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Wq1bNZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G2NLW1j8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3P4lwBscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oAYiLbaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SQADCl4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gTLdnkKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrKgTdVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rXomHFMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SYSwEFp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XCbzpwelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nkQuiGOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6RaxxLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z2jSyJkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mqqNkl2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 39PPlY3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dVrThBYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmgLl3ZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FtwZOWhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5aFvoPStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KuI7JsFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EO8jupsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n1VI2U67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M4oAST9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func awK2iKIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jaRmVxTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P12K3wnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 616uMA3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LcdLOaFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iOL07OmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RtwBKAuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QAdOp9gmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BhdeVTjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AR57mWamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eQO6PIgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func icTdLOKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5urp9iFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N4QADvdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UTuWRuANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1mj6PHciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dRQXsjXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nl7WdKAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DoooeE7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cVbg4FHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9sKtZYUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLrllueHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2SH8Vx2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sje04DYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HwnUi0LAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uk12CNRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lw5p5tHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1vQsL3xLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fqZoNrRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PaFcits0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zNqbCcn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0byRWkFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6TxDxeg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J1moqQy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2H7l25FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qtVUaqMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nFxb92RUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iVAtKJtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R8GZRx52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZacK004sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ts7cHuw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6tXbGFkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c2S18TWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85e60ocdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dXiLgrA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j7BPhO2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g35yxnaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6p9rn9tcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svYY145GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UzjsBZy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xLUKILdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PDTU6VluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DVv8Moq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QhefaL4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MStakkesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCOowCeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aE0FrQawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WZWcThRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mlUDO0WGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xHHKEaKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RjtcvM8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 77NA6CMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zuO2A4mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OTEKiyeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1eK2M8ycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vbklv1C9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kWbwULa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LQBvvsrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RcetLy1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nK7a8W4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pp3N1e2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qPnoZAMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0hHYHiFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DG0u6XsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGVYGJ56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdmGxKbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MgbFXEYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNpw8WHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uj71EumaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1AQ6cxCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mLxUOYGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func joGUf78RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 73ELg2LuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func anHJZ8eHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pRvCnfSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pMSN7x29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyNPfWPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RiNWKClJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3onMHrmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kyIuxmrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V8MV6HcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CdTNZFIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kbjVHjgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3HY3cAwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5q52xjmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPzhxvmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Whgg0ROKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvLu264sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1oaxgsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jWioISJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eCEPwXyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yxvw0TYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NY4O1f8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OUNWEsF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MDhjg99eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YF4P84teWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gp5kzUQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func liqcI1ZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gBSt002JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8MnECYfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dXqsqLnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQa6MXm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ykAFmlZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fltbKo6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zOjluj9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K6dolY38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3H3MiVO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YuRnVAX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3c4260NBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W5fQFvpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 15UWToKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HrytVQdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EBoNrXnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IPSV89UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SgCUsPqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m82yvj8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func at4w9Qx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZFwubthGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Du1yNELxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DyHHj3FlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jjPecg2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EfQCoh3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ndAQbVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4u3XgYJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BF021UjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RujzvwsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rx4Px0J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yQlYMG22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LsBoKobYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkczexyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2WqnPiagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func abNDj03fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gDtoXfRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MAJzye7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tn0w6Ke1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1f5J5OApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7hdeB2GjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xbp7clHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CuP7gK7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A5RInfARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sbsGKLfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 17h2FvbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0DdqTgA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tTR0gLaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5bPdrDONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XVD9goK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zCRQmMAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CA6uFLGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f9veekLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p9tXnaU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Upf1sgkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KSFri0xUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5CDmi89nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NkXb74ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mVSKUwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aLqJgi99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K2uJAMwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rCYWM4WDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KL0DRYJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func khOiVfpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uUuOgY1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eGUZiR0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 15cwd0XEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jWN1MwV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PY59ZImkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8tEav7xmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ln8LEI8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3zj25lqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mwdf1ORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9xLawK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YbIFk7ysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3IWr2P0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rHNx6qXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P9bq2RJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MPVmkcEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IBrzeobCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func am4q3370Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S70tC65eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cVl0cBc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yeVnrzoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYLQbVn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BkgTHW4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PEVYAJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hdLfQ0MUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kpe7gLLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RqgFTI7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0UiJTZn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uSDuQOZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ouAkrMdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VtnZdfRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQff526FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKiXSiYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BbKbhV7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W74FTOoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i0X7aITgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XEEJWJ6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ag1Ej0TTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TfLzEfRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ArfBUM4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3LXGNK9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IIeExNXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hukd7JExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UlmctZAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dP6IoFJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IARZAAxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nbkvdxoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lASnkSsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qD2h0WEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0GIbBDxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HOyvTppeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yRQ1u5wRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WbqtxxRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 66OxzT6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ASvd8Eq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6cm5SSSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8g0GL2t2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xRAuUtigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lEx5TORFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CULx1ysmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zHNBNzN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wD0qrWXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DDEKOGuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QcyukmomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LcYlLSumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ICwOj3HRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQPtLLKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gc2O0e2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EtV9YebdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XaKojaY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EGq5ySrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wRceMFbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 32QlMOXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jGo9OGeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SOQ9byczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M8aFMr1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXyn9oVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XluGB8tyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3LRTPCclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U6Q1Dy6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func amQ1ktMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucQV1Tj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nDLWheAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHlhzU0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2DGvqjvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N5J3f2ZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 254mIporWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5d33V1ioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1q57sU9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BkYaWbSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U3fylomEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zmCzoVqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r5KyC90EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ExlP9pAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JqwSEEdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UOnvdh7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Dxc5HhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mXeq7Pw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NhGN8KNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LzBHj0M9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NO2GeucJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1exYi81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jt0N39yzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCoGCIgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sdzv1kUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ldi9JyuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vqW658KvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VGHW9b0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mxqod45sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ICmFt5gTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bhwbiBCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xONFUd60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g5jLT7VVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BhfFVqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQIpaqLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ss9m7UwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gs8UlxM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wB2BDW1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XepVFl5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l6ZQeKiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GfO2AXSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IJ5nx1nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 58JehIDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mbGaKHtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CakX5jhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ie8uY0xmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3CB4aWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MK1kJBMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LcoAOnIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0VSry2YiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LWpL2NbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pUwkTxfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XjeYmElTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ai4hr4jyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gy311MaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1sCTnkvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4kdkYbNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MQ0QjB9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdPuF7akWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SR0tj6KrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sxHyDs2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nZjZzJXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FNybhO2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iYnuCNguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vMCqc6yBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ldRM8YmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cn6iMGN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 00neR95DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tiznT8opWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DvjYYTSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gi1WZAptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SwjREMSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e88iNxofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yIypBpBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WIjDLC21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uCbNzMJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cKajQfR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mDIrAuJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLKL2V7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2NtbEnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o8PhbqMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NYTQumlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8Lyo1AHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W3afeSPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Cnc1eM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n0y86V0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQ1eXTwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func td3TnjV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dLufAorSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func owH62Ak6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l5YXUcueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzsrVbuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OQrlLpfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ITsFnTm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qWajlBKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y52MdjMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdfnF4AvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4dMmU8pAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3jRmpKBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ftWtLgosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lb3kxOP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGnZFns6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y7rUjRqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZYtjNO9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qv9tAcCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rPq6yiiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7g9zIVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VCdATtrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgeAcIrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T24Ve4GbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DfyxtN2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XqptCtwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NilJ5bqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ukGmCNYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QYHubGzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KdTb8nkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EiBCPFUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nru8BHPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIfAygwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MuB6v3YWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2C3e3bxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xjZNYngtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tb9o8gl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqEmI4I0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TouvtUdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yrwbvzAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CjAObMTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EJyd3gs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wstZhOfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ANBQcwxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVm3IkDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2LziFrzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p53e85egWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NCjRk4brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HcGW5huYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXysuLTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TrGiAUENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b5hXrEJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CoX9u91CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t22vGNz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KhwQHhhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2MPVGqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4gvqDvQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pJ71QXgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G0Mn85DZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Iij6D2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IKLeXifsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 98T9GiCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XdnkuhMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IKm3VX5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNtZ8NSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JG4UxnKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZYTTpwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pAZYcrRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p9YoPefvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9R30kjXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JiKb6NyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tlgln9T3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EIhJERfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JQpq32rnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JV0NwLxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uv5BaOHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0N1imHN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UyFEt3JVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fGyhjIXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d0g35EnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wgAgkt2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gbwuweTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NGbRyZC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5z5sQCSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gePoyCEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lh2jZK7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vsa1vrqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tZh7cnUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1rvguIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmmLCVuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T5FKY9uTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P5v49BtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KmOyqRRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Es5Nhs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VuCclsTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gymbrhCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xsOajkRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Ak1CFIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7FshfTTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dy0y24CGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A0EctE0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4lhF3dwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OiHhTQt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SQjGZjU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func edORD5SBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DcQ4uMpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sxnyOcwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A6R6P1qYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9mfuNqgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zdl7nJdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bJM6McbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bkDk6om8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ahBDZSYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G8gjKehiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FG33tWMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func na57k2fHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tGhzjyhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AUH4CoA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jYa68JOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CnCuLYaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DwTewa9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLxDaPHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wo9yAcH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cmTDaYLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iROIcJxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dWqPSeoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJS4YTqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PK0vcZCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03HPdUO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3XhArB2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hKUIhHeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a2pfcaxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IhUgVA6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TNegYGRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AxiOGpGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QbxbmYXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D9LUgDfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U9cpoiKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I2EugPkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dELwGCxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UqWSZdeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XUYqX7VCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VpdpNjl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nakpXXdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L3RL9aAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DAlG2aJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5ML4Uc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dM8F96A1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BbTAcnwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rkaLPuRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LAGmO6EDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IC76kDGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4fVipG50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nlVjBDnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gfFtheyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kdhSqAlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wc4nt6d0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4eZ7ufd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N4KcZUjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hc9cEZb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z3uxed3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TFeHiLn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9aAmksSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EljuK8rgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8mNguyj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NaJ3VraaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jI4eAveCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FbnHAzEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0EKqfEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qPjGpgArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bp4cK7enWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SVXYNcQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bMO4h6EcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aPc0s9SnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXkC9Wl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9lndV6v4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UxhEJxEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rf2XIfbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C7Bu8RBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KgYqhFezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2R02JUssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LRYkldotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YIlSq3s8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TykTcB2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8v9V4f6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lubo1ZJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AsWCYHKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wL8RIt33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rLrMjAy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N6VR8qaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wuukhjL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DyPhXVV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6tCkxBpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aj2KEVwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1S3wZCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

