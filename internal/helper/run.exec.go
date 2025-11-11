package helper

import "gorm.io/gorm"

func RunExecs(db *gorm.DB) {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Exec(`
			DO $$
			BEGIN
				IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_type') THEN
					CREATE TYPE payment_type AS ENUM ('daily', 'weekly', 'monthly');
				END IF;
			END$$;
			`)
}
