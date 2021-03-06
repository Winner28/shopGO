CREATE TABLE users
(
  id INT SERIAL PRIMARY KEY,
  firstname varchar(255),
  lastname varchar(255),
  passwordHash varchar(60),
  email varchar(255), 
);

CREATE TABLE products
(
	id INT SERIAL PRIMARY KEY,
	name varchar(255),
	price decimal,
	description varchar(255)
);

CREATE TABLE categories
(
	id SERIAL NOT NULL,
	name varchar(255)
);

CREATE TABLE product_categories 
(
	id INT SERIAL PRIMARY KEY NOT NULL , (serial now???)
	product_id INT NOT NULL,
	category_id INT NOT NULL,
	CONSTRAINT product_category_product_id_fk FOREIGN KEY(product_id) REFERENCES product (id)
	ON DELETE CASCADE
	CONSTRAINT product_category_category_id_fk FOREIGN KEY(category_id) REFERENCES categories (id)
	ON DELETE CASCADE
);

CREATE TABLE orders
(
	id INT SERIAL PRIMARY KEY NOT NULL,
	user_id INT,
	amount INT,
	date TIMESTAMP,
	CONSTRAINT orders_user_id_fk FOREIGN KEY(user_id) REFERENCES users (id)
	ON DELETE CASCADE
);

CREATE TABLE product_orders 
(
	id SERIAL(!!!) NOT NULL,
	order_id INT NOT NULL, 	
	product_id INT NOT NULL, 
	CONSTRAINT product_orders_order_id_fk FOREIGN KEY(order_id) REFERENCES orders (id)
	ON DELETE CASCADE
	CONSTRAINT product_orders_product_id_fk FOREIGN KEY(product_id) REFERENCES product (id)
	ON DELETE CASCADE
);

CREATE TABLE roles
(
	id SERIAL PRIMARY KEY,
	user_id INT,
	name varchar(20),
	CONSTRAINT role_user_id_fk FOREIGN KEY(user_id) REFERENCES users (id)
	ON DELETE CASCADE
);


