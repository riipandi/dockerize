INSERT INTO "_selfservice_settings_flows_tmp" (id, request_url, issued_at, expires_at, identity_id, created_at, updated_at, active_method, state, type, ui) SELECT id, request_url, issued_at, expires_at, identity_id, created_at, updated_at, active_method, state, type, ui FROM "selfservice_settings_flows";