package mapper

func ToDatabaseFilters(passportNumber, surname, name, patronymic, address string) map[string]string {
	filters := make(map[string]string)

	if passportNumber != "" {
		filters["passport_number"] = passportNumber
	}

	if surname != "" {
		filters["surname"] = surname
	}

	if name != "" {
		filters["name"] = name
	}

	if patronymic != "" {
		filters["patronymic"] = patronymic
	}

	if address != "" {
		filters["address"] = address
	}

	return filters
}
