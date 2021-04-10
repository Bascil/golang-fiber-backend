-- Adminer 4.7.8 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` longtext,
  `last_name` longtext,
  `email` longtext,
  `updated_at` longtext,
  `created_at` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `orders` (`id`, `first_name`, `last_name`, `email`, `updated_at`, `created_at`) VALUES
(1,	'Basil',	'Ndonga',	'basilndonga@gmail.com',	'2021-02-06 18:50',	'2021-02-06 18:50'),
(2,	'Josphine',	'Nzila',	'test@example.com',	'2021-02-08 18:50',	'2021-02-08 18:50');

DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) unsigned DEFAULT NULL,
  `product_title` longtext,
  `price` float DEFAULT NULL,
  `quantity` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orders_order_items` (`order_id`),
  CONSTRAINT `fk_orders_order_items` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `order_items` (`id`, `order_id`, `product_title`, `price`, `quantity`) VALUES
(1,	1,	'Cookies',	39,	2),
(2,	2,	'Burgers',	30,	15);

DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `permissions` (`id`, `name`) VALUES
(1,	'view_users'),
(2,	'edit_users'),
(3,	'view_roles'),
(4,	'edit_roles'),
(5,	'view_products'),
(6,	'edit_products'),
(7,	'view_orders'),
(8,	'edit_orders');

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` longtext,
  `description` longtext,
  `image` longtext,
  `price` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `products` (`id`, `title`, `description`, `image`, `price`) VALUES
(2,	'test',	'test description',	'http://localhost:5000/api/uploads/2021-02-1016:42:28.jpg',	20),
(3,	'test 2',	'test desc',	'http://localhost:5000/api/uploads/2021-03-2312:01:02.jpg',	30),
(4,	'Pos Device',	'Point of sale',	'http://localhost:5000/api/uploads/25580322c9c531a833b1c11723b78bb.jpg',	200);

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `roles` (`id`, `name`) VALUES
(1,	'admin'),
(2,	'staff'),
(5,	'clerk'),
(6,	'customer');

DROP TABLE IF EXISTS `role_permissions`;
CREATE TABLE `role_permissions` (
  `role_id` bigint(20) unsigned NOT NULL,
  `permission_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`),
  KEY `fk_role_permissions_permission` (`permission_id`),
  CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`),
  CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `role_permissions` (`role_id`, `permission_id`) VALUES
(1,	1),
(5,	1),
(6,	1),
(5,	3),
(6,	3);

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `firstname` longtext,
  `lastname` longtext,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext,
  `role_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `fk_users_role` (`role_id`),
  CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `users` (`id`, `firstname`, `lastname`, `email`, `password`, `role_id`) VALUES
(1,	'Johns',	'Magufuli',	'test@example.com',	'$2a$14$Lk6DMTx.n7Nfs37cV4qseOwaQ9R6MlV0wyGNVKSgoH7vGyZNZdmJ6',	1),
(2,	'Basil',	'Ndonga',	'basil@example.com',	'$2a$14$hIcPbzXgWKDTerGAzwBHTuIdzJYQmE/V2OBQ1VqvN1dn463DCoFG.',	2),
(8,	'Stone',	'Man',	'st0ne@example.com',	'$2a$14$/p7/201GBF0YO7JQz44m/eWswtp6.2q5F2ynKGGg.HKyfvu8P7H1i',	2),
(12,	'Stone',	'Man',	'st0nee@example.com',	'$2a$14$FXw/KECNRgKSdTghzBRBzuIvHaoCwH.90/maTfp6194hDxt9SXy9y',	2),
(13,	'Anzila',	'People',	'anzila@example.com',	'$2a$14$CHmPWDHVaBE6Q/tNdHUpo.Ln/C/iMzc7dBMalfcp7Askwf4xBnZNS',	2),
(15,	'Extreme',	'Rules',	'extreme@example.com',	'$2a$14$5hmWfwc4fChb04B2TrwiOeRIq6TwwIv5wIB//zXo/ioRo4Nq.O1sO',	2);

-- 2021-04-10 20:07:00
