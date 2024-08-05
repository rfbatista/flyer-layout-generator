-- Create enum type "roles"
CREATE TYPE "roles" AS ENUM ('admin', 'colab', 'gm');
-- Create "companies" table
CREATE TABLE "companies" ("id" bigserial NOT NULL, "name" text NOT NULL, "enabled" boolean NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, PRIMARY KEY ("id"));
-- Modify "advertisers" table
ALTER TABLE "advertisers" ADD COLUMN "company_id" integer NULL, ADD CONSTRAINT "advertisers_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "clients" table
ALTER TABLE "clients" ADD COLUMN "company_id" integer NULL, ADD CONSTRAINT "clients_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Create "companies_api_credentials" table
CREATE TABLE "companies_api_credentials" ("id" bigserial NOT NULL, "name" text NULL, "api_key" text NOT NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, "company_id" integer NULL, PRIMARY KEY ("id"), CONSTRAINT "companies_api_credentials_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE);
-- Modify "design" table
ALTER TABLE "design" ADD COLUMN "company_id" integer NULL, ADD CONSTRAINT "design_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "layout" table
ALTER TABLE "layout" ADD COLUMN "company_id" integer NULL, ADD CONSTRAINT "layout_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "layout_requests" table
ALTER TABLE "layout_requests" ADD COLUMN "company_id" integer NULL, ADD CONSTRAINT "layout_requests_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "projects" table
ALTER TABLE "projects" ADD COLUMN "company_id" integer NULL, ADD CONSTRAINT "projects_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "templates" table
ALTER TABLE "templates" ADD COLUMN "company_id" integer NULL, ADD CONSTRAINT "templates_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Create "users" table
CREATE TABLE "users" ("id" bigserial NOT NULL, "name" text NOT NULL, "email" text NULL, "username" text NULL, "role" "roles" NULL DEFAULT 'colab', "company_id" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, PRIMARY KEY ("id"), CONSTRAINT "users_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE CASCADE ON DELETE CASCADE);
