package entities

func GetAllModels() []interface{} {
	models := []interface{}{
		&Todo{},
	}

	return models
}
