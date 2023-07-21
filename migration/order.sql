CREATE TABLE IF NOT EXISTS orders (
    id int AUTO_INCREMENT PRIMARY KEY,
    widget_id int,
    transaction_id int,
    status_id int,
    customer_id int,
    quantity int,
    amount int,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
)

GO

ALTER TABLE orders
    ADD COLUMN customer_id INT AFTER status_id

GO

ALTER TABLE orders
    CHANGE COLUMN customer_id customer_id INT AFTER status_id

GO

ALTER TABLE orders
    ADD FOREIGN KEY (customer_id) REFERENCES customers(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE