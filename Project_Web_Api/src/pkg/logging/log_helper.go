package logging

func mapToZapParams(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0)

	for k, v := range extra {
		params = append(params, string(k))
		params = append(params, v)
	}
	return params

}

func prepareLogKeys(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}


// ============================== Map For Zero ===============================

func logParamsToZeroParams(keys map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}

	for k, v := range keys {
		params[string(k)] = v
	}

	return params
}