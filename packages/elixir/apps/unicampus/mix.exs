defmodule Unicampus.Umbrella.MixProject do
  use Mix.Project

  def project do
    [
      app: :unicampus,
      version: "0.1.0",
      build_path: "../../_build",
      config_path: "../../config/config.exs",
      deps_path: "../../deps",
      lockfile: "../../mix.lock",
      elixir: "~> 1.7",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  def application do
    [
      extra_applications: [:logger]
    ]
  end

  defp deps do
    [
      {:grpc, github: "tony612/grpc-elixir"}
    ]
  end
end
