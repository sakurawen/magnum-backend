package req

type NewForm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ReleaseFormFieldConfig struct {
	FieldID         string `json:"field_id"`
	Key             string `json:"key"`
	Text            string `json:"text"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	JSONStringValue string `json:"json_string_value"`
	OrderIndex      int    `json:"order_index"`
}

type ReleaseFormField struct {
	FieldType  string                   `json:"field_type"`
	FieldName  string                   `json:"field_name"`
	OrderIndex int                      `json:"order_index"`
	Config     []ReleaseFormFieldConfig `json:"config"`
}

type ReleaseForm struct {
	FormID string             `json:"form_id"`
	Fields []ReleaseFormField `json:"fields"`
}
