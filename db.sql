-- MySQL dump 10.13  Distrib 5.7.18, for osx10.12 (x86_64)
--
-- Host: localhost    Database: qiyue_api_server
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `book`
--

DROP TABLE IF EXISTS `book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book`
--

LOCK TABLES `book` WRITE;
/*!40000 ALTER TABLE `book` DISABLE KEYS */;
INSERT INTO `book` VALUES (1,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,7,'[美]保罗·格雷厄姆','https://img3.doubanio.com/lpic/s4669554.jpg','黑客与画家','9787115249494'),(2,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,65,'MarkPilgrim','https://img3.doubanio.com/lpic/s4059293.jpg','Dive Into Python 3','9781430224150'),(3,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,183,'MagnusLieHetland','https://img3.doubanio.com/lpic/s4387251.jpg','Python基础教程','9787115230270'),(4,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1002,'[哥伦比亚]加西亚·马尔克斯','https://img3.doubanio.com/lpic/s6384944.jpg','百年孤独','9787544253994'),(5,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1049,'[日]岩井俊二','https://img1.doubanio.com/view/subject/l/public/s29775868.jpg','情书','9787201048161'),(6,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1061,'[美]乔治·R·R·马丁','https://img3.doubanio.com/lpic/s1358984.jpg','冰与火之歌（卷一）','9787536671256'),(7,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1120,'[日]东野圭吾','https://img3.doubanio.com/lpic/s4610502.jpg','白夜行','9787544242516'),(8,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1166,'金庸','https://img1.doubanio.com/lpic/s23632058.jpg','天龙八部','9787108006721'),(9,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1308,'[日]东野圭吾','https://img3.doubanio.com/lpic/s3814606.jpg','恶意','9787121224683'),(10,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1339,'[英]J·K·罗琳','https://img3.doubanio.com/lpic/s1074376.jpg','哈利·波特与阿兹卡班的囚徒','9787020033454'),(11,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1383,'韩寒','https://img1.doubanio.com/lpic/s3557848.jpg','他的国','9787807592099'),(12,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1398,'[英]J·K·罗琳','https://img1.doubanio.com/lpic/s2752367.jpg','哈利·波特与死亡圣器','9787020063659'),(13,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,1560,'王小波','https://img1.doubanio.com/lpic/s3463069.jpg','三十而立','9787545201475'),(14,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,7821,'[伊朗]玛赞·莎塔碧','https://img3.doubanio.com/lpic/s6144591.jpg','我在伊朗长大','9787108033215'),(15,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,8854,'[日]村上春树','https://img1.doubanio.com/lpic/s29494718.jpg','远方的鼓声','9787532754533'),(16,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,8866,'三毛','https://img3.doubanio.com/lpic/s2393243.jpg','梦里花落知多少','9787531325093'),(17,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,15198,'韩寒','https://img1.doubanio.com/lpic/s1080179.jpg','像少年啦飞驰','9787506322522'),(18,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,15984,'鲁迅','https://img3.doubanio.com/lpic/s27970504.jpg','朝花夕拾','9787533914196'),(19,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,21050,'[日]井上雄彦','https://img3.doubanio.com/lpic/s2853431.jpg','灌篮高手31','9787806649343'),(20,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,51664,'[日]新井一二三','https://img3.doubanio.com/lpic/s29034294.jpg','东京时味记','9787544762069');
/*!40000 ALTER TABLE `book` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `book_comment`
--

DROP TABLE IF EXISTS `book_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book_comment`
--

LOCK TABLES `book_comment` WRITE;
/*!40000 ALTER TABLE `book_comment` DISABLE KEYS */;
INSERT INTO `book_comment` VALUES (1,'2018-10-01 13:46:13','2018-10-01 13:46:13',NULL,'123',2,1),(2,'2018-10-01 13:46:32','2018-10-01 13:46:39',NULL,'456',2,2),(3,'2018-10-01 13:46:53','2018-10-01 13:46:53',NULL,'456',23,1),(4,'2018-10-02 06:04:58','2018-10-02 06:04:58',NULL,'程序员也有艺术的人生',7,1001),(5,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'黑客and',7,475),(6,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'七月老师好',7,107),(7,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'a',7,28),(8,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'一二三四五六七八九十十一',7,26),(9,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'666',7,10),(10,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'好好的',7,8),(11,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'艺术',7,7),(12,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'555',7,6),(13,'2018-10-02 06:04:59','2018-10-02 06:07:57',NULL,'ABC',7,6),(14,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'来个沙发！',65,478),(15,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'888',65,167),(16,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'测试一下',65,102),(17,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'444',65,2),(18,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'我喜欢Python很久了',65,2),(19,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'a',65,2),(20,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'测试',65,2),(21,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'刑警本色',65,1),(22,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'qwewqe',65,1),(23,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'你的',65,1),(24,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'人生苦短，我用Py',183,292),(25,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'I like Py',183,113),(26,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'life',183,53),(27,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'lifeisshort',183,2),(28,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'你好，七月',183,2),(29,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'七月，好python ',183,1),(30,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'经典致敬',183,1),(31,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'还不错',183,1),(32,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'cool',183,1),(33,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'公司',183,1),(34,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'魔幻第一书',1002,173),(35,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'永恒的经典',1002,171),(36,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'一百年，一世纪',1002,39),(37,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'Nice',1002,1),(38,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'我很喜欢',1002,1),(39,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'推图',1002,1),(40,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'啦咯啦咯啦咯啦咯啦咯啦',1002,1),(41,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'凝聚态',1002,1),(42,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'怎么了？',1002,1),(43,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'我写作业了',1002,1),(44,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'致敬经典',1049,109),(45,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'天亦老',1049,60),(46,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'人生亦相逢',1049,26),(47,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'爱而不得',1049,3),(48,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'眼泪成诗',1049,2),(49,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'程序员',1049,2),(50,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'大师经典之作',1049,1),(51,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'123',1049,1),(52,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'岩井俊二',1049,1),(53,'2018-10-02 06:04:59','2018-10-02 06:04:59',NULL,'大师',1049,1),(54,'2018-10-02 06:07:43','2018-10-02 06:07:49',NULL,'wutong',7,2);
/*!40000 ALTER TABLE `book_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `classic`
--

DROP TABLE IF EXISTS `classic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `classic`
--

LOCK TABLES `classic` WRITE;
/*!40000 ALTER TABLE `classic` DISABLE KEYS */;
INSERT INTO `classic` VALUES (1,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'谁念过 千字文章 秋收冬已藏','music.7.png','http://music.163.com/song/media/outer/url?id=29719651.mp3',7,'不才 《参商》',200),(2,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'心上无垢，林间有风','sentence.6.png','',6,'未名',300),(3,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'许多人来来去去，相聚又别离','music.5.png','http://music.163.com/song/media/outer/url?id=26427662.mp3',5,'好妹妹 《一个人的北京》',200),(4,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'在变换的生命里，岁月原来是最大的小偷','movie.4.png','',4,'罗启锐《岁月神偷》',100),(5,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'你陪我步入蝉夏 越过城市喧嚣','music.1.png','http://music.163.com/song/media/outer/url?id=557581284.mp3',3,'花粥 《纸短情长》',200),(6,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'这个夏天又是一个毕业季','sentence.2.png','',2,'未名',300),(7,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'岁月长，衣裳薄','music.3.png','http://music.163.com/song/media/outer/url?id=317245.mp3',1,'杨千嬅《再见二丁目》',200),(8,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'人生不能像做菜，把所有的料准备好才下锅','movie.8.png','',8,'李安《饮食男女 》',100);
/*!40000 ALTER TABLE `classic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `favor`
--

DROP TABLE IF EXISTS `favor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `favor`
--

LOCK TABLES `favor` WRITE;
/*!40000 ALTER TABLE `favor` DISABLE KEYS */;
INSERT INTO `favor` VALUES (15,'2018-10-01 12:35:13','2018-10-01 12:35:13','2018-10-01 12:37:59','admin',400,1002),(16,'2018-10-01 12:35:26','2018-10-01 12:35:26',NULL,'admin1',400,1002),(17,'2018-10-01 12:38:19','2018-10-01 12:38:19','2018-10-01 12:39:25','admin',400,1002);
/*!40000 ALTER TABLE `favor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hot_keyword`
--

DROP TABLE IF EXISTS `hot_keyword`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hot_keyword`
--

LOCK TABLES `hot_keyword` WRITE;
/*!40000 ALTER TABLE `hot_keyword` DISABLE KEYS */;
INSERT INTO `hot_keyword` VALUES (1,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'Python'),(2,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'哈利·波特'),(3,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'村上春树'),(4,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'东野圭吾'),(5,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'白夜行'),(6,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'韩寒'),(7,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'金庸'),(8,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'王小波');
/*!40000 ALTER TABLE `hot_keyword` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'admin1','admin1'),(2,'2018-10-01 11:32:22','2018-10-01 11:32:22',NULL,'admin2','admin2'),(11,'2018-10-02 07:40:26','2018-10-02 07:40:26',NULL,'cyj','$2a$10$HErYgwOmrv09z9Twy/z8wOF/2I5VrH2SxCEGq5yKZlqJOQ5D/LFKS');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-10-09 21:06:34
