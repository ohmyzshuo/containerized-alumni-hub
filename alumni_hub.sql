/*
 Navicat Premium Data Transfer

 Source Server         : localhost_5432
 Source Server Type    : PostgreSQL
 Source Server Version : 170002 (170002)
 Source Host           : localhost:5432
 Source Catalog        : alumni_hub
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 170002 (170002)
 File Encoding         : 65001

 Date: 24/01/2025 11:46:19
*/


-- ----------------------------
-- Type structure for gender
-- ----------------------------
DROP TYPE IF EXISTS "public"."gender";
CREATE TYPE "public"."gender" AS ENUM (
  '',
  'Other',
  'Female',
  'Male'
);
ALTER TYPE "public"."gender" OWNER TO "shawn";

-- ----------------------------
-- Type structure for publication_type
-- ----------------------------
DROP TYPE IF EXISTS "public"."publication_type";
CREATE TYPE "public"."publication_type" AS ENUM (
  'ISI WOS',
  'Scopus',
  'Book',
  'Chapter of Book',
  'Journal A',
  'Journal B',
  'Others'
);
ALTER TYPE "public"."publication_type" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for attachments_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."attachments_id_seq";
CREATE SEQUENCE "public"."attachments_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."attachments_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for contents_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."contents_id_seq";
CREATE SEQUENCE "public"."contents_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."contents_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for email_attachments_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."email_attachments_id_seq";
CREATE SEQUENCE "public"."email_attachments_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."email_attachments_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for emails_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."emails_id_seq";
CREATE SEQUENCE "public"."emails_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."emails_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for faculty_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."faculty_id_seq";
CREATE SEQUENCE "public"."faculty_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."faculty_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for honors_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."honors_id_seq";
CREATE SEQUENCE "public"."honors_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."honors_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for publication_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."publication_id_seq";
CREATE SEQUENCE "public"."publication_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."publication_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for staff_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."staff_id_seq";
CREATE SEQUENCE "public"."staff_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."staff_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for studies_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."studies_id_seq";
CREATE SEQUENCE "public"."studies_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."studies_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for tags_sequence
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."tags_sequence";
CREATE SEQUENCE "public"."tags_sequence" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."tags_sequence" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."users_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Sequence structure for work_experiences_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."work_experiences_id_seq";
CREATE SEQUENCE "public"."work_experiences_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."work_experiences_id_seq" OWNER TO "shawn";

-- ----------------------------
-- Table structure for alumni_tags
-- ----------------------------
DROP TABLE IF EXISTS "public"."alumni_tags";
CREATE TABLE "public"."alumni_tags" (
  "alumni_id" int4 NOT NULL,
  "tag_id" int4 NOT NULL,
  "weight" float8 DEFAULT 1.0
)
;
ALTER TABLE "public"."alumni_tags" OWNER TO "shawn";

-- ----------------------------
-- Records of alumni_tags
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for alumnis
-- ----------------------------
DROP TABLE IF EXISTS "public"."alumnis";
CREATE TABLE "public"."alumnis" (
  "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "nationality" varchar(50) COLLATE "pg_catalog"."default",
  "ethnicity" varchar(50) COLLATE "pg_catalog"."default",
  "dob" date,
  "gender" varchar(10) COLLATE "pg_catalog"."default",
  "marital" varchar(20) COLLATE "pg_catalog"."default",
  "address" varchar(255) COLLATE "pg_catalog"."default",
  "email" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "matric_no" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "phone" varchar(20) COLLATE "pg_catalog"."default",
  "is_hidden" bool DEFAULT false,
  "has_verified" bool DEFAULT false,
  "location" varchar(255) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  "token" varchar(500) COLLATE "pg_catalog"."default",
  "linked_in" varchar(100) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."alumnis" OWNER TO "shawn";

-- ----------------------------
-- Records of alumnis
-- ----------------------------
BEGIN;
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (31, 'Hoo Wei Han', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', 'U24102749', '$2a$10$XZ7JuUkMFEiyPcTqhnTXu.2HuRXK8twj.FxI.xDclUber65heTH9.', '+60 102141419', 'f', 'f', '', '2024-12-05 19:58:11.54938', '2024-12-05 19:58:11.54938', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (32, 'Lo Zi Yang', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', 'U24334145', '$2a$10$ruLZzAq4W4YmCY3WCrhwE.zNNh40tayRsd8MS4XzDJw4B5/1R9aEC', '+60 102141420', 'f', 'f', '', '2024-12-05 19:58:11.629383', '2024-12-05 19:58:11.629383', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (6, 'Shuo Li', '', '', '0001-01-01', '', '', '', 'lishuo0527@gmail.com', '1', '$2a$10$siW02yjTrUhWeVY5nrTbXO2KEyNgKLShrmRBnRavOakk2HrWRSs8G', '', 't', 'f', '', '2024-11-26 15:45:52.808816', '2024-12-04 21:34:18.704951', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (2, 'Shawn Ko', 'N/A', 'N/A', '0001-01-01', 'Male', 'N/A', 'N/A', 'lishuo0527@gmail.com', 'T2001314', '$2a$10$B9y7JR0bD2UgVLljG0q1uuQ6MJr0E6kLDhMlmAUxoAEIX3xMo0mgO', 'N/A', 'f', 'f', 'N/A', '2024-11-11 20:05:29.287616', '2025-01-08 14:34:08.762026', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJyb2xlIjoiYWx1bW5pIiwiZXhwIjoxNzM2Njg3NDQ5LCJpYXQiOjE3MzQwOTU0NDksImlzcyI6ImFsdW1uaV9odWIifQ.XfQHo2r2Iopt_oq76mh6hJ4WeBnNBJLBlMPAetbvtg4', 'N/A');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (27, 'Tan Zi Rong', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', 'U24132552', '$2a$10$voyd35hqlQLhDtv1l.VsguxVdJGnBZqUCUfbeS/s8z500pXDyc1j6', '+60 102141415', 'f', 'f', '', '2024-12-05 19:58:11.262526', '2024-12-05 19:58:11.262526', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (5, 'Shuo Li ', 'China', 'Chinese', '2001-03-14', 'Male', 'Single', 'Kota Damansara, Pataling Jaya, Selangor, Malaysia', '3146681@gmail.com', 'S2001314', '$2a$10$SCEBvQ9F7kmmxYsGGHqwmeQyeRyIo86G7vJnLQsxK9BKZvNhsC6ci', '+86 13131257871', 'f', 'f', 'Kuala Lumpur', '2024-11-17 15:58:49.755492', '2025-01-10 16:36:41.211195', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJyb2xlIjoiYWx1bW5pIiwiZXhwIjoxNzM5MDkwMjAxLCJpYXQiOjE3MzY0OTgyMDEsImlzcyI6ImFsdW1uaV9odWIifQ.6XPKdJdIkITTWTXEaBd33k7z1SxChuaOPEviwzsuj3k', 'oh-my-zshuo1');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (11, 'Hoo Wei Han', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', '81724102749', '$2a$10$KGk7RBEv1RMrAhtfGyu0YOesUp8TG59UmVZ0c/SnL0E9414lLgpQ6', '109130784', 't', 'f', '', '2024-11-29 13:12:50.513212', '2024-12-05 19:57:09.68104', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (9, 'shawn', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', '81724110761', '$2a$10$5HRzUI1H9tkaxAPodNjK8.x3ZkQsf2SxNKrJBSusbfi0sLjBGKcJ.', '109130784', 'f', 'f', '', '2024-11-29 13:12:50.374554', '2024-11-29 16:21:59.361524', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (12, 'Lo Zi Yang', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', '81724334145', '$2a$10$NB64Lh/vsCS.LqaGhZReReCZc5use2rqQsmWj2PlwKAARW6QPsjZm', '109130784', 'f', 'f', '', '2024-11-29 13:12:50.581665', '2024-11-29 13:12:50.581665', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (20, 'Leong  Fang Loong', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', '81724139214', '$2a$10$ILI31FbCFHdZE9u/ouM9FeE9vqRYO1BbFPcqRRVbb4LzWC9tzJ57.', '109130784', 'f', 'f', '', '2024-12-04 23:29:29.548972', '2024-12-04 23:29:29.548972', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (8, 'Leong ', 'Malaysia', 'N/A', '2019-12-31', 'Male', 'N/A', 'N/A', 'lishuo0527@gmail.com', 'S2223232', '$2a$10$lmceCANG67dsNgp/YyXEfeEAnNkaEIeLhJjAWOEeZE0A2gmqpToY.', '109130784', 'f', 'f', 'N/A', '2024-11-29 13:12:50.310368', '2024-12-13 22:54:58.748759', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4LCJyb2xlIjoiYWx1bW5pIiwiZXhwIjoxNzM2NjkzNjk4LCJpYXQiOjE3MzQxMDE2OTgsImlzcyI6ImFsdW1uaV9odWIifQ.3MZdwyWiRVxj1Ml0uyBeMU0rSGiMGA5DKwZqCEHNBxs', 'N/A');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (4, 'shawn', '', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', 'T2003464', '$2a$10$w7geAe8AlBF4S9kXWRClnOMq1Eo0.gnoEWfTCu8LVWAvvG3ENdbXS', '', 'f', 'f', '', '2024-11-11 20:08:01.44755', '2024-11-29 16:22:26.86881', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (7, 'Tan Ah Rong', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', '81724132552', '$2a$10$.Vg/7VrxOIeRIID9.RCSjeltZQg9cuHi9gbCTwocNp45yAEy9VkF6', '109130784', 'f', 'f', '', '2024-11-29 13:12:50.237808', '2024-11-29 13:12:50.237808', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (30, 'Koay Teng Fai', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', 'U24102760', '$2a$10$0mgg84Sg4jzUqttG3Z4peOfrIDXdnftuRSnltN0adq.fYF.2K/LJ6', '+60 102141418', 'f', 'f', '', '2024-12-05 19:58:11.480211', '2024-12-13 21:26:23.547647', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozMCwicm9sZSI6ImFsdW1uaSIsImV4cCI6MTczNjY4ODM4MywiaWF0IjoxNzM0MDk2MzgzLCJpc3MiOiJhbHVtbmlfaHViIn0.mTiMqeeoZBAdplX8-0VnRKCasuduIqzk5JJxyBK4zfI', '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (28, 'Leong Fang Loong', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', 'U24139214', '$2a$10$jPjRwhX1lc3i4WhHsBfBDOYI1tOtE9Qc.7xqiVVwJ0z6kUjLJoL7C', '+60 102525274', 'f', 'f', '', '2024-12-05 19:58:11.338291', '2024-12-05 19:58:11.338291', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (29, 'Tan Choong Wei', 'Malaysia', '', '0001-01-01', 'Female', '', '', 'lishuo0527@gmail.com', 'U24110761', '$2a$10$lAY/qiiJwhdAhwnKzwlztexmGLbUrTpjamINmNmlWvcfQfLHyB61O', '+60 14534453', 'f', 'f', '', '2024-12-05 19:58:11.406838', '2024-12-05 19:58:11.406838', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (37, 'Koay Teng Fai', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'test123631@gmail.com', 'U242760', '$2a$10$ViGRFlXwWslAohjLAJSO.ef/vMyrONs3gQ8a83nR/J202x80OepVq', '+60 102141418', 'f', 'f', '', '2025-01-08 20:02:53.450969', '2025-01-08 21:29:49.719833', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (19, 'T1213131', 'china', 'chinese', '2020-01-10', 'Female', 'Divorced', 'kk7', 'lishuo@gmail.com', 's2037452', '$2a$10$zk36.dw4QEbgALlrcmccAulhVfUcHgMk7ARRKaWNgFseBi2yTo4K2', '12345678', 'f', 'f', 'pj', '2024-12-04 21:46:41.056755', '2025-01-10 16:09:05.46847', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxOSwicm9sZSI6ImFsdW1uaSIsImV4cCI6MTczOTA4ODU0NSwiaWF0IjoxNzM2NDk2NTQ1LCJpc3MiOiJhbHVtbmlfaHViIn0.V4awlqlpQeAH2kR2ZBj54MbEUNFCga1xRYj_15Hgv1c', '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (34, 'Tan Zi Rong', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'test12131@gmail.com', 'U241332552', '$2a$10$wf8oN8XtFoYS.zKYhPFoyODC/swYR3.6Xni4V9fOXyFhdsD1770x2', '+60 102141415', 'f', 'f', '', '2025-01-08 20:02:53.108246', '2025-01-08 21:29:49.70892', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (35, 'Leong Fang Loong', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'test1331@gmail.com', 'U241239214', '$2a$10$/fD1NEJwfpM24eL/evKXKO43fawvPigqRoXi9jpHWOcYA5diqWv/C', '+60 102525274', 'f', 'f', '', '2025-01-08 20:02:53.223513', '2025-01-08 21:29:49.712406', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (36, 'Tan Choong Wei', 'Malaysia', '', '0001-01-01', 'Female', '', '', 'test121521@gmail.com', 'U24161', '$2a$10$z2Zgm65lIzx6k5eHeN5qcu2dy4L9qoesTNWkb72V9ZsNKdxBO.npm', '+60 14534453', 'f', 'f', '', '2025-01-08 20:02:53.332627', '2025-01-08 21:29:49.715436', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (38, 'Hoo Wei Han', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'test12331@gmail.com', 'U241024749', '$2a$10$clRTCukdyFv7sEM30iV8Fur9wccdYv/4vOnFtLwlp1RxBWrKUjVvC', '+60 102141419', 'f', 'f', '', '2025-01-08 20:02:53.56162', '2025-01-08 21:29:49.721457', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (39, 'Lo Zi Yang', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'test1555@gmail.com', 'U243344145', '$2a$10$roA9ZSCWHlrPXfcPegwv2u.pZ2SInQwiEQpiR2eMlI.mFApQFauRC', '+60 102141420', 'f', 'f', '', '2025-01-08 20:02:53.672557', '2025-01-08 21:29:49.723088', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (40, 'LEE KHUAY KHIANG ', '', '', '0001-01-01', '', '', '', '', '17020741', '$2a$10$8tZ2lCYKUQ87BfxjDEXYqungclQ8kZ4ebt.3DXoDstQbC0AZaGYf2', '', 'f', 'f', '', '2025-01-09 14:36:01.939955', '2025-01-09 14:37:06.304269', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (41, 'NUR AYEESHA QISTEENA BINTI MUZIR ', '', '', '0001-01-01', '', '', '', '', '17013368', '$2a$10$ZsZnwFBohAxjqPq7rZIkZOqH5.AJs6cdzKwdo4dLGrOTGfAoIhc9m', '', 'f', 'f', '', '2025-01-09 14:36:02.034297', '2025-01-09 14:37:06.310092', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (42, 'MUTHU KUMARAN GUNASEGARAN', '', '', '0001-01-01', '', '', '', '', '17004842', '$2a$10$2Tvne17SfwfFOPOytA.vKOSKeC.A4//nph0ZvcXIZAf9KiypBqrie', '', 'f', 'f', '', '2025-01-09 14:36:02.105379', '2025-01-09 14:37:06.317853', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (43, 'NARENDREN A/L RENGASAMY', '', '', '0001-01-01', '', '', '', '', '17031734', '$2a$10$rv1tBoJk/Bq5Y4ajOVNKMe/UsCD83NApUSEVAfe1PqXGU5tgD3FMm', '', 'f', 'f', '', '2025-01-09 14:36:02.175098', '2025-01-09 14:37:06.384035', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (44, 'SARRAF RAJESH KUMAR', '', '', '0001-01-01', '', '', '', '', '17035326', '$2a$10$mahogy1Hvdes0cg8XNp4E.3THo6ScC0t/OaDfF/lWA/rGdNWcduJe', '', 'f', 'f', '', '2025-01-09 14:36:02.242897', '2025-01-09 14:37:06.388614', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (45, 'MUHAMAD HILMI BIN ABDUL RAHMAN', '', '', '0001-01-01', '', '', '', '', '17049995', '$2a$10$3mMIs/J2arzPYPrUmPzmMOVaS/m8Xs4s9NTzq1p5kv7QxEI5EjX8O', '', 'f', 'f', '', '2025-01-09 14:36:02.310842', '2025-01-09 14:37:06.391092', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (46, 'ONG CHIU LING', '', '', '0001-01-01', '', '', '', '', '17058647', '$2a$10$dvXsYSAe9WNIAb3W3ErQae6CZRSWA1L9Pq2Bwhh11V/D2jCZ7ab1i', '', 'f', 'f', '', '2025-01-09 14:36:02.377683', '2025-01-09 14:37:06.392751', NULL, NULL, '');
INSERT INTO "public"."alumnis" ("id", "name", "nationality", "ethnicity", "dob", "gender", "marital", "address", "email", "matric_no", "password", "phone", "is_hidden", "has_verified", "location", "created_at", "updated_at", "deleted_at", "token", "linked_in") VALUES (10, 'Koay Teng Fai', 'Malaysia', '', '0001-01-01', 'Male', '', '', 'lishuo0527@gmail.com', '81724102760', '$2a$10$H5DHrRLH31V4pl16xRdOrOGK2XEFtreP0z/3f3vt2YlpcHSCQu7US', '109130784', 'f', 'f', '', '2024-11-29 13:12:50.440209', '2023-12-03 17:14:59.271281', NULL, NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for attachments
-- ----------------------------
DROP TABLE IF EXISTS "public"."attachments";
CREATE TABLE "public"."attachments" (
  "id" int4 NOT NULL DEFAULT nextval('attachments_id_seq'::regclass),
  "content_id" int4 NOT NULL,
  "attachment_path" text COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "original_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."attachments" OWNER TO "shawn";

-- ----------------------------
-- Records of attachments
-- ----------------------------
BEGIN;
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (1, 1, '/path/to/test.txt', '2024-11-29 10:11:43.783455', '2024-11-29 10:11:43.783455', '2024-11-29 10:11:43.783455', 'test.txt');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (14, 5, 'uploads/content/qs60.png', '2024-12-04 01:08:42', '2024-12-04 01:08:46', '2024-12-04 01:08:49', 'qs60.png');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (5, 181, 'uploads/content/event1.jpeg', '2024-12-04 01:12:39', '2024-12-04 01:12:37', '2024-12-04 01:12:35', 'event1.jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (15, 38, 'uploads/content/Snipaste_2024-12-05_16-38-28.png', '2024-12-05 16:38:38.621645', '2024-12-05 16:38:38.621645', '2024-12-05 16:38:38.610189', 'Snipaste_2024-12-05_16-38-28.png');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (17, 39, 'uploads/content/qs60.png', '2024-12-05 16:48:04.944458', '2024-12-05 16:48:04.944458', '2024-12-05 16:48:04.941917', 'qs60.png');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (18, 311, 'uploads/content/images (1).jpeg', '2024-12-05 18:09:01.012112', '2024-12-05 18:09:01.012112', '2024-12-05 18:09:01.009838', 'images (1).jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (19, 213, 'uploads/content/event1.jpeg', '2024-12-05 18:57:15.753298', '2024-12-05 18:57:15.753298', '2024-12-05 18:57:15.750899', 'event1.jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (20, 213, 'uploads/content/images (1).jpeg', '2024-12-05 18:57:15.753298', '2024-12-05 18:57:15.753298', '2024-12-05 18:57:15.750899', 'images (1).jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (22, 301, 'uploads/content/images.jpeg', '2024-12-05 20:00:56.403482', '2024-12-05 20:00:56.403482', '2024-12-05 20:00:56.402729', 'images.jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (31, 52, 'uploads/content/event5.jpeg', '2025-01-10 15:25:00.555809', '2025-01-10 15:25:00.555809', '2025-01-10 15:25:00.553536', 'event5.jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (32, 51, 'uploads/content/event7.jpeg', '2025-01-10 15:25:09.222407', '2025-01-10 15:25:09.222407', '2025-01-10 15:25:09.220736', 'event7.jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (33, 50, 'uploads/content/event3.jpeg', '2025-01-10 15:25:23.487664', '2025-01-10 15:25:23.487664', '2025-01-10 15:25:23.485895', 'event3.jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (34, 49, 'uploads/content/event2.jpeg', '2025-01-10 15:25:39.264634', '2025-01-10 15:25:39.264634', '2025-01-10 15:25:39.261564', 'event2.jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (35, 49, 'uploads/content/event3.jpeg', '2025-01-10 15:25:39.264634', '2025-01-10 15:25:39.264634', '2025-01-10 15:25:39.261564', 'event3.jpeg');
INSERT INTO "public"."attachments" ("id", "content_id", "attachment_path", "created_at", "updated_at", "deleted_at", "original_name") VALUES (36, 48, 'uploads/content/event1.jpg', '2025-01-10 15:25:47.880725', '2025-01-10 15:25:47.880725', '2025-01-10 15:25:47.878722', 'event1.jpg');
COMMIT;

-- ----------------------------
-- Table structure for content_tags
-- ----------------------------
DROP TABLE IF EXISTS "public"."content_tags";
CREATE TABLE "public"."content_tags" (
  "content_id" int4 NOT NULL,
  "tag_id" int4 NOT NULL
)
;
ALTER TABLE "public"."content_tags" OWNER TO "shawn";

-- ----------------------------
-- Records of content_tags
-- ----------------------------
BEGIN;
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (38, 19);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (38, 20);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (38, 21);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (39, 22);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (39, 23);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (48, 24);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (48, 25);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (48, 26);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (48, 27);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (48, 28);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (48, 29);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (48, 30);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (48, 31);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (49, 32);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (49, 33);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (49, 19);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (49, 34);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (49, 35);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (49, 36);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (50, 37);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (50, 38);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (50, 11);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (50, 39);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (50, 40);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (51, 41);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (51, 42);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (51, 43);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (51, 44);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (51, 45);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (51, 46);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (52, 47);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (52, 48);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (52, 49);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (52, 50);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (52, 51);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (52, 52);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (52, 53);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (52, 54);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (55, 55);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (55, 56);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (55, 57);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (55, 58);
INSERT INTO "public"."content_tags" ("content_id", "tag_id") VALUES (55, 59);
COMMIT;

-- ----------------------------
-- Table structure for contents
-- ----------------------------
DROP TABLE IF EXISTS "public"."contents";
CREATE TABLE "public"."contents" (
  "id" int8 NOT NULL DEFAULT nextval('contents_id_seq'::regclass),
  "title" varchar(255) COLLATE "pg_catalog"."default",
  "description" text COLLATE "pg_catalog"."default",
  "contact" text COLLATE "pg_catalog"."default",
  "created_by" int4,
  "published_at" timestamp(6),
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "venue" varchar(255) COLLATE "pg_catalog"."default",
  "faculty_id" int4,
  "start_time" timestamp(6),
  "end_time" timestamp(6),
  "is_hidden" bool,
  "num_of_likes" int4,
  "deleted_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "is_pinned" bool,
  "updated_by" int4,
  "participant_quota" int4,
  "content_type" int2
)
;
ALTER TABLE "public"."contents" OWNER TO "shawn";

-- ----------------------------
-- Records of contents
-- ----------------------------
BEGIN;
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (11, 'Quantum C]', 'Gain insight into the principles of quantum mechanics and how they are leveraged to build quantum computers. Learn about the potential applications and implications of this groundbreaking technology.', '', 7, '0001-01-01 06:55:25', '2024-06-11 19:14:04.906253', '2024-12-01 23:19:09.918911', 'FSKTM', NULL, '0001-01-01 00:00:00', '0001-01-01 00:00:00', NULL, NULL, '2024-11-11 17:44:39.65401', NULL, NULL, NULL, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (8, 'Women in Tech Panel Discussion', 'Hear from successful women leaders in tech as they share their experiences, challenges, and insights on building diverse and inclusive tech communities.', NULL, 7, '0001-01-01 06:55:25', '2024-06-11 19:13:02.971797', '2024-06-11 19:13:02.971797', 'FSKTM', NULL, '0001-01-01 00:00:00', '0001-01-01 00:00:00', NULL, NULL, '2024-11-11 17:44:39.65401', NULL, NULL, NULL, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (10, 'Building Solutions for Social Good', 'Collaborate with peers to brainstorm, design, and prototype tech-based solutions that have a positive impact on society and the planet.', NULL, 7, '0001-01-01 06:55:25', '2024-06-11 19:13:43.061732', '2024-06-11 19:13:43.061732', 'FSKTM', NULL, '0001-01-01 00:00:00', '0001-01-01 00:00:00', NULL, NULL, '2024-11-11 17:44:39.65401', NULL, NULL, NULL, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (161, 'Digital Marketing Strategies', 'Learn effective digital marketing strategies to boost your online presence and engagement.', '', 12, '2024-05-05 09:00:00', '2024-06-16 09:00:00', '2024-12-01 23:24:00', 'Business Faculty', NULL, '2024-06-01 09:00:00', '2024-06-01 17:00:00', NULL, 35, NULL, NULL, 13, 300, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (7, 'Programming Challenge Competition', 'This lecture series will cover the latest developments in areas such as artificial intelligence, machine learning, data science, and their applications in society and industry.', NULL, 7, '0001-01-01 06:55:25', '2024-06-11 19:12:06.726699', '2024-06-11 19:12:06.726699', 'FSKTM', NULL, '0001-01-01 00:00:00', '0001-01-01 00:00:00', NULL, NULL, '2024-11-11 17:44:39.65401', NULL, NULL, NULL, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (141, 'Sustainable Energy Solutions', 'Discover innovative approaches to sustainable energy and the future of renewable resources.', '', 10, '2024-03-25 14:00:00', '2024-06-14 14:00:00', '2024-12-01 23:22:00', 'Science Faculty', NULL, '2024-04-01 09:00:00', '2024-04-01 17:00:00', NULL, 25, NULL, NULL, 11, 200, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (131, 'Blockchain Basics', 'Understand the fundamentals of blockchain technology and its applications beyond cryptocurrencies.', '', 9, '2024-02-20 11:30:00', '2024-06-13 11:30:00', '2024-12-01 23:21:00', 'Engineering Faculty', NULL, '2024-03-01 10:00:00', '2024-03-01 16:00:00', NULL, 20, NULL, NULL, 10, 150, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (121, 'AI in Healthcare', 'Explore the transformative impact of artificial intelligence on healthcare, from diagnostics to treatment.', '', 8, '2024-01-15 10:00:00', '2024-06-12 10:00:00', '2024-12-01 23:20:00', 'Medical Faculty', NULL, '2024-02-01 09:00:00', '2024-02-01 17:00:00', NULL, 15, NULL, NULL, 9, 100, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (151, 'Cybersecurity Trends', 'Stay ahead of the latest trends and threats in cybersecurity, and learn how to protect your digital assets.', '', 11, '2024-04-30 16:00:00', '2024-06-15 16:00:00', '2024-12-01 23:23:00', 'IT Faculty', NULL, '2024-05-01 10:00:00', '2024-05-01 15:00:00', NULL, 30, NULL, NULL, 12, 250, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (5, 'UNIVERSITI MALAYA RISES TO TOP 60 IN THE QS-WORLD UNIVERSITY RANKINGS 2025', 'Universiti Malaya (UM) has made a significant achievement by securing a position among the top 60 universities globally in the Quacquarelli Symonds World University Rankings (QS-WUR) for 2025. This accomplishment marks a 5-place improvement from the previous year and a notable 10-place climb since the 2023 rankings. ', NULL, 7, '0001-01-01 06:55:25', '2024-06-11 18:56:57.654106', '2024-06-11 18:56:57.654106', '', NULL, '0001-01-01 00:00:00', '0001-01-01 00:00:00', NULL, NULL, '2024-11-11 17:44:39.65401', NULL, NULL, NULL, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (181, 'Entrepreneurship 101', 'Get insights into starting your own business, from ideation to execution.', '', 14, '2024-07-15 11:00:00', '2024-06-18 11:00:00', '2024-12-01 23:26:00', 'Business Faculty', NULL, '2024-08-01 09:00:00', '2024-08-01 17:00:00', NULL, 45, NULL, NULL, 15, 400, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (201, 'Robotics and Automation', 'Explore the latest advancements in robotics and automation and their applications in various industries.', '', 16, '2024-09-25 13:00:00', '2024-06-20 13:00:00', '2024-12-01 23:28:00', 'Engineering Faculty', NULL, '2024-10-01 09:00:00', '2024-10-01 17:00:00', NULL, 55, NULL, NULL, 17, 500, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (211, 'Artificial Intelligence Ethics', 'Discuss the ethical considerations and challenges associated with the development and deployment of AI.', '', 17, '2024-10-30 14:00:00', '2024-06-21 14:00:00', '2024-12-01 23:29:00', 'Philosophy Faculty', NULL, '2024-11-01 10:00:00', '2024-11-01 12:00:00', NULL, 60, NULL, NULL, 18, 550, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (261, 'Introduction to Bioinformatics', 'Discover the basics of bioinformatics and its importance in modern biological research.', '', 22, '2025-01-05 09:00:00', '2024-06-26 09:00:00', '2024-12-01 23:34:00', 'Biology Faculty', NULL, '2025-03-01 09:00:00', '2025-03-01 17:00:00', NULL, 85, NULL, NULL, 23, 800, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (281, 'Environmental Conservation', 'Learn about environmental conservation efforts and how you can contribute to protecting our planet.', '', 24, '2025-01-15 11:00:00', '2024-06-28 11:00:00', '2024-12-01 23:36:00', 'Environmental Faculty', NULL, '2025-05-01 09:00:00', '2025-05-01 17:00:00', NULL, 95, NULL, NULL, 25, 900, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (271, 'Creative Writing Workshop', 'Participate in a creative writing workshop to enhance your writing skills and explore different genres.', '', 23, '2025-01-10 10:00:00', '2024-06-27 10:00:00', '2024-12-01 23:35:00', 'Literature Faculty', NULL, '2025-04-01 09:00:00', '2025-04-01 12:00:00', 't', 90, NULL, NULL, 24, 850, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (171, 'Climate Change Awareness', 'Raise awareness about climate change and discuss actionable steps to mitigate its effects.', '', 13, '2024-06-10 10:00:00', '2024-06-17 10:00:00', '2024-12-01 23:25:00', 'Environmental Faculty', NULL, '2024-07-01 09:00:00', '2024-07-01 16:00:00', NULL, 40, NULL, 't', 14, 350, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (251, 'Public Speaking Skills', 'Enhance your public speaking skills and learn techniques to communicate effectively.', '', 21, '2024-12-20 18:00:00', '2024-06-25 18:00:00', '2024-12-05 18:48:43.866089', 'Communication Faculty', NULL, '2025-02-01 09:00:00', '2025-02-01 12:00:00', NULL, 0, NULL, NULL, 22, 0, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (213, 'Mental Health Awareness', 'Promote mental health awareness and provide resources and strategies for maintaining mental well-being.', '', 19, '2024-12-10 16:00:00', '2024-06-23 16:00:00', '2024-12-05 18:57:15.751041', 'Psychology Faculty', NULL, '2024-12-15 09:00:00', '2024-12-15 12:00:00', NULL, 70, NULL, NULL, 20, 650, NULL);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (52, 'Student Fitness Memberships', 'Sign up for a student fitness membership and get access to all gym facilities, classes, and personal training sessions at a special student rate!', '', 2, NULL, '2025-01-10 12:24:02.073212', '2025-01-10 15:25:00.553924', '', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', 'f', 0, '2025-01-10 12:24:02.07286', 'f', 2, 0, 3);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (38, 'Modern Art Appreciation', 'Explore the world of modern art and gain a deeper understanding of its various forms and expressions.', '', 25, NULL, '2024-06-29 12:00:00', '2024-12-05 16:45:37.975346', 'Arts Faculty', 0, '2025-06-01 09:00:00', '2025-06-01 12:00:00', 'f', 100, '2024-12-05 16:38:38.610189', 'f', 26, 95, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (51, 'Library Renovation', 'The university library will undergo renovations starting next month. Please check our website for updates on services and temporary location', '', 2, NULL, '2025-01-10 12:23:41.075679', '2025-01-10 15:25:09.220894', '', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', 'f', 0, '2025-01-10 12:23:41.07549', 'f', 2, 0, 2);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (50, 'Guest Lecture Series: Innovations in Technology', 'Attend a thought-provoking lecture featuring industry leaders discussing the latest trends and innovations in technology and their impact on society.', '', 2, NULL, '2025-01-10 12:23:10.930005', '2025-01-10 15:25:23.486058', 'Fac of Eng', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', 'f', 0, '2025-01-10 12:23:10.929798', 'f', 2, 0, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (49, 'International Culture Night', 'Experience a vibrant evening celebrating the diverse cultures on campus with food, performances, and interactive activities from around the world.', '', 2, NULL, '2025-01-10 12:22:40.556035', '2025-01-10 15:25:39.261848', '', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', 'f', 0, '2025-01-10 12:22:40.555857', 'f', 2, 130, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (48, 'Spring Career Fair', 'Join us for our annual Spring Career Fair, where students can meet potential employers, network, and explore internship and job opportunities.
', '', 2, NULL, '2025-01-10 12:22:10.048403', '2025-01-10 15:25:47.878941', 'DTC', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', 'f', 0, '2025-01-10 12:22:10.048007', 'f', 2, 230, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (301, 'Introduction to Philosophy', 'Delve into the fundamental questions of philosophy and explore different philosophical perspectives.', '', 26, '2025-01-25 13:00:00', '2024-06-30 13:00:00', '2024-12-05 20:01:45.298578', 'Philosophy Faculty', NULL, '2025-07-01 09:00:00', '2025-07-01 12:00:00', NULL, 105, NULL, NULL, 27, 1000, 3);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (311, 'Ethical Hacking', 'Learn the basics of ethical hacking and how to identify and mitigate security vulnerabilities.', '', 27, '2025-01-30 14:00:00', '2024-07-01 14:00:00', '2024-12-05 18:48:00.084778', 'FSKTM', NULL, '2025-08-01 09:00:00', '2025-08-01 17:00:00', NULL, 110, NULL, NULL, 28, 55, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (191, 'Data Science Workshop', 'Hands-on workshop to learn data science techniques and tools for data analysis and visualization.', '', 15, '2024-08-20 12:00:00', '2024-06-19 12:00:00', '2024-12-05 18:48:10.992716', 'FSKTM', NULL, '2024-09-01 09:00:00', '2024-09-01 17:00:00', NULL, 50, NULL, NULL, 16, 45, 1);
INSERT INTO "public"."contents" ("id", "title", "description", "contact", "created_by", "published_at", "created_at", "updated_at", "venue", "faculty_id", "start_time", "end_time", "is_hidden", "num_of_likes", "deleted_at", "is_pinned", "updated_by", "participant_quota", "content_type") VALUES (39, 'UM Ranked 60th in QS Ranking this year', 'Universiti Malaya (UM) has made a significant achievement by securing a position among the top 60 universities globally in the Quacquarelli Symonds World University Rankings (QS-WUR) for 2025.', '', 24, NULL, '2024-06-28 11:00:00', '2024-12-05 16:53:44.969122', '-', 0, '2025-05-01 09:00:00', '2025-05-01 17:00:00', 'f', 95, '2024-12-05 16:39:59.14334', 'f', 25, 1, 2);
COMMIT;

-- ----------------------------
-- Table structure for email_attachments
-- ----------------------------
DROP TABLE IF EXISTS "public"."email_attachments";
CREATE TABLE "public"."email_attachments" (
  "id" int4 NOT NULL DEFAULT nextval('email_attachments_id_seq'::regclass),
  "email_id" int4 NOT NULL,
  "original_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "attachment_path" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."email_attachments" OWNER TO "shawn";

-- ----------------------------
-- Records of email_attachments
-- ----------------------------
BEGIN;
INSERT INTO "public"."email_attachments" ("id", "email_id", "original_name", "attachment_path", "created_at", "updated_at") VALUES (1, 2, 'WIF3008 Midsem Attendance Proof.JPG', 'uploads/WIF3008 Midsem Attendance Proof.JPG', '2024-11-29 20:10:45.11407+08', '2024-11-29 20:10:45.11407+08');
INSERT INTO "public"."email_attachments" ("id", "email_id", "original_name", "attachment_path", "created_at", "updated_at") VALUES (2, 2, '阳光下的野草 2021年9月25日.jpg', 'uploads/阳光下的野草 2021年9月25日.jpg', '2024-11-29 20:10:45.11407+08', '2024-11-29 20:10:45.11407+08');
INSERT INTO "public"."email_attachments" ("id", "email_id", "original_name", "attachment_path", "created_at", "updated_at") VALUES (3, 3, 'WIF3008 Midsem Attendance Proof.JPG', 'uploads/WIF3008 Midsem Attendance Proof.JPG', '2024-11-29 20:19:37.952255+08', '2024-11-29 20:19:37.952255+08');
INSERT INTO "public"."email_attachments" ("id", "email_id", "original_name", "attachment_path", "created_at", "updated_at") VALUES (4, 3, '阳光下的野草 2021年9月25日.jpg', 'uploads/阳光下的野草 2021年9月25日.jpg', '2024-11-29 20:19:37.952255+08', '2024-11-29 20:19:37.952255+08');
COMMIT;

-- ----------------------------
-- Table structure for emails
-- ----------------------------
DROP TABLE IF EXISTS "public"."emails";
CREATE TABLE "public"."emails" (
  "id" int4 NOT NULL DEFAULT nextval('emails_id_seq'::regclass),
  "to" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "from" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "subject" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "body" text COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."emails" OWNER TO "shawn";

-- ----------------------------
-- Records of emails
-- ----------------------------
BEGIN;
INSERT INTO "public"."emails" ("id", "to", "from", "subject", "body", "created_at", "updated_at") VALUES (1, 'lishuo0527@gmail.com', 's2001314@siswa.um.edu.my', 's2001314', 's2001314', '2024-11-29 20:08:21.991199+08', '2024-11-29 20:08:21.991199+08');
INSERT INTO "public"."emails" ("id", "to", "from", "subject", "body", "created_at", "updated_at") VALUES (2, 'lishuo0527@gmail.com', 's2001314@siswa.um.edu.my', 's2001314', 's2001314', '2024-11-29 20:10:45.111388+08', '2024-11-29 20:10:45.111388+08');
INSERT INTO "public"."emails" ("id", "to", "from", "subject", "body", "created_at", "updated_at") VALUES (3, 'lishuo0527@gmail.com', 's2001314@siswa.um.edu.my', 's2001314', 's2001314', '2024-11-29 20:19:37.92877+08', '2024-11-29 20:19:37.92877+08');
COMMIT;

-- ----------------------------
-- Table structure for event_participants
-- ----------------------------
DROP TABLE IF EXISTS "public"."event_participants";
CREATE TABLE "public"."event_participants" (
  "content_id" int4 NOT NULL,
  "alumni_id" int4 NOT NULL,
  "id" int8 NOT NULL DEFAULT nextval('contents_id_seq'::regclass),
  "status" int2 NOT NULL,
  "comment" varchar(100) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."event_participants" OWNER TO "shawn";
COMMENT ON COLUMN "public"."event_participants"."status" IS '0, 1, 2, 3, 4, 5';

-- ----------------------------
-- Records of event_participants
-- ----------------------------
BEGIN;
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (5, 2, 36, 3, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (181, 2, 37, 1, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (131, 10, 27, 4, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (141, 5, 42, 5, NULL);
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (151, 5, 43, 4, NULL);
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (121, 5, 45, 4, NULL);
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (131, 5, 26, 1, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (311, 5, 39, 1, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (38, 5, 38, 5, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (49, 5, 54, 3, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (50, 9, 56, 3, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (251, 5, 41, 3, '');
INSERT INTO "public"."event_participants" ("content_id", "alumni_id", "id", "status", "comment") VALUES (50, 5, 53, 2, '');
COMMIT;

-- ----------------------------
-- Table structure for faculties
-- ----------------------------
DROP TABLE IF EXISTS "public"."faculties";
CREATE TABLE "public"."faculties" (
  "id" int4 NOT NULL DEFAULT nextval('faculty_id_seq'::regclass),
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."faculties" OWNER TO "shawn";

-- ----------------------------
-- Records of faculties
-- ----------------------------
BEGIN;
INSERT INTO "public"."faculties" ("id", "name") VALUES (2, 'Test Faculty 1');
INSERT INTO "public"."faculties" ("id", "name") VALUES (3, 'Institute For Advanced Studies (IAS)');
INSERT INTO "public"."faculties" ("id", "name") VALUES (4, 'Test Faculty 2');
INSERT INTO "public"."faculties" ("id", "name") VALUES (26, 'Asia-Europe Institute');
INSERT INTO "public"."faculties" ("id", "name") VALUES (27, 'International Institute Of Public Policy And Management');
INSERT INTO "public"."faculties" ("id", "name") VALUES (28, 'Institute Of China Studies');
INSERT INTO "public"."faculties" ("id", "name") VALUES (5, 'Academy Of Islamic Studies');
INSERT INTO "public"."faculties" ("id", "name") VALUES (6, 'Academy Of Malay Studies');
INSERT INTO "public"."faculties" ("id", "name") VALUES (7, 'Faculty Of Built Environment');
INSERT INTO "public"."faculties" ("id", "name") VALUES (8, 'Faculty Of Languages And Linguistics');
INSERT INTO "public"."faculties" ("id", "name") VALUES (9, 'Faculty Of Pharmacy');
INSERT INTO "public"."faculties" ("id", "name") VALUES (10, 'Faculty Of Engineering');
INSERT INTO "public"."faculties" ("id", "name") VALUES (11, 'Faculty Of Education');
INSERT INTO "public"."faculties" ("id", "name") VALUES (12, 'Faculty Of Dentistry');
INSERT INTO "public"."faculties" ("id", "name") VALUES (13, 'Faculty Of Business And Economics');
INSERT INTO "public"."faculties" ("id", "name") VALUES (14, 'Faculty Of Medicine');
INSERT INTO "public"."faculties" ("id", "name") VALUES (15, 'Faculty Of Science');
INSERT INTO "public"."faculties" ("id", "name") VALUES (16, 'Faculty Of Computer Science And Information Technology');
INSERT INTO "public"."faculties" ("id", "name") VALUES (17, 'Faculty Of Arts And Social Sciences');
INSERT INTO "public"."faculties" ("id", "name") VALUES (18, 'Faculty Of Creative Arts');
INSERT INTO "public"."faculties" ("id", "name") VALUES (19, 'Faculty Of Law');
INSERT INTO "public"."faculties" ("id", "name") VALUES (20, 'Faculty Of Sport & Exercise Sciences');
INSERT INTO "public"."faculties" ("id", "name") VALUES (21, 'Academic Enhancement & Leadership Development Centre (ADEC)');
INSERT INTO "public"."faculties" ("id", "name") VALUES (22, 'Centre For Internship Training And Academic Enrichment (CITrA)');
INSERT INTO "public"."faculties" ("id", "name") VALUES (23, 'Universiti Malaya Centre For Continuing Education (UMCCed)');
INSERT INTO "public"."faculties" ("id", "name") VALUES (24, 'Centre For Foundation Studies');
INSERT INTO "public"."faculties" ("id", "name") VALUES (25, 'Centre For Civilisational Dialogue');
INSERT INTO "public"."faculties" ("id", "name") VALUES (1, 'Other/Not Specified');
COMMIT;

-- ----------------------------
-- Table structure for honors
-- ----------------------------
DROP TABLE IF EXISTS "public"."honors";
CREATE TABLE "public"."honors" (
  "id" int8 NOT NULL DEFAULT nextval('honors_id_seq'::regclass),
  "title" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "alumni_id" int8 NOT NULL,
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "date" date NOT NULL,
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."honors" OWNER TO "shawn";

-- ----------------------------
-- Records of honors
-- ----------------------------
BEGIN;
INSERT INTO "public"."honors" ("id", "title", "alumni_id", "description", "date", "created_at", "updated_at", "deleted_at") VALUES (2, 'sample ttttt', 4, 'adkjfgadjffaf', '0001-01-01', '2024-11-11 20:30:01.096792', '2024-11-11 20:30:01.096792', '2024-11-11 20:30:01.096334');
INSERT INTO "public"."honors" ("id", "title", "alumni_id", "description", "date", "created_at", "updated_at", "deleted_at") VALUES (4, 'sample tttt爾爾t', 4, 'adkjfgadjffaf', '0001-01-01', '2024-11-12 14:10:33.45046', '2024-11-12 14:10:33.45046', '2024-11-12 14:10:33.45022');
INSERT INTO "public"."honors" ("id", "title", "alumni_id", "description", "date", "created_at", "updated_at", "deleted_at") VALUES (5, 'sample tttt爾爾t', 5, 'adkjfgadjffaf', '0001-01-01', '2024-11-25 17:19:06.013665', '2024-11-25 17:19:06.013665', '2024-11-25 17:19:06.013556');
COMMIT;

-- ----------------------------
-- Table structure for preferences
-- ----------------------------
DROP TABLE IF EXISTS "public"."preferences";
CREATE TABLE "public"."preferences" (
  "id" int4 NOT NULL,
  "content_id" int4,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "alumni_id" int4,
  "like" bool,
  "dislike" bool
)
;
ALTER TABLE "public"."preferences" OWNER TO "shawn";

-- ----------------------------
-- Records of preferences
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for publications
-- ----------------------------
DROP TABLE IF EXISTS "public"."publications";
CREATE TABLE "public"."publications" (
  "id" int4 NOT NULL DEFAULT nextval('publication_id_seq'::regclass),
  "alumni_id" int4,
  "sequence_no" int8 NOT NULL,
  "article_title" varchar(1000) COLLATE "pg_catalog"."default" NOT NULL,
  "journal_article" varchar(255) COLLATE "pg_catalog"."default",
  "quartile" varchar(10) COLLATE "pg_catalog"."default",
  "authors" varchar(1000) COLLATE "pg_catalog"."default" NOT NULL,
  "corresponding_authors" varchar(1000) COLLATE "pg_catalog"."default" NOT NULL,
  "accepted_date" date,
  "journal_title" varchar(1000) COLLATE "pg_catalog"."default" NOT NULL,
  "publication_type" varchar(25) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "status" varchar(25) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."publications" OWNER TO "shawn";

-- ----------------------------
-- Records of publications
-- ----------------------------
BEGIN;
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (26, 46, 0, 'Ling, Ong & ghaffari khaligh, Nader & Juan, Joon Ching. (2020). Recent Catalytic Advance to the Synthesis of Organic Symmetric Disulfides. Current Organic Chemistry. 24. 10.2174/1385272824666200221111120.', NULL, '', '', '', '0001-01-01', '', '', '2025-01-09 14:36:02.379811', '2025-01-09 14:37:06.402699', '2025-01-09 14:36:02.379769', '');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (16, 5, 0, 'Advances in Quantum Computing', NULL, 'Q4', 'Shuo Li, Alice Johnson, Bob Brown', 'Eve White, Frank Green', '2024-06-11', 'IEEE Access', 'Journal', '2024-12-04 22:55:57.25179', '2025-01-03 21:09:25.680822', '2024-12-04 22:55:57.25159', 'Published');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (15, 5, 0, 'Machine Learning for Healthcare', NULL, 'Q2', 'Shuo Li, Grace Lee, Henry Martin', 'Charlie Davis', '2024-08-22', 'Journal of Medical Systems', 'Scopus', '2024-12-04 22:55:32.812692', '2025-01-03 21:33:48.774068', '2024-12-04 22:55:32.812197', 'Accepted');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (21, 5, 0, 'viva test', NULL, 'Q2', 'emma', '', '2024-12-31', '', 'Scopus', '2025-01-06 23:21:26.564378', '2025-01-06 23:21:26.564378', '2025-01-06 23:21:26.56111', 'Published');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (22, 40, 0, 'Khiang, Lee. (2009). The Role of Chinese Clan Associations for Singapore''s Economic Development.', NULL, '', '', '', '0001-01-01', '', 'Others', '2025-01-09 14:36:01.965636', '2025-01-09 14:36:01.965636', '2025-01-09 14:36:01.965579', '');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (23, 41, 0, 'Muzir, N.A.Q.; Mojumder, M.R.H.; Hasanuzzaman, M.; Selvaraj, J. Challenges of Electric Vehicles and Their Prospects in Malaysia: A Comprehensive Review. Sustainability 2022, 14, 8320. https://doi.org/10.3390/su14148320', NULL, 'Q2', '', '', '0001-01-01', '', '', '2025-01-09 14:36:02.03757', '2025-01-09 14:37:06.313854', '2025-01-09 14:36:02.037407', '');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (27, 41, 0, 'Muzir, N. A. Q., Hasanuzzaman, M., & Selvaraj, J. (2023). Modeling and Analyzing the Impact of Different Operating Conditions for Electric and Conventional Vehicles in Malaysia on Energy, Economic, and the Environment. Energies, 16(13), 5048. https://doi.org/10.3390/en16135048', NULL, 'Q3', '', '', '0001-01-01', '', '', '2025-01-09 14:37:06.316282', '2025-01-09 14:37:06.316282', '2025-01-09 14:37:06.316211', '');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (28, 42, 0, 'Gunasegaran, Muthu Kumaran & Hasanuzzaman, Md & Tan, Chia Kwang & Abu Bakar, Ab Halim & Ponniah, Vignes. (2022). Energy Analysis, Building Energy Index and Energy Management Strategies for Fast-Food Restaurants in Malaysia. Sustainability. 14. 10.3390/su142013515.', NULL, '', '', '', '0001-01-01', '', '', '2025-01-09 14:37:06.383077', '2025-01-09 14:37:06.383077', '2025-01-09 14:37:06.383019', '');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (29, 43, 0, 'Rengasamy, N., Othman, R. Y., Che, H. S., & Harikrishna, J. A. (2022). Beyond the PAR spectra: impact of light quality on the germination, flowering, and metabolite content of Stevia rebaudiana (Bertoni). Journal of the Science of Food and Agriculture, 102(1), 299-311', NULL, 'Q1', '', '', '0001-01-01', '', 'Others', '2025-01-09 14:37:06.387183', '2025-01-09 14:37:06.387183', '2025-01-09 14:37:06.386948', '');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (24, 43, 0, 'Rengasamy, N., Othman, R. Y., Che, H. S., & Harikrishna, J. A. (2022). Artificial Lighting Photoperiod Manipulation Approach to Improve Productivity and Energy Use Efficacies of Plant Factory Cultivated Stevia rebaudiana. Agronomy, 12(8), 1787', NULL, 'Q1', '', '', '0001-01-01', '', 'Lain-lain', '2025-01-09 14:36:02.178788', '2025-01-09 14:37:06.388024', '2025-01-09 14:36:02.178742', '');
INSERT INTO "public"."publications" ("id", "alumni_id", "sequence_no", "article_title", "journal_article", "quartile", "authors", "corresponding_authors", "accepted_date", "journal_title", "publication_type", "created_at", "updated_at", "deleted_at", "status") VALUES (25, 44, 0, 'Kumar, Sarraf & Hamid, Suraya. (2017). Analysis of Learning Analytics in Higher Educational Institutions: A Review.', NULL, 'Q1', '', '', '0001-01-01', '', 'Others', '2025-01-09 14:36:02.244974', '2025-01-09 14:37:06.390539', '2025-01-09 14:36:02.244893', '');
COMMIT;

-- ----------------------------
-- Table structure for staffs
-- ----------------------------
DROP TABLE IF EXISTS "public"."staffs";
CREATE TABLE "public"."staffs" (
  "id" int4 NOT NULL DEFAULT nextval('staff_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default",
  "email" text COLLATE "pg_catalog"."default",
  "is_super_admin" bool NOT NULL DEFAULT false,
  "faculty_id" int4 DEFAULT 1,
  "phone" text COLLATE "pg_catalog"."default",
  "username" text COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "is_hidden" bool,
  "gender" varchar(10) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "token" varchar(500) COLLATE "pg_catalog"."default",
  "position" varchar(25) COLLATE "pg_catalog"."default",
  "deleted_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."staffs" OWNER TO "shawn";

-- ----------------------------
-- Records of staffs
-- ----------------------------
BEGIN;
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (29, 'super admin', 's2001314@siswa.um.edu.my2', 't', 13, '', 'admin (shawn1)', '$2a$10$ICM4PtLCOSVrZkV1NiQ3xOJ8Zskkj/a7EMBv1POETSgecVA1woGMa', 't', '', '2024-11-11 17:51:58.310181', '2024-11-28 10:28:40.148728', NULL, NULL, NULL);
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (30, 'super admin', 's2001314@siswa.um.edu.my3', 't', 13, '', 'admin (shawn1test)', '$2a$10$1CZHLYEXGsVEZKsT49676.5hHV//T5RNl18utn5/Ohr4MjtPo/y6G', 't', '', '2024-11-11 17:52:43.503569', '2024-11-28 10:29:52.706362', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozMCwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzMxODQyMzUyLCJpYXQiOjE3MzE4MzE1NTIsImlzcyI6ImFsdW1uaV9odWIifQ._IiM3zhO4d6tbtHsO4O2-4sYWlbcpEIcWj30wUd8HMo', NULL, '2024-11-11 17:52:43.503334');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (16, 'lishuo', 'lishuo@live.o', 'f', 7, '60-13134264246', 'ls6', '$2a$10$3se0BgMJo4CP8pFrR0/iguAK/o5SHUfY70t0PtLX0yD4eYXojD.Ay', 't', '', '2024-06-08 22:37:31.676053', '2024-11-28 10:36:12.438641', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzE4MjExNjUyLCJpYXQiOjE3MTgyMDA4NTIsImlzcyI6ImFsdW1uaV9odWIifQ.0ekdNIqjMF-8MLO-HGRZQE0nJJbncAnnATp0tFQepUU', NULL, '2024-11-11 17:45:22.735106');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (18, 'male', 'li-shuo@live.com2', 't', 1, 's', 's', '$2a$10$/FvJvB8c2KxAakQvTlX30egQWQWogVXXhtAS5y7l/692f4hkJr98q', 't', 'Male', '2024-06-09 12:49:49.092959', '2024-11-11 18:37:47.751042', NULL, NULL, '2024-11-11 17:45:22.735106');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (41, '', '', 'f', 1, '', '', '$2a$10$6KMySA8AjmQitw8naED7POmIvpQ/LV4MRKaH/XZ/i2IV8YSvasio6', 't', '', '2025-01-04 19:07:57.097848', '2025-01-04 19:08:01.367887', NULL, '', '2025-01-04 19:07:57.096872');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (42, 'Viva', 'test@ss.com', 't', 1, '', 'vivatest', '$2a$10$OwAXUX2yidE0tWRE7Edznuw1C8tfpr7wfZctAlGQx/2RmRjFtbXJC', 't', 'Female', '2025-01-06 21:51:58.357738', '2025-01-06 21:52:47.703055', NULL, '', '2025-01-06 21:51:58.357508');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (2, 'shawn', 'li-shuo@live.com', 'f', 15, '60-10913044', 'test', 'asd123', 'f', 'Female', '2024-06-07 06:24:02.681879', '2024-11-30 22:04:35.368983', NULL, NULL, '2024-11-11 17:45:22.735106');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (26, 'shawnli', 'shuo@live.', 'f', 18, '', 'shawnli27', '$2a$10$LMeWb4VqW/7iJzfU0vASCeZPffha7Cthjiib02XMFSJUUBNtqgSrS', 'f', '', '2024-06-12 22:04:14.91231', '2024-12-13 23:53:24.048554', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyNiwicm9sZSI6InN0YWZmIiwiZXhwIjoxNzM2NjkzNTY2LCJpYXQiOjE3MzQxMDE1NjYsImlzcyI6ImFsdW1uaV9odWIifQ.xp6LpGAvkrJdJXfcY5xL5uv59M_rFHJTairFeD5Bojw', NULL, '2024-11-11 17:45:22.735106');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (28, 'super admin', 's2001314@siswa.um.edu.my', 't', 13, '', 'admin (shawn)', '$2a$10$HTUt00Kd87OoXZ44iOFyquDN9.60UAd6ndPlVlpHFsGn63Wj/13LC', 'f', 'Male', '2024-11-11 17:47:19.750586', '2024-12-02 21:42:46.826091', NULL, NULL, NULL);
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (7, 'qingziyu', 'qingziyu', 't', 12, '60-13134264246', 'qingzi71', 'asd123', 'f', 'Male', '2024-06-07 06:24:02.681879', '2024-12-02 21:42:56.3417', NULL, NULL, '2024-11-11 17:45:22.735106');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (40, 'Shawn Li', 'shwawn@live.com', 'f', 1, '1313131', 'shawntest0103', '$2a$10$ciZG9gUS3WYHfU4oKhWFzeLRMQ4k3PuFDIDzCBwVDkPtiNX5SDl7y', 'f', 'Male', '2025-01-03 21:38:32.312501', '2025-01-10 12:29:20.178666', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0MCwicm9sZSI6InN0YWZmIiwiZXhwIjoxNzM5MDc1MzYwLCJpYXQiOjE3MzY0ODMzNjAsImlzcyI6ImFsdW1uaV9odWIifQ.IPEGKlrTcsUibNF2AihwvNjjGgnsMLmP-WgNhZdLSDo', 'some', '2025-01-03 21:38:32.311499');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (20, 'shawn 121313patch', 's@hawn.com', 't', 7, '13131', 'shawnlilili', '$2a$10$ddD9MiSvxNe3g1m0pk90b.6bSlvBIOAsWKIg.WovfzzFxj9NMTGKC', 'f', 'Female', '2024-06-09 21:32:37.757977', '2024-12-02 21:43:17.151927', NULL, NULL, '2024-11-11 17:45:22.735106');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (33, '31', '131', 't', 1, '121', 'dddddddd', '$2a$10$6qm7U0JBqCAh4kR2KrJ.6uBx9a4k8FRmfLwVB.i.DWsdle6yw.sPO', 'f', '', '2024-11-28 10:48:59.648535', '2024-12-02 21:45:47.835043', NULL, '', '2024-11-28 10:48:59.648326');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (36, 'sad', 'dsa2222222', 'f', 1, '13131', 'dsadasd2234', '$2a$10$6nRJhWCgmpoeq3IAfHOwrucd91se6UBZhLs.wE6rj3PFo.9gxvK86', 'f', '', '2024-11-28 10:53:46.754342', '2024-12-02 21:46:46.708493', NULL, '', '2024-11-28 10:53:46.754149');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (38, 'Shawn Ko', 'shawn.ko@gmail.com', 't', 1, '0138967452', 'shawnko', '$2a$10$224hHpAN9jYT3x5nwn0uI.XHcSJ5hD5wwV.ukmdkKEyXFhEc1/v.q', 'f', 'Male', '2024-12-02 21:47:49.7775', '2024-12-02 21:47:49.7775', NULL, 'IAS Admin', '2024-12-02 21:47:49.776411');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (6, 'Test2', 'li-shuo@live1.com', 't', 2, '60-109130784', 'test4', '$2a$10$HODoY6zbCvTPBh2Xa8902O1fZE6L7wZAbBgy5Jq1ok/K93vqzre7u', 't', NULL, '2024-06-07 06:24:02.681879', '2024-12-13 21:01:49.642559', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJyb2xlIjoic3RhZmYiLCJleHAiOjE3MzY2ODY5MDksImlhdCI6MTczNDA5NDkwOSwiaXNzIjoiYWx1bW5pX2h1YiJ9.xWbkpTFOVnMctlOHUQ_G1BWH8DbKfxGEdyQnTgTllLo', NULL, '2024-11-11 17:45:22.735106');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (39, 'Whi', 'Whi@live.com', 'f', 1, '121313', 'fdadafdafadf', '$2a$10$ym9aJ9Q5/M0mHax7qCHjveyPZgcBO4e1J91bs6FjMipjKX.SVtdvO', 't', 'Female', '2024-12-03 14:09:40.21472', '2025-01-05 18:39:33.978554', NULL, 's', '2024-12-03 14:09:40.211276');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (34, 'ssadasasd', 'da@live.com', 'f', 3, '22131232', 'saddsadsadasd', '$2a$10$1P3W5syA./.TMY4LmNLfa.hfhXCOp0y0SDkWGk3k3O7/zHLJaaKCK', 't', 'Female', '2024-11-28 10:51:42.206427', '2025-01-05 18:39:37.824302', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozNCwicm9sZSI6InN0YWZmIiwiZXhwIjoxNzM2Njg4Njg4LCJpYXQiOjE3MzQwOTY2ODgsImlzcyI6ImFsdW1uaV9odWIifQ.JTUm-_k31FdCetoC5-TGv-MCqk19klxwPtNxhG1RD5I', 'sd', '2024-11-28 10:51:42.206252');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (35, 'we', 'wew', 'f', 1, '', 'weewqeqwe', '$2a$10$e.xB.CjCZsPPV.jS6yBu9.F5yBMq1Tfk7nW.E1ty2V2LSDlHTUZU6', 't', '', '2024-11-28 10:52:27.996378', '2025-01-05 19:58:48.589055', NULL, '', '2024-11-28 10:52:27.99619');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (37, 'laohu', 'ti@ger.com', 'f', 2, '', 'laohu', '$2a$10$Ly/bLOmd6g1J2.7XZos/u.nx8VDQOw4.HGjEPqblv90FG0.BXshCy', 't', '', '2024-11-28 11:02:07.730526', '2025-01-05 19:58:53.00284', NULL, '', '2024-11-28 11:02:07.730372');
INSERT INTO "public"."staffs" ("id", "name", "email", "is_super_admin", "faculty_id", "phone", "username", "password", "is_hidden", "gender", "created_at", "updated_at", "token", "position", "deleted_at") VALUES (32, 'super admin', 's20013154@siswa.um.edu.my', 't', 13, '', 'shawn', '$2a$10$HwOngL338xnKrItww5XKXu5YBMM1L08n1qTemYEvsHTV0clH0V8A.', 'f', '', '2024-11-26 15:16:28.112996', '2025-01-20 18:27:13.762335', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozMiwicm9sZSI6InN0YWZmIiwiZXhwIjoxNzM5OTYwODMzLCJpYXQiOjE3MzczNjg4MzMsImlzcyI6ImFsdW1uaV9odWIifQ.kTP2gagx8GbSFpG8Y19KJh3jdPYhnPLKJT9aIeW2jmM', 'backend develop', '2024-11-26 15:16:28.112803');
COMMIT;

-- ----------------------------
-- Table structure for studies
-- ----------------------------
DROP TABLE IF EXISTS "public"."studies";
CREATE TABLE "public"."studies" (
  "id" int8 NOT NULL DEFAULT nextval('studies_id_seq'::regclass),
  "alumni_id" int8 NOT NULL,
  "level_of_study" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "faculty_id" int8 NOT NULL,
  "programme" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "intake_year" int8 NOT NULL,
  "intake_session" varchar(10) COLLATE "pg_catalog"."default" NOT NULL,
  "convocation_year" int8,
  "status" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "title_of_thesis" varchar(255) COLLATE "pg_catalog"."default",
  "supervisor" varchar(100) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."studies" OWNER TO "shawn";

-- ----------------------------
-- Records of studies
-- ----------------------------
BEGIN;
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (27, 5, 'Master', 16, 'Master of Computer Science', 2023, 'Spring', 2024, 'Completed', 'Advanced Deep Learning Techniques for Enhancing Natural Language Processing', ' Dr. Grace Goh Jia Li', '2024-12-04 22:14:11.230858', '2024-12-05 17:39:37.743165', '2024-12-04 22:14:11.230665');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (50, 38, 'Master', 15, 'Master of Applied Mathematics', 2021, '2021', 2029, '', '', '', '2025-01-08 20:02:53.563654', '2025-01-08 21:29:49.722007', '2025-01-08 20:02:53.563541');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (51, 39, 'Master', 14, 'Master of Medicine Science', 2021, '2021', 2029, '', '', '', '2025-01-08 20:02:53.675042', '2025-01-08 21:29:49.723718', '2025-01-08 20:02:53.674893');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (14, 7, 'Master', 15, 'Bachelor of CS', 2021, '2021', 2029, '', '', '', '2024-11-29 13:12:50.242523', '2024-11-29 13:12:50.242523', '2024-11-29 13:12:50.242469');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (15, 8, 'Master', 15, 'Bachelor of CS', 2021, '2021', 2029, '', '', '', '2024-11-29 13:12:50.311571', '2024-11-29 13:12:50.311571', '2024-11-29 13:12:50.311519');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (16, 9, 'Master', 14, 'Bachelor of CS', 2021, '2021', 2029, '', '', '', '2024-11-29 13:12:50.375942', '2024-11-29 13:12:50.375942', '2024-11-29 13:12:50.375841');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (17, 10, 'Master', 12, 'Bachelor of CS', 2021, '2021', 2029, '', '', '', '2024-11-29 13:12:50.44248', '2024-11-29 13:12:50.44248', '2024-11-29 13:12:50.442361');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (18, 11, 'Master', 15, 'Bachelor of CS', 2021, '2021', 2029, '', '', '', '2024-11-29 13:12:50.514341', '2024-11-29 13:12:50.514341', '2024-11-29 13:12:50.51427');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (19, 12, 'Master', 14, 'Bachelor of CS', 2021, '2021', 2029, '', '', '', '2024-11-29 13:12:50.58251', '2024-11-29 13:12:50.58251', '2024-11-29 13:12:50.582398');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (20, 8, 'Master', 15, 'Bachelor of CS', 2022, '2022', 2029, '', '', '', '2024-11-29 13:30:03.1438', '2024-11-29 13:30:03.1438', '2024-11-29 13:30:03.143664');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (5, 5, 'Bachelor', 15, 'Bachelor of Mathematics', 2021, 'Spring', 2023, 'Completed', 'Mathematical Modeling and Simulation of Complex Distributed Networks', 'Prof. Dr. Andrew Chan ', '2024-11-26 16:15:21.076538', '2024-12-05 19:53:09.448006', '2024-11-26 16:15:21.076244');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (37, 27, 'Master', 16, 'Master of Computer Science', 2021, '2021', 2029, '', '', '', '2024-12-05 19:58:11.270237', '2024-12-05 19:58:11.270237', '2024-12-05 19:58:11.270149');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (38, 28, 'Master', 16, 'Master of Computer Science', 2022, '2022', 2029, '', '', '', '2024-12-05 19:58:11.34128', '2024-12-05 19:58:11.34128', '2024-12-05 19:58:11.341213');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (39, 29, 'Master', 14, 'Master of Medicine Science', 2021, '2021', 2029, '', '', '', '2024-12-05 19:58:11.409826', '2024-12-05 19:58:11.409826', '2024-12-05 19:58:11.409649');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (40, 30, 'Master', 12, 'Master of Dentistry', 2021, '2021', 2029, '', '', '', '2024-12-05 19:58:11.481315', '2024-12-05 19:58:11.481315', '2024-12-05 19:58:11.481242');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (41, 31, 'Master', 15, 'Master of Applied Mathematics', 2021, '2021', 2029, '', '', '', '2024-12-05 19:58:11.550555', '2024-12-05 19:58:11.550555', '2024-12-05 19:58:11.550422');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (42, 32, 'Master', 14, 'Master of Medicine Science', 2021, '2021', 2029, '', '', '', '2024-12-05 19:58:11.6305', '2024-12-05 19:58:11.6305', '2024-12-05 19:58:11.630423');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (46, 34, 'Master', 16, 'Master of Computer Science', 2021, '2021', 2029, '', '', '', '2025-01-08 20:02:53.111894', '2025-01-08 21:29:49.710876', '2025-01-08 20:02:53.111845');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (47, 35, 'Master', 16, 'Master of Computer Science', 2022, '2022', 2029, '', '', '', '2025-01-08 20:02:53.224679', '2025-01-08 21:29:49.713131', '2025-01-08 20:02:53.224618');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (48, 36, 'Master', 14, 'Master of Medicine Science', 2021, '2021', 2029, '', '', '', '2025-01-08 20:02:53.334204', '2025-01-08 21:29:49.717672', '2025-01-08 20:02:53.334124');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (49, 37, 'Master', 12, 'Master of Dentistry', 2021, '2021', 2029, '', '', '', '2025-01-08 20:02:53.451804', '2025-01-08 21:29:49.720407', '2025-01-08 20:02:53.451744');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (54, 40, 'PhD', 0, '', 0, '', 0, '', '', '', '2025-01-09 14:36:01.956972', '2025-01-09 14:37:06.306869', '2025-01-09 14:36:01.956892');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (55, 41, 'PhD', 0, '', 0, '', 0, '', '', '', '2025-01-09 14:36:02.035658', '2025-01-09 14:37:06.311667', '2025-01-09 14:36:02.0356');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (56, 42, 'PhD', 0, '', 0, '', 0, '', '', '', '2025-01-09 14:36:02.107735', '2025-01-09 14:37:06.381648', '2025-01-09 14:36:02.107644');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (57, 43, 'PhD', 0, '', 0, '', 0, '', '', '', '2025-01-09 14:36:02.176545', '2025-01-09 14:37:06.386143', '2025-01-09 14:36:02.176472');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (58, 44, 'PhD', 0, '', 0, '', 0, '', '', '', '2025-01-09 14:36:02.244137', '2025-01-09 14:37:06.389567', '2025-01-09 14:36:02.244084');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (59, 45, 'PhD', 0, '', 0, '', 0, '', '', '', '2025-01-09 14:36:02.31254', '2025-01-09 14:37:06.392226', '2025-01-09 14:36:02.312404');
INSERT INTO "public"."studies" ("id", "alumni_id", "level_of_study", "faculty_id", "programme", "intake_year", "intake_session", "convocation_year", "status", "title_of_thesis", "supervisor", "created_at", "updated_at", "deleted_at") VALUES (60, 46, 'PhD', 0, '', 0, '', 0, '', '', '', '2025-01-09 14:36:02.379142', '2025-01-09 14:37:06.401239', '2025-01-09 14:36:02.379092');
COMMIT;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS "public"."tags";
CREATE TABLE "public"."tags" (
  "id" int4 NOT NULL DEFAULT nextval('tags_sequence'::regclass),
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "name" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."tags" OWNER TO "shawn";

-- ----------------------------
-- Records of tags
-- ----------------------------
BEGIN;
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (1, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'management');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (2, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'marketing');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (3, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'evaluation');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (4, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'article');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (5, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'methods');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (6, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'ai');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (7, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'aspect');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (8, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'capabilities');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (9, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'key');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (10, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'events');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (11, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'technology');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (12, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'planning');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (13, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'execution');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (14, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'reading');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (15, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'takeaways');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (16, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'power');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (17, '2024-11-14 20:47:50.449212', '2024-11-14 20:47:50.449212', 'event');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (18, '2024-12-02 15:16:22.412341', '2024-12-02 15:16:22.412341', 'sample');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (19, '2024-12-05 16:38:38.610189', '2024-12-05 16:38:38.610189', 'world');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (20, '2024-12-05 16:38:38.610189', '2024-12-05 16:38:38.610189', 'art');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (21, '2024-12-05 16:38:38.610189', '2024-12-05 16:38:38.610189', 'understanding');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (22, '2024-12-05 16:39:59.14334', '2024-12-05 16:39:59.14334', 'conservation');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (23, '2024-12-05 16:39:59.14334', '2024-12-05 16:39:59.14334', 'planet');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (24, '2025-01-10 12:22:10.048007', '2025-01-10 12:22:10.048007', 'career fair');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (25, '2025-01-10 12:22:10.048007', '2025-01-10 12:22:10.048007', 'join');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (26, '2025-01-10 12:22:10.048007', '2025-01-10 12:22:10.048007', 'spring');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (27, '2025-01-10 12:22:10.048007', '2025-01-10 12:22:10.048007', 'career');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (28, '2025-01-10 12:22:10.048007', '2025-01-10 12:22:10.048007', 'fair');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (29, '2025-01-10 12:22:10.048007', '2025-01-10 12:22:10.048007', 'network');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (30, '2025-01-10 12:22:10.048007', '2025-01-10 12:22:10.048007', 'internship');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (31, '2025-01-10 12:22:10.048007', '2025-01-10 12:22:10.048007', 'job');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (32, '2025-01-10 12:22:40.555857', '2025-01-10 12:22:40.555857', 'campus');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (33, '2025-01-10 12:22:40.555857', '2025-01-10 12:22:40.555857', 'food');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (34, '2025-01-10 12:22:40.555857', '2025-01-10 12:22:40.555857', 'experience');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (35, '2025-01-10 12:22:40.555857', '2025-01-10 12:22:40.555857', 'evening');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (36, '2025-01-10 12:22:40.555857', '2025-01-10 12:22:40.555857', 'diverse');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (37, '2025-01-10 12:23:10.929798', '2025-01-10 12:23:10.929798', 'lecture');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (38, '2025-01-10 12:23:10.929798', '2025-01-10 12:23:10.929798', 'industry');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (39, '2025-01-10 12:23:10.929798', '2025-01-10 12:23:10.929798', 'impact');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (40, '2025-01-10 12:23:10.929798', '2025-01-10 12:23:10.929798', 'society');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (41, '2025-01-10 12:23:41.07549', '2025-01-10 12:23:41.07549', 'university');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (42, '2025-01-10 12:23:41.07549', '2025-01-10 12:23:41.07549', 'library');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (43, '2025-01-10 12:23:41.07549', '2025-01-10 12:23:41.07549', 'month');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (44, '2025-01-10 12:23:41.07549', '2025-01-10 12:23:41.07549', 'please');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (45, '2025-01-10 12:23:41.07549', '2025-01-10 12:23:41.07549', 'website');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (46, '2025-01-10 12:23:41.07549', '2025-01-10 12:23:41.07549', 'location');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (47, '2025-01-10 12:24:02.07286', '2025-01-10 12:24:02.07286', 'access');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (48, '2025-01-10 12:24:02.07286', '2025-01-10 12:24:02.07286', 'gym');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (49, '2025-01-10 12:24:02.07286', '2025-01-10 12:24:02.07286', 'training');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (50, '2025-01-10 12:24:02.07286', '2025-01-10 12:24:02.07286', 'rate');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (51, '2025-01-10 12:24:02.07286', '2025-01-10 12:24:02.07286', 'sign');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (52, '2025-01-10 12:24:02.07286', '2025-01-10 12:24:02.07286', 'student');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (53, '2025-01-10 12:24:02.07286', '2025-01-10 12:24:02.07286', 'fitness');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (54, '2025-01-10 12:24:02.07286', '2025-01-10 12:24:02.07286', 'membership');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (55, '2025-01-10 16:28:36.12302', '2025-01-10 16:28:36.12302', 'descrpdescrp');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (56, '2025-01-10 16:28:36.12302', '2025-01-10 16:28:36.12302', 'descrpdescrpdescrp');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (57, '2025-01-10 16:28:36.12302', '2025-01-10 16:28:36.12302', 'descrpdescrpdescrpdescrpdescrpdescrp');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (58, '2025-01-10 16:28:36.12302', '2025-01-10 16:28:36.12302', 'crpdescrpdescrpdescrpdescrpdescrpdescrpdescrpdescrpdescrpdescrpdescrpd');
INSERT INTO "public"."tags" ("id", "created_at", "updated_at", "name") VALUES (59, '2025-01-10 16:28:36.12302', '2025-01-10 16:28:36.12302', 'escrpdescrpdescrpdescrpdescrpdescrpdescrp');
COMMIT;

-- ----------------------------
-- Table structure for work_experiences
-- ----------------------------
DROP TABLE IF EXISTS "public"."work_experiences";
CREATE TABLE "public"."work_experiences" (
  "id" int8 NOT NULL DEFAULT nextval('work_experiences_id_seq'::regclass),
  "alumni_id" int8 NOT NULL,
  "workplace" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "position" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "country" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "city" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date,
  "status" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "occupation_field" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."work_experiences" OWNER TO "shawn";

-- ----------------------------
-- Records of work_experiences
-- ----------------------------
BEGIN;
INSERT INTO "public"."work_experiences" ("id", "alumni_id", "workplace", "position", "country", "city", "start_date", "end_date", "status", "occupation_field", "created_at", "updated_at", "deleted_at") VALUES (1, 4, 'alibaba', 'backend', 'Singapore', 'Singapore', '0001-01-01', '0001-01-01', 'on', '', '2024-11-11 20:45:57.939275', '2024-11-11 20:45:57.939275', '2024-11-11 20:45:57.938983');
INSERT INTO "public"."work_experiences" ("id", "alumni_id", "workplace", "position", "country", "city", "start_date", "end_date", "status", "occupation_field", "created_at", "updated_at", "deleted_at") VALUES (3, 4, 'bytedance', 'backend', 'Singapore', 'Singapore', '0001-01-01', '0001-01-01', 'on', '', '2024-11-11 20:46:17.56111', '2024-11-11 20:46:17.56111', '2024-11-11 20:46:17.560993');
INSERT INTO "public"."work_experiences" ("id", "alumni_id", "workplace", "position", "country", "city", "start_date", "end_date", "status", "occupation_field", "created_at", "updated_at", "deleted_at") VALUES (6, 5, 'Bytedance', 'Full Stack Developer', 'Singapore', 'Singapore', '2020-03-19', '2024-07-11', 'Past', 'Cloud-Native applications using Go, Docker and Kubernetes', '2024-11-27 14:01:48.680728', '2024-12-05 17:46:10.859541', '2024-11-27 14:01:48.679645');
INSERT INTO "public"."work_experiences" ("id", "alumni_id", "workplace", "position", "country", "city", "start_date", "end_date", "status", "occupation_field", "created_at", "updated_at", "deleted_at") VALUES (4, 5, 'Grab', 'Backend Developer', 'Malaysia', 'Kuala Lumpur', '2019-01-09', '2020-01-15', 'Past', 'Automating deployment, DIstributed systems', '2024-11-25 17:18:45.872625', '2024-12-05 17:47:44.758146', '2024-11-25 17:18:45.87242');
INSERT INTO "public"."work_experiences" ("id", "alumni_id", "workplace", "position", "country", "city", "start_date", "end_date", "status", "occupation_field", "created_at", "updated_at", "deleted_at") VALUES (11, 19, 'w', 'p', 'c', 'ci', '2025-01-07', '0001-01-01', 'Current', 'o', '2025-01-10 16:07:22.026511', '2025-01-10 16:08:28.120947', '2025-01-10 16:07:22.026315');
COMMIT;

-- ----------------------------
-- Function structure for update_updated_at_column
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."update_updated_at_column"();
CREATE OR REPLACE FUNCTION "public"."update_updated_at_column"()
  RETURNS "pg_catalog"."trigger" AS $BODY$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION "public"."update_updated_at_column"() OWNER TO "shawn";

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."attachments_id_seq"', 37, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."contents_id_seq"', 56, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."email_attachments_id_seq"
OWNED BY "public"."email_attachments"."id";
SELECT setval('"public"."email_attachments_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."emails_id_seq"
OWNED BY "public"."emails"."id";
SELECT setval('"public"."emails_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."faculty_id_seq"
OWNED BY "public"."faculties"."id";
SELECT setval('"public"."faculty_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."honors_id_seq"
OWNED BY "public"."honors"."id";
SELECT setval('"public"."honors_id_seq"', 5, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."publication_id_seq"
OWNED BY "public"."publications"."id";
SELECT setval('"public"."publication_id_seq"', 31, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."staff_id_seq"
OWNED BY "public"."staffs"."id";
SELECT setval('"public"."staff_id_seq"', 42, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."studies_id_seq"
OWNED BY "public"."studies"."id";
SELECT setval('"public"."studies_id_seq"', 63, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."tags_sequence"', 59, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."alumnis"."id";
SELECT setval('"public"."users_id_seq"', 46, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."work_experiences_id_seq"
OWNED BY "public"."work_experiences"."id";
SELECT setval('"public"."work_experiences_id_seq"', 11, true);

-- ----------------------------
-- Primary Key structure for table alumni_tags
-- ----------------------------
ALTER TABLE "public"."alumni_tags" ADD CONSTRAINT "alumni_tags_pkey" PRIMARY KEY ("alumni_id", "tag_id");

-- ----------------------------
-- Indexes structure for table alumnis
-- ----------------------------
CREATE UNIQUE INDEX "idx_users_matric_no" ON "public"."alumnis" USING btree (
  "matric_no" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table alumnis
-- ----------------------------
ALTER TABLE "public"."alumnis" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table attachments
-- ----------------------------
ALTER TABLE "public"."attachments" ADD CONSTRAINT "event_images_copy1_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table content_tags
-- ----------------------------
ALTER TABLE "public"."content_tags" ADD CONSTRAINT "content_tags_pkey" PRIMARY KEY ("content_id", "tag_id");

-- ----------------------------
-- Primary Key structure for table contents
-- ----------------------------
ALTER TABLE "public"."contents" ADD CONSTRAINT "events_copy1_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table email_attachments
-- ----------------------------
ALTER TABLE "public"."email_attachments" ADD CONSTRAINT "email_attachments_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table emails
-- ----------------------------
ALTER TABLE "public"."emails" ADD CONSTRAINT "emails_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table event_participants
-- ----------------------------
ALTER TABLE "public"."event_participants" ADD CONSTRAINT "unique" UNIQUE ("alumni_id", "content_id");

-- ----------------------------
-- Primary Key structure for table event_participants
-- ----------------------------
ALTER TABLE "public"."event_participants" ADD CONSTRAINT "content_tags_copy1_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table faculties
-- ----------------------------
ALTER TABLE "public"."faculties" ADD CONSTRAINT "faculty_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table honors
-- ----------------------------
ALTER TABLE "public"."honors" ADD CONSTRAINT "honors_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table preferences
-- ----------------------------
ALTER TABLE "public"."preferences" ADD CONSTRAINT "attachments_copy1_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table publications
-- ----------------------------
ALTER TABLE "public"."publications" ADD CONSTRAINT "publication_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Triggers structure for table staffs
-- ----------------------------
CREATE TRIGGER "update_staffs_updated_at" BEFORE UPDATE ON "public"."staffs"
FOR EACH ROW
EXECUTE PROCEDURE "public"."update_updated_at_column"();

-- ----------------------------
-- Uniques structure for table staffs
-- ----------------------------
ALTER TABLE "public"."staffs" ADD CONSTRAINT "staff_email" UNIQUE ("email");
ALTER TABLE "public"."staffs" ADD CONSTRAINT "staff_username_key" UNIQUE ("username");

-- ----------------------------
-- Primary Key structure for table staffs
-- ----------------------------
ALTER TABLE "public"."staffs" ADD CONSTRAINT "staff_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table studies
-- ----------------------------
ALTER TABLE "public"."studies" ADD CONSTRAINT "studies_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table tags
-- ----------------------------
ALTER TABLE "public"."tags" ADD CONSTRAINT "preferences_copy1_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table work_experiences
-- ----------------------------
ALTER TABLE "public"."work_experiences" ADD CONSTRAINT "work_experiences_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table alumni_tags
-- ----------------------------
ALTER TABLE "public"."alumni_tags" ADD CONSTRAINT "alumni_tags_alumni_id_fkey" FOREIGN KEY ("alumni_id") REFERENCES "public"."alumnis" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "public"."alumni_tags" ADD CONSTRAINT "alumni_tags_tag_id_fkey" FOREIGN KEY ("tag_id") REFERENCES "public"."tags" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table email_attachments
-- ----------------------------
ALTER TABLE "public"."email_attachments" ADD CONSTRAINT "email_attachments_email_id_fkey" FOREIGN KEY ("email_id") REFERENCES "public"."emails" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table event_participants
-- ----------------------------
ALTER TABLE "public"."event_participants" ADD CONSTRAINT "content_tags_copy1_content_id_fkey" FOREIGN KEY ("content_id") REFERENCES "public"."contents" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table honors
-- ----------------------------
ALTER TABLE "public"."honors" ADD CONSTRAINT "fk_users_honor" FOREIGN KEY ("alumni_id") REFERENCES "public"."alumnis" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table publications
-- ----------------------------
ALTER TABLE "public"."publications" ADD CONSTRAINT "pub-alu" FOREIGN KEY ("alumni_id") REFERENCES "public"."alumnis" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table staffs
-- ----------------------------
ALTER TABLE "public"."staffs" ADD CONSTRAINT "staff_faculty_id_fkey" FOREIGN KEY ("faculty_id") REFERENCES "public"."faculties" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table studies
-- ----------------------------
ALTER TABLE "public"."studies" ADD CONSTRAINT "fk_users_studies" FOREIGN KEY ("alumni_id") REFERENCES "public"."alumnis" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table work_experiences
-- ----------------------------
ALTER TABLE "public"."work_experiences" ADD CONSTRAINT "fk_users_work_experiences" FOREIGN KEY ("alumni_id") REFERENCES "public"."alumnis" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
