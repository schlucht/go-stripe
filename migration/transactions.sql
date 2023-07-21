
CREATE TABLE IF NOT EXISTS transactions(
    id INT AUTO_INCREMENT PRIMARY KEY,
    amount INT,
    currency VARCHAR(255),
    last_four VARCHAR(255),
    bank_return_code VARCHAR(255),
    transaction_status_id INT,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)ENGINE=INNODB

ALTER TABLE transactions
    ADD COLUMN expiry_month INT DEFAULT 0 AFTER transaction_status_id,
    ADD COLUMN expiry_year INT DEFAULT 0 AFTER transaction_status_id
