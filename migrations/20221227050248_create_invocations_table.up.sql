CREATE TYPE invocation_status  AS ENUM ('pending', 'running', 'failed', 'succeeded');
CREATE TABLE invocations (
    id uuid DEFAULT uuid_generate_v4 (),
    status invocation_status DEFAULT 'pending',
    function_name TEXT NOT NULL,
    function_id uuid,
    image TEXT NOT NULL,
    input TEXT,
    output TEXT,
    error_message TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP WITHOUT TIME ZONE,
    ended_at TIMESTAMP WITHOUT TIME ZONE
);
CREATE TRIGGER invocations_updated_at
  BEFORE UPDATE
  ON invocations
  FOR EACH ROW
    EXECUTE FUNCTION moddatetime(updated_at);
