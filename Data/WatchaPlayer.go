package Data

import (
	"github.com/jinzhu/gorm"
)

//WatchaPlayer is the data model for a player
type WatchaPlayer struct {
	gorm.Model
	DiscordID          string
	DiscordDisplayName string
	TwitchID           string
	TwitchDisplayName  string
	WatchaCoins        int
}
