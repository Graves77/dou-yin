-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `record_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '记录编号,自增',
  `favorite_video_id` int(11) DEFAULT NULL COMMENT '被点赞视频id',
  `favorite_user_id` int(11) DEFAULT NULL COMMENT '点赞账号id',
  `record_time` datetime DEFAULT NULL COMMENT '记录时间',
  PRIMARY KEY (`record_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow_and_follower_list` (
  `record_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '记录编号，自增',
  `user_id` int(11) NOT NULL COMMENT '被关注的账号id',
  `follower_id` int(11) NOT NULL COMMENT '粉丝账号id',
  `record_time` datetime DEFAULT NULL COMMENT '关注时间',
  PRIMARY KEY (`record_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for user
-- ----------------------------
CREATE TABLE users (
    user_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户唯一标识', -- 用户id, 主键, 自增
    username VARCHAR(255) NOT NULL COMMENT '用户名', -- 用户名, 唯一键
    password VARCHAR(255) NOT NULL COMMENT '密码', -- 密码
    follow_count BIGINT UNSIGNED DEFAULT 0 COMMENT '用户关注总数', -- 用户关注总数
    follower_count BIGINT UNSIGNED DEFAULT 0 COMMENT '用户粉丝总数', -- 用户粉丝总数
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用户注册时间', -- 用户注册时间
    PRIMARY KEY (user_id),
    UNIQUE KEY (username) -- 保证用户名唯一
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储用户信息';


-- ----------------------------
-- Table structure for videos
-- ----------------------------
CREATE TABLE videos(
    video_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '视频唯一标识', -- 视频id, 主键, 自增
    user_id BIGINT UNSIGNED NOT NULL COMMENT '视频作者id', -- 视频作者id, 外键
    play_url VARCHAR(255) NOT NULL COMMENT '视频播放地址', -- 视频播放地址
    cover_url VARCHAR(255) NOT NULL COMMENT '视频封面地址', -- 视频封面地址
    favorite_count BIGINT UNSIGNED DEFAULT 0 COMMENT '视频的点赞总数', -- 视频的点赞总数
    comment_count BIGINT UNSIGNED DEFAULT 0 COMMENT '视频的评论总数', -- 视频的评论总数
    title	VARCHAR(30) NOT NULL COMMENT '视频标题', -- 视频标题
    PRIMARY KEY (video_id)
 )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储视频信息';
