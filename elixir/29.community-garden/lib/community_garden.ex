# Use the Plot struct as it is provided
defmodule Plot do
  @enforce_keys [:plot_id, :registered_to]
  defstruct [:plot_id, :registered_to]
end

defmodule CommunityGarden do
  def start(opts \\ []) do
    Agent.start(fn -> [] end, opts)
  end

  def list_registrations(pid) do
    Agent.get(pid, fn state -> state end)
  end

  def register(pid, register_to) do
    Agent.update(pid, fn state ->
      if state[:plot_id] == nil do
        %Plot{plot_id: 0, registered_to: register_to}
      else
        %Plot{plot_id: state + 1, registered_to: register_to}
      end
    end)

    # Agent.
    list_registrations(pid)
  end

  def release(pid, plot_id) do
    # Please implement the release/2 function
  end

  def get_registration(pid, plot_id) do
    # Please implement the get_registration/2 function
  end
end
