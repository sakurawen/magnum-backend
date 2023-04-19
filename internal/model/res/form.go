package res

import "magnum/ent"

type FormFieldsAndConfigs struct {
	Fields  []*ent.FormField       `json:"fields"`
	Configs []*ent.FormFieldConfig `json:"configs"`
}
type FormTemplate struct {
	Form     *ent.Form             `json:"form"`
	Template *FormFieldsAndConfigs `json:"template"`
}
