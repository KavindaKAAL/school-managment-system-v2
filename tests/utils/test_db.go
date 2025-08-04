package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func ClearDatabase(db *gorm.DB) error {
	var tables []string
	err := db.Raw(`
        SELECT tablename
        FROM pg_tables
        WHERE schemaname = 'public'
    `).Scan(&tables).Error

	if err != nil {
		return err
	}

	db.Exec("SET session_replication_role = replica")

	for _, table := range tables {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
	}

	db.Exec("SET session_replication_role = DEFAULT")

	return nil
}
