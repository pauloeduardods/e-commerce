DROP SCHEMA IF EXISTS Go_Ecommerce;
CREATE SCHEMA Go_Ecommerce;
CREATE TABLE Go_Ecommerce.Users (
  id INTEGER AUTO_INCREMENT PRIMARY KEY NOT NULL,
  username TEXT NOT NULL,
  classe TEXT NOT NULL,
  level INTEGER NOT NULL,
  password TEXT NOT NULL
);
CREATE TABLE Go_Ecommerce.Orders (
  id INTEGER AUTO_INCREMENT PRIMARY KEY NOT NULL,
  userId INTEGER,
  FOREIGN KEY (userId) REFERENCES Go_Ecommerce.Users (id)
);
CREATE TABLE Go_Ecommerce.Products (
  id INTEGER AUTO_INCREMENT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  quantity TEXT NOT NULL,
  price DECIMAL(10, 2) NOT NULL
);

INSERT INTO Go_Ecommerce.Products (name, quantity, price) VALUES ('Product 1', '10', '10.00');
INSERT INTO Go_Ecommerce.Products (name, quantity, price) VALUES ('Product 2', '20', '20.00');
INSERT INTO Go_Ecommerce.Products (name, quantity, price) VALUES ('Product 3', '30', '30.00');
INSERT INTO Go_Ecommerce.Products (name, quantity, price) VALUES ('Product 4', '40', '40.00');
