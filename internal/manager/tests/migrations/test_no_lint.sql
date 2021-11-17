-- nolint:rule_2
CREATE INDEX "email_no_lint_idx" ON "app_user" ("email");
-- must error with rule_2
CREATE INDEX "email_must_lint_idx" ON "app_user" ("email");
