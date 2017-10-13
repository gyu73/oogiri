CREATE TABLE `oogiris` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(255) COMMENT 'title',
  `body` varchar(255) COMMENT 'body',
  `image` varchar(255) COMMENT 'image',
  `user_id` int(11) COMMENT 'UserId',
  `created` timestamp DEFAULT NOW() COMMENT 'when created',
  `updated` timestamp DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'when last updated',
  `deleted` timestamp DEFAULT NOW(),
  PRIMARY KEY (`id`),
  CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO oogiris (
  title,
  body,
  image,
  user_id
)
VALUES
(  'Title!_7',
   'Body_7',
   'http://media-groove.com/wp-content/uploads/2016/12/aeeec022e5bab8ef4f8298daa50e697c.jpg',
   '1'
);
