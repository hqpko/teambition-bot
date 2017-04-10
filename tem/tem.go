package tem

import "github.com/hqpko/teambition-bot/teambition"

var teambitionAPI *teambition.TeambitionAPI

//InitTeambitionAPI _
func InitTeambitionAPI(appKey, appSecret, redirectURL string) {
	teambitionAPI = teambition.CreateTeambitionAPI(appKey, appSecret, redirectURL)
}

var (
	temHost string
	temPort string
)
