-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: foundationService
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Credentials`
--

DROP TABLE IF EXISTS `Credentials`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Credentials` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Type` varchar(100) NOT NULL,
  `Issuer` varchar(100) NOT NULL,
  `IssuanceDate` timestamp NOT NULL,
  `ExpirationDate` timestamp NOT NULL,
  `Description` varchar(100) NOT NULL,
  `Revoked` tinyint NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Credentials`
--

LOCK TABLES `Credentials` WRITE;
/*!40000 ALTER TABLE `Credentials` DISABLE KEYS */;
INSERT INTO `Credentials` VALUES (4,'PermanentResidentCard','did:metablox:sampleIssuer','2022-04-08 22:17:09','2032-04-08 22:17:09','Government of Example Permanent Resident Card',0),(5,'WifiAccess','did:metablox:sampleIssuer','2022-04-08 22:57:21','2032-04-08 22:57:21','Example Wifi Access Credential',0),(6,'WifiAccess','did:metablox:sampleIssuer','2022-04-08 22:58:59','2032-04-08 22:58:59','Example Wifi Access Credential',0),(7,'WifiAccess','did:metablox:sampleIssuer','2022-04-08 22:59:13','2032-04-08 22:59:13','Example Wifi Access Credential',0),(8,'WifiAccess','did:metablox:sampleIssuer','2022-04-08 23:08:50','2032-04-08 23:08:50','Example Wifi Access Credential',0),(11,'WifiAccess','did:metablox:sampleIssuer','2022-04-08 23:18:09','2032-04-08 23:18:09','Example Wifi Access Credential',0),(12,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-08 23:29:43','2032-04-08 23:29:43','Example Wifi Access Credential',0),(13,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-11 22:26:45','2032-04-11 22:26:45','Example Wifi Access Credential',0),(14,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-11 22:27:50','2032-04-11 22:27:50','Example Wifi Access Credential',0),(15,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-11 22:31:55','2032-04-11 22:31:55','Example Wifi Access Credential',0),(17,'MiningLicense','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-13 21:36:38','2032-04-13 21:36:38','Example Mining License Credential',0),(18,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-13 21:36:58','2032-04-13 21:36:58','Example Wifi Access Credential',0),(19,'MiningLicense','did:metablox:sampleIssuer','2022-04-14 18:19:04','2032-04-14 18:19:04','Example Mining License Credential',0),(20,'MiningLicense','did:metablox:sampleIssuer','2022-04-14 18:20:11','2032-04-14 18:20:11','Example Mining License Credential',0),(21,'WifiAccess','did:metablox:sampleIssuer','2022-04-14 18:20:34','2032-04-14 18:20:34','Example Wifi Access Credential',0),(22,'WifiAccess','did:metablox:sampleIssuer','2022-04-14 18:20:51','2032-04-14 18:20:51','Example Wifi Access Credential',0),(23,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-14 19:13:26','2032-04-14 19:13:26','Example Wifi Access Credential',0),(24,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-14 19:16:46','2032-04-14 19:16:46','Example Wifi Access Credential',0),(25,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-14 19:30:19','2032-04-14 19:30:19','Example Wifi Access Credential',0),(26,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-14 19:30:38','2032-04-14 19:30:38','Example Wifi Access Credential',0),(27,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-14 19:30:43','2033-04-14 19:30:43','Example Wifi Access Credential',1),(28,'MiningLicense','did:metablox:sampleIssuer','2022-04-18 18:32:58','2032-04-18 18:32:58','Example Mining License Credential',0),(29,'MiningLicense','did:metablox:sampleIssuer','2022-04-18 18:32:58','2032-04-18 18:32:58','Example Mining License Credential',0),(30,'MiningLicense','did:metablox:sampleIssuer','2022-04-18 18:36:11','2032-04-18 18:36:11','Example Mining License Credential',0),(31,'MiningLicense','did:metablox:sampleIssuer','2022-04-18 18:36:11','2032-04-18 18:36:11','Example Mining License Credential',0),(32,'MiningLicense','did:metablox:sampleIssuer','2022-04-18 18:36:34','2032-04-18 18:36:34','Example Mining License Credential',0),(33,'MiningLicense','did:metablox:sampleIssuer','2022-04-18 18:36:34','2032-04-18 18:36:34','Example Mining License Credential',0),(34,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-19 19:00:02','2032-04-19 19:00:02','Example Wifi Access Credential',0),(35,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-19 20:31:30','2032-04-19 20:31:30','Example Wifi Access Credential',0),(36,'MiningLicense','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-19 20:40:58','2032-04-19 20:40:58','Example Mining License Credential',0),(37,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-19 23:14:47','2032-04-19 23:14:47','Example Wifi Access Credential',0),(38,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-19 23:36:54','2032-04-19 23:36:54','Example Wifi Access Credential',0),(39,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-19 23:47:58','2032-04-19 23:47:58','Example Wifi Access Credential',0),(40,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-20 18:18:25','2032-04-20 18:18:25','Example Wifi Access Credential',0),(41,'MiningLicense','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-21 18:28:25','2032-04-21 18:28:25','Example Mining License Credential',0),(42,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-21 21:36:35','2032-04-21 21:36:35','Example Wifi Access Credential',0),(43,'MiningLicense','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-21 22:16:03','2032-04-21 22:16:03','Example Mining License Credential',0),(44,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-26 18:18:30','2032-04-26 18:18:30','Example Wifi Access Credential',0),(45,'MiningLicense','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-04-26 18:19:02','2032-04-26 18:19:02','Example Mining License Credential',0),(46,'MiningLicense','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-02 23:09:20','2032-05-02 23:09:20','Example Mining License Credential',0),(47,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-02 23:21:00','2032-05-02 23:21:00','Example Wifi Access Credential',0),(48,'MiningLicense','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-02 23:21:49','2032-05-02 23:21:49','Example Mining License Credential',0),(51,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-03 22:23:13','2032-05-03 22:23:13','Example Wifi Access Credential',0),(52,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-03 22:40:09','2032-05-03 22:40:09','Example Wifi Access Credential',0),(53,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-03 22:41:32','2032-05-03 22:41:32','Example Wifi Access Credential',0),(54,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-03 22:44:57','2032-05-03 22:44:57','Example Wifi Access Credential',0),(55,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-03 23:52:30','2032-05-03 23:52:30','Example Wifi Access Credential',0),(56,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-03 23:58:37','2032-05-03 23:58:37','Example Wifi Access Credential',0),(57,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:02:09','2032-05-04 00:02:09','Example Wifi Access Credential',0),(58,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:06:18','2032-05-04 00:06:18','Example Wifi Access Credential',0),(59,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:18:41','2032-05-04 00:18:41','Example Wifi Access Credential',0),(60,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:26:01','2032-05-04 00:26:01','Example Wifi Access Credential',0),(61,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:33:33','2032-05-04 00:33:33','Example Wifi Access Credential',0),(62,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:36:47','2032-05-04 00:36:47','Example Wifi Access Credential',0),(63,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:41:35','2032-05-04 00:41:35','Example Wifi Access Credential',0),(64,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:43:35','2032-05-04 00:43:35','Example Wifi Access Credential',0),(65,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:49:08','2032-05-04 00:49:08','Example Wifi Access Credential',0),(66,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 00:57:01','2032-05-04 00:57:01','Example Wifi Access Credential',0),(67,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:05:01','2032-05-04 01:05:01','Example Wifi Access Credential',0),(68,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:14:03','2032-05-04 01:14:03','Example Wifi Access Credential',0),(69,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:17:08','2032-05-04 01:17:08','Example Wifi Access Credential',0),(70,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:26:42','2032-05-04 01:26:42','Example Wifi Access Credential',0),(71,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:28:42','2032-05-04 01:28:42','Example Wifi Access Credential',0),(72,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:37:38','2032-05-04 01:37:38','Example Wifi Access Credential',0),(73,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:38:41','2032-05-04 01:38:41','Example Wifi Access Credential',0),(74,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:40:37','2032-05-04 01:40:37','Example Wifi Access Credential',0),(75,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:40:49','2032-05-04 01:40:49','Example Wifi Access Credential',0),(76,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:44:52','2032-05-04 01:44:52','Example Wifi Access Credential',0),(77,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:46:08','2032-05-04 01:46:08','Example Wifi Access Credential',0),(78,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:48:56','2032-05-04 01:48:56','Example Wifi Access Credential',0),(79,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:54:23','2032-05-04 01:54:23','Example Wifi Access Credential',0),(80,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 01:59:57','2032-05-04 01:59:57','Example Wifi Access Credential',0),(81,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 02:02:28','2032-05-04 02:02:28','Example Wifi Access Credential',0),(82,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 02:04:36','2032-05-04 02:04:36','Example Wifi Access Credential',0),(83,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 02:11:26','2032-05-04 02:11:26','Example Wifi Access Credential',0),(84,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 02:20:51','2032-05-04 02:20:51','Example Wifi Access Credential',0),(85,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 02:23:03','2032-05-04 02:23:03','Example Wifi Access Credential',0),(86,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 02:24:54','2032-05-04 02:24:54','Example Wifi Access Credential',0),(87,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 02:25:31','2032-05-04 02:25:31','Example Wifi Access Credential',0),(88,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 19:47:42','2032-05-04 19:47:42','Example Wifi Access Credential',0),(89,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-04 21:34:59','2032-05-04 21:34:59','Example Wifi Access Credential',0),(90,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-06 20:36:56','2032-05-06 20:36:56','Example Wifi Access Credential',0),(91,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-09 18:14:04','2032-05-09 18:14:04','Example Wifi Access Credential',0),(92,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-09 18:35:03','2032-05-09 18:35:03','Example Wifi Access Credential',0),(93,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-09 18:55:10','2032-05-09 18:55:10','Example Wifi Access Credential',0),(94,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-10 02:12:58','2032-05-10 02:12:58','Example Wifi Access Credential',0),(95,'WifiAccess','did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','2022-05-10 02:43:49','2032-05-10 02:43:49','Example Mining License Credential',0),(97,'StakingVC','did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX','2022-05-26 04:29:43','2032-05-26 04:29:43','Example Mining License Credential',0),(98,'StakingVC','did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX','2022-05-26 04:30:37','2032-05-26 04:30:37','Example Mining License Credential',0),(99,'WifiAccess','did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX','2022-05-28 00:51:09','2032-05-28 00:51:09','Example Wifi Access Credential',0);
/*!40000 ALTER TABLE `Credentials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ExchangeRate`
--

DROP TABLE IF EXISTS `ExchangeRate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ExchangeRate` (
  `ID` int NOT NULL,
  `ExchangeRate` float NOT NULL,
  `CreateTime` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ExchangeRate`
--

LOCK TABLES `ExchangeRate` WRITE;
/*!40000 ALTER TABLE `ExchangeRate` DISABLE KEYS */;
INSERT INTO `ExchangeRate` VALUES (1,25.3,'2022-05-27 21:55:00.000');
/*!40000 ALTER TABLE `ExchangeRate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `MinerInfo`
--

DROP TABLE IF EXISTS `MinerInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `MinerInfo` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(100) NOT NULL,
  `SSID` varchar(100) DEFAULT NULL,
  `BSSID` varchar(100) DEFAULT NULL,
  `CreateTime` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `Longitude` float DEFAULT NULL,
  `Latitude` float DEFAULT NULL,
  `OnlineStatus` tinyint NOT NULL DEFAULT '1',
  `MiningPower` float DEFAULT NULL,
  `IsMinable` tinyint NOT NULL,
  `DID` varchar(100) NOT NULL,
  `Host` varchar(100) NOT NULL,
  `IsVirtual` tinyint NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MinerInfo`
--

LOCK TABLES `MinerInfo` WRITE;
/*!40000 ALTER TABLE `MinerInfo` DISABLE KEYS */;
INSERT INTO `MinerInfo` VALUES (1,'testName','testSSID','testBSSID','2022-04-19 07:00:00.000',50,100,1,42,1,'sampleDID','sampleHost',0),(2,'testName2','testSSID2','testBSSID2','2022-05-20 07:00:00.000',50,60,1,75,0,'sampleDID2','sampleHost2',0),(3,'Virtual Miner 1','testSSIDV1','testBSSIDV1','2022-06-02 07:00:00.000',0,0,1,500,1,'did:metablox:V612aF8NKhDm2KEq9V43WhemArcRnx5LvxK35fQyzmX','http://54.226.46.25:2060',1);
/*!40000 ALTER TABLE `MinerInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `MinerManufacturer`
--

DROP TABLE IF EXISTS `MinerManufacturer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `MinerManufacturer` (
  `Name` varchar(100) NOT NULL,
  `Email` varchar(100) NOT NULL,
  `Address` varchar(100) NOT NULL,
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MinerManufacturer`
--

LOCK TABLES `MinerManufacturer` WRITE;
/*!40000 ALTER TABLE `MinerManufacturer` DISABLE KEYS */;
/*!40000 ALTER TABLE `MinerManufacturer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `MiningLicenseInfo`
--

DROP TABLE IF EXISTS `MiningLicenseInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `MiningLicenseInfo` (
  `CredentialID` int NOT NULL,
  `ID` varchar(100) NOT NULL,
  `Name` varchar(100) NOT NULL,
  `Model` varchar(100) NOT NULL,
  `Serial` varchar(100) NOT NULL,
  PRIMARY KEY (`CredentialID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MiningLicenseInfo`
--

LOCK TABLES `MiningLicenseInfo` WRITE;
/*!40000 ALTER TABLE `MiningLicenseInfo` DISABLE KEYS */;
INSERT INTO `MiningLicenseInfo` VALUES (17,'','check','',''),(19,'','ThisIsAPlaceholder2','',''),(20,'','ThisIsAPlaceholder2','',''),(28,'did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','ThisIsAPlaceholder2','',''),(29,'did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','ThisIsAPlaceholder2','',''),(30,'did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','ThisIsAPlaceholder2','',''),(31,'did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','ThisIsAPlaceholder2','',''),(32,'did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','ThisIsAPlaceholder2','',''),(33,'did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo','ThisIsAPlaceholder2','',''),(36,'did:metablox:hgsduijgbwd','check','',''),(41,'did:metablox:hgsduijgbwk','check','',''),(43,'did:metablox:hgsduijgbww','check','',''),(45,'did:metablox:hgsduijgbwxxx','check','',''),(46,'did:metablox:hgsduijgbwxxxx','','',''),(48,'did:metablox:hgsduijgbwxxxxx','check2','check','check3'),(95,'did:metablox:hgsduijgbwxxxxxxxxxx','check2','check','check3'),(101,'did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX','sampleName','sampleModel','sampleSerial');
/*!40000 ALTER TABLE `MiningLicenseInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `MiningRole`
--

DROP TABLE IF EXISTS `MiningRole`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `MiningRole` (
  `DID` varchar(128) NOT NULL,
  `WalletAddress` varchar(128) NOT NULL,
  `Type` varchar(64) NOT NULL,
  PRIMARY KEY (`DID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MiningRole`
--

LOCK TABLES `MiningRole` WRITE;
/*!40000 ALTER TABLE `MiningRole` DISABLE KEYS */;
INSERT INTO `MiningRole` VALUES ('did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX','0xBE1e1dB948CC1f441514aFb8924B67891f1c6889','sampleType');
/*!40000 ALTER TABLE `MiningRole` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OrderInterest`
--

DROP TABLE IF EXISTS `OrderInterest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `OrderInterest` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `OrderID` int NOT NULL,
  `Time` datetime NOT NULL,
  `APY` float NOT NULL,
  `InterestGain` float NOT NULL,
  `TotalInterestGain` float NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OrderInterest`
--

LOCK TABLES `OrderInterest` WRITE;
/*!40000 ALTER TABLE `OrderInterest` DISABLE KEYS */;
INSERT INTO `OrderInterest` VALUES (1,1,'2022-05-12 00:00:00',10,50,30),(2,2,'2022-05-12 01:00:00',20,20,20),(3,3,'2022-05-12 02:00:00',10,50,0);
/*!40000 ALTER TABLE `OrderInterest` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Orders`
--

DROP TABLE IF EXISTS `Orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Orders` (
  `OrderID` int NOT NULL AUTO_INCREMENT,
  `ProductID` int NOT NULL,
  `UserDID` varchar(45) NOT NULL,
  `Type` enum('Pending','Holding','Complete') NOT NULL,
  `Term` int DEFAULT NULL,
  `AccumulatedInterest` float NOT NULL DEFAULT '0',
  `TotalInterestGained` float NOT NULL DEFAULT '0',
  `PaymentAddress` varchar(45) NOT NULL,
  `Amount` float NOT NULL,
  `UserAddress` varchar(80) NOT NULL,
  PRIMARY KEY (`OrderID`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Orders`
--

LOCK TABLES `Orders` WRITE;
/*!40000 ALTER TABLE `Orders` DISABLE KEYS */;
INSERT INTO `Orders` VALUES (1,1,'test','Holding',1,30,30,'',100,'userAddress'),(4,0,'','Pending',1,0,0,'placeholder',0,''),(5,0,'testDID','Pending',1,0,0,'placeholder',0,''),(6,0,'','Pending',1,0,0,'placeholder',0,''),(7,0,'','Pending',1,0,0,'placeholder',0,''),(8,0,'','Pending',1,0,0,'placeholder',0,''),(9,0,'testDID','Pending',1,0,0,'placeholder',0,''),(10,0,'testDID','Pending',1,0,0,'placeholder',0,''),(11,1,'testDID','Pending',1,0,0,'placeholder',0,''),(12,1,'testDID','Holding',1,0,0,'placeholder',30,'0xBE1e1dB948CC1f441514aFb8924B67891f1c6889'),(13,1,'did:metablox:test','Holding',1,50,50,'placeholder',30,'0xBE1e1dB948CC1f441514aFb8924B67891f1c6889'),(14,1,'did:metablox:test','Pending',1,0,0,'testProductAddress',1234,'testAddress'),(15,1,'did:metablox:test','Pending',1,0,0,'testProductAddress',1234,'testAddress'),(16,1,'did:metablox:test','Pending',1,0,0,'testProductAddress',1234,'testAddress'),(17,1,'did:metablox:test','Pending',1,0,0,'testProductAddress',1234,'testAddress'),(18,1,'did:metablox:test','Pending',1,0,0,'placeholder',1234,'testAddress');
/*!40000 ALTER TABLE `Orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PrincipalUpdates`
--

DROP TABLE IF EXISTS `PrincipalUpdates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PrincipalUpdates` (
  `ID` int NOT NULL,
  `ProductID` int NOT NULL,
  `Time` datetime NOT NULL,
  `TotalPrincipal` float NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PrincipalUpdates`
--

LOCK TABLES `PrincipalUpdates` WRITE;
/*!40000 ALTER TABLE `PrincipalUpdates` DISABLE KEYS */;
/*!40000 ALTER TABLE `PrincipalUpdates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SeedExchangeHistory`
--

DROP TABLE IF EXISTS `SeedExchangeHistory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SeedExchangeHistory` (
  `VcID` int NOT NULL,
  `UserDID` varchar(100) NOT NULL,
  `ExchangeRate` float NOT NULL,
  `Amount` float NOT NULL,
  `CreateTime` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`VcID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SeedExchangeHistory`
--

LOCK TABLES `SeedExchangeHistory` WRITE;
/*!40000 ALTER TABLE `SeedExchangeHistory` DISABLE KEYS */;
INSERT INTO `SeedExchangeHistory` VALUES (1,'test',30,10,'2022-05-20 07:00:00.000'),(2,'test',20,45,'2022-05-20 07:00:00.000'),(27,'did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX',30,900,'2022-05-24 18:44:12.866'),(28,'did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX',30,900,'2022-05-27 21:44:20.942'),(29,'did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX',30,900,'2022-05-31 22:32:28.737'),(30,'did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX',30,900,'2022-05-31 22:56:58.992'),(31,'did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX',30,900,'2022-05-31 23:16:06.502'),(32,'did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX',30,900,'2022-05-31 23:24:44.597');
/*!40000 ALTER TABLE `SeedExchangeHistory` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `StakeEvents`
--

DROP TABLE IF EXISTS `StakeEvents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `StakeEvents` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Address` varchar(100) NOT NULL,
  `Amount` int NOT NULL,
  `Time` timestamp(3) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StakeEvents`
--

LOCK TABLES `StakeEvents` WRITE;
/*!40000 ALTER TABLE `StakeEvents` DISABLE KEYS */;
INSERT INTO `StakeEvents` VALUES (3,'0x00000000000000000000000056bdbb8ecb54570b5a3971aaacf85040e7ac3b4f',30000,'2022-06-02 01:00:01.000'),(4,'0x00000000000000000000000056bdbb8ecb54570b5a3971aaacf85040e7ac3b4f',30000,'2022-06-02 01:00:27.000');
/*!40000 ALTER TABLE `StakeEvents` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `StakingProducts`
--

DROP TABLE IF EXISTS `StakingProducts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `StakingProducts` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ProductName` varchar(45) NOT NULL,
  `MinOrderValue` int NOT NULL,
  `TopUpLimit` float NOT NULL,
  `MinRedeemValue` int NOT NULL,
  `LockUpPeriod` int NOT NULL,
  `DefaultAPY` float NOT NULL,
  `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `StartDate` timestamp(3) NOT NULL,
  `Term` int NOT NULL,
  `BurnedInterest` float NOT NULL DEFAULT '0',
  `Status` tinyint NOT NULL DEFAULT '1',
  `PaymentAddress` varchar(45) NOT NULL,
  `CurrencyType` varchar(10) NOT NULL,
  `Network` varchar(45) NOT NULL,
  `NextProductID` int DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StakingProducts`
--

LOCK TABLES `StakingProducts` WRITE;
/*!40000 ALTER TABLE `StakingProducts` DISABLE KEYS */;
INSERT INTO `StakingProducts` VALUES (1,'TestProduct',10,15,5,3,30,'2022-05-18 07:00:00.000','2022-12-25 08:00:00.000',1,50,1,'testProductAddress','MBLX','Ethereum',NULL);
/*!40000 ALTER TABLE `StakingProducts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `StakingVCInfo`
--

DROP TABLE IF EXISTS `StakingVCInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `StakingVCInfo` (
  `CredentialID` int NOT NULL,
  `ID` varchar(100) NOT NULL,
  PRIMARY KEY (`CredentialID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StakingVCInfo`
--

LOCK TABLES `StakingVCInfo` WRITE;
/*!40000 ALTER TABLE `StakingVCInfo` DISABLE KEYS */;
INSERT INTO `StakingVCInfo` VALUES (97,'did:metablox:hgsduijgbwxxxxxxxxxx'),(98,'did:metablox:hgsduijgbwxxxxxxxxxxw');
/*!40000 ALTER TABLE `StakingVCInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TXInfo`
--

DROP TABLE IF EXISTS `TXInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `TXInfo` (
  `PaymentNo` int NOT NULL AUTO_INCREMENT,
  `OrderID` int NOT NULL,
  `TXCurrencyType` varchar(10) NOT NULL,
  `TXType` enum('BuyIn','OrderClosure','InterestOnly') NOT NULL,
  `TXHash` varchar(66) NOT NULL,
  `Principal` float NOT NULL,
  `Interest` float NOT NULL,
  `UserAddress` varchar(45) NOT NULL,
  `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `RedeemableTime` timestamp(3) NOT NULL,
  PRIMARY KEY (`PaymentNo`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TXInfo`
--

LOCK TABLES `TXInfo` WRITE;
/*!40000 ALTER TABLE `TXInfo` DISABLE KEYS */;
INSERT INTO `TXInfo` VALUES (27,13,'MBLX','BuyIn','0xd6a12b3f112c640c8c1a0e868d21341a79070e00f8ef0f95753ca925f7920f94',30,0,'0xBE1e1dB948CC1f441514aFb8924B67891f1c6889','2022-05-27 20:25:20.955','2022-11-22 00:00:00.000'),(28,13,'MBLX','OrderClosure','placeholderHash',0,0,'0xBE1e1dB948CC1f441514aFb8924B67891f1c6889','2022-05-27 20:32:04.061','2022-11-22 00:00:00.000'),(29,13,'MBLX','OrderClosure','placeholderHash',0,0,'0xBE1e1dB948CC1f441514aFb8924B67891f1c6889','2022-05-27 20:32:04.065','2022-11-22 00:00:00.000');
/*!40000 ALTER TABLE `TXInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Users`
--

DROP TABLE IF EXISTS `Users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Users` (
  `DID` varchar(45) NOT NULL,
  `Currency` text NOT NULL,
  `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`DID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Users`
--

LOCK TABLES `Users` WRITE;
/*!40000 ALTER TABLE `Users` DISABLE KEYS */;
/*!40000 ALTER TABLE `Users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `WifiAccessInfo`
--

DROP TABLE IF EXISTS `WifiAccessInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `WifiAccessInfo` (
  `CredentialID` int NOT NULL,
  `ID` varchar(100) NOT NULL,
  `Type` enum('User','Validator') NOT NULL,
  PRIMARY KEY (`CredentialID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `WifiAccessInfo`
--

LOCK TABLES `WifiAccessInfo` WRITE;
/*!40000 ALTER TABLE `WifiAccessInfo` DISABLE KEYS */;
INSERT INTO `WifiAccessInfo` VALUES (47,'did:metablox:hgsduijgbwxxxx','User'),(51,'did:metablox:hgsduijgbwxxxxx','User'),(52,'did:metablox:hgsduijgbwxxxxxx','User'),(53,'did:metablox:hgsduijgbwxxxxxxx','User'),(54,'did:metablox:hgsduijgbwv','User'),(55,'did:metablox:hgsduijgbwz','User'),(56,'did:metablox:hgsduijgbwzz','User'),(57,'did:metablox:hgsduijgbwzzz','User'),(58,'did:metablox:hgsduijgbwa','User'),(59,'did:metablox:hgsduijgbwb','User'),(60,'did:metablox:hgsduijgbwd','User'),(61,'did:metablox:hgsduijgbww','User'),(62,'did:metablox:hgsduijgbwww','User'),(63,'did:metablox:hgsduijgbwwww','User'),(64,'did:metablox:hgsduijgbwwaw','User'),(65,'did:metablox:hgsduijgbwwawa','User'),(66,'did:metablox:hgsduijgbwwawwa','User'),(67,'did:metablox:hgsduijgbwwawawa','User'),(68,'did:metablox:hgsduijgbwawawawa','User'),(69,'did:metablox:hgsduijgbwawawawaw','User'),(70,'did:metablox:hgsduijgbwawawawawa','User'),(71,'did:metablox:hgsduijgbwawawawawaw','User'),(72,'did:metablox:hgsduijgbwawawawawawa','User'),(73,'did:metablox:hgsduijgbwawawawawawaw','User'),(74,'did:metablox:hgsduijgbwawawawawawawa','User'),(75,'did:metablox:hgsduijgbwawawawawawawaw','User'),(76,'did:metablox:hgsduijgbwawawawawawawawa','User'),(77,'did:metablox:hgsduijgbwawawawawawawawaw','User'),(78,'did:metablox:hgsduijgbwawawawawawawawawa','User'),(79,'did:metablox:hgsduijgbwawawawawawawawawaw','User'),(80,'did:metablox:hgsduijgbwawawawawawawawawawa','User'),(81,'did:metablox:hgsduijgbwawawawawawawawawawaw','User'),(82,'did:metablox:hgsduijgbwawawawawawawawawawawa','User'),(83,'did:metablox:hgsduijgbwawawawawawawawawawawaw','User'),(84,'did:metablox:hgsduijgbwawawawawawawawawawawawa','User'),(85,'did:metablox:hgsduijgbwawawawawawawawawawawawaw','User'),(86,'did:metablox:hgsduijgbwawawawawawawawawawawawawa','User'),(87,'did:metablox:hgsduijgbwawawawawawawawawawawawawaw','User'),(88,'did:metablox:hgsduijgbwawawawawawawawawawawawawawa','User'),(89,'did:metablox:hgsduijgbwawawawawawawawawawawawawawaw','User'),(90,'did:metablox:test','User'),(91,'did:metablox:test1','User'),(92,'did:metablox:test11','User'),(93,'did:metablox:test3','User'),(94,'did:metablox:test5','User'),(99,'did:metablox:test500','User'),(100,'did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX','Validator');
/*!40000 ALTER TABLE `WifiAccessInfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-07 16:48:46
