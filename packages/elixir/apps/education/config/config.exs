use Mix.Config

config :education,
  ecto_repos: [Education.Repo]

import_config "#{Mix.env()}.exs"
