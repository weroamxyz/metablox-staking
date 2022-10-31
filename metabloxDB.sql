-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: localhost    Database: foundationService
-- ------------------------------------------------------
-- Server version	8.0.29-0ubuntu0.20.04.3

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
) ENGINE=InnoDB AUTO_INCREMENT=160 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ExchangeHistory`
--

DROP TABLE IF EXISTS `ExchangeHistory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ExchangeHistory` (
  `ID` int NOT NULL,
  `ExchangeRate` float NOT NULL,
  `CreateTime` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ExchangeHistory`
--

DROP TABLE IF EXISTS `ExchangeHistory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ExchangeHistory` (
  `ID` int NOT NULL,
  `ExchangeRate` float NOT NULL,
  `CreateTime` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ExchangeHistory`
--

LOCK TABLES `ExchangeHistory` WRITE;
/*!40000 ALTER TABLE `ExchangeHistory` DISABLE KEYS */;
/*!40000 ALTER TABLE `ExchangeHistory` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ExchangeRate`
--

DROP TABLE IF EXISTS `ExchangeRate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ExchangeRate` (
  `ID` int NOT NULL,
  `ExchangeRate` decimal(20,0) NOT NULL,
  `CreateTime` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  `Password` varchar(100) NOT NULL,
  `DeviceName` varchar(100) NOT NULL,
  `Address` varchar(100) NOT NULL,
  `RewardEarned` varchar(100) NOT NULL,
  `SignalStrength` varchar(100) NOT NULL,
  `Availability` tinyint(1) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=667 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `OrderInterest`
--

DROP TABLE IF EXISTS `OrderInterest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `OrderInterest` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `OrderID` int NOT NULL,
  `Time` timestamp(3) NOT NULL,
  `APY` float NOT NULL,
  `InterestGain` decimal(20,0) NOT NULL,
  `TotalInterestGain` decimal(20,0) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=12951 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `Orders`
--

DROP TABLE IF EXISTS `Orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Orders` (
  `OrderID` int NOT NULL AUTO_INCREMENT,
  `ProductID` int NOT NULL,
  `UserDID` varchar(100) NOT NULL,
  `Type` enum('Pending','Holding','Complete') NOT NULL,
  `Term` int DEFAULT NULL,
  `AccumulatedInterest` decimal(20,0) NOT NULL DEFAULT '0',
  `TotalInterestGained` decimal(20,0) NOT NULL DEFAULT '0',
  `PaymentAddress` varchar(45) NOT NULL,
  `Amount` decimal(20,0) NOT NULL,
  `UserAddress` varchar(80) NOT NULL,
  PRIMARY KEY (`OrderID`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `PrincipalUpdates`
--

DROP TABLE IF EXISTS `PrincipalUpdates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PrincipalUpdates` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ProductID` int NOT NULL,
  `Time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `TotalPrincipal` decimal(20,0) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `SeedExchangeHistory`
--

DROP TABLE IF EXISTS `SeedExchangeHistory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SeedExchangeHistory` (
  `UserDID` varchar(100) NOT NULL,
  `TargetDID` varchar(100) NOT NULL,
  `Challenge` varchar(45) NOT NULL,
  `ExchangeRate` float NOT NULL,
  `Amount` decimal(20,0) NOT NULL,
  `CreateTime` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`UserDID`,`TargetDID`,`Challenge`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `StakingProducts`
--

DROP TABLE IF EXISTS `StakingProducts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `StakingProducts` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ProductName` varchar(45) NOT NULL,
  `MinOrderValue` decimal(20,0) NOT NULL,
  `TopUpLimit` decimal(20,0) NOT NULL,
  `MinRedeemValue` decimal(20,0) NOT NULL,
  `LockUpPeriod` int NOT NULL,
  `DefaultAPY` float NOT NULL,
  `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `StartDate` timestamp(3) NOT NULL,
  `Term` int NOT NULL,
  `BurnedInterest` decimal(20,0) NOT NULL DEFAULT '0',
  `Status` tinyint NOT NULL DEFAULT '1',
  `PaymentAddress` varchar(45) NOT NULL,
  `CurrencyType` varchar(10) NOT NULL,
  `Network` varchar(45) NOT NULL,
  `NextProductID` int DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  `Principal` decimal(20,0) NOT NULL,
  `Interest` decimal(20,0) NOT NULL,
  `UserAddress` varchar(45) NOT NULL,
  `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `RedeemableTime` timestamp(3) NOT NULL,
  PRIMARY KEY (`PaymentNo`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  `SSID` varchar(100) NOT NULL,
  `UserName` varchar(100) NOT NULL,
  `Password` varchar(100) NOT NULL,
  PRIMARY KEY (`CredentialID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-04 17:23:31
