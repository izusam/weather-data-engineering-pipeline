CREATE TABLE weather_data (
    id SERIAL PRIMARY KEY,
    location VARCHAR(255),
    temperature FLOAT,
    recorded_at TIMESTAMP
);
