-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vehicle(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    brand VARCHAR(150) NOT NULL,
    model VARCHAR(150) NOT NULL,
    year int NOT NULL CHECK (year > 0),
    price bigint NOT NULL CHECK (price > 0),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE IF EXISTS vehicle OWNER TO oncarapp;

CREATE TABLE IF NOT EXISTS lead(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(180) NOT NULL,
    email VARCHAR(150) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    vehicle_id uuid NOT NULL REFERENCES vehicle(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE IF EXISTS lead OWNER TO oncarapp;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE lead;
DROP TABLE vehicle;
-- +goose StatementEnd


