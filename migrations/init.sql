CREATE TABLE IF NOT EXISTS urls (
                                    id SERIAL PRIMARY KEY,
                                    original_url TEXT UNIQUE NOT NULL,
                                    short_url VARCHAR(10) UNIQUE NOT NULL
    );
