package service

import (
	"context"
	"magnum/ent"
	"magnum/ent/form"
	"magnum/internal/app"
	"magnum/internal/model/req"
	"time"
)

type Form struct {
}

func (Form) New(nf *req.NewForm, userId string) (*ent.Form, error) {
	return app.DB.Form.Create().SetTitle(nf.Title).SetDescription(nf.Description).SetUserID(userId).SetCreateAt(time.Now()).SetUpdateAt(time.Now()).Save(context.Background())
}

func (Form) ListAll(uid string) ([]*ent.Form, error) {
	return app.DB.Form.Query().Where(form.UserID(uid)).Order(ent.Desc("create_at")).All(context.Background())
}
