use Mix.Config

config :logger, level: :warn

config :education, Education.Repo,
  username: "unicampus",
  password: "unicampus",
  database: "education_test",
  hostname: "localhost",
  pool: Ecto.Adapters.SQL.Sandbox
