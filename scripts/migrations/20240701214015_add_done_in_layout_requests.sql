-- Modify "layout_requests" table
ALTER TABLE "layout_requests" ADD COLUMN "done" integer NOT NULL DEFAULT 0, ADD COLUMN "total" integer NULL, ADD COLUMN "updated_at" timestamp NULL;
