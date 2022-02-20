package solution

import (
        "github.com/kyokomi/emoji/v2"
)


func GetMessage() string {
        worldMessage := emoji.Sprint("Hello :world_map:!")
	return worldMessage
}
