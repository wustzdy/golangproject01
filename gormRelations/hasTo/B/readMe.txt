如果你想要使用另一个字段来保存该关系，你同样可以使用标签 foreignKey 来更改它，例如：

type User struct {
  gorm.Model
  CreditCard CreditCard `gorm:"foreignKey:UserName"`
  // use UserName as foreign key
}

type CreditCard struct {
  gorm.Model
  Number   string
  UserName string
}

//也就是表CreditCard 的UserName 存的就是Users表中的主键ID
可以看到，CreditCard表最后的字段确实是变成了user_name，但是user_name还是1 ，也就是说，
还是把User的ID字段直接拿过来放在了里面，我们可以再确认一下，结构体定义不变，看下方代码
show create table users
CREATE TABLE `users` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`created_at` datetime(3) DEFAULT NULL,
`updated_at` datetime(3) DEFAULT NULL,
`deleted_at` datetime(3) DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

show create table credit_cards
CREATE TABLE `credit_cards` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`created_at` datetime(3) DEFAULT NULL,
`updated_at` datetime(3) DEFAULT NULL,
`deleted_at` datetime(3) DEFAULT NULL,
`number` longtext,
`user_name` bigint unsigned DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `idx_credit_cards_deleted_at` (`deleted_at`),
KEY `fk_users_credit_card` (`user_name`),
CONSTRAINT `fk_users_credit_card` FOREIGN KEY (`user_name`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
