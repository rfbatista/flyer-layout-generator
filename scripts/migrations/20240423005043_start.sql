-- Create enum type "component_type"
CREATE TYPE "component_type" AS ENUM ('background', 'logotipo_marca', 'logotipo_produto', 'texto_cta');
-- Create "images" table
CREATE TABLE "images" ("id" bigserial NOT NULL, "url" text NOT NULL, "photoshop_id" integer NULL, "template_id" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "photoshop" table
CREATE TABLE "photoshop" ("id" serial NOT NULL, "name" text NOT NULL, "image_url" text NULL, "file_url" text NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, PRIMARY KEY ("id"));
-- Create "template" table
CREATE TABLE "template" ("id" serial NOT NULL, "name" text NOT NULL, "width" integer NULL, "height" integer NULL, "slots_x" integer NULL, "slots_y" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, PRIMARY KEY ("id"));
-- Create "photoshop_element" table
CREATE TABLE "photoshop_element" ("id" serial NOT NULL, "photoshop_id" integer NOT NULL, "name" text NULL, "layer_id" text NULL, "text" text NULL, "xi" integer NULL, "xii" integer NULL, "yi" integer NULL, "yii" integer NULL, "width" integer NULL, "height" integer NULL, "is_group" boolean NULL, "group_id" integer NULL, "level" integer NULL, "kind" text NULL, "component_id" text NULL, "component_type" "component_type" NULL, "image_url" text NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, PRIMARY KEY ("id"), CONSTRAINT "photoshop_element_photoshop_id_fkey" FOREIGN KEY ("photoshop_id") REFERENCES "photoshop" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
