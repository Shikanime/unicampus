defmodule Education.Service do
  use GRPC.Server, service: Education.Api.V1alpha1.EducationService.Service

  alias Education.Api.V1alpha1.{
    Region,
    School,
    Location,
    Link,
    Student
  }

  def list_schools(schools, stream) do
    schools
    |> EducationPostgres.list_schools()
    |> Enum.each(&(GRPC.Server.send_reply(stream, &1)))
  end

  def register_school(school, _stream) do
    with {:ok, _} <- EducationPostgres.create_school(school) do
      :ok
    end
  end

  def update_school(school, _stream) do
    school
    |> EducationPostgres.get_school()
    |> EducationPostgres.update_school(school)
  end

  def unregister_school(school, _stream) do
    school
    |> EducationPostgres.delete_school()
  end
end
