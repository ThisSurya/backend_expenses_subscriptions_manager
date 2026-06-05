CREATE TABLE categories(
    id serial PRIMARY KEY,
    user_id integer NOT NULL,
    name varchar(100) NOT NULL,
    color varchar(7) NOT NULL,
    icon_url varchar(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  ,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)