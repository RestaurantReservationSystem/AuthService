CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       user_name VARCHAR(50) UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                     token text,
                      created_at timestamp default current_timestamp,
                      updated_at timestamp default current_timestamp,
                      deleted_at timestamp
);
