-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: metabloxStaking
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
INSERT INTO `OrderInterest` VALUES (1,1,'2022-05-12 00:00:00',10,10,0),(2,2,'2022-05-12 01:00:00',20,20,20),(3,3,'2022-05-12 02:00:00',10,10,0);
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
  `UserAddress` varchar(45) NOT NULL,
  PRIMARY KEY (`OrderID`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Orders`
--

LOCK TABLES `Orders` WRITE;
/*!40000 ALTER TABLE `Orders` DISABLE KEYS */;
INSERT INTO `Orders` VALUES (1,1,'test','Holding',1,30,30,'',100,'userAddress'),(4,0,'','Pending',1,0,0,'placeholder',0,''),(5,0,'testDID','Pending',1,0,0,'placeholder',0,''),(6,0,'','Pending',1,0,0,'placeholder',0,''),(7,0,'','Pending',1,0,0,'placeholder',0,''),(8,0,'','Pending',1,0,0,'placeholder',0,''),(9,0,'testDID','Pending',1,0,0,'placeholder',0,''),(10,0,'testDID','Pending',1,0,0,'placeholder',0,''),(11,1,'testDID','Pending',1,0,0,'placeholder',0,''),(12,1,'testDID','Pending',1,0,0,'placeholder',1234,'testAddress');
/*!40000 ALTER TABLE `Orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PaymentInfo`
--

DROP TABLE IF EXISTS `PaymentInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PaymentInfo` (
  `PaymentAddress` varchar(45) NOT NULL,
  `Tag` varchar(20) NOT NULL,
  `CurrencyType` varchar(10) NOT NULL,
  `Network` varchar(45) NOT NULL,
  PRIMARY KEY (`PaymentAddress`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PaymentInfo`
--

LOCK TABLES `PaymentInfo` WRITE;
/*!40000 ALTER TABLE `PaymentInfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `PaymentInfo` ENABLE KEYS */;
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
-- Table structure for table `StakingProducts`
--

DROP TABLE IF EXISTS `StakingProducts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `StakingProducts` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ProductName` varchar(45) NOT NULL,
  `MinOrderValue` int NOT NULL,
  `TopUpLimit` bigint NOT NULL,
  `MinRedeemValue` int NOT NULL,
  `LockUpPeriod` int NOT NULL,
  `DefaultAPY` float NOT NULL,
  `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `StartDate` timestamp(3) NOT NULL,
  `Term` int NOT NULL,
  `BurnedInterest` float NOT NULL DEFAULT '0',
  `Status` tinyint NOT NULL DEFAULT '1',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StakingProducts`
--

LOCK TABLES `StakingProducts` WRITE;
/*!40000 ALTER TABLE `StakingProducts` DISABLE KEYS */;
INSERT INTO `StakingProducts` VALUES (1,'TestProduct',10,15,5,3,30,'2022-05-18 07:00:00.000','2022-12-25 08:00:00.000',1,50,1);
/*!40000 ALTER TABLE `StakingProducts` ENABLE KEYS */;
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
  `TXType` enum('BuyIn','Redeem','Harvest') NOT NULL,
  `TXHash` varchar(66) DEFAULT NULL,
  `Principal` float NOT NULL,
  `Interest` float NOT NULL,
  `UserAddress` varchar(45) NOT NULL,
  `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `RedeemableTime` timestamp(3) NOT NULL,
  PRIMARY KEY (`PaymentNo`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TXInfo`
--

LOCK TABLES `TXInfo` WRITE;
/*!40000 ALTER TABLE `TXInfo` DISABLE KEYS */;
INSERT INTO `TXInfo` VALUES (6,1,'MBLX','BuyIn','hfeguyewh',40,10,'sampleAddress','2022-05-18 23:58:47.000','2022-05-20 07:00:00.000'),(7,2,'MBLX','BuyIn','hfeguyewhh',0,0,'userAddress','2022-05-19 17:51:09.000','2022-05-19 17:51:09.000'),(8,3,'MBLX','BuyIn','hfeguyewbhh',0,0,'userAddress','2022-05-19 17:52:37.000','2022-11-14 18:52:37.000'),(9,4,'MBLX','BuyIn','hfeguyefwbhh',100,20,'userAddress','2022-05-19 18:09:03.000','2022-11-14 19:09:03.000'),(10,1,'MBLX','Harvest',NULL,0,0,'userAddress','2022-05-19 23:51:57.000','2022-05-19 23:51:57.000'),(11,1,'MBLX','Harvest',NULL,0,0,'userAddress','2022-05-20 00:18:53.000','2022-05-20 00:18:53.000'),(12,1,'MBLX','Harvest',NULL,0,0,'userAddress','2022-05-20 00:19:19.000','2022-05-20 00:19:19.000'),(13,1,'MBLX','Harvest',NULL,0,0,'userAddress','2022-05-20 00:19:28.000','2022-05-20 00:19:28.000'),(14,1,'MBLX','Harvest',NULL,0,0,'userAddress','2022-05-20 00:19:37.000','2022-05-20 00:19:37.000'),(15,1,'MBLX','Redeem','placeholderHash',0,0,'userAddress','2022-05-20 00:59:29.000','2022-05-20 07:00:00.000'),(16,1,'MBLX','Redeem','placeholderHash',0,0,'userAddress','2022-05-20 01:02:26.000','2022-05-20 07:00:00.000'),(17,1,'MBLX','Harvest','placeholderHash',0,0,'userAddress','2022-05-20 01:07:37.000','2022-05-20 01:07:37.000'),(18,1,'MBLX','Redeem','placeholderHash',0,0,'userAddress','2022-05-20 01:20:05.371','2022-05-20 07:00:00.000'),(19,1,'MBLX','Redeem','placeholderHash',0,0,'userAddress','2022-05-20 01:20:12.178','2022-05-20 07:00:00.000');
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
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-24 15:51:57
