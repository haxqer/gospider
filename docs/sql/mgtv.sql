SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for spider_mgtv
-- ----------------------------
DROP TABLE IF EXISTS `spider_mgtv`;
CREATE TABLE `spider_mgtv` (
  `episode_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '集id',
  `channel_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '频道id',
  `drama_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '剧 id',
  `drama_title` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '剧标题',
  `title1` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集标题1',
  `title2` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集标题2',
  `title3` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集标题3',
  `title4` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集标题4',
  `episode_url` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '集 url',
  `duration` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '时长',
  `content_type` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '内容类型',
  `image` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '缩略图url',
  `is_intact` char(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否完整',
  `is_new` char(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否最新',
  `is_vip` char(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否vip',
  `play_counter` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '播放数量',
  `ts` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频上传时间',
  `next_id` char(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '下一集id',
  `src_clip_id` char(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '来源id',
  PRIMARY KEY (`episode_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='芒果 TV 剧集';

SET FOREIGN_KEY_CHECKS = 1;
