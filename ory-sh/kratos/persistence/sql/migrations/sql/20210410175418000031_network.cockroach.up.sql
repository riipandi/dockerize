ALTER TABLE "selfservice_errors" ADD CONSTRAINT "selfservice_errors_nid_fk_idx" FOREIGN KEY ("nid") REFERENCES "networks" ("id") ON UPDATE RESTRICT ON DELETE CASCADE;