package service

import (
	"context"
	"magnum/ent"
	"magnum/ent/form"
	"magnum/ent/formfield"
	"magnum/ent/formfieldconfig"
	"magnum/internal/app"
	"magnum/internal/model/req"
	"magnum/internal/model/res"
	"time"
)

type Form struct {
}

func (Form) New(nf *req.NewForm, userId string) (*ent.Form, error) {
	return app.DB.Form.Create().SetTitle(nf.Title).SetDescription(nf.Description).SetUserID(userId).SetCreateAt(time.Now()).SetUpdateAt(time.Now()).Save(context.Background())
}

func (Form) ListAllByUid(uid string) ([]*ent.Form, error) {
	return app.DB.Form.Query().Where(form.UserID(uid), form.Disabled(0)).Order(ent.Desc("create_at")).All(context.Background())
}

func (Form) FindById(id string) (*ent.Form, error) {
	return app.DB.Form.Query().Where(form.ID(id)).Only(context.Background())
}

func (Form) FindFormTemplateById(id string) (*res.FormFieldsAndConfigs, error) {
	ctx := context.Background()
	fileds, err := app.DB.FormField.Query().
		Where(formfield.FormID(id), formfield.Disabled(0)).
		Order(ent.Asc("order_index")).
		All(ctx)
	if err != nil {
		return nil, err
	}

	configs, err := app.DB.FormFieldConfig.Query().
		Where(formfieldconfig.FormID(id), formfieldconfig.Disabled(0)).
		Order(ent.Asc("order_index")).
		All(ctx)
	if err != nil {
		return nil, err
	}

	temp := &res.FormFieldsAndConfigs{
		Configs: configs,
		Fields:  fileds,
	}
	return temp, nil
}

func (Form) QueryRelease(id string) (bool, error) {
	isRelease := true
	count, err := app.DB.Form.Query().Where(form.ID(id), form.IsRelease(1)).Count(context.Background())
	if err != nil {
		return false, err
	}
	if count == 0 {
		isRelease = false
	}
	return isRelease, nil
}

type ReleaseResult struct {
	Fields  []*ent.FormField       `json:"fields"`
	Configs []*ent.FormFieldConfig `json:"configs"`
}

func saveFieldAndConfig(data *req.ReleaseForm) (*ReleaseResult, error) {
	ctx := context.Background()
	tx, err := app.DB.Tx(ctx)
	ret, err := createFieldsAndConfig(ctx, tx, data)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return ret, nil
}

func createFieldsAndConfig(ctx context.Context, tx *ent.Tx, data *req.ReleaseForm) (*ReleaseResult, error) {
	var configs []*ent.FormFieldConfig
	var fields []*ent.FormField
	for i := 0; i < len(data.Fields); i++ {
		f := data.Fields[i]
		field, err := tx.FormField.Create().
			SetFormID(data.FormID).
			SetFieldType(f.FieldType).
			SetFieldName(f.FieldName).
			SetOrderIndex(f.OrderIndex).
			SetCreateAt(time.Now()).
			SetUpdateAt(time.Now()).Save(ctx)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)

		for ci := 0; ci < len(f.Config); ci++ {
			c := f.Config[ci]
			config, err := tx.FormFieldConfig.Create().
				SetFormID(data.FormID).
				SetOrderIndex(c.OrderIndex).
				SetFieldID(field.ID).
				SetJSONStringValue(c.JSONStringValue).
				SetValue(c.Value).
				SetKey(c.Key).
				SetType(c.Type).
				SetText(c.Text).
				Save(ctx)

			if err != nil {
				return nil, err
			}
			configs = append(configs, config)
		}
	}
	ret := &ReleaseResult{
		Fields:  fields,
		Configs: configs,
	}
	return ret, nil
}

func updateFieldAndConfig(data *req.ReleaseForm) (*ReleaseResult, error) {
	ctx := context.Background()
	tx, err := app.DB.Tx(ctx)
	if err != nil {
		return nil, err
	}
	_, err = tx.FormField.Update().Where(formfield.FormID(data.FormID), formfield.Disabled(0)).SetDisabled(1).SetUpdateAt(time.Now()).Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	ret, err := createFieldsAndConfig(ctx, tx, data)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return ret, nil
}

func (Form) UpdateForm(data *req.ReleaseForm) (*ReleaseResult, error) {
	return updateFieldAndConfig(data)
}

func (Form) ReleaseForm(data *req.ReleaseForm) (*ReleaseResult, error) {
	_, err := app.DB.Form.Update().Where(form.ID(data.FormID)).SetIsRelease(1).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return saveFieldAndConfig(data)
}
