defmodule Education.Education.RegionSchema do
  use Ecto.Schema
  import Ecto.Changeset

  schema "regions" do
    field :country
    field :state
    field :zipcode
    field :city
  end

  @doc false
  def changeset(region, attrs) do
    region
    |> cast(attrs, [])
    |> validate_required([])
  end
end
