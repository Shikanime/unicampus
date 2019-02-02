defmodule Education.Api.V1alpha1.School do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          UUID: String.t(),
          name: String.t(),
          description: String.t(),
          phone: String.t(),
          email: String.t(),
          pictures: [Education.Api.V1alpha1.Link.t()],
          regions: [Education.Api.V1alpha1.Link.t()],
          locations: [Education.Api.V1alpha1.Location.t()]
        }
  defstruct [:UUID, :name, :description, :phone, :email, :pictures, :regions, :locations]

  field :UUID, 1, type: :string
  field :name, 2, type: :string
  field :description, 3, type: :string
  field :phone, 4, type: :string
  field :email, 5, type: :string
  field :pictures, 6, repeated: true, type: Education.Api.V1alpha1.Link
  field :regions, 7, repeated: true, type: Education.Api.V1alpha1.Link
  field :locations, 8, repeated: true, type: Education.Api.V1alpha1.Location
end

defmodule Education.Api.V1alpha1.Student do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          UUID: String.t(),
          first_name: String.t(),
          last_name: String.t(),
          phone: String.t(),
          email: String.t()
        }
  defstruct [:UUID, :first_name, :last_name, :phone, :email]

  field :UUID, 1, type: :string
  field :first_name, 2, type: :string
  field :last_name, 3, type: :string
  field :phone, 4, type: :string
  field :email, 5, type: :string
end

defmodule Education.Api.V1alpha1.Link do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          Type: String.t(),
          Reference: String.t()
        }
  defstruct [:Type, :Reference]

  field :Type, 1, type: :string
  field :Reference, 2, type: :string
end

defmodule Education.Api.V1alpha1.Location do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          address: String.t(),
          latitude: float,
          longitude: float,
          region: Education.Api.V1alpha1.Region.t()
        }
  defstruct [:address, :latitude, :longitude, :region]

  field :address, 1, type: :string
  field :latitude, 2, type: :double
  field :longitude, 3, type: :double
  field :region, 4, type: Education.Api.V1alpha1.Region
end

defmodule Education.Api.V1alpha1.Region do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          city: String.t(),
          state: String.t(),
          country: String.t(),
          zipcode: String.t()
        }
  defstruct [:city, :state, :country, :zipcode]

  field :city, 1, type: :string
  field :state, 2, type: :string
  field :country, 3, type: :string
  field :zipcode, 4, type: :string
end

defmodule Education.Api.V1alpha1.EducationService.Service do
  @moduledoc false
  use GRPC.Service, name: "education.api.v1alpha1.EducationService"

  rpc :ListSchools, stream(Education.Api.V1alpha1.School), stream(Education.Api.V1alpha1.School)
  rpc :RegisterSchool, Education.Api.V1alpha1.School, Education.Api.V1alpha1.School
  rpc :UpdateSchool, Education.Api.V1alpha1.School, Education.Api.V1alpha1.School
  rpc :UnregisterSchool, Education.Api.V1alpha1.School, Education.Api.V1alpha1.School
  rpc :RegisterStudent, Education.Api.V1alpha1.Student, Education.Api.V1alpha1.Student
  rpc :UpdateStudent, Education.Api.V1alpha1.Student, Education.Api.V1alpha1.Student
  rpc :UnregisterStudent, Education.Api.V1alpha1.Student, Education.Api.V1alpha1.Student
end

defmodule Education.Api.V1alpha1.EducationService.Stub do
  @moduledoc false
  use GRPC.Stub, service: Education.Api.V1alpha1.EducationService.Service
end
