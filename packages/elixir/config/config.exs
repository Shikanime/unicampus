use Mix.Config

import_config "../apps/*/config/config.exs"

config :logger, :console,
  format: "$time $metadata[$level] $message\n"

import_config "#{Mix.env()}.exs"
