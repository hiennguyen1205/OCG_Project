-- Drop table
DROP TABLE IF EXISTS `address`;
DROP TABLE IF EXISTS `order_items`;
DROP TABLE IF EXISTS `order_details`;
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `products`;
DROP TABLE IF EXISTS `categories`;
-- Create table
CREATE TABLE `categories` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL
);
CREATE TABLE `products` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `price` INT NOT NULL,
  `category_id` INT,
  `image` VARCHAR(255) NOT NULL,
  `sale` INT DEFAULT 0,
  `is_feature` TINYINT(1) DEFAULT 0
);
CREATE TABLE `users` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `password` TEXT NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `role` TINYINT(1) DEFAULT 0
);
CREATE TABLE `address` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `detail` NVARCHAR(255) NOT NULL
);
CREATE TABLE `order_details` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `total_price` FLOAT NOT NULL DEFAULT 0,
  `payment` VARCHAR(100) NOT NULL,
  `discount` TINYINT(100) NOT NULL
);
CREATE TABLE `order_items` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `order_id` INT NOT NULL,
  `product_id` INT NOT NULL,
  `quantity` INT NOT NULL,
`active` TINYINT(1) DEFAULT 0
);
-- ADD FOREIGN KEY
-- table products
ALTER TABLE
  `products`
ADD
  FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);
-- table order_items
ALTER TABLE
  `order_items`
ADD
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
ALTER TABLE
  `order_items`
ADD
  FOREIGN KEY (`order_id`) REFERENCES `order_details` (`id`);
-- table order_details
ALTER TABLE
  `order_details`
ADD
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
-- table address
ALTER TABLE
  `address`
ADD
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
-- table
  -- insert data
  -- table categories
INSERT INTO
  `categories` (name)
VALUES
  ("Armchair"),
  ("Sofa"),
  ("Table"),
  ("Tủ Tivi"),
  ("Kệ trưng bày"),
  ("Ghế thư giãn"),
  ("Bàn ăn");
-- table products
INSERT INTO
  `products` (
    `name`,
    `description`,
    price,
    category_id,
    image,
    sale,
    is_feature
  )
VALUES
  (
    "ABC",
    "được vl bạn êi",
    101,
    1,
    "static/image/armchair_1.jpg",
    10,
    1
  ),
  (
    "XYZ",
    "mua làm cc gì",
    103,
    1,
    "static/image/armchair_2.jpg",
    0,
    0
  ),
  (
    "HIHI",
    "Hàng như loz vậy. Nhưng t thích loz, nên t thích hàng",
    150,
    1,
    "static/image/armchair_3.jpg",
    10,
    1
  ),
  (
    "AAHQ",
    "ádsdsad",
    1067,
    1,
    "static/image/armchair_4.jpg",
    15,
    0
  ),
  (
    "ADRF",
    "ádasdsadsad",
    10890,
    1,
    "static/image/armchair_5.jpg",
    35,
    1
  ),
  (
    "ACBHT",
    "sadsadsadsa",
    10000,
    1,
    "static/image/armchair_6.jpg",
    20,
    1
  ),
  (
    "CXS",
    "ffjgjhjkjh",
    1076,
    2,
    "static/image/armchair_7.jpg",
    50,
    0
  ),
  (
    "XYH",
    "khkjhjhj",
    1550,
    1,
    "static/image/armchair_8.jpg",
    90,
    1
  ),
  (
    "QEVB",
    "hjhgjghjghj",
    130,
    3,
    "static/image/armchair_1.jpg",
    12,
    0
  ),
  (
    "HRAB",
    "ghjhgjhgj",
    1220,
    2,
    "static/image/table_3.jpg",
    60,
    1
  );
-- table users
INSERT INTO
  `users` (username, password, email)
VALUES
  ("user1", "1234", "user1@gmail.com"),
  ("user2", "abc", "user2@gmail.com"),
  ("user3", "xyz", "user3@gmail.com");
-- table address
INSERT INTO
  `address` (user_id, detail)
VALUES
  (1, "Hà Nội, Việt Nam"),
  (2, "Ở đâu còn lâu mới nói");

-- table order details
INSERT INTO
  order_details (
    `user_id`,
    `total_price`,
    `payment`,
    `discount`
  )
VALUES
  (1, 1234, "Tiền mặt", 12),
  (2, 3456, "Thẻ Master Card", 40);


-- table order items
INSERT INTO
  order_items (`order_id`, `product_id`, `quantity`)
VALUES
  (1, 2, 12),
  (2, 6, 40);
  
-- test query
SELECT * FROM products
SELECT
  *
FROM
  products
  WHERE
  id > 0
  AND name LIKE "%A%"
ORDER BY
  price ASC
  LIMIT 4;
SELECT * FROM categories

SELECT * FROM products WHERE id > 0 AND name LIKE "A" LIMIT 6



SELECT * FROM products WHERE id > 0 AND is_feature = 1 LIMIT 6

  SELECT od.user_id, od.id, od.discount, od.total_price, od.payment,oi.`active`, oi.product_id, oi.quantity, p.`name`, p.price, p.image FROM  order_items oi JOIN order_details od ON od.id = oi.order_id JOIN products p ON p.id = oi.product_id WHERE od.user_id = 1 AND oi.`active` = 0


SELECT od.user_id, od.discount, od.total_price, od.payment FROM order_items oi JOIN order_details od ON od.id = oi.order_id WHERE od.user_id = 1

SELECT od.user_id, od.id, od.discount, od.total_price, od.payment,oi.`active`, oi.product_id, oi.quantity, p.`name`, p.price, p.image FROM order_items oi JOIN order_details od ON od.id = oi.order_id JOIN products p ON p.id = oi.product_id WHERE od.user_id = 1 AND oi.`active` = 0