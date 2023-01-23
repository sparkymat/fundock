ALTER TABLE functions
  ADD COLUMN environment jsonb NOT NULL DEFAULT '{}'::jsonb,
  ADD COLUMN secrets jsonb NOT NULL DEFAULT '{}'::jsonb;
