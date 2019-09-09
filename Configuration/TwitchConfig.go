package configuration

type TwitchConfig struct {
	StreamerName      string `json:"streamer_name"`       //The name of the streamer to track activity data for.
	ClientID          string `json:"client_id"`           //The client ID of the application
	WatchInterval     int    `json:"watch_interval"`      //The time in minutes a viewer should be watching to get a reward.
	WatchRewardAmount int    `json:"watch_reward_amount"` //The amount of coins user gets per WatchInterval.
}
