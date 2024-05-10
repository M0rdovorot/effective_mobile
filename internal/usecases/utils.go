package usecases

func validateCarMap(carMap map[string]any) (map[string]any, error) {
	filter := map[string]any{}
	
	if regNum, ok := carMap["regNum"]; ok{
		filter["regnum"] = regNum
	}
	if mark, ok := carMap["mark"]; ok{
		filter["mark"] = mark
	}
	if model, ok := carMap["model"]; ok{
		filter["model"] = model
	}
	if _, ok := carMap["year"]; ok{
		year, ok := carMap["year"].(int)
		if !ok {
			return filter, ErrBadID
		}
		filter["year"] = year
	}
	if name, ok := carMap["name"]; ok{
		filter["name"] = name
	}
	if surname, ok := carMap["surname"]; ok{
		filter["surname"] = surname
	}
	if patronymic, ok := carMap["patronymic"]; ok{
		filter["patronymic"] = patronymic
	}
	return filter, nil
}