CREATE TABLE messages_data (
  message_id string NOT NULL,
  data_id    string NOT NULL,
  data_hash  string NOT NULL,
  data_idx   int64  NOT NULL
);

CREATE UNIQUE INDEX messages_data_primary ON messages_data(message_id,data_id)