CREATE TABLE `answers` (
  `id` int(11) AUTO_INCREMENT COMMENT 'ID',
  `body` varchar(255) COMMENT 'body',
  `user_id` int(11) COMMENT 'UserId',
  `oogiri_id` int(11) COMMENT 'OogiriId',
  `created` timestamp DEFAULT NOW() COMMENT 'when created',
  `updated` timestamp DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'when last updated',
  `deleted` timestamp DEFAULT NOW(),
  PRIMARY KEY (`id`),
  CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT FOREIGN KEY (oogiri_id) REFERENCES oogiris(id)
);

INSERT INTO answers (
  body,
  user_id,
  oogiri_id
)
VALUES
(   'Body_1',
    '1',
    '7'
);
