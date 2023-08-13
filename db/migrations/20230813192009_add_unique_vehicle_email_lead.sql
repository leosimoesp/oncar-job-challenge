-- +goose Up
-- +goose StatementBegin

-- +goose StatementEnd
BEGIN;
ALTER TABLE lead ADD CONSTRAINT lead_vehicle_mail_ukey UNIQUE(vehicle_id, email);
COMMIT;
-- +goose Down
-- +goose StatementBegin
BEGIN;
ALTER TABLE lead DROP CONSTRAINt lead_vehicle_mail_ukey;
COMMIT;
-- +goose StatementEnd
