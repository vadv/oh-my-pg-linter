return {
  { sql = [[ CREATE TABLE "app_user" (
    "name" varchar(100)
); ]], passed = false },
   { sql = [[ CREATE TABLE "app_user" (
     "name" text
 ); ]], passed = true },
}
