defmodule Education.SchoolSchema do
  use Ecto.Schema
  import Ecto.Changeset

  alias Education.LinkSchema
  alias Education.LocationSchema

  schema "schools" do
    field :name
    field :description
    field :phone
    field :email
    embeds_many :redirections, LinkSchema
    embeds_many :pictures, LinkSchema
    has_many :locations, LocationSchema

    timestamps()
  end

  @doc false
  def changeset(school, attrs) do
    school
    |> cast(attrs, [])
    |> validate_required([])
  end
end
