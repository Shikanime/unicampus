defmodule Education.Repo do
  use Ecto.Repo,
    otp_app: :education,
    adapter: Ecto.Adapters.Postgres
end
