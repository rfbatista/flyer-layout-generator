-- Modify "layout_jobs" table
ALTER TABLE "layout_jobs" ADD CONSTRAINT "layout_jobs_created_layout_id_fkey" FOREIGN KEY ("created_layout_id") REFERENCES "layout" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
