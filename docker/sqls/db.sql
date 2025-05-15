-- Adminer 4.8.1 MySQL 8.0.25 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `rbac` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `rbac`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `jwt` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `users` (`id`, `username`, `password`, `jwt`, `created_at` ,`updated_at`) VALUES
(1,	'admin',	'21232f297a57a5a743894a0e4a801fc3',	'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50IjoiYWRtaW4iLCJleHAiOjE2MzkxNDU1NTQsImlzcyI6Ik1lIn0.VCp98bVMlPNYfrY9KiDIm7QRu-SfQ9WM7KmFZetNWT8',	'2021-05-11 07:07:26', '2021-05-11 07:07:26'),
(2,	'jared',	'b620e68b3bf4387bf7c43d51bd12910b',	NULL,	'2021-05-11 07:07:26', '2021-05-11 07:07:26'),
(3,	'derek',	'7815696ecbf1c96e6894b779456d330e',	NULL,	'2021-12-11 19:16:45', '2021-05-11 19:16:45');
