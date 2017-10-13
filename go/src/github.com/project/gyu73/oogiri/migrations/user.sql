CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) COMMENT 'name',
  `email` varchar(255) COMMENT 'email',
  `salt` varchar(255) COMMENT  'salt',
  `created` timestamp DEFAULT NOW() COMMENT 'when created',
  `updated` timestamp DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'when last updated',
  `deleted` timestamp DEFAULT NOW(),
  PRIMARY KEY (`id`),
  UNIQUE KEY (`email`)
);

INSERT INTO users (
  name,
  email,
  salt
)
VALUES (
  'name_1',
  'email_3',
  'passwordA'
);
