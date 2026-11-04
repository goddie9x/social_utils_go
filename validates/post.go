package validates

import (
	"github.com/goddie9x/social_utils_go/constants"
	"github.com/goddie9x/social_utils_go/middlewares"
)

func CanModifyTarget(currentUser middlewares.UserAuth, targetOwnerId string) bool {
	return currentUser.Role != constants.User || targetOwnerId == currentUser.UserId
}
