package main

import (
        "github.com/kyokomi/emoji/v2"
)

func main() {
        GetMessage()
}

func GetMessage() string {
        worldMessage := emoji.Sprint("Hello :world_map:!")

        return worldMessage
}
