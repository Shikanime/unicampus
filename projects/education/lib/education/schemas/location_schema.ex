defmodule Education.LocationSchema do
  use Ecto.Schema
  import Ecto.Changeset

  alias Education.GeoPointSchema
  alias Education.RegionSchema

  schema "locations" do
    field :address
    field :longitude, :float
    field :latitude, :float
    has_one :region, RegionSchema
  end

  @doc false
  def changeset(location, attrs) do
    location
    |> cast(attrs, [])
    |> validate_required([])
  end
end
