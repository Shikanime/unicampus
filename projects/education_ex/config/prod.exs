use Mix.Config

config :logger, level: :info

config :education, Education.Repo,
  username: "unicampus",
  password: "unicampus",
  database: "education",
  pool_size: 15
