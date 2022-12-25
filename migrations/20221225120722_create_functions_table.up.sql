CREATE EXTENSION moddatetime;
CREATE EXTENSION "uuid-ossp";
CREATE TABLE functions (
    id uuid DEFAULT uuid_generate_v4 (),
    name TEXT NOT NULL UNIQUE,
    image TEXT NOT NULL,
    skip_logging BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER functions_updated_at
  BEFORE UPDATE
  ON functions
  FOR EACH ROW
    EXECUTE FUNCTION moddatetime(updated_at);
