-- 000002_add_users.down.sql

-- Drop foreign key constraint
ALTER TABLE IF EXISTS accounts DROP CONSTRAINT IF EXISTS accounts_owner_fkey;

-- Drop unique constraint
ALTER TABLE IF EXISTS accounts DROP CONSTRAINT IF EXISTS owner_currency_key;

-- Drop users table
DROP TABLE IF EXISTS users;
