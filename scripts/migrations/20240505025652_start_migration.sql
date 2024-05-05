-- Create enum type "component_type"
CREATE TYPE "component_type" AS ENUM ('background', 'logotipo_marca', 'logotipo_produto', 'texto_cta');
-- Create enum type "template_type"
CREATE TYPE "template_type" AS ENUM ('slots', 'distortion');
-- Create "images" table
CREATE TABLE "images" ("id" bigserial NOT NULL, "url" text NOT NULL, "photoshop_id" integer NULL, "template_id" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "photoshop" table
CREATE TABLE "photoshop" ("id" serial NOT NULL, "name" text NOT NULL, "image_url" text NULL, "image_extension" text NULL, "file_url" text NULL, "file_extension" text NULL, "width" integer NULL, "height" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, PRIMARY KEY ("id"));
-- Create "photoshop_components" table
CREATE TABLE "photoshop_components" ("id" serial NOT NULL, "photoshop_id" integer NOT NULL, "width" integer NULL, "height" integer NULL, "color" text NULL, "type" "component_type" NULL, "xi" integer NULL, "xii" integer NULL, "yi" integer NULL, "yii" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "photoshop_components_photoshop_id_fkey" FOREIGN KEY ("photoshop_id") REFERENCES "photoshop" ("id") ON UPDATE CASCADE ON DELETE CASCADE);
-- Create "photoshop_element" table
CREATE TABLE "photoshop_element" ("id" serial NOT NULL, "photoshop_id" integer NOT NULL, "name" text NULL, "layer_id" text NULL, "text" text NULL, "xi" integer NULL, "xii" integer NULL, "yi" integer NULL, "yii" integer NULL, "width" integer NULL, "height" integer NULL, "is_group" boolean NULL, "group_id" integer NULL, "level" integer NULL, "kind" text NULL, "component_id" integer NULL, "image_url" text NULL, "image_extension" text NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, PRIMARY KEY ("id"), CONSTRAINT "fk_photoshop_element_component_id" FOREIGN KEY ("component_id") REFERENCES "photoshop_components" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "fk_photoshop_element_photoshop_id" FOREIGN KEY ("photoshop_id") REFERENCES "photoshop" ("id") ON UPDATE CASCADE ON DELETE CASCADE);
-- Create "templates" table
CREATE TABLE "templates" ("id" serial NOT NULL, "name" text NOT NULL, "type" "template_type" NULL, "width" integer NULL, "height" integer NULL, "slots_x" integer NULL, "slots_y" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, PRIMARY KEY ("id"));
-- Create "templates_distortions" table
CREATE TABLE "templates_distortions" ("id" serial NOT NULL, "x" integer NULL, "y" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, "template_id" integer NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "templates_distortions_template_id_fkey" FOREIGN KEY ("template_id") REFERENCES "templates" ("id") ON UPDATE CASCADE ON DELETE CASCADE);
-- Create "templates_slots" table
CREATE TABLE "templates_slots" ("id" serial NOT NULL, "xi" integer NULL, "yi" integer NULL, "width" integer NULL, "height" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, "template_id" integer NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "templates_slots_template_id_fkey" FOREIGN KEY ("template_id") REFERENCES "templates" ("id") ON UPDATE CASCADE ON DELETE CASCADE);
