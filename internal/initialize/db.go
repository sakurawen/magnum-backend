package initialize

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"magnum/ent"
)

func DB() *ent.Client {
	client, err := ent.Open("mysql", "root:wisakura@(localhost:3306)/magnum?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// 运行自动迁移工具。
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
