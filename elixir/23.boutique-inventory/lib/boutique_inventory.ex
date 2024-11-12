defmodule BoutiqueInventory do
  def sort_by_price(inventory), do: Enum.sort(inventory, fn a, b -> a.price <= b.price end)

  def with_missing_price(inventory), do: Enum.filter(inventory, fn item -> item.price == nil end)

  def update_names(inventory, old_word, new_word) do
    Enum.map(inventory, fn %{name: name}=item ->
      %{item | name: String.replace(name, old_word, new_word)}
    end)
  end

  def increase_quantity(item, count) do
    # Please implement the increase_quantity/2 function
  end

  def total_quantity(item) do
    # Please implement the total_quantity/1 function
  end
end
