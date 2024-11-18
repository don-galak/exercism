defmodule BoutiqueSuggestions do
  @default_max_price 100.00
  def get_combinations(tops, bottoms, options \\ []) do
    max_price = Keyword.get(options, :maximum_price, @default_max_price)

    for top <- tops,
        bottom <- bottoms,
        top.base_color != bottom.base_color and
          top.price + bottom.price < max_price do
      {top, bottom}
    end
  end
end
