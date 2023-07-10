-- MySQL Script generated by MySQL Workbench
-- Wed Jul  6 10:01:09 2022
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering
<<<<<<< HEAD
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
-- -----------------------------------------------------
-- Schema melisprint
-- -----------------------------------------------------
=======

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema melisprint
-- -----------------------------------------------------

>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Schema melisprint
-- -----------------------------------------------------
DROP DATABASE IF EXISTS `melisprint` ;
CREATE SCHEMA IF NOT EXISTS `melisprint` DEFAULT CHARACTER SET utf8 ;
USE `melisprint` ;
<<<<<<< HEAD
=======

>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`countries`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`countries` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `country_name` VARCHAR(255) NOT NULL,
<<<<<<< HEAD
  CONSTRAINT `country_name_UNIQUE` UNIQUE (`country_name`),
  PRIMARY KEY (`id`))
ENGINE = InnoDB;
=======
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`provinces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`provinces` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `province_name` VARCHAR(255) NOT NULL,
  `country_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `country_id_idx` (`country_id` ASC) VISIBLE,
<<<<<<< HEAD
  CONSTRAINT `province_UNIQUE` UNIQUE (`province_name`, `country_id`),
=======
>>>>>>> origin/feature-warehouse-carries-2
  CONSTRAINT `fk_country_provinces`
    FOREIGN KEY (`country_id`)
    REFERENCES `melisprint`.`countries` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`localities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`localities` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `locality_name` VARCHAR(255) NOT NULL,
  `province_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `province_id_idx` (`province_id` ASC) VISIBLE,
<<<<<<< HEAD
  CONSTRAINT `locality_UNIQUE` UNIQUE (`locality_name`, `province_id`),
=======
>>>>>>> origin/feature-warehouse-carries-2
  CONSTRAINT `fk_province_localities`
    FOREIGN KEY (`province_id`)
    REFERENCES `melisprint`.`provinces` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`sellers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`sellers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cid` VARCHAR(255) NOT NULL,
  `company_name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `locality_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `locality_id_idx` (`locality_id` ASC) VISIBLE,
  UNIQUE INDEX `cid_UNIQUE` (`cid` ASC) VISIBLE,
  CONSTRAINT `fk_locality_sellers`
    FOREIGN KEY (`locality_id`)
    REFERENCES `melisprint`.`localities` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`product_types`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`product_types` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`products` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `product_code` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `width` VARCHAR(45) NOT NULL,
  `height` DECIMAL(19,2) NOT NULL,
  `length` DECIMAL(19,2) NOT NULL,
  `net_weight` DECIMAL(19,2) NOT NULL,
  `expiration_rate` DECIMAL(19,2) NOT NULL,
  `recommended_freezing_temperature` DECIMAL(19,2) NOT NULL,
  `freezing_rate` DECIMAL(19,2) NOT NULL,
  `product_type_id` INT NOT NULL,
  `seller_id` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `seller_id_idx` (`seller_id` ASC) VISIBLE,
  INDEX `product_type_id_idx` (`product_type_id` ASC) VISIBLE,
  UNIQUE INDEX `product_code_UNIQUE` (`product_code` ASC) VISIBLE,
  CONSTRAINT `fk_seller_products`
    FOREIGN KEY (`seller_id`)
    REFERENCES `melisprint`.`sellers` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_product_type_products`
    FOREIGN KEY (`product_type_id`)
    REFERENCES `melisprint`.`product_types` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`warehouses`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`warehouses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `warehouse_code` VARCHAR(255) NOT NULL,
<<<<<<< HEAD
  `minimum_capacity` INT NOT NULL,
  `minimum_temperature` DECIMAL(19,2) NOT NULL,
=======
  `minimun_capacity` INT NOT NULL,
  `minimun_temperature` DECIMAL(19,2) NOT NULL,
>>>>>>> origin/feature-warehouse-carries-2
  `locality_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `locality_id_idx` (`locality_id` ASC) VISIBLE,
  UNIQUE INDEX `warehouse_code_UNIQUE` (`warehouse_code` ASC) VISIBLE,
  CONSTRAINT `fk_locality_warehouse`
    FOREIGN KEY (`locality_id`)
    REFERENCES `melisprint`.`localities` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`sections`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`sections` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `section_number` INT NOT NULL,
  `current_temperature` DECIMAL(19,2) NOT NULL,
  `minimum_temperature` DECIMAL(19,2) NOT NULL,
  `current_capacity` INT NOT NULL,
  `minimum_capacity` INT NOT NULL,
  `maximum_capacity` INT NOT NULL,
  `warehouse_id` INT NOT NULL,
  `product_type_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `product_type_id_idx` (`product_type_id` ASC) VISIBLE,
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  UNIQUE INDEX `section_number_UNIQUE` (`section_number` ASC) VISIBLE,
  CONSTRAINT `fk_product_type_sections`
    FOREIGN KEY (`product_type_id`)
    REFERENCES `melisprint`.`product_types` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_warehouse_sections`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `melisprint`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`product_batches`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`product_batches` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `batch_number` INT NOT NULL,
  `current_quantity` INT NOT NULL,
  `current_temperature` DECIMAL(19,2) NOT NULL,
  `due_date` DATETIME(6) NOT NULL,
  `initial_quantity` INT NOT NULL,
  `manufacturing_date` DATETIME(6) NOT NULL,
  `manufacturing_hour` INT NOT NULL,
  `minimum_temperature` DECIMAL(19,2) NOT NULL,
  `product_id` INT NOT NULL,
  `section_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `product_id_idx` (`product_id` ASC) VISIBLE,
  INDEX `section_id_idx` (`section_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_product_batches`
    FOREIGN KEY (`product_id`)
    REFERENCES `melisprint`.`products` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_section_product_batches`
    FOREIGN KEY (`section_id`)
    REFERENCES `melisprint`.`sections` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`product_records`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`product_records` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `last_update_date` DATETIME(6) NOT NULL,
  `purchase_price` DECIMAL(19,2) NOT NULL,
  `sale_price` DECIMAL(19,2) NOT NULL,
  `product_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `product_id_idx` (`product_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_product_records`
    FOREIGN KEY (`product_id`)
    REFERENCES `melisprint`.`products` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`buyers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`buyers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `card_number_id` VARCHAR(255) NOT NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `card_number_id_UNIQUE` (`card_number_id` ASC) VISIBLE)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`carriers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`carriers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cid` VARCHAR(255) NOT NULL,
  `company_name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `locality_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `locality_id_idx` (`locality_id` ASC) VISIBLE,
  CONSTRAINT `fk_locality_carrier`
    FOREIGN KEY (`locality_id`)
    REFERENCES `melisprint`.`localities` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`order_status`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`order_status` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`purchase_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`purchase_orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `order_number` VARCHAR(255) NOT NULL,
  `order_date` DATETIME(6) NOT NULL,
  `tracking_code` VARCHAR(255) NOT NULL,
  `buyer_id` INT NOT NULL,
  `carrier_id` INT NULL,
  `order_status_id` INT NOT NULL,
  `warehouse_id` INT NULL,
  `product_record_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `buyer_id_idx` (`buyer_id` ASC) VISIBLE,
  INDEX `carrier_id_idx` (`carrier_id` ASC) VISIBLE,
  INDEX `order_status_id_idx` (`order_status_id` ASC) VISIBLE,
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  INDEX `fk_product_record_orders_idx` (`product_record_id` ASC) VISIBLE,
  CONSTRAINT `fk_buyer_purchase_orders`
    FOREIGN KEY (`buyer_id`)
    REFERENCES `melisprint`.`buyers` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_carrier_purchase_orders`
    FOREIGN KEY (`carrier_id`)
    REFERENCES `melisprint`.`carriers` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_order_status_purchase_orders`
    FOREIGN KEY (`order_status_id`)
    REFERENCES `melisprint`.`order_status` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_warehouse_purchase_orders`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `melisprint`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_product_record_orders`
    FOREIGN KEY (`product_record_id`)
    REFERENCES `melisprint`.`product_records` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`order_details`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`order_details` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `clean_liness_status` VARCHAR(255) NOT NULL,
  `quantity` INT NOT NULL,
  `temperature` DECIMAL(19,2) NOT NULL,
  `product_record_id` INT NOT NULL,
  `purchase_order_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `product_record_id_idx` (`product_record_id` ASC) VISIBLE,
  INDEX `purchase_order_id_idx` (`purchase_order_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_record_order_details`
    FOREIGN KEY (`product_record_id`)
    REFERENCES `melisprint`.`product_records` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_purchase_order_order_details`
    FOREIGN KEY (`purchase_order_id`)
    REFERENCES `melisprint`.`purchase_orders` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`employees`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`employees` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `card_number_id` VARCHAR(255) NOT NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `warehouse_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  UNIQUE INDEX `card_number_id_UNIQUE` (`card_number_id` ASC) VISIBLE,
  CONSTRAINT `fk_warehouse_employees`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `melisprint`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`inbound_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`inbound_orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `order_date` DATETIME(6) NOT NULL,
  `order_number` VARCHAR(255) NOT NULL,
  `employee_id` INT NOT NULL,
  `product_batch_id` INT NOT NULL,
  `warehouse_id` INT NOT NULL,
<<<<<<< HEAD
  CONSTRAINT `order_number` UNIQUE (`order_number`),
=======
>>>>>>> origin/feature-warehouse-carries-2
  PRIMARY KEY (`id`),
  INDEX `employee_id_idx` (`employee_id` ASC) VISIBLE,
  INDEX `product_batch_id_idx` (`product_batch_id` ASC) VISIBLE,
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  CONSTRAINT `fk_employee_inbound_orders`
    FOREIGN KEY (`employee_id`)
    REFERENCES `melisprint`.`employees` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_product_batch_inbound_orders`
    FOREIGN KEY (`product_batch_id`)
    REFERENCES `melisprint`.`product_batches` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_warehouse_inbound_orders`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `melisprint`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`roles`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`roles` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  `rol_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `passoword` VARCHAR(255) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `melisprint`.`user_rol`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`user_rol` (
  `usuario_id` INT NOT NULL AUTO_INCREMENT,
  `rol_id` INT NOT NULL,
  INDEX `usuario_id_idx` (`usuario_id` ASC) VISIBLE,
  INDEX `rol_id_idx` (`rol_id` ASC) VISIBLE,
  CONSTRAINT `fk_usuario_user_rol`
    FOREIGN KEY (`usuario_id`)
    REFERENCES `melisprint`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_rol_user_rol`
    FOREIGN KEY (`rol_id`)
    REFERENCES `melisprint`.`roles` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
<<<<<<< HEAD
=======


>>>>>>> origin/feature-warehouse-carries-2
-- -----------------------------------------------------
-- Table `logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`logs` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `method` VARCHAR(255) NOT NULL,
    `label` VARCHAR(255) NOT NULL,
    `level` VARCHAR(255) NOT NULL,
    `message` VARCHAR(255) NOT NULL,
    `status` INT NOT NULL,
    `insert_date` DATETIME(6) NOT NULL,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;
<<<<<<< HEAD
SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
INSERT INTO `melisprint`.`countries` (`country_name`) VALUES ('Brazil');
INSERT INTO `melisprint`.`countries` (`country_name`) VALUES ('United States');
INSERT INTO `melisprint`.`provinces` (`province_name`, `country_id`) VALUES ('São Paulo', 1);
INSERT INTO `melisprint`.`provinces` (`province_name`, `country_id`) VALUES ('California', 2);
INSERT INTO `melisprint`.`localities` (`locality_name`, `province_id`) VALUES ('São Paulo City', 1);
INSERT INTO `melisprint`.`localities` (`locality_name`, `province_id`) VALUES ('Los Angeles', 2);
INSERT INTO `melisprint`.`sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('123456789', 'Seller 1', 'Address 1', '123456789', 1);
INSERT INTO `melisprint`.`sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('987654321', 'Seller 2', 'Address 2', '987654321', 2);
INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 1');
INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 2');
INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 3');
INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 4');
INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 5');
INSERT INTO `melisprint`.`products` (`product_code`, `description`, `width`, `height`, `length`, `net_weight`, `expiration_rate`, `recommended_freezing_temperature`, `freezing_rate`, `product_type_id`, `seller_id`) VALUES ('P001', 'Product 1', '10', 5.5, 8.2, 100.25, 0.8, -18, 0.5, 1, 1);
INSERT INTO `melisprint`.`products` (`product_code`, `description`, `width`, `height`, `length`, `net_weight`, `expiration_rate`, `recommended_freezing_temperature`, `freezing_rate`, `product_type_id`, `seller_id`) VALUES ('P002', 'Product 2', '7.5', 3.2, 6.7, 75.5, 0.9, -15, 0.3, 2, 2);
INSERT INTO `melisprint`.`warehouses` (`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature`, `locality_id`) VALUES ('Warehouse 1 Address', '111111111', 'W001', 100, -20, 1);
INSERT INTO `melisprint`.`warehouses` (`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature`, `locality_id`) VALUES ('Warehouse 2 Address', '222222222', 'W002', 150, -18, 2);
INSERT INTO `melisprint`.`sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (1, -18, -20, 50, 20, 100, 1, 1);
INSERT INTO `melisprint`.`sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (2, -15, -18, 60, 30, 150, 2, 2);
INSERT INTO `melisprint`.`product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES (1, 200, -18, '2023-07-31 00:00:00', 300, '2023-07-01 00:00:00', 8, -20, 1, 1);
INSERT INTO `melisprint`.`product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES (2, 150, -15, '2023-08-15 00:00:00', 200, '2023-07-10 00:00:00', 9, -18, 2, 2);
INSERT INTO `melisprint`.`product_records` (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES ('2023-07-05 10:00:00', 10.50, 15.00, 1);
INSERT INTO `melisprint`.`product_records` (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES ('2023-07-05 10:00:00', 8.75, 12.50, 2);
INSERT INTO `melisprint`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES ('987654321', 'John', 'Doe');
INSERT INTO `melisprint`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES ('123456789', 'Jane', 'Smith');
INSERT INTO `melisprint`.`carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('111111', 'Carrier 1', 'Carrier Address 1', '111111111', 1);
INSERT INTO `melisprint`.`carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('222222', 'Carrier 2', 'Carrier Address 2', '222222222', 2);
INSERT INTO `melisprint`.`order_status` (`description`) VALUES ('Completed');
INSERT INTO `melisprint`.`order_status` (`description`) VALUES ('Pending');
INSERT INTO `melisprint`.`order_status` (`description`) VALUES ('Processing');
INSERT INTO `melisprint`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `carrier_id`, `order_status_id`, `warehouse_id`, `product_record_id`) VALUES ('PO001', '2023-07-01 10:00:00', 'TRACK001', 1, 1, 1, 1, 1);
INSERT INTO `melisprint`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `carrier_id`, `order_status_id`, `warehouse_id`, `product_record_id`) VALUES ('PO002', '2023-07-02 11:00:00', 'TRACK002', 2, 2, 2, 2, 2);
INSERT INTO `melisprint`.`order_details` (`clean_liness_status`, `quantity`, `temperature`, `product_record_id`, `purchase_order_id`) VALUES ('Clean', 10, -18, 1, 1);
INSERT INTO `melisprint`.`order_details` (`clean_liness_status`, `quantity`, `temperature`, `product_record_id`, `purchase_order_id`) VALUES ('Not clean', 20, -15, 2, 2);
INSERT INTO `melisprint`.`employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES ('123456', 'John', 'Smith', 1);
INSERT INTO `melisprint`.`employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES ('654321', 'Jane', 'Doe', 2);
INSERT INTO `melisprint`.`inbound_orders` (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES ('2023-07-05 14:00:00', 'INB001', 1, 1, 1);
INSERT INTO `melisprint`.`inbound_orders` (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES ('2023-07-06 15:00:00', 'INB002', 2, 2, 2);
INSERT INTO `melisprint`.`roles` (`description`, `rol_name`) VALUES ('Administrator', 'admin');
INSERT INTO `melisprint`.`roles` (`description`, `rol_name`) VALUES ('Employee', 'employee');
INSERT INTO `melisprint`.`users` (`passoword`, `username`) VALUES ('password1', 'user1');
INSERT INTO `melisprint`.`users` (`passoword`, `username`) VALUES ('password2', 'user2');
INSERT INTO `melisprint`.`user_rol` (`usuario_id`, `rol_id`) VALUES (1, 1);
INSERT INTO `melisprint`.`user_rol` (`usuario_id`, `rol_id`) VALUES (2, 2);
=======



SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

INSERT INTO `melisprint`.`countries` (`country_name`) VALUES ('Brazil');
INSERT INTO `melisprint`.`countries` (`country_name`) VALUES ('United States');

INSERT INTO `melisprint`.`provinces` (`province_name`, `country_id`) VALUES ('São Paulo', 1);
INSERT INTO `melisprint`.`provinces` (`province_name`, `country_id`) VALUES ('California', 2);

INSERT INTO `melisprint`.`localities` (`locality_name`, `province_id`) VALUES ('São Paulo City', 1);
INSERT INTO `melisprint`.`localities` (`locality_name`, `province_id`) VALUES ('Los Angeles', 2);

INSERT INTO `melisprint`.`sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('123456789', 'Seller 1', 'Address 1', '123456789', 1);
INSERT INTO `melisprint`.`sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('987654321', 'Seller 2', 'Address 2', '987654321', 2);

INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 1');
INSERT INTO `melisprint`.`product_types` (`description`) VALUES ('Type 2');

INSERT INTO `melisprint`.`products` (`product_code`, `description`, `width`, `height`, `length`, `net_weight`, `expiration_rate`, `recommended_freezing_temperature`, `freezing_rate`, `product_type_id`, `seller_id`) VALUES ('P001', 'Product 1', '10', 5.5, 8.2, 100.25, 0.8, -18, 0.5, 1, 1);
INSERT INTO `melisprint`.`products` (`product_code`, `description`, `width`, `height`, `length`, `net_weight`, `expiration_rate`, `recommended_freezing_temperature`, `freezing_rate`, `product_type_id`, `seller_id`) VALUES ('P002', 'Product 2', '7.5', 3.2, 6.7, 75.5, 0.9, -15, 0.3, 2, 2);

INSERT INTO `melisprint`.`warehouses` (`address`, `telephone`, `warehouse_code`, `minimun_capacity`, `minimun_temperature`, `locality_id`) VALUES ('Warehouse 1 Address', '111111111', 'W001', 100, -20, 1);
INSERT INTO `melisprint`.`warehouses` (`address`, `telephone`, `warehouse_code`, `minimun_capacity`, `minimun_temperature`, `locality_id`) VALUES ('Warehouse 2 Address', '222222222', 'W002', 150, -18, 2);

INSERT INTO `melisprint`.`sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (1, -18, -20, 50, 20, 100, 1, 1);
INSERT INTO `melisprint`.`sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (2, -15, -18, 60, 30, 150, 2, 2);

INSERT INTO `melisprint`.`product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES (1, 200, -18, '2023-07-31 00:00:00', 300, '2023-07-01 00:00:00', 8, -20, 1, 1);
INSERT INTO `melisprint`.`product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES (2, 150, -15, '2023-08-15 00:00:00', 200, '2023-07-10 00:00:00', 9, -18, 2, 2);

INSERT INTO `melisprint`.`product_records` (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES ('2023-07-05 10:00:00', 10.50, 15.00, 1);
INSERT INTO `melisprint`.`product_records` (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES ('2023-07-05 10:00:00', 8.75, 12.50, 2);

INSERT INTO `melisprint`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES ('987654321', 'John', 'Doe');
INSERT INTO `melisprint`.`buyers` (`card_number_id`, `first_name`, `last_name`) VALUES ('123456789', 'Jane', 'Smith');

INSERT INTO `melisprint`.`carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('111111', 'Carrier 1', 'Carrier Address 1', '111111111', 1);
INSERT INTO `melisprint`.`carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES ('222222', 'Carrier 2', 'Carrier Address 2', '222222222', 2);

INSERT INTO `melisprint`.`order_status` (`description`) VALUES ('Pending');
INSERT INTO `melisprint`.`order_status` (`description`) VALUES ('Processing');

INSERT INTO `melisprint`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `carrier_id`, `order_status_id`, `warehouse_id`, `product_record_id`) VALUES ('PO001', '2023-07-01 10:00:00', 'TRACK001', 1, 1, 1, 1, 1);
INSERT INTO `melisprint`.`purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `carrier_id`, `order_status_id`, `warehouse_id`, `product_record_id`) VALUES ('PO002', '2023-07-02 11:00:00', 'TRACK002', 2, 2, 2, 2, 2);

INSERT INTO `melisprint`.`order_details` (`clean_liness_status`, `quantity`, `temperature`, `product_record_id`, `purchase_order_id`) VALUES ('Clean', 10, -18, 1, 1);
INSERT INTO `melisprint`.`order_details` (`clean_liness_status`, `quantity`, `temperature`, `product_record_id`, `purchase_order_id`) VALUES ('Not clean', 20, -15, 2, 2);

INSERT INTO `melisprint`.`employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES ('123456', 'John', 'Smith', 1);
INSERT INTO `melisprint`.`employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES ('654321', 'Jane', 'Doe', 2);

INSERT INTO `melisprint`.`inbound_orders` (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES ('2023-07-05 14:00:00', 'INB001', 1, 1, 1);
INSERT INTO `melisprint`.`inbound_orders` (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES ('2023-07-06 15:00:00', 'INB002', 2, 2, 2);

INSERT INTO `melisprint`.`roles` (`description`, `rol_name`) VALUES ('Administrator', 'admin');
INSERT INTO `melisprint`.`roles` (`description`, `rol_name`) VALUES ('Employee', 'employee');

INSERT INTO `melisprint`.`users` (`passoword`, `username`) VALUES ('password1', 'user1');
INSERT INTO `melisprint`.`users` (`passoword`, `username`) VALUES ('password2', 'user2');

INSERT INTO `melisprint`.`user_rol` (`usuario_id`, `rol_id`) VALUES (1, 1);
INSERT INTO `melisprint`.`user_rol` (`usuario_id`, `rol_id`) VALUES (2, 2);

>>>>>>> origin/feature-warehouse-carries-2
INSERT INTO `melisprint`.`logs` (`method`, `label`, `level`, `message`, `status`, `insert_date`) VALUES ('GET', 'API Request', 'Info', 'API request received', 200, '2023-07-05 16:00:00');
INSERT INTO `melisprint`.`logs` (`method`, `label`, `level`, `message`, `status`, `insert_date`) VALUES ('POST', 'Data Update', 'Warning', 'Data update failed', 500, '2023-07-05 17:00:00');