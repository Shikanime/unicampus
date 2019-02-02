defmodule Education.Repo.Migrations.CreateLocations do
  use Ecto.Migration

  def change do
    create table(:locations) do
      add :address
      add :longitude, :float
      add :latitude, :float
      add :region, references(:regions)
    end
  end
end
