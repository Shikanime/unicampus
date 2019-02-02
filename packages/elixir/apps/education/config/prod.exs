use Mix.Config

config :education, Education.Repo,
  username: {:system, "POSTGRES_USERNAME"},
  password: {:system, "POSTGRES_PASSWORD"},
  database: "education",
  pool_size: 15
