defmodule Education.Repo.Migrations.CreateSchools do
  use Ecto.Migration

  def change do
    create table(:schools) do
      add :name
      add :description
      add :phone
      add :email
      add :redirections, {:array, :map}
      add :pictures, {:array, :map}
      add :locations, references(:locations)

      timestamps()
    end
  end
end
