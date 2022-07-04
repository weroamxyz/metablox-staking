use metabloxstakingtest;

INSERT INTO `StakingProducts` VALUES (1,'TestProduct1',10000000,500000000000,5000000,180,0.2,'2022-12-27 00:00:00.000','2022-12-27 00:00:00.000',1,0,1,'testPaymentAddress','MBLX','Ethereum',NULL);
INSERT INTO `StakingProducts` VALUES (2,'TestProduct2',10000000,500000000000,4000000,180,0.2,'2022-06-30 00:00:00.000','2022-06-30 00:00:00.000',1,0,1,'testPaymentAddress2','MBLX','Ethereum',NULL);
INSERT INTO `StakingProducts` VALUES (3,'TestProduct3',10000000,500000000000,5000000,180,0.2,'2022-12-27 00:00:00.000','2022-12-27 00:00:00.000',1,0,1,'testPaymentAddress3','MBLX','Ethereum',NULL);

INSERT INTO `TXInfo` VALUES (1, 1, 'MBLX', 'BuyIn', 'placeholderHash', 30, 50, 'placeholderUserAddress', '2022-05-27 13:25:20.955', '2022-11-21 16:00:00.000');
INSERT INTO `TXInfo` VALUES (2, 2, 'MBLX', 'BuyIn', 'placeholderHash2', 30, 50, 'placeholderUserAddress', '2022-05-27 13:25:20.955', '2022-11-21 16:00:00.000');
INSERT INTO `TXInfo` VALUES (3, 1, 'MBLX', 'InterestOnly', 'placeholderHash3', 30, 50, 'placeholderUserAddress', '2022-05-27 13:25:20.955', '2022-11-21 16:00:00.000');
INSERT INTO `TXInfo` VALUES (4, 3, 'MBLX', 'OrderClosure', 'placeholderHash4', 30, 50, 'placeholderUserAddress', '2022-05-27 13:25:20.955', '2022-11-21 16:00:00.000');

INSERT INTO `Orders` VALUES (1, 1, "did:metablox:test", "Holding", 1, 100, 50, "placeholder", 100, "placeholderUserAddress");
INSERT INTO `Orders` VALUES (2, 1, "did:metablox:test2", "Holding", 1, 100, 50, "placeholder2", 100, "placeholderUserAddress");
INSERT INTO `Orders` VALUES (3, 1, "did:metablox:test3", "Holding", 1, 100, 50, "placeholder3", 100, "placeholderUserAddress3");
INSERT INTO `Orders` VALUES (4, 1, "did:metablox:test4", "Pending", 1, 100, 50, "placeholder4", 100, "placeholderUserAddress4");
INSERT INTO `Orders` VALUES (5, 2, "did:metablox:test5", "Holding", 1, 100, 50, "placeholder5", 100, "placeholderUserAddress5");

INSERT INTO `PrincipalUpdates` VALUES (1, 3, '2022-12-27 00:00:01', 777);
INSERT INTO `PrincipalUpdates` VALUES (2, 3, '2022-12-27 00:00:02', 888);
INSERT INTO `PrincipalUpdates` VALUES (3, 3, '2022-12-27 00:00:00', 999);

INSERT INTO `OrderInterest` VALUES (1, 1, '2022-11-06 00:00:00', 111, 15, 10);
INSERT INTO `OrderInterest` VALUES (2, 1, '2022-12-06 00:00:00', 222, 30, 20);
INSERT INTO `OrderInterest` VALUES (3, 1, '2022-10-06 00:00:00', 333, 45, 30);

INSERT INTO `MinerInfo` VALUES (1, "testName", "testSSID", "testBSSID", "2022-04-19 00:00:00.000", 50, 100, 1, 25, 1, "did:metablox:test", "sampleHost", 0);
INSERT INTO `MinerInfo` VALUES (2, "testName2", "testSSID2", "testBSSID2", "2022-04-20 00:00:00.000", 75, 25, 1, 50, 1, "did:metablox:test2", "sampleHost", 0);
INSERT INTO `MinerInfo` VALUES (3, "testName3", "testSSID3", "testBSSID3", "2022-04-21 00:00:00.000", NULL, NULL, 1, 75, 1, "did:metablox:test3", "sampleHost", 1);
INSERT INTO `MinerInfo` VALUES (4, "testName4", "testSSID4", "testBSSID4", "2022-04-22 00:00:00.000", NULL, NULL, 1, 100, 1, "did:metablox:test4", "sampleHost", 1);

INSERT INTO `SeedExchangeHistory` VALUES ("did:metablox:sampleUser", "did:metablox:sampleTarget", "sampleChallenge", 50, 123, "2022-04-19 00:00:00.000");
INSERT INTO `SeedExchangeHistory` VALUES ("did:metablox:sampleUser", "did:metablox:sampleTarget2", "sampleChallenge2", 500, 1234, "2022-04-20 00:00:00.000");
INSERT INTO `SeedExchangeHistory` VALUES ("did:metablox:sampleUser2", "did:metablox:sampleTarget2", "sampleChallenge3", 70, 12345, "2022-04-21 00:00:00.000");

INSERT INTO `ExchangeRate` VALUES ("1", 721, "2022-04-19 00:00:00.000");

INSERT INTO `WifiAccessInfo` VALUES ("1", "did:metablox:test", "Validator");
INSERT INTO `WifiAccessInfo` VALUES ("2", "did:metablox:test2", "User");

INSERT INTO `MiningLicenseInfo` VALUES ("1", "did:metablox:test", "sampleName", "sampleModel", "sampleSerial");