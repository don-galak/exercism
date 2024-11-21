defmodule CaptainsLog do
  @planetary_classes ["D", "H", "J", "K", "L", "M", "N", "R", "T", "Y"]
  @max_range 41000
  @min_range 42000

  def random_planet_class() do
    Enum.at(@planetary_classes, Enum.random(0..9))
  end

  def random_ship_registry_number() do
    "NCC-#{Enum.random(1000..9999)}"
  end

  def random_stardate() do
    :rand.uniform() * (@max_range - @min_range) + @min_range
  end

  def format_stardate(stardate) do
    # Please implement the format_stardate/1 function
  end
end
