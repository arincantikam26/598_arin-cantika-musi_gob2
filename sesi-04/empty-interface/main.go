package main

func main() {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Airell", "Nanda"}

	//type assertion
	var v interface{}
	v = 20
	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	// map & slice with wmpty interface
	rs := []interface{}{1, "Airell", true, "Ananda", true}

	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}
	_, _ = rs, rm
}
