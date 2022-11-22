CREATE TABLE IF NOT EXISTS public.user (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  name varchar(100) NOT NULL,
  role varchar(15) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);

CREATE TABLE IF NOT EXISTS price (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  admin_id VARCHAR (15) NOT NULL,
  price_topup DECIMAL(12, 3) NOT NULL,
  price_buyback DECIMAL(12, 3) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);  


CREATE TABLE IF NOT EXISTS account (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  account_no varchar(20) NOT NULL,
  balance DECIMAL(12, 3) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);

CREATE TABLE IF NOT EXISTS topup (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  price varchar(100) NOT NULL,
  gram varchar(100) NOT NULL,
  account_no varchar(20) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);

CREATE TABLE IF NOT EXISTS transaction (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  type varchar(20) NOT NULL,
  account_no varchar(20) NOT NULL,
  balance DECIMAL(12, 3) NOT NULL,
  gram DECIMAL(12, 3) NOT NULL,
  price_topup DECIMAL(12, 3) NOT NULL,
  price_buyback DECIMAL(12, 3) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);

INSERT INTO public.user (reff_id, name, role, created_at, updated_at) 
VALUES ('a001','admin1', 'admin', EXTRACT(epoch FROM NOW()), EXTRACT(epoch FROM NOW())),
('c001','customer1', 'customer', EXTRACT(epoch FROM NOW()), EXTRACT(epoch FROM NOW()));

INSERT INTO account (reff_id, account_no, balance, created_at, updated_at) 
VALUES ('account_no_1','ao001', 0, EXTRACT(epoch FROM NOW()), EXTRACT(epoch FROM NOW()));