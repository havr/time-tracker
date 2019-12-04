CREATE TABLE work_sessions
(
  id UUID,
  name VARCHAR(20) NOT NULL,
  start_time TIMESTAMPTZ NOT NULL,
  duration INTERVAL SECOND(0) NOT NULL,

  CONSTRAINT "pk_work_sessions_id" PRIMARY KEY(id)
);

INSERT INTO work_sessions (id, name, start_time, duration) VALUES
    ('e8ac8621-46bf-4e84-bff5-a3f7561b3a80', 'recently', now() - interval '5 minutes', interval '10 minutes'),
    ('a8ecbcf8-ef72-4f86-976f-ce383f2c386e', 'last hour', now() - interval '1 hour', interval '30 minutes'),
    ('5b78ad58-b488-4812-b65f-8f2da203d082', 'three days ago', now() - interval '3 days', interval '3 hours'),
    ('66a4adae-216e-46da-90a3-60db0b9deb54', 'two weeks ago', now() - interval '2 weeks', interval '2 hours');
