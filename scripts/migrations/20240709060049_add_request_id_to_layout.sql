-- Modify "layout" table
ALTER TABLE "layout" ADD COLUMN "request_id" integer NULL, ADD CONSTRAINT "fk_layout_request_id" FOREIGN KEY ("request_id") REFERENCES "layout_requests" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
