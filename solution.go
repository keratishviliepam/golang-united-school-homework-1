package solution

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	worldMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(worldMessage)
}
