return {
  { sql = [[ ALTER TABLE "a" ADD CONSTRAINT "positive_balance" CHECK ("balance" >= 0); ]], passed = false },
  { sql = [[ ALTER TABLE "a" ADD CONSTRAINT "positive_balance" CHECK ("balance" >= 0) NOT VALID;]], passed = true },
}
