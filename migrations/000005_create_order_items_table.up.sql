CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    product_image TEXT,
    variant_name VARCHAR(255),
    portion_size VARCHAR(255),
    quantity INT NOT NULL,
    unit_base_price INT NOT NULL,
    variant_price INT DEFAULT 0,
    portion_price INT DEFAULT 0,
    unit_final_price INT NOT NULL,
    total_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,

    CONSTRAINT fk_order
        FOREIGN KEY(order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE
);