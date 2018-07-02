

CREATE TABLE IF NOT EXISTS `txvin`(
   `Txid` VARCHAR(40) not null,
   `txHash` VARCHAR(40) NOT NULL,
   `pretxid` VARCHAR(40) NOT NULL,
   `Vout` int,
	 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY ( `Txid` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC;


CREATE TABLE IF NOT EXISTS `txvout`(
   `Txid` VARCHAR(40) not null,
   `txHash` VARCHAR(40) NOT NULL,
   `Value` varchar(255) NOT NULL,
   `N` int,
	 `address` VARCHAR(40) NOT NULL,
	 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY ( `Txid`,`address` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC;



CREATE TABLE IF NOT EXISTS `btcTransaction`(
   `Txid` VARCHAR(40) not null,
   `txHash` VARCHAR(40) NOT NULL,
   `Version` int NOT NULL,
   `Size` int,
	 `Vsize` int,
	 `Locktime` int,
	 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY ( `Txid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC;


CREATE TABLE IF NOT EXISTS `btcblockinfo`(
   `Hash` VARCHAR(40) not null,
   `Version` int NOT NULL,
	 `Time` int,
   `Size` int,
	 `Vsize` int,
	 `Nonce` int,
   `Weight` int,
	 `Difficulty` int,
	 `Merkleroot` VARCHAR(40) not null,
	 `NextBlockhash` VARCHAR(40) not null,
	 `Confirmations` int not null,
	 `Height`int not null,	 
	 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY ( `Hash`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC;



