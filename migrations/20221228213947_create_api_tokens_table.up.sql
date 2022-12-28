CREATE TABLE api_tokens (
  id uuid DEFAULT uuid_generate_v4 (),
  client_name TEXT NOT NULL,
  token TEXT NOT NULL UNIQUE,
  last_used_at TIMESTAMP WITHOUT TIME ZONE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER api_tokens_updated_at
  BEFORE UPDATE
  ON api_tokens
  FOR EACH ROW
    EXECUTE FUNCTION moddatetime(updated_at);
