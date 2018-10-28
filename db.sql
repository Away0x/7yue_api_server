/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table book
# ------------------------------------------------------------

DROP TABLE IF EXISTS `book`;

CREATE TABLE `book` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `book_id` int(11) NOT NULL,
  `author` varchar(128) NOT NULL,
  `image` varchar(128) NOT NULL,
  `title` varchar(128) NOT NULL,
  `isbn` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_book_deleted_at` (`deleted_at`),
  KEY `isbn` (`isbn`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

LOCK TABLES `book` WRITE;
/*!40000 ALTER TABLE `book` DISABLE KEYS */;

INSERT INTO `book` (`id`, `created_at`, `updated_at`, `deleted_at`, `book_id`, `author`, `image`, `title`, `isbn`)
VALUES
	(1,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,6021440,'[美]保罗·格雷厄姆','https://img3.doubanio.com/lpic/s4669554.jpg','黑客与画家','9787115249494'),
	(2,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,3628911,'MarkPilgrim','https://img3.doubanio.com/lpic/s4059293.jpg','Dive Into Python 3','9781430224150'),
	(3,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,4866934,'MagnusLieHetland','https://img3.doubanio.com/lpic/s4387251.jpg','Python基础教程','9787115230270'),
	(4,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,6082808,'[哥伦比亚]加西亚·马尔克斯','https://img3.doubanio.com/lpic/s6384944.jpg','百年孤独','9787544253994'),
	(5,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,1080370,'[日]岩井俊二','https://img1.doubanio.com/view/subject/l/public/s29775868.jpg','情书','9787201048161'),
	(6,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,1336330,'[美]乔治·R·R·马丁','https://img3.doubanio.com/lpic/s1358984.jpg','冰与火之歌（卷一）','9787536671256'),
	(7,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,3259440,'[日]东野圭吾','https://img3.doubanio.com/lpic/s4610502.jpg','白夜行','9787544242516'),
	(8,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,1255625,'金庸','https://img1.doubanio.com/lpic/s23632058.jpg','天龙八部','9787108006721'),
	(9,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,3646172,'[日]东野圭吾','https://img3.doubanio.com/lpic/s3814606.jpg','恶意','9787544244428'),
	(10,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,1071241,'[英]J·K·罗琳','https://img3.doubanio.com/lpic/s1074376.jpg','哈利·波特与阿兹卡班的囚徒','9787020033454'),
	(11,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,3394338,'韩寒','https://img1.doubanio.com/lpic/s3557848.jpg','他的国','9787807592099'),
	(12,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,2295163,'[英]J·K·罗琳','https://img1.doubanio.com/lpic/s2752367.jpg','哈利·波特与死亡圣器','9787020063659'),
	(13,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,3264665,'王小波','https://img1.doubanio.com/lpic/s3463069.jpg','三十而立','9787545201475'),
	(14,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,4071859,'[伊朗]玛赞·莎塔碧','https://img3.doubanio.com/lpic/s6144591.jpg','我在伊朗长大','9787108033215'),
	(15,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,6126821,'[日]村上春树','https://img1.doubanio.com/lpic/s29494718.jpg','远方的鼓声','9787532754533'),
	(16,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,2070844,'三毛','https://img3.doubanio.com/lpic/s2393243.jpg','梦里花落知多少','9787530208922'),
	(17,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,1009160,'韩寒','https://img1.doubanio.com/lpic/s1080179.jpg','像少年啦飞驰','9787506322522'),
	(18,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,1100152,'鲁迅','https://img3.doubanio.com/lpic/s27970504.jpg','朝花夕拾','9787533914196'),
	(19,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,1221515,'[日]井上雄彦','https://img3.doubanio.com/lpic/s2853431.jpg','灌篮高手31','9787806649343'),
	(20,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,26871297,'[日]新井一二三','https://img3.doubanio.com/lpic/s29034294.jpg','东京时味记','9787544762069');

/*!40000 ALTER TABLE `book` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table book_comment
# ------------------------------------------------------------

DROP TABLE IF EXISTS `book_comment`;

CREATE TABLE `book_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `content` varchar(255) NOT NULL,
  `book_id` int(11) NOT NULL,
  `nums` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `book_id` (`book_id`),
  KEY `idx_book_comment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8;

LOCK TABLES `book_comment` WRITE;
/*!40000 ALTER TABLE `book_comment` DISABLE KEYS */;

INSERT INTO `book_comment` (`id`, `created_at`, `updated_at`, `deleted_at`, `content`, `book_id`, `nums`)
VALUES
	(4,'2018-10-02 14:04:58','2018-10-27 17:59:16',NULL,'程序员也有艺术的人生',6021440,1004),
	(5,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'黑客and',6021440,475),
	(6,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'七月老师好',6021440,107),
	(7,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'a',6021440,28),
	(8,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'一二三四五六七八九十十一',6021440,26),
	(9,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'666',6021440,10),
	(10,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'好好的',6021440,8),
	(11,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'艺术',6021440,7),
	(12,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'555',6021440,6),
	(13,'2018-10-02 14:04:59','2018-10-02 14:07:57',NULL,'ABC',6021440,6),
	(14,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'来个沙发！',3628911,478),
	(15,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'888',3628911,167),
	(16,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'测试一下',3628911,102),
	(17,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'444',3628911,2),
	(18,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'我喜欢Python很久了',3628911,2),
	(19,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'a',3628911,2),
	(20,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'测试',3628911,2),
	(21,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'刑警本色',3628911,1),
	(22,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'qwewqe',3628911,1),
	(23,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'你的',3628911,1),
	(24,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'人生苦短，我用Py',4866934,292),
	(25,'2018-10-02 14:04:59','2018-10-27 17:50:56',NULL,'I like Py',4866934,114),
	(26,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'life',4866934,53),
	(27,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'lifeisshort',4866934,2),
	(28,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'你好，七月',4866934,2),
	(29,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'七月，好python ',4866934,1),
	(30,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'经典致敬',4866934,1),
	(31,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'还不错',4866934,1),
	(32,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'cool',4866934,1),
	(33,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'公司',4866934,1),
	(34,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'魔幻第一书',6082808,173),
	(35,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'永恒的经典',6082808,171),
	(36,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'一百年，一世纪',6082808,39),
	(37,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'Nice',6082808,1),
	(38,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'我很喜欢',6082808,1),
	(39,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'推图',6082808,1),
	(40,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'啦咯啦咯啦咯啦咯啦咯啦',6082808,1),
	(41,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'凝聚态',6082808,1),
	(42,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'怎么了？',6082808,1),
	(43,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'我写作业了',6082808,1),
	(44,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'致敬经典',1080370,109),
	(45,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'天亦老',1080370,60),
	(46,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'人生亦相逢',1080370,26),
	(47,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'爱而不得',1080370,3),
	(48,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'眼泪成诗',1080370,2),
	(49,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'程序员',1080370,2),
	(50,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'大师经典之作',1080370,1),
	(51,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'123',1080370,1),
	(52,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'岩井俊二',1080370,1),
	(53,'2018-10-02 14:04:59','2018-10-02 14:04:59',NULL,'大师',1080370,1),
	(54,'2018-10-02 14:07:43','2018-10-02 14:07:49',NULL,'wutong',6021440,2),
	(55,'2018-10-27 17:44:59','2018-10-27 17:44:59',NULL,'测试',1071241,1),
	(56,'2018-10-27 18:05:40','2018-10-27 18:05:40',NULL,'123',1100152,1);

/*!40000 ALTER TABLE `book_comment` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table classic
# ------------------------------------------------------------

DROP TABLE IF EXISTS `classic`;

CREATE TABLE `classic` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `content` varchar(255) NOT NULL,
  `image` varchar(128) NOT NULL,
  `url` varchar(128) NOT NULL,
  `index` int(11) NOT NULL,
  `title` varchar(128) NOT NULL,
  `type` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_classic_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

LOCK TABLES `classic` WRITE;
/*!40000 ALTER TABLE `classic` DISABLE KEYS */;

INSERT INTO `classic` (`id`, `created_at`, `updated_at`, `deleted_at`, `content`, `image`, `url`, `index`, `title`, `type`)
VALUES
	(1,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'谁念过 千字文章 秋收冬已藏','music.7.png','http://music.163.com/song/media/outer/url?id=29719651.mp3',7,'不才 《参商》',200),
	(2,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'心上无垢，林间有风','sentence.6.png','',6,'未名',300),
	(3,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'许多人来来去去，相聚又别离','music.5.png','http://music.163.com/song/media/outer/url?id=26427662.mp3',5,'好妹妹 《一个人的北京》',200),
	(4,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'在变换的生命里，岁月原来是最大的小偷','movie.4.png','',4,'罗启锐《岁月神偷》',100),
	(5,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'你陪我步入蝉夏 越过城市喧嚣','music.1.png','http://music.163.com/song/media/outer/url?id=557581284.mp3',3,'花粥 《纸短情长》',200),
	(6,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'这个夏天又是一个毕业季','sentence.2.png','',2,'未名',300),
	(7,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'岁月长，衣裳薄','music.3.png','http://music.163.com/song/media/outer/url?id=317245.mp3',1,'杨千嬅《再见二丁目》',200),
	(8,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'人生不能像做菜，把所有的料准备好才下锅','movie.8.png','',8,'李安《饮食男女 》',100);

/*!40000 ALTER TABLE `classic` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table favor
# ------------------------------------------------------------

DROP TABLE IF EXISTS `favor`;

CREATE TABLE `favor` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_key` varchar(255) NOT NULL,
  `type` int(11) NOT NULL,
  `target_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_favor_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8;

LOCK TABLES `favor` WRITE;
/*!40000 ALTER TABLE `favor` DISABLE KEYS */;

INSERT INTO `favor` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_key`, `type`, `target_id`)
VALUES
	(15,'2018-10-01 20:35:13','2018-10-01 20:35:13','2018-10-01 20:37:59','admin',400,1002),
	(16,'2018-10-01 20:35:26','2018-10-01 20:35:26',NULL,'admin1',400,1002),
	(17,'2018-10-01 20:38:19','2018-10-01 20:38:19','2018-10-01 20:39:25','admin',400,1002),
	(18,'2018-10-27 17:07:59','2018-10-27 17:07:59','2018-10-27 17:28:26','admin',400,1080370),
	(19,'2018-10-27 17:19:24','2018-10-27 17:19:24','2018-10-27 17:22:54','admin',400,6021440),
	(20,'2018-10-27 17:24:46','2018-10-27 17:24:46','2018-10-27 17:24:50','admin',400,6021440),
	(21,'2018-10-27 17:24:52','2018-10-27 17:24:52','2018-10-27 17:24:56','admin',400,6021440),
	(22,'2018-10-27 17:26:13','2018-10-27 17:26:13','2018-10-27 17:26:20','admin',400,3628911),
	(23,'2018-10-27 17:26:57','2018-10-27 17:26:57','2018-10-27 17:27:15','admin',400,4866934),
	(24,'2018-10-27 17:27:51','2018-10-27 17:27:51',NULL,'admin',400,6082808),
	(25,'2018-10-27 17:31:18','2018-10-27 17:31:18','2018-10-27 17:31:53','admin',400,3646172),
	(26,'2018-10-27 17:34:37','2018-10-27 17:34:37','2018-10-27 17:34:48','admin',400,2070844),
	(27,'2018-10-27 17:35:50','2018-10-27 17:35:50',NULL,'admin',400,26871297),
	(28,'2018-10-27 17:36:34','2018-10-27 17:36:34','2018-10-27 17:36:38','admin',400,2295163),
	(29,'2018-10-28 17:54:42','2018-10-28 17:54:42',NULL,'admin',100,8),
	(30,'2018-10-28 17:58:00','2018-10-28 17:58:00',NULL,'admin',200,1),
	(31,'2018-10-28 18:28:46','2018-10-28 18:28:46',NULL,'admin',300,2);

/*!40000 ALTER TABLE `favor` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table hot_keyword
# ------------------------------------------------------------

DROP TABLE IF EXISTS `hot_keyword`;

CREATE TABLE `hot_keyword` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `text` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `text` (`text`),
  KEY `idx_hot_keyword_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

LOCK TABLES `hot_keyword` WRITE;
/*!40000 ALTER TABLE `hot_keyword` DISABLE KEYS */;

INSERT INTO `hot_keyword` (`id`, `created_at`, `updated_at`, `deleted_at`, `text`)
VALUES
	(1,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'Python'),
	(2,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'哈利·波特'),
	(3,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'村上春树'),
	(4,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'东野圭吾'),
	(5,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'白夜行'),
	(6,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'韩寒'),
	(7,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'金庸'),
	(8,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'王小波');

/*!40000 ALTER TABLE `hot_keyword` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `key` (`key`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `key`)
VALUES
	(1,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'wutong','admin'),
	(2,'2018-10-01 19:32:22','2018-10-01 19:32:22',NULL,'wutong1','admin1'),
	(7,'2018-10-02 22:08:57','2018-10-02 22:08:57',NULL,'cyj','$2a$10$O1CJOJNTyp4K2HYU8e37OerDgItPr.EDOZt1NGeT/QbWWJM3ASbwe'),
	(8,'2018-10-03 10:45:24','2018-10-03 10:45:24',NULL,'111','$2a$10$UZKE39yi2HJbVHJ7cTNEhukpwN4hVxiizA5g4nVkwULcVTDq3H4.q'),
	(10,'2018-10-03 10:45:28','2018-10-03 10:45:28',NULL,'1','$2a$10$7J/Jdhe2mIsmyk5yU9gCjO6b9EC9pVbmrkh/M5pBW6RBWlkZojc9i'),
	(13,'2018-10-06 15:38:28','2018-10-06 15:38:28',NULL,'xxx','$2a$10$6CwS1ZWysRQswCslghPI3.gJsSDVBJfEjO4aiW5KHxBJn26rPwB/K'),
	(14,'2018-10-06 23:49:23','2018-10-06 23:49:23',NULL,'xtx','$2a$10$SHpoArIkg/quuroPnBopceDE4PlruiY1JIRXtzc2TGzv8va/ygYdm'),
	(15,'2018-10-28 12:45:48','2018-10-28 12:45:48',NULL,'王境泽','$2a$10$01szjVjcBbsbQwTaCHglR.W.xrI.m6iUgZS20.0Bn4mbtYAyBa0ZW');

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
