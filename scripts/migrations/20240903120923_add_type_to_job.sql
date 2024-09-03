-- Create enum type "job_type"
CREATE TYPE "job_type" AS ENUM ('adaptation', 'replication', 'unknown');
-- Modify "adaptation_batch" table
ALTER TABLE "adaptation_batch" ADD COLUMN "type" "job_type" NULL;
