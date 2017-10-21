CREATE TABLE `azan` (
	`id` INT(10) NOT NULL AUTO_INCREMENT,
	`city` VARCHAR(255) NULL DEFAULT NULL,
	`dt` DATE NULL DEFAULT NULL,
	`fajr` VARCHAR(5) NULL DEFAULT NULL,
	`sunrise` VARCHAR(5) NULL DEFAULT NULL,
	`zuhr` VARCHAR(5) NULL DEFAULT NULL,
	`asr` VARCHAR(5) NULL DEFAULT NULL,
	`maghrib` VARCHAR(5) NULL DEFAULT NULL,
	`isya` VARCHAR(5) NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
);