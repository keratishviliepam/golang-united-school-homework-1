package main

import (
        "github.com/kyokomi/emoji/v2"
        "fmt"
)

func main() {
        GetMessage()
}

func GetMessage() {
        worldMessage := emoji.Sprint("Hello :world_map:!")

        fmt.Println(worldMessage)
}
