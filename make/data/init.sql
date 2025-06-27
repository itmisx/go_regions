CREATE TABLE `regions` (
  `id` bigint NOT NULL /*编码*/,
  `pid` bigint NOT NULL /*父级编码*/,
  `level` tinyint(4) NOT NULL /*1-省 2-市 3-区 4-街道*/,
  `pinyin_prefix` varchar(1) NOT NULL /*拼音前缀*/,
  `pinyin` varchar(50) NOT NULL /*拼音*/,
  `name` varchar(100) DEFAULT NULL /*完整编码*/,
  PRIMARY KEY (`id`)
);
CREATE INDEX `idx_pid` ON `regions` (`pid`);