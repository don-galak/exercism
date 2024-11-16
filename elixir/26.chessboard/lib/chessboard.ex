defmodule Chessboard do
  def rank_range, do: 1..8

  def file_range, do: ?A..?H

  def ranks, do: Enum.to_list(rank_range())

  def files, do: Enum.map(Enum.to_list(file_range()), fn el -> <<el::utf8>> end)
end
