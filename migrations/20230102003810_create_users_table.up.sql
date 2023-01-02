CREATE TABLE users (
  id uuid DEFAULT uuid_generate_v4 (),
  username TEXT NOT NULL UNIQUE,
  encrypted_password TEXT NOT NULL,
  email TEXT,
  name TEXT,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER users_updated_at
  BEFORE UPDATE
  ON users
  FOR EACH ROW
  EXECUTE FUNCTION moddatetime(updated_at);
