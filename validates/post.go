package validates

import (
	"post_service/pkg/constants"
	"post_service/pkg/middlewares"
)

func CanModifyTarget(currentUser middlewares.UserAuth, targetOwnerId string) bool {
	return currentUser.Role != constants.User || targetOwnerId == currentUser.UserId
}
