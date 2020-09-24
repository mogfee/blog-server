CREATE TABLE `blog_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '文章简述',
  `cover_image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片地址',
  `content` longtext COMMENT '文章内容',
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否删除 0未删除 1删除',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0禁用 1启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章'

CREATE TABLE `blog_article_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
  `tag_id` int unsigned NOT NULL DEFAULT '0' COMMENT '标签id',
  `cover_image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片地址',
  `content` longtext COMMENT '文章内容',
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否删除 0未删除 1删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章标签关联表'

CREATE TABLE `blog_tag` (
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否删除 0未删除 1删除',
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0禁用 1启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='标签'

