-- Drop table
DROP TABLE IF EXISTS `order_items`;
DROP TABLE IF EXISTS `order_details`;
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `products`;
DROP TABLE IF EXISTS `categories`;
-- Create table
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
  `address` TEXT NOT NULL,
  `role` TINYINT(1) DEFAULT 0
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
  `active` TINYINT(1) DEFAULT 0,
  `send_email` TINYINT(1) DEFAULT 0
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
    "Vintage Luxury Chair",
    "Đối với những tín đồ của phong cách bán cổ điển thì chiếc ghế armchair Polo này rất đáng để thử. Nhờ vào thiết kế với những đường cong nhẹ nhàng, vừa đủ và phần vải bọc trung tính, armchair có thể đi kèm với rất nhiều phong cách và không gian khác nhau, từ phòng khách đến phòng ngủ và cả phòng đọc sách.",
    890000,
    1,
    "static/image/armchair_1.jpg",
    10,
    1
  ),
  (
    "Armchair Tudor Pink",
    "Nhờ vào thiết kế tôn vinh đường cong tuyệt mỹ, các điểm nhấn hình kim cương ở phần tựa lưng và đệm ngồi cùng với chất liệu vải nhung cao cấp, Armchair Tudor mang lại cảm giác sang trọng và đẳng cấp. Đặc biệt khi được kết hợp với không gian mở, phá cách thì chiếc ghế bành này toát lên sự thư giãn và tận hưởng. Còn trong không gian phòng ngủ hoặc phòng đọc sách, chiếc ghế bành này còn tạo ra sự thoải mái vô cùng.",
    1900000,
    1,
    "static/image/armchair_2.jpg",
    0,
    0
  ),
  (
    "Armchair Maxine",
    "Armchair Maxine với thiết kế tinh tế, trang nhã và nhỏ gọn phù hợp với căn hộ có diện tích vừa phải. Phần lưng tựa bằng da bò kết hợp gỗ Xà cừ màu nâu mang lại cảm giác ấm cúng cho không gian phòng khách. Maxine – Nét nâu trầm sang trọng Maxine, mang thiết kế vượt thời gian, gửi gắm và gói gọn lại những nét đẹp của thiên nhiên và con người nhưng vẫn đầy tính ứng dụng cao trong suốt hành trình phiêu lưu của nhà thiết kế người Pháp Dominique Moal. Bộ sưu tập nổi bật với màu sắc nâu trầm sang trọng, là sự kết hợp giữa gỗ, da bò và kim loại vàng bóng. Đặc biệt trên mỗi sản phẩm, những nét bo viền, chi tiết kết nối được sử dụng khéo léo tạo ra điểm nhất rất riêng cho bộ sưu tập. Maxine thể hiện nét trầm tư, thư giãn để tận hưởng không gian sống trong nhịp sống hiện đại. Sản phẩm thuộc BST Maxine được sản xuất tại Việt Nam.",
    1390000,
    1,
    "static/image/armchair_3.jpg",
    10,
    1
  ),
  (
    "Armchair Vải Santos",
    "Kích thước nhỏ gọn phù hợp cho các căn hộ diện tích hẹp - nệm ngồi êm , chức năng xoay nhẹ nhàng nhờ chân khung kim loại , vải bọc màu trung tính dễ dàng phối với các mẫu sofa .",
    1740000,
    1,
    "static/image/armchair_4.jpg",
    15,
    0
  ),
  (
    "Armchair Hantz",
    "Armchair Hantz có phần chân kim loại sơn đen và phần đệm ngồi bọc vải cao cấp, mang lại cảm giác nhẹ nhàng nhưng lại rất vừa vặn và thoải mài. Armchair phù hợp với những không gian nội thất hiện đại và chuộng sắc màu. Nhờ thiết kế nhỏ gọn, sản phẩm có thể được dùng trong cả phòng khách lẫn phòng ngủ.",
    490000,
    1,
    "static/image/armchair_5.jpg",
    35,
    1
  ),
  (
    "Armchair Garbo",
    "Armchair Garbo là sự kết hợp mới mẻ giữa phần chân kim loại sơn đen với phần đệm ngồi bọc vải cao cấp, mang lại cảm giác thanh mảnh nhưng rất chắc chắn. Armchair phù hợp với những không gian phòng khách hiện đại và phong cách. Với những gam màu sáng, sản phẩm sẽ là điểm nhấn cho nội thất nhà bạn.",
    1090000,
    1,
    "static/image/armchair_6.jpg",
    20,
    1
  ),
  (
    "Armchair Tudor Pink",
    "Coco là chiếc ghế được thiết kế hình cong phễu ở phần lưng với khung thép kim loại dầy 20mm, phần đệm ngồi và phần lưng tựa được bọc hoàn toàn bởi chất liệu là nhung và da cho từng phiên bản. Calligaris đã ấn định màu sắc riêng biệt của mình vượt mọi thời đại với biểu tượng của thập niên 50. Cá tính nổi bật bởi phần khung kim loại sáng bóng và phần ghế ngồi hình vỏ xò hết sức lôi cuốn thì chiếc ghế này cũng có thể phối với nhiều phần vật liệu cũng như màu sắc khác nhau",
    2900000,
    1,
    "static/image/armchair_7.jpg",
    50,
    0
  ),
  (
    "Armchair String",
    "Chiếc ghế Armchair được thiết kế đặc biệt với điểm nhấn là phần tay vịn được đan dây mang lại phong cách Retro cho phòng khách nhà bạn.",
    1650000,
    1,
    "static/image/armchair_8.jpg",
    90,
    1
  ),
  (
    "Sofa Miami",
    "Sofa Miami 2 chỗ là một thiết kế tối giản cho không gian phòng khách hiện đại, chất liệu sofa vải cao cấp trên tông màu xám nhạt rất dễ dàng phối hợp cùng các sản phẩm khác. Kích thước nhỏ gọn 2 chỗ sẽ phù hợp cho các căn hộ, hoặc những góc nhỏ trong ngôi nhà của bạn.",
    1250000,
    2,
    "static/image/sofa_1.jpg",
    12,
    0
  ),
  (
    "Sofa Lovely",
    "Sofa Lovely 3 chỗ với thiết kế đường nét thanh mảnh nhẹ nhàng. Chất liệu vải cao cấp, màu sắc tươi mới. Sofa Lovely là lựa chọn tối ưu cho phòng khách căn hộ có diện tích nhỏ.",
    1600000,
    2,
    "static/image/sofa_2.jpg",
    60,
    1
  ),
  (
    "Sofa John",
    "Sofa John 3 chỗ với thiết kế mạnh mẽ - khỏe khoắn với hình khối vuông cạnh là điểm cộng của mẫu sofa John. Phần nệm ngồi rộng cho phép nằm thư giãn hoặc một giấc ngủ sâu sau một ngày bận rộn. Sofa có phần phối vải kẻ sọc cho sofa hình khối không những không bị cứng mà còn là điểm nhấn riêng biệt. Tại Nhà Xinh có đa dạng các mẫu sofa đẹp hiện đại, đa dạng kiểu dáng, phù hợp cho từng không gian nhà bạn.",
    2590000,
    2,
    "static/image/sofa_3.jpg",
    5,
    0
  ),
  (
    "Sofa Twoback",
    "Sofa Twoback 2.5 chỗ với gam màu xanh tươi mát, nhã nhặn. Kết cấu khung làm từ gỗ thông của Newzerland, bọc nệm vải cao cấp tạo cảm giác thoải mái. Sofa Twoback là 1 lựa chọn tối ưu cho không gian phòng khách hiện đại.",
    3180000,
    2,
    "static/image/sofa_4.jpg",
    0,
    0
  ),
  (
    "Sofa Dubai",
    "Sofa Dubai 2.5 chỗ với đường nét mỏng đảm bảo cái nhìn nhẹ nhàng và thanh thoát. Thiết kế sofa 2 chỗ nhưng vẫn mang lại cảm giác chỗ ngồi rộng hơn. Sofa Dubai có 2 màu nâu và kem để chọn lựa phù hợp cho không gian phòng khách hiện đại của gia đình bạn.",
    16500000,
    2,
    "static/image/sofa_5.jpg",
    0,
    0
  ),
  (
    "Sofa Bridge 3 chỗ",
    "Sofa Bridge 3 chỗ với thiết kế vượt thời gian, sử dụng chất liệu gỗ sồi đặc và da bò tự nhiên, sofa Bridge là điểm nhấn đẳng cấp trong phòng khách của bạn. Đặc biệt, phần tay vịn được hoàn thiện vô cùng tinh xảo kết hợp với kết cấu vô cùng chắn chắn giúp cho sofa Bridge tạo cảm xúc gần gũi, tự nhiên và thoái mái khi sử dụng. Sản phẩm có các màu sắc hoàn thiện gỗ sồi sáng và trầm và nhiều màu da khác nhau để lựa chọn. Sofa Bridge 3 chỗ là 1 lựa chọn sáng gia cho phong cách nội thất Bắc Âu.",
    8400000,
    2,
    "static/image/sofa_6.jpg",
    12,
    0
  ),
  (
    "Bàn nước Pio",
    "Bàn nước PIO thu hút ánh nhìn với mặt bàn bằng chất liệu hiện đại melamine marble. Thiết kế với kiểu dáng oval giúp, bàn nước Pio tạo điểm nhấn khác biệt cho không gian phòng khách của bạn. Việc bổ sung ngăn bên dưới là khu chứa đồ cũng như trưng bày các vật dụng trang trí. PIO – Vẻ đẹp yên bình giữa lối sống đô thị Pha trộn giữa phong cách scandinavian và sự mới lạ trong chọn lựa màu sắc, bộ sưu tập PIO toát lên vẻ đẹp nhẹ nhàng, thanh lịch và cũng rất gần gũi với thiên nhiên. PIO thể hiện lối sống của những người trẻ, biết chiêm nghiệm và thưởng ngoạn sự trở về bình yên giữa nhịp sống hiện đại. Thiết kế bởi những đường cong, điểm xuyến các chi tiết nhấn nhá bên cạnh sử dụng các vật liệu như gỗ beech, melamine marble.. giúp PIO trở nên cá tính và thu hút trong không gian hiện đại. Sản phẩm được thiết kế bởi đội ngũ Nhà Xinh và sản xuất tại Việt Nam.",
    6900000,
    3,
    "static/image/table_1.jpg",
    20,
    0
  ),
  (
    "Bàn nước Jazz",
    "Bàn nước Jazz được ghép từ những thanh gỗ sồi già tự nhiên. Bề mặt đặc trưng với những đường nứt tét gỗ tự nhiên được xử lý khéo léo, kết hợp với chân sắt sơn tĩnh điện đầy mạnh mẽ sẽ mang lại nét cá tính độc đáo cho gia chủ.",
    9500000,
    3,
    "static/image/table_2.jpg",
    20,
    0
  ),
  (
    "Bàn nước Mây",
    "Một chiếc bàn nước kết hợp nhịp nhàng bởi 2 khối hình khác nhau về độ cao. Bàn nước Mây giúp cho không gian phòng khách trở nên cá tính hơn. Sản phẩm sử dụng chất liệu đã marble cho phần mặt bàn, được bao quanh bởi kết cấu khung gỗ và nhấn nhá với phần chân kim loại đồng hiện đại.",
    10900000,
    3,
    "static/image/table_3.jpg",
    40,
    0
  ),(
    "Bàn nước Daylight",
    "Bàn nước Daylight với thiết kế hiện đại, kết hợp giữa mặt đá và chân inox mạ sang trọng sẽ là điểm nhấn độc đáo cho phòng khách nhà bạn.",
    12500000,
    3,
    "static/image/table_4.jpg",
    30,
    0
  );
-- table users
INSERT INTO
  `users` (username, password, email, address)
VALUES
  ("user1", "1234", "user1@gmail.com", "hà nội");
-- table order details
INSERT INTO
  order_details (
    `user_id`,
    `total_price`,
    `payment`,
    `discount`
  )
VALUES
  (1, 1234, "Tiền mặt", 12);
-- table order items
INSERT INTO
  order_items (`order_id`, `product_id`, `quantity`)
VALUES
  (1, 2, 12);