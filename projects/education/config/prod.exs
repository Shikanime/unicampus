use Mix.Config

config :logger, level: :info

config :education, Education.Repo,
  username: "postgres",
  password: "postgres",
  database: "education",
  pool_size: 15
