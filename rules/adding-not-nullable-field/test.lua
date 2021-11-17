return {
  { sql = [[ ALTER TABLE "core_recipe" ADD COLUMN "foo" integer DEFAULT 10 NOT NULL; ]], passed = false },
  { sql = [[ ALTER TABLE "core_recipe" ADD COLUMN "foo" integer DEFAULT 10; ]], passed = true },
}
