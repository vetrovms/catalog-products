CREATE TABLE IF NOT EXISTS products(
   id serial PRIMARY KEY,
   title VARCHAR (255) NOT NULL,
   price DECIMAL NOT NULL,
   color VARCHAR (255),
   currency VARCHAR (255) NOT NULL,
   best_before TIMESTAMP,
   manufacturer VARCHAR (255),
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   deleted_at TIMESTAMP
);

CREATE INDEX ix_products_title ON products (title);
CREATE INDEX ix_products_price ON products (price);
CREATE INDEX ix_products_color ON products (color);
CREATE INDEX ix_products_best_before ON products (best_before);
CREATE INDEX ix_products_manufacturer ON products (manufacturer);