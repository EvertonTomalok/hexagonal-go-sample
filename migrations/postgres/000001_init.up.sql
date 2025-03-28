CREATE TABLE ports (
    id SERIAL PRIMARY KEY,
    identifier TEXT UNIQUE,
    name TEXT NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    province TEXT,
    timezone TEXT,
    code TEXT,
    alias TEXT[],
    regions TEXT[],
    coordinates DOUBLE PRECISION[],
    unlocs TEXT[]
);