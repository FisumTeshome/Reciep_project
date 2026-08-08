package utils

import (
"log"
"os"
)

var Log = log.New(os.Stdout, "[recipe-api] ", log.LstdFlags|log.Lshortfile)
