package mysql

import "database/sql"

func Migrate(db *sql.DB) error {
	schema := `
CREATE TABLE IF NOT EXISTS shifts (
  id VARCHAR(36) PRIMARY KEY,
  user_id VARCHAR(36) NOT NULL,
  opened_at DATETIME NOT NULL,
  closed_at DATETIME NULL,
  INDEX idx_shifts_user_id (user_id),
  INDEX idx_shifts_open (user_id, closed_at)
);

CREATE TABLE IF NOT EXISTS orders (
  id VARCHAR(36) PRIMARY KEY,
  user_id VARCHAR(36) NOT NULL,
  shift_id VARCHAR(36) NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_at DATETIME NOT NULL,
  INDEX idx_orders_shift (shift_id)
);

CREATE TABLE IF NOT EXISTS order_items (
  id VARCHAR(36) PRIMARY KEY,
  order_id VARCHAR(36) NOT NULL,
  name VARCHAR(255) NOT NULL,
  price BIGINT NOT NULL,
  quantity INT NOT NULL,
  INDEX idx_items_order (order_id)
);
`
	_, err := db.Exec(schema)
	return err
}
