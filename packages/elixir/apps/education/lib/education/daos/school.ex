defmodule Education.SchoolDao do
  use Ecto.Schema
  import Ecto.Changeset

  schema "schools" do
    field :name
    field :description

    timestamps()
  end

  @doc false
  def changeset(bite, attrs) do
    bite
    |> cast(attrs, [])
    |> validate_required([])
  end
end
