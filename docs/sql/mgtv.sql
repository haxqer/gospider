
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for spider_mgtv
-- ----------------------------
DROP TABLE IF EXISTS `spider_mgtv`;
CREATE TABLE `spider_mgtv` (
  `episode_id` int(11) NOT NULL COMMENT '集 id',
  `channel_id` int(11) DEFAULT '0' COMMENT '频道 id',
  `drama_id` int(11) DEFAULT '0' COMMENT '剧 id',
  `drama_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '剧标题',
  `title1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集标题 1',
  `title2` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集标题 2',
  `title3` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集标题 3',
  `title4` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集标题 4',
  `episode_url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集 url',
  `duration` int(11) NOT NULL COMMENT '时长, 单位为 秒',
  `content_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '内容类型',
  `image` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '缩略图 url',
  `is_intact` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否完整',
  `is_new` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否最新',
  `is_vip` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否 vip',
  `play_counter` bigint(22) DEFAULT '0' COMMENT '播放数量',
  `ts` timestamp NULL DEFAULT NULL COMMENT '视频上传时间',
  `next_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '下一集 id',
  `src_clip_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '来源 id',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modify_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`episode_id`),
  KEY `idx_mtime` (`modify_time`) USING BTREE,
  KEY `idx_ts` (`ts`) USING BTREE,
  KEY `idx_episode` (`episode_id`),
  KEY `idx_drama` (`drama_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='芒果 TV 剧集';

SET FOREIGN_KEY_CHECKS = 1;
