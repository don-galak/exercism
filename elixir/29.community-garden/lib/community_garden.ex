# Use the Plot struct as it is provided
defmodule Plot do
  @enforce_keys [:plot_id, :registered_to]
  defstruct [:plot_id, :registered_to]
end

defmodule CommunityGarden do
  @not_registered {:not_found, "plot is unregistered"}

  def start(_opts \\ []), do: Agent.start_link(fn -> %{plots: [], id: 0} end)

  def list_registrations(pid) do
    Agent.get(pid, fn %{plots: plots} -> plots end)
  end

  def register(pid, register_to) do
    Agent.get_and_update(pid, fn %{plots: plots, id: id} ->
      id = id + 1
      plot = %Plot{plot_id: id, registered_to: register_to}
      {plot, %{plots: [plot | plots], id: id}}
    end)
  end

  def release(pid, plot_id) do
    Agent.cast(pid, fn %{plots: plots} = state ->
      plots = Enum.filter(plots, fn %{plot_id: p} -> p != plot_id end)
      %{state | plots: plots}
    end)
  end

  def get_registration(pid, plot_id) do
    Agent.get(pid, fn %{plots: plots} ->
      Enum.find(plots, @not_registered, fn %{plot_id: p} -> p === plot_id end)
    end)
  end
end
