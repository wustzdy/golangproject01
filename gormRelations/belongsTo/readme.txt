// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"`
	//// 使用 CompanyRefer 作为外键
}

type Company struct {
	ID   int
	Name string
}

# 转储表 companies
# ------------------------------------------------------------

CREATE TABLE `companies` (
`id` bigint NOT NULL AUTO_INCREMENT,
`name` longtext,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

# 转储表 users
# ------------------------------------------------------------

CREATE TABLE `users` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`created_at` datetime(3) DEFAULT NULL,
`updated_at` datetime(3) DEFAULT NULL,
`deleted_at` datetime(3) DEFAULT NULL,
`name` longtext,
`company_refer` bigint DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `idx_users_deleted_at` (`deleted_at`),
KEY `fk_users_company` (`company_refer`),
CONSTRAINT `fk_users_company` FOREIGN KEY (`company_refer`) REFERENCES `companies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;