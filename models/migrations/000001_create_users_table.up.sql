CREATE TYPE user_role AS ENUM ('basic', 'premium', 'admin');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    google_id varchar(255) UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255),
    reminder_enabled boolean DEFAULT false,
    reminder_days integer DEFAULT 0,
    role user_role DEFAULT 'basic',
    timezone VARCHAR(50) DEFAULT 'UTC',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  
)

