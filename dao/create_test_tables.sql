-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: localhost    Database: metabloxstakingtest
-- ------------------------------------------------------
-- Server version	8.0.27-0ubuntu0.20.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `metabloxstakingtest`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `metabloxstakingtest` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `metabloxstakingtest`;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
-- Table structure for table `PrincipalUpdates`
--

DROP TABLE IF EXISTS `PrincipalUpdates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PrincipalUpdates` (
                                    `ID` int NOT NULL AUTO_INCREMENT,
                                    `ProductID` int NOT NULL,
                                    `Time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
                                    `TotalPrincipal` float NOT NULL,
                                    PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
                                   `MinOrderValue` int NOT NULL,
                                   `TopUpLimit` bigint NOT NULL,
                                   `MinRedeemValue` int NOT NULL,
                                   `LockUpPeriod` int NOT NULL,
                                   `DefaultAPY` float NOT NULL,
                                   `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
                                   `StartDate` timestamp(3) NOT NULL,
                                   `Term` int NOT NULL,
                                   `BurnedInterest` float NOT NULL DEFAULT '0',
                                   `NextProductID` int NULL,
                                   `Status` tinyint NOT NULL DEFAULT '1',
                                   PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
                          `TXType` enum('BuyIn','Redeem','OrderClosure') NOT NULL,
                          `TXHash` varchar(66) DEFAULT NULL,
                          `Principal` float NOT NULL,
                          `Interest` float NOT NULL,
                          `UserAddress` varchar(45) NOT NULL,
                          `CreateDate` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
                          `RedeemableTime` timestamp(3) NOT NULL,
                          PRIMARY KEY (`PaymentNo`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `OrderInterest`;
CREATE TABLE `OrderInterest` (
                                 `ID` int NOT NULL AUTO_INCREMENT,
                                 `OrderID` int NOT NULL,
                                 `Time` timestamp(3) NOT NULL,
                                 `APY` float NOT NULL,
                                 `InterestGain` float NOT NULL,
                                 `TotalInterestGain` float NOT NULL,
                                 PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-02 17:25:35
