defmodule BoutiqueInventory do
  def sort_by_price(inventory), do: Enum.sort_by(inventory, & &1.price)

  def with_missing_price(inventory), do: Enum.filter(inventory, fn item -> item.price == nil end)

  def update_names(inventory, old_word, new_word) do
    Enum.map(inventory, fn %{name: name} = item ->
      %{item | name: String.replace(name, old_word, new_word)}
    end)
  end

  def increase_quantity(item, count),
    do: %{
      item
      | quantity_by_size: Map.new(item.quantity_by_size, fn {key, val} -> {key, val + count} end)
    }

  def total_quantity(item),
    do: Enum.reduce(item.quantity_by_size, 0, fn {_, val}, acc -> acc + val end)
end
