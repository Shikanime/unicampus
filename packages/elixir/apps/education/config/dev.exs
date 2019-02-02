use Mix.Config

config :education, Education.Repo,
  username: "postgres",
  password: "postgres",
  database: "education_dev",
  hostname: "localhost",
  pool_size: 10
