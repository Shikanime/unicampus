use Mix.Config

config :logger, level: :warn

config :education, Education.Repo,
  username: "postgres",
  password: "postgres",
  database: "education_test",
  hostname: "localhost",
  pool: Ecto.Adapters.SQL.Sandbox
