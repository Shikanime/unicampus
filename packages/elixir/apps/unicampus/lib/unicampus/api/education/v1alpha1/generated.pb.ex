defmodule Unicampus.Api.Education.V1alpha1.Query do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          content: String.t()
        }
  defstruct [:content]

  field :content, 1, type: :string
end

defmodule Unicampus.Api.Education.V1alpha1.Critera do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          sector: String.t()
        }
  defstruct [:sector]

  field :sector, 1, type: :string
end

defmodule Unicampus.Api.Education.V1alpha1.School do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          UUID: String.t(),
          name: String.t(),
          description: String.t(),
          phone: String.t(),
          email: String.t(),
          pictures: [Unicampus.Api.Education.V1alpha1.Link.t()],
          regions: [Unicampus.Api.Education.V1alpha1.Link.t()],
          locations: [Unicampus.Api.Education.V1alpha1.Location.t()]
        }
  defstruct [:UUID, :name, :description, :phone, :email, :pictures, :regions, :locations]

  field :UUID, 1, type: :string
  field :name, 2, type: :string
  field :description, 3, type: :string
  field :phone, 4, type: :string
  field :email, 5, type: :string
  field :pictures, 6, repeated: true, type: Unicampus.Api.Education.V1alpha1.Link
  field :regions, 7, repeated: true, type: Unicampus.Api.Education.V1alpha1.Link
  field :locations, 8, repeated: true, type: Unicampus.Api.Education.V1alpha1.Location
end

defmodule Unicampus.Api.Education.V1alpha1.Student do
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

defmodule Unicampus.Api.Education.V1alpha1.Link do
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

defmodule Unicampus.Api.Education.V1alpha1.Location do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          address: String.t(),
          geo_point: Unicampus.Api.Education.V1alpha1.GeoPoint.t(),
          region: Unicampus.Api.Education.V1alpha1.Region.t()
        }
  defstruct [:address, :geo_point, :region]

  field :address, 1, type: :string
  field :geo_point, 2, type: Unicampus.Api.Education.V1alpha1.GeoPoint
  field :region, 3, type: Unicampus.Api.Education.V1alpha1.Region
end

defmodule Unicampus.Api.Education.V1alpha1.GeoPoint do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          latitude: float,
          longitude: float
        }
  defstruct [:latitude, :longitude]

  field :latitude, 1, type: :double
  field :longitude, 2, type: :double
end

defmodule Unicampus.Api.Education.V1alpha1.Region do
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

defmodule Unicampus.Api.Education.V1alpha1.AdmissionService.Service do
  @moduledoc false
  use GRPC.Service, name: "unicampus.api.education.v1alpha1.AdmissionService"

  rpc :ListSchools,
      stream(Unicampus.Api.Education.V1alpha1.School),
      stream(Unicampus.Api.Education.V1alpha1.School)

  rpc :ListSchoolsByQuery,
      Unicampus.Api.Education.V1alpha1.Query,
      stream(Unicampus.Api.Education.V1alpha1.School)

  rpc :ListSchoolsByCritera,
      Unicampus.Api.Education.V1alpha1.Critera,
      stream(Unicampus.Api.Education.V1alpha1.School)

  rpc :RegisterSchool,
      Unicampus.Api.Education.V1alpha1.School,
      Unicampus.Api.Education.V1alpha1.School

  rpc :UpdateSchool,
      Unicampus.Api.Education.V1alpha1.School,
      Unicampus.Api.Education.V1alpha1.School

  rpc :UnregisterSchool,
      Unicampus.Api.Education.V1alpha1.School,
      Unicampus.Api.Education.V1alpha1.School

  rpc :RegisterStudent,
      Unicampus.Api.Education.V1alpha1.Student,
      Unicampus.Api.Education.V1alpha1.Student

  rpc :UpdateStudent,
      Unicampus.Api.Education.V1alpha1.Student,
      Unicampus.Api.Education.V1alpha1.Student

  rpc :UnregisterStudent,
      Unicampus.Api.Education.V1alpha1.Student,
      Unicampus.Api.Education.V1alpha1.Student
end

defmodule Unicampus.Api.Education.V1alpha1.AdmissionService.Stub do
  @moduledoc false
  use GRPC.Stub, service: Unicampus.Api.Education.V1alpha1.AdmissionService.Service
end
