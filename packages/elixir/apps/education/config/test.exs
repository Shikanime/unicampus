use Mix.Config

config :education, Education.Repo,
  username: "postgres",
  password: "postgres",
  database: "education_test",
  hostname: "localhost",
  pool: Ecto.Adapters.SQL.Sandbox
