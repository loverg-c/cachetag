package GetTag

import (
	"strconv"
	"tags-finder/Domain/model"
	repository "tags-finder/Infrastructure/Database/Repository"
)

func HandleGetTag(tagQuery GetTag) *model.Tag {
	id, _ := strconv.Atoi(tagQuery.TagId)
	tag, err := repository.FindTagById(id)

	if err != nil {
		return nil
	}

	return tag
}
