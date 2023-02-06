package GetValidatedByTagAndUser

import (
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
)

func HandleGetValidatedByTagAndUser(tagQuery GetValidatedByTagAndUser) *model.PlayerHasValidateTag {
	tag, err := repository.FindByPlayerAndTagId(tagQuery.PlayerId, tagQuery.TagId)

	if err != nil {
		return nil
	}

	return tag
}
