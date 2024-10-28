defmodule Username do
  def sanitize([]), do: []

  def sanitize([head | tail]) do
    username =
    case head do
      ?ä -> ~c"ae"
      ?ö -> ~c"oe"
      ?ü -> ~c"ue"
      ?ß -> ~c"ss"
      ?_ -> ~c"_"
      head when head >= ?a and head <= ?z -> [head]
      _ -> []
    end
    username ++ sanitize(tail)
  end
end
