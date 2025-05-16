-- Adminer 5.2.1 MySQL 8.4.5 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `rbac`;
CREATE DATABASE `rbac` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `rbac`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `jwt` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `users` (`id`, `username`, `password`, `jwt`, `created_at`, `updated_at`) VALUES
(1,	'admin',	'$2a$10$KMRUJ8PA8f77GsjJ9M4Y1OBn7uFMZ4nyGnNVLt/j6BPD/0fE5Xy7e',	'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50IjoiYWRtaW4iLCJleHAiOjE2MzkxNDU1NTQsImlzcyI6Ik1lIn0.VCp98bVMlPNYfrY9KiDIm7QRu-SfQ9WM7KmFZetNWT8',	'2025-05-11 07:07:26',	'2025-05-11 07:07:26'),
(2,	'jared',	'$2a$10$duiWjUH4WOZkK/OoXO78aOewuzTaI.6yaH42MDvoIG6HDcy3XuCdy',	NULL,	'2025-05-11 07:07:26',	'2025-05-11 07:07:26'),
(3,	'derek',	'$2a$10$HMATJI7/j1TurK7RzfPO8.yxWv9p4XBV1DXPGNhJRPI4IbuivwRHq',	NULL,	'2025-05-11 19:16:45',	'2025-05-11 19:16:45');

-- 2025-05-16 08:57:03 UTC