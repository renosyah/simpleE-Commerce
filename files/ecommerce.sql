
DROP DATABASE IF EXISTS ecommerce CASCADE;

CREATE DATABASE ecommerce;

CREATE TABLE ecommerce.admin (
	id_admin SERIAL,
	name STRING,
	address STRING,
	email STRING,
	username STRING,
	password STRING,
	PRIMARY KEY (id_admin)
);

INSERT INTO ecommerce.admin (name,address,email,username,password) VALUES ('reno','jln.janti','reno@gmail.com','renosyah','12345');

CREATE TABLE ecommerce.product_category (
	id_product_category SERIAL,
	category_name STRING,
	PRIMARY KEY (id_product_category)
);

INSERT INTO ecommerce.product_category (id_product_category,category_name) VALUES (13133332367668,'Food'),(131414230598,'Cloth'),(13234534956759,'Electronic'),(46745674645,'Tool'),(57945734579,'Furniture');

CREATE TABLE ecommerce.product (
	id_product SERIAL,
	id_product_category INT64 NULL REFERENCES ecommerce.product_category (id_product_category),
	product_name STRING,
	stock INT,
	price INT,
	curency STRING,
	PRIMARY KEY (id_product)
);

INSERT INTO ecommerce.product (id_product,id_product_category,product_name,stock,price,curency) VALUES (4857343545453,13133332367668,'burger',41,3000,'IDR'),(48573573957353,131414230598,'baju kaos',45,50000,'IDR'),(48535335211,57945734579,'lemari',45,50000,'IDR');
INSERT INTO ecommerce.product (id_product,id_product_category,product_name,stock,price,curency) VALUES (1121415151331,13133332367668,'nuget',34,2500,'IDR'),(443235688776554,131414230598,'celana jean',98,72000,'IDR'),(7766554411,57945734579,'meja Makan',12,890000,'IDR');

CREATE TABLE ecommerce.product_detail (
	id_product_detail SERIAL,
	id_product INT64 NULL REFERENCES ecommerce.product (id_product),
	description STRING,
	PRIMARY KEY (id_product_detail)
);

INSERT INTO ecommerce.product_detail (id_product,description) VALUES (48535335211,'lemari yg mantap'),(4857343545453,'burger enak'),(48573573957353,'kaos bagus');
INSERT INTO ecommerce.product_detail (id_product,description) VALUES (7766554411,'meja yg menarik dan mantap'),(1121415151331,'nugget pedas dan enak'),(443235688776554,'celana mahal');

CREATE TABLE ecommerce.product_image (
	id_product_image SERIAL,
	id_product INT64 NULL REFERENCES ecommerce.product (id_product),		
	url_image STRING,
	PRIMARY KEY (id_product_image)
);

INSERT INTO ecommerce.product_image (id_product,url_image) VALUES (48535335211,'data/wardrop.jpeg'),(4857343545453,'data/burger.png'),(48573573957353,'data/kaos.jpeg');
INSERT INTO ecommerce.product_image (id_product,url_image) VALUES (7766554411,'data/dinner_table.jpg'),(1121415151331,'data/nugget.jpg'),(443235688776554,'data/jeans.jpeg');

CREATE TABLE ecommerce.order (
	id_order SERIAL,
	order_date TIMESTAMP,
	total INT,
	PRIMARY KEY (id_order)
);

CREATE TABLE ecommerce.detail_order (
	id_detail_order SERIAL,
	id_order INT64 NULL REFERENCES ecommerce.order (id_order),
	id_product INT64 NULL REFERENCES ecommerce.product (id_product),
	quantity INT,
	sub_total INT,
	PRIMARY KEY (id_detail_order)
);

