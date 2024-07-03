CREATE TABLE IF NOT EXISTS products (
	id int NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    description VARCHAR(255),
    price DECIMAL(10, 2),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS product_types (
	id int NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    active BOOLEAN,
    PRIMARY KEY (id)
);

ALTER TABLE products
ADD type_id int,
ADD FOREIGN KEY (type_id) REFERENCES product_types(id);