defmodule Education.LinkSchema do
  use Ecto.Schema

  embedded_schema do
    field :name, :string
    field :age, :integer
    field :email, :string
    field :accepts_conditions, :boolean
  end
end
