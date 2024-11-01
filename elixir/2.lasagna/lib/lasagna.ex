defmodule Lasagna do
  def expected_minutes_in_oven(), do: 40
  def remaining_minutes_in_oven(x), do: expected_minutes_in_oven() - x
  def preparation_time_in_minutes(x), do: 2 * x
  def total_time_in_minutes(x, y), do: preparation_time_in_minutes(x) + y
  def alarm(), do: "Ding!"
end
