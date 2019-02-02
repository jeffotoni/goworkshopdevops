// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "time"
)

func writeFile(fSize int64) error {
    fName := `stress` // test file
    defer os.Remove(fName)
    f, err := os.Create(fName)
    if err != nil {
        return err
    }
    const defaultBufSize = 4096
    buf := make([]byte, defaultBufSize)
    buf[len(buf)-1] = '\n'
    w := bufio.NewWriterSize(f, len(buf))

    start := time.Now()
    written := int64(0)
    for i := int64(0); i < fSize; i += int64(len(buf)) {
        nn, err := w.Write(buf)
        written += int64(nn)
        if err != nil {
            return err
        }
    }
    err = w.Flush()
    if err != nil {
        return err
    }
    err = f.Sync()
    if err != nil {
        return err
    }
    since := time.Since(start)

    err = f.Close()
    if err != nil {
        return err
    }
    fmt.Printf("written: %dB %dns %.2fGB %.2fs %.2fMB/s\n",
        written, since,
        float64(written)/1000000000, float64(since)/float64(time.Second),
        (float64(written)/1000000)/(float64(since)/float64(time.Second)),
    )
    return nil
}

var size = flag.Int("size", 8, "file size in GiB")

func main() {
    flag.Parse()
    fSize := int64(*size) * (1024 * 1024 * 1024)
    err := writeFile(fSize)
    if err != nil {
        fmt.Fprintln(os.Stderr, fSize, err)
    }
}
