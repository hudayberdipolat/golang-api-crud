CREATE TABLE IF NOT EXISTS posts(
    id serial PRIMARY KEY,
    post_title TEXT NOT NULL,
    post_slug TEXT NOT NULL,
    post_desc text NOT NULL,
    post_image VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
    );


