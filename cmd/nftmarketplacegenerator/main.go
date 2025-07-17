// cmd/nftmarketplacegenerator/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacegenerator/internal/nftmarketplacegenerator"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacegenerator.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
