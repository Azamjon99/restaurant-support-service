-- eater schema
CREATE SCHEMA IF NOT EXISTS eater;

-- eaters
CREATE TABLE IF NOT EXISTS eater.eaters (
	id varchar(36) PRIMARY KEY,
    phone_number VARCHAR(16),
    password_hash VARCHAR(255),
    password_salt VARCHAR(255),
	created_at timestamp,
	updated_at timestamp
);

-- idx_eater_phone_number
CREATE INDEX IF NOT EXISTS idx_eater_phone_number ON eater.eaters(phone_number);

-- eater_sms_codes
CREATE TABLE IF NOT EXISTS eater.eater_sms_codes (
	id serial PRIMARY KEY,
	eater_id varchar(36),
	code varchar(5),
	expires_in integer,
	created_at timestamp
);

-- idx_sms_code_by_eater
CREATE INDEX IF NOT EXISTS idx_sms_code_by_eater ON eater.eater_sms_codes(eater_id,code);

-- eaters_profile
CREATE TABLE IF NOT EXISTS eater.eater_profiles (
	eater_id varchar(36) PRIMARY KEY,
    phone_number VARCHAR(16),
    name VARCHAR(128),
    image_url VARCHAR(255),
	is_phone_number_confirmed BOOLEAN,
	created_at timestamp,
	updated_at timestamp
);


-- addresses
CREATE TABLE IF NOT EXISTS eater.addresses (
	id varchar(36) PRIMARY KEY,
	eater_id varchar(36) NOT NULL,
	name varchar(255),
	location jsonb,
	created_at timestamp,
	updated_at timestamp
);

-- idx_address_by_eater
CREATE INDEX IF NOT EXISTS idx_address_by_eater ON eater.addresses(eater_id);

-- payment_cards
CREATE TABLE IF NOT EXISTS eater.payment_cards (
	id varchar(36) PRIMARY KEY,
	eater_id varchar(36) NOT NULL,
	number varchar(100),
	card_token varchar(255),
	is_verified boolean,
	created_at timestamp
);

-- idx_address_by_eater
CREATE INDEX IF NOT EXISTS idx_card_by_eater ON eater.payment_cards(eater_id);

-- orders
CREATE TABLE IF NOT EXISTS eater.orders (
	id varchar(36) PRIMARY KEY,
	eater_id varchar(36) NOT NULL,
	instruction varchar(255),
	restaurant_id varchar(36) NOT NULL,
	restaurant jsonb,
	delivery jsonb,
	payment jsonb,
	items jsonb,
	status varchar(36),
	payment_status varchar(36),
	created_at timestamp,
	updated_at timestamp
);

-- idx_order_by_eater
CREATE INDEX IF NOT EXISTS idx_order_by_eater ON eater.orders(eater_id);
-- idx_order_by_restaurant
CREATE INDEX IF NOT EXISTS idx_order_by_eater ON eater.orders(restaurant_id);

-- restaurant_ratings
CREATE TABLE IF NOT EXISTS eater.restaurant_ratings (
	id varchar(36) PRIMARY KEY,
	eater_id varchar(36) NOT NULL,
	restaurant_id varchar(36) NOT NULL,
	rating smallint,
	comment text,
	created_at timestamp,
	updated_at timestamp
);

-- idx_r_rating_by_eater
CREATE INDEX IF NOT EXISTS idx_r_rating_by_eater ON eater.restaurant_ratings(eater_id);
-- idx_r_rating_by_restaurant
CREATE INDEX IF NOT EXISTS idx_r_rating_by_restaurant ON eater.restaurant_ratings(restaurant_id);

-- delivery_ratings
CREATE TABLE IF NOT EXISTS eater.delivery_ratings (
	id varchar(36) PRIMARY KEY,
	eater_id varchar(36) NOT NULL,
	order_id varchar(36) NOT NULL,
	rating smallint,
	comment text,
	created_at timestamp,
	updated_at timestamp
);

-- idx_d_rating_by_eater
CREATE INDEX IF NOT EXISTS idx_d_rating_by_eater ON eater.delivery_ratings(eater_id);
-- idx_d_rating_by_order
CREATE INDEX IF NOT EXISTS idx_d_rating_by_order ON eater.delivery_ratings(order_id);