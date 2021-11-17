return {
  { sql = [[ CREATE TABLE "app_user" (
    "name" char(100)
); ]], passed = false },
   { sql = [[ CREATE TABLE "app_user" (
     "id" serial NOT NULL PRIMARY KEY,
     "name" varchar(100) NOT NULL,
     "email" TEXT NOT NULL
 ); ]], passed = true },
}
