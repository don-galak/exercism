defmodule KitchenCalculator do
  def get_volume(volume_pair), do: elem(volume_pair, 1)

  def to_milliliter({:cup, quantity}), do: {:milliliter, quantity * 240}
  def to_milliliter({:fluid_ounce, quantity}), do: {:milliliter, quantity * 30}
  def to_milliliter({:teaspoon, quantity}), do: {:milliliter, quantity * 5}
  def to_milliliter({:tablespoon, quantity}), do: {:milliliter, quantity * 15}
  def to_milliliter({:milliliter, quantity}), do: {:milliliter, quantity}

  def from_milliliter(volume_pair, unit) when unit == :cup, do: {unit, elem(volume_pair, 1)/240}
  def from_milliliter(volume_pair, unit) when unit == :fluid_ounce, do: {unit, elem(volume_pair, 1)/30}
  def from_milliliter(volume_pair, unit) when unit == :teaspoon, do: {unit, elem(volume_pair, 1)/5}
  def from_milliliter(volume_pair, unit) when unit == :tablespoon, do: {unit, elem(volume_pair, 1)/15}
  def from_milliliter(volume_pair, unit) when unit == :milliliter, do: {unit, elem(volume_pair, 1)}

  def convert(volume_pair, unit) do
    # Please implement the convert/2 function
  end
end
