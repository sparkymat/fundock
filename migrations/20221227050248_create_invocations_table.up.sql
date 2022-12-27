CREATE TABLE invocations (
    id uuid DEFAULT uuid_generate_v4 (),
    function_name TEXT NOT NULL,
    function_id uuid,
    image TEXT NOT NULL,
    input TEXT,
    output TEXT,
    exec_duration_ms BIGINT,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    executed_at TIMESTAMP WITHOUT TIME ZONE
);
CREATE TRIGGER invocations_updated_at
  BEFORE UPDATE
  ON invocations
  FOR EACH ROW
    EXECUTE FUNCTION moddatetime(updated_at);
