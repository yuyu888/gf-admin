CREATE TABLE `admin_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `menu_name` varchar(100) NOT NULL COMMENT '菜单名',
  `menu_path` varchar(200) NOT NULL COMMENT '菜单地址',
  `menu_type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '菜单类型：1、导航；2、页面；3、接口；4、页面元素',
  `fid` int unsigned NOT NULL DEFAULT '1' COMMENT '父菜单id',
  `sort_no` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `fid` (`fid`)
) ENGINE=InnoDB AUTO_INCREMENT=960 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单表';

CREATE TABLE `admin_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(30) NOT NULL COMMENT '角色名称',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=312 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色表';

CREATE TABLE `admin_role_menu_relation` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` int unsigned NOT NULL COMMENT '角色id',
  `menu_id` int unsigned NOT NULL COMMENT '菜单id',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `rolled_menuid` (`role_id`,`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2315 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色权限表';


CREATE TABLE `admin_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '电话号码',
  `email` varchar(64) DEFAULT NULL COMMENT '电子邮箱',
  `real_name` varchar(16) NOT NULL COMMENT '真实姓名',
  `avatar` varchar(200) DEFAULT NULL COMMENT '用户头像',
  `password` char(32) NOT NULL COMMENT '密码md5加密',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态：1，有效；2、删除',
  `department` varchar(100) NOT NULL DEFAULT '' COMMENT '所属部门',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2007 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

CREATE TABLE `admin_user_role_relation` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uid` int unsigned NOT NULL COMMENT '用户id',
  `role_id` int unsigned NOT NULL COMMENT '角色id',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid_roleid` (`role_id`,`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2170 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色关联表';

CREATE TABLE `sys_future_job` (
  `id` int unsigned NOT NULL COMMENT '自增主键',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '创建人',
  `job_code` varchar(30) NOT NULL DEFAULT '' COMMENT '任务标识',
  `params` text NOT NULL COMMENT '参数',
  `exec_time` int NOT NULL COMMENT '执行时间',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1、未开始；2、已执行；3、已作废；4、执行失败',
  `job_desc` text NOT NULL COMMENT '任务描述',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `username` (`username`),
  KEY `job_code` (`job_code`),
  KEY `status` (`status`,`exec_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


