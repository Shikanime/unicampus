use Mix.Config

config :logger, :console, format: "[$level] $message\n"

config :education, Education.Repo,
  username: "postgres",
  password: "postgres",
  database: "education_dev",
  hostname: "localhost",
  pool_size: 10
