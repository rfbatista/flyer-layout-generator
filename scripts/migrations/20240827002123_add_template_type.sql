-- Add value to enum type: "template_type"
ALTER TYPE "template_type" ADD VALUE 'unknown';
-- Add value to enum type: "layout_job_status"
ALTER TYPE "layout_job_status" ADD VALUE 'canceled';
-- Modify "templates" table
ALTER TABLE "templates" ADD COLUMN "type" "template_type" NULL;
-- Modify "layout_jobs" table
ALTER TABLE "layout_jobs" DROP CONSTRAINT "layout_jobs_layout_id_fkey", ADD COLUMN "created_layout_id" integer NULL, ADD COLUMN "user_id" integer NULL, ADD COLUMN "config" text NULL, ADD COLUMN "updated_at" timestamp NULL, ADD CONSTRAINT "layout_jobs_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Rename a column from "layout_id" to "based_on_layout_id"
ALTER TABLE "layout_jobs" RENAME COLUMN "layout_id" TO "based_on_layout_id";
