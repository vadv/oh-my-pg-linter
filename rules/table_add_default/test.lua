return {
  { sql = [[ ALTER TABLE "core_recipe" ADD COLUMN "foo_add_default" integer DEFAULT 10 NOT NULL; ]], passed = false },
  { sql = [[ ALTER TABLE "core_recipe" ADD COLUMN "foo_not_default" integer; ]], passed = true },
}
