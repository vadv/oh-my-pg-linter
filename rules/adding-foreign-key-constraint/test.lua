return {
  { sql = [[ ALTER TABLE "email" ADD CONSTRAINT "fk_user"
    FOREIGN KEY ("user_id") REFERENCES "user" ("id"); ]], passed = false },
  { sql = [[ ALTER TABLE "email" ADD CONSTRAINT "fk_user"
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") NOT VALID;]], passed = true },
}
