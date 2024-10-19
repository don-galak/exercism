defmodule KitchenCalculator do
  def get_volume(volume_pair), do: elem(volume_pair, 1)

  def to_milliliter({:milliliter, quantity}), do: {:milliliter, quantity}
  def to_milliliter({:cup, quantity}), do: {:milliliter, quantity * 240}
  def to_milliliter({:fluid_ounce, quantity}), do: {:milliliter, quantity * 30}
  def to_milliliter({:teaspoon, quantity}), do: {:milliliter, quantity * 5}
  def to_milliliter({:tablespoon, quantity}), do: {:milliliter, quantity * 15}

  def from_milliliter(volume_pair, :milliliter), do: volume_pair
  def from_milliliter(volume_pair, :cup), do: {:cup, get_volume(volume_pair) / 240}
  def from_milliliter(volume_pair, :fluid_ounce), do: {:fluid_ounce, get_volume(volume_pair) / 30}
  def from_milliliter(volume_pair, :teaspoon), do: {:teaspoon, get_volume(volume_pair) / 5}
  def from_milliliter(volume_pair, :tablespoon), do: {:tablespoon, get_volume(volume_pair) / 15}

  def convert(volume_pair, unit) do
    from_milliliter(to_milliliter(volume_pair), unit)
  end
end
