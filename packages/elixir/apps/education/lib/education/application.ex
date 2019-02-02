defmodule Education.Application do
  use Application

  def start(_type, _args) do
    children = [
      {Education.Repo, []},
    ]

    opts = [strategy: :one_for_one, name: Education.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
