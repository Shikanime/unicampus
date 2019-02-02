defmodule Education.Service do
  use GRPC.Server, service: Unicampus.Api.Education.V1alpha1.AdmissionService.Service

  alias Unicampus.Api.Education.V1alpha1.AdmissionService

  @spec list_schools(AdmissionService.School.t(), GRPC.Server.Stream.t()) :: AdmissionService.Shool.t()
  def list_schools(request, _stream) do
  end
end
