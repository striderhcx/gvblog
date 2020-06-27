-- MySQL dump 10.13  Distrib 5.7.30, for Linux (x86_64)
--
-- Host: localhost    Database: blog
-- ------------------------------------------------------
-- Server version	5.7.30-0ubuntu0.16.04.1

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
-- Current Database: `blog`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;

USE `blog`;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `pstatus` tinyint(4) NOT NULL DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Python','2020-06-20 18:20:50','2020-06-20 18:20:50',0,NULL),(2,'Golang','2020-06-20 18:21:00','2020-06-20 18:21:00',0,NULL),(3,'K8s','2020-06-20 18:21:08','2020-06-20 18:21:08',0,NULL),(4,'Docker','2020-06-20 18:21:16','2020-06-20 18:21:16',0,NULL),(5,'Anything','2020-06-20 18:22:29','2020-06-20 18:22:29',0,NULL),(6,'Test','2020-06-23 00:38:53','2020-06-23 00:38:53',0,NULL);
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comments` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `post_id` int(11) NOT NULL,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `pstatus` tinyint(4) NOT NULL DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_tags`
--

DROP TABLE IF EXISTS `post_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `post_tags` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `post_id` int(11) unsigned NOT NULL,
  `tag_id` int(11) unsigned NOT NULL,
  `pstatus` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_tags`
--

LOCK TABLES `post_tags` WRITE;
/*!40000 ALTER TABLE `post_tags` DISABLE KEYS */;
INSERT INTO `post_tags` VALUES (1,508,9,0,'2020-06-23 01:35:00','2020-06-23 01:35:00',NULL);
/*!40000 ALTER TABLE `post_tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `posts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `pstatus` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=509 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES (504,3,'k8s一种解决启动依赖的办法','最近在自己耍k8s，尝试搞一些实际的部署测试，发现一些问题，记录一下。web服务实际访问的时候再k8s里一般直接可以通过定义的service 名称来直接访问，比如这里的web服务连接mysql会直接通过名称\"mysql\"，端口3306进行连接访问，k8s自动帮我们进行解析到了相应的service中。本次坑中，web服务先于这个\"mysql\" service被k8s解析完成之前启动，导致了我掉到坑中。\n\n现象是这样的：\nweb pod 看到running状态， kubectl logs 看也没报错，一阵好找啊。最终是通过docker-compose工具来写个yaml文件对比了下，才发现可能是依赖启动顺序的问题。\ndocker-compose一般会有这样的配置：\n```yaml\nlinks:\n      - mysql\ndepends_on:\n      - mysql\n```\n\n最终k8s资源文件长这样:\n```yaml\napiVersion: extensions/v1beta1\n    kind: Ingress\n    metadata:\n    name: ingress-nginx\n    namespace: env-cx\n    annotations:\n        kubernetes.io/ingress.class: nginx\n    spec:\n    rules:\n        - host: ihxn.envcx.cn\n        http:\n            paths:\n            - path: /\n                backend:\n                serviceName: web-svc\n                servicePort: 80\n\n    ---\n    apiVersion: v1\n    kind: Service\n    metadata:\n    name: mysql\n    namespace: \"env-cx\"\n    labels:\n        run: mysql\n    spec:\n    type: NodePort\n    selector:\n        run: mysql\n    ports:\n        - port: 3306\n        name: mysql\n\n    ---\n\n    apiVersion: apps/v1\n    kind: Deployment\n    metadata:\n    name: mysql\n    namespace: \"env-cx\"\n    spec:\n    replicas: 1\n    selector:\n        matchLabels:\n        run: mysql\n    template:\n        metadata:\n        namespace: \"env-cx\"\n        labels:\n            run: mysql\n        spec:\n        containers:\n            - name: mysql\n            image: web_tpl-mysql\n            imagePullPolicy: IfNotPresent\n            resources:\n                limits:\n                memory: \"1Gi\"\n                requests:\n                memory: \"500Mi\"\n            env:\n                - name: MYSQL_ALLOW_EMPTY_PASSWORD\n                value: \"no\"\n                - name: MYSQL_ROOT_PASSWORD\n                value: \"666\"\n\n    ---\n    apiVersion: apps/v1\n    kind: Deployment \n    metadata: \n    name: web-deployment\n    namespace: env-cx\n    spec:\n    replicas: 2\n    selector:\n        matchLabels:\n        run: web-app\n    template:\n        metadata:\n        namespace: env-cx\n        labels:\n            run: web-app\n        spec:\n        containers:\n        - name: web-app\n            imagePullPolicy: IfNotPresent\n            image: web_tpl\n        initContainers:\n            - name: init-mysql\n            image: busybox\n            imagePullPolicy: IfNotPresent\n            command: [\'sh\', \'-c\', \'until nslookup mysql; do echo waiting for mysql; sleep 1; done;\']\n\n    ---\n    apiVersion: v1\n    kind: Service\n    metadata:\n    namespace: env-cx\n    name: web-svc\n    spec:\n    type: NodePort\n    selector:\n        run: web-app\n    ports:\n        - protocol: TCP\n        port: 80\n```',0,'2020-06-20 19:46:02','2020-06-20 19:46:02',NULL),(505,6,'这是一篇专门用来测试的blog','###  测试代码高亮\n```python\ndef main(a, b)\n    return a + b\n```\n\n```javascript\nfunction main(a, b) {\n    return a + b\n}\n```',0,'2020-06-23 00:42:21','2020-06-23 00:42:21',NULL),(508,6,'这是第二篇测试博客','我们都知道事务的几种性质，数据库为了维护这些性质，尤其是一致性和隔离性，一般使用加锁这种方式。同时数据库又是个高并发的应用，同一时间会有大量的并发访问，如果加锁过度，会极大的降低并发处理能力。所以对于加锁的处理，可以说就是数据库对于事务处理的精髓所在。这里通过分析MySQL中InnoDB引擎的加锁机制，来抛砖引玉，让读者更好的理解，在事务处理中数据库到底做了什么我们都知道事务的几种性质，数据库为了维护这些性质，尤其是一致性和隔离性，一般使用加锁这种方式。同时数据库又是个高并发的应用，同一时间会有大量的并发访问，如果加锁过度，会极大的降低并发处理能力。所以对于加锁的处理，可以说就是数据库对于事务处理的精髓所在。这里通过分析MySQL中InnoDB引擎的加锁机制，来抛砖引玉，让读者更好的理解，在事务处理中数据库到底做了什么',0,'2020-06-23 01:35:00','2020-06-23 01:35:00',NULL);
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tags` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pstatus` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (1,'Markdown',0,'2020-06-20 18:23:32','2020-06-20 18:23:32',NULL),(2,'Ubuntu',0,'2020-06-20 18:23:43','2020-06-20 18:23:43',NULL),(3,'Vue',0,'2020-06-20 18:24:11','2020-06-20 18:24:11',NULL),(4,'Django',0,'2020-06-20 18:24:31','2020-06-20 18:24:31',NULL),(5,'Flask',0,'2020-06-20 18:24:38','2020-06-20 18:24:38',NULL),(6,'Celery',0,'2020-06-20 18:24:45','2020-06-20 18:24:45',NULL),(7,'Mq',0,'2020-06-20 18:24:53','2020-06-20 18:24:53',NULL),(8,'leetcode',0,'2020-06-20 18:25:05','2020-06-20 18:25:05',NULL),(9,'Test',0,'2020-06-23 00:39:19','2020-06-23 00:39:19',NULL);
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pstatus` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'test','15926517215e21e586321582b26db15b39f37886a0','1225427292@qq.com',0,'2020-06-20 19:15:21','2020-06-20 19:15:21',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-06-24 14:33:21
