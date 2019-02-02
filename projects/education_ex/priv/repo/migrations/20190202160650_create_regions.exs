defmodule Education.Repo.Migrations.CreateRegions do
  use Ecto.Migration

  def change do
    create table(:regions) do
      add :country
      add :state
      add :zipcode
      add :city
    end
  end
end
