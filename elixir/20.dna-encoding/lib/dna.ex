defmodule DNA do
  def encode_nucleotide(?\s), do: 0
  def encode_nucleotide(?A), do: 1
  def encode_nucleotide(?C), do: 2
  def encode_nucleotide(?G), do: 4
  def encode_nucleotide(?T), do: 8

  def decode_nucleotide(0), do: ?\s
  def decode_nucleotide(1), do: ?A
  def decode_nucleotide(2), do: ?C
  def decode_nucleotide(4), do: ?G
  def decode_nucleotide(8), do: ?T

  def encode(dna), do: do_encode(dna, "")

  defp do_encode([], acc), do: acc
  defp do_encode([head | tail], acc), do: do_encode(tail, "")


  def decode(dna) do
    # Please implement the decode/1 function
  end
end
