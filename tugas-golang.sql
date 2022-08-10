-- Adminer 4.5.0 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task` varchar(100) NOT NULL,
  `assignee` varchar(100) NOT NULL,
  `deadline` date NOT NULL,
  `status` enum('true','false','open') NOT NULL DEFAULT 'true',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `task` (`id`, `task`, `assignee`, `deadline`, `status`) VALUES
(1,	'Identify Project Stakeholders and send out memo to entire project team',	'Maruna',	'2022-08-09',	'false'),
(2,	'Identify Project Risk',	'Avicenna',	'2022-08-09',	'false'),
(3,	'Plan Risk Management',	'Wilujeng',	'2022-08-09',	'true'),
(4,	'Test',	'Wilujeng',	'2022-08-18',	'false');

-- 2022-08-10 14:11:34
