defmodule Education.Service do
  use GRPC.Server, service: Education.Api.V1alpha1.EducationService.Service

  alias Education.Api.V1alpha1.{
    Region,
    School,
    Location,
    Link,
    Student
  }

  @spec list_schools(Enumerable.t(), GRPC.Server.Stream.t()) :: any
  def list_schools(school, stream) do
    GRPC.Server.send_reply(stream, school)
  end
end
