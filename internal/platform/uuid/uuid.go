package uuid

import "github.com/google/uuid"

func NewID() string {
	id := uuid.New()

	return id.String()
}

func IsValidUUID(id string) error {
	_, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	return nil
}
