CREATE TABLE t_license_url_suffix(
    id INT AUTO_INCREMENT PRIMARY KEY COMMENt "ID",
    url_suffix VARCHAR(1023) NOT NULL COMMENT "license_url suffix"
)ENGINE=InnoDB CHARSET=utf8mb4

CREATE TABLE t_source_web_page_license_keyword(
    id INT AUTO_INCREMENT PRIMARY KEY COMMENt "ID",
    keyword VARCHAR(1023) NOT NULL COMMENT "license link in webpage keyword"
)ENGINE=InnoDB CHARSET=utf8mb4



CREATE TABLE t_source_web_page_copyright_keyword(
    id INT AUTO_INCREMENT PRIMARY KEY COMMENt "ID",
    keyword VARCHAR(1023) NOT NULL COMMENT "copyrignt flag keyword"
)ENGINE=InnoDB CHARSET=utf8mb4


CREATE TABLE `t_license_mes` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `license_name` varchar(1023) NOT NULL COMMENT 'license name',
  `license_url` varchar(1023) NOT NULL COMMENT 'license url',
  `license_type` int(11) DEFAULT NULL COMMENT 'license type 10:data license 20:specific license 30:TOU',
  `copyright_flag` varchar(1023) NOT NULL COMMENT 'copyrigt flag in webpage',
  `licensor` varchar(1023) NOT NULL COMMENT 'licenseor',
  `license_content` longtext COMMENT 'license content',
  `source_url` varchar(1023) DEFAULT NULL COMMENT 'source url',
  `aibom_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;