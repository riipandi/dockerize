ALTER TABLE hydra_oauth2_consent_request ADD COLUMN amr TEXT NOT NULL DEFAULT '';
ALTER TABLE hydra_oauth2_authentication_request_handled ADD COLUMN amr TEXT NOT NULL DEFAULT '';

