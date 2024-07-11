-- MySQL dump 10.13  Distrib 5.5.62, for Win64 (AMD64)
--
-- Host: localhost    Database: invoice
-- ------------------------------------------------------
-- Server version	5.7.44-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer` (
  `name` varchar(100) CHARACTER SET latin1 NOT NULL,
  `address` varchar(100) CHARACTER SET latin1 NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES ('Customer 1','Address 1'),('Customer 2','Address 2'),('Customer 3','Address 3'),('Customer 4','Address 4'),('Customer 5','Address 5');
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_trx`
--

DROP TABLE IF EXISTS `invoice_trx`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoice_trx` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `issue_date` varchar(100) NOT NULL,
  `subject` varchar(100) NOT NULL,
  `total_item` int(11) NOT NULL,
  `customer_name` varchar(100) NOT NULL,
  `due_date` varchar(100) NOT NULL,
  `status` varchar(100) NOT NULL,
  `customer_address` varchar(100) DEFAULT NULL,
  `sub_total` int(11) DEFAULT NULL,
  `tax` int(11) DEFAULT NULL,
  `grand_total` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_trx`
--

LOCK TABLES `invoice_trx` WRITE;
/*!40000 ALTER TABLE `invoice_trx` DISABLE KEYS */;
INSERT INTO `invoice_trx` VALUES (1,'13-12-2023','Subject 1',2,'Customer 1','16-12-2023','Unpaid','Address 1',20,2,22),(20,'14-12-2023','Subject 2',3,'Customer 2','17-12-2023','Unpaid','Address 2',40,4,44);
/*!40000 ALTER TABLE `invoice_trx` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_trx_dtl`
--

DROP TABLE IF EXISTS `invoice_trx_dtl`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoice_trx_dtl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_invoice` int(11) NOT NULL,
  `item_name` varchar(100) NOT NULL,
  `item_type` varchar(100) NOT NULL,
  `qty` int(11) NOT NULL,
  `unit_price` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `invoice_trx_dtl_FK` (`id_invoice`),
  CONSTRAINT `invoice_trx_dtl_FK` FOREIGN KEY (`id_invoice`) REFERENCES `invoice_trx` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_trx_dtl`
--

LOCK TABLES `invoice_trx_dtl` WRITE;
/*!40000 ALTER TABLE `invoice_trx_dtl` DISABLE KEYS */;
INSERT INTO `invoice_trx_dtl` VALUES (1,1,'Item 1','Type 1',1,10,10),(2,1,'Item 2','Type 2',1,10,10),(4,20,'Item 1','Type 1',1,10,10),(5,20,'Item 3','Type 3',1,10,10),(6,20,'Item 4','Type 4',2,10,20);
/*!40000 ALTER TABLE `invoice_trx_dtl` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item`
--

DROP TABLE IF EXISTS `item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item` (
  `item_name` varchar(100) NOT NULL,
  `item_type` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item`
--

LOCK TABLES `item` WRITE;
/*!40000 ALTER TABLE `item` DISABLE KEYS */;
INSERT INTO `item` VALUES ('Item 1','Type 1'),('Item 2','Type 2'),('Item 3','Type 3'),('Item 4','Type 4');
/*!40000 ALTER TABLE `item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'invoice'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-16  2:28:56
